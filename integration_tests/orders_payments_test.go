package integration_tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fern-demo/square-go-sdk"
	"github.com/google/uuid"
)

func TestOrderPaymentLifecycle(t *testing.T) {
	squareClient := initSquareClient()

	// Grab necessary information
	// Location ID
	listLocationsResp, err := squareClient.Locations.List(context.TODO())
	if err != nil {
		t.Fatalf("failed to list locations: %v", err)
	}
	locationID := *listLocationsResp.Locations[0].ID

	// Step 1: Create an order with a line item. not attached to a catalog item
	createOrderResp, err := squareClient.Orders.Create(context.TODO(), &square.CreateOrderRequest{
		IdempotencyKey: square.String(uuid.New().String()),
		Order: &square.Order{
			LocationID: locationID,
			LineItems: []*square.OrderLineItem{
				{
					Quantity: "1",
					BasePriceMoney: &square.Money{
						Amount:   square.Int64(1000),
						Currency: square.CurrencyUsd.Ptr(),
					},
					CatalogObjectID: nil, // not attached to a catalog item
					Name:            square.String("test-line-item-name"),
				},
			},
			Fulfillments: []*square.Fulfillment{
				{
					Type: square.FulfillmentTypePickup.Ptr(),
					PickupDetails: &square.FulfillmentPickupDetails{
						Recipient: &square.FulfillmentRecipient{
							DisplayName: square.String("test-recipient-name"),
						},
						PickupAt: square.String(time.Now().Add(time.Hour).Format(time.RFC3339)),
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create order: %v", err)
	}

	fmt.Printf("Created order: %v\n", createOrderResp.Order)

	// Step 2: Pay for the order
	createPaymentResp, err := squareClient.Payments.Create(context.TODO(), &square.CreatePaymentRequest{
		LocationID:     &locationID,
		IdempotencyKey: uuid.New().String(),
		AmountMoney: &square.Money{
			Amount:   square.Int64(1000),
			Currency: square.CurrencyUsd.Ptr(),
		},
		SourceID: "cnon:card-nonce-ok",
		OrderID:  createOrderResp.Order.ID,
	})
	if err != nil {
		t.Fatalf("failed to create payment: %v", err)
	}

	fmt.Printf("Created payment: %v\n", createPaymentResp.Payment)

	// Step 3: Fulfill the order
	getOrderResp, err := squareClient.Orders.Get(context.TODO(), *createOrderResp.Order.ID)
	if err != nil {
		t.Fatalf("failed to get order: %v", err)
	}
	orderToFulfill := getOrderResp.Order

	orderToFulfill.Fulfillments[0].State = square.FulfillmentStateCompleted.Ptr()
	orderToFulfill.State = square.OrderStateCompleted.Ptr()

	updateOrderResp, err := squareClient.Orders.Update(context.TODO(), *orderToFulfill.ID, &square.UpdateOrderRequest{
		Order:          orderToFulfill,
		IdempotencyKey: square.String(uuid.New().String()),
	})
	if err != nil {
		t.Fatalf("failed to update order: %v", err)
	}

	fmt.Printf("Updated order: %v\n", updateOrderResp.Order)

	// Step 4: Refund the payment
	createRefundResp, err := squareClient.Refunds.RefundPayment(context.TODO(), &square.RefundPaymentRequest{
		AmountMoney: &square.Money{
			Amount:   square.Int64(1000),
			Currency: square.CurrencyUsd.Ptr(),
		},
		PaymentID:      createPaymentResp.Payment.ID,
		IdempotencyKey: uuid.New().String(),
	})
	if err != nil {
		t.Fatalf("failed to refund payment: %v", err)
	}

	fmt.Printf("Refunded payment: %v\n", createRefundResp.Refund)

	// Step 5: Search for the completed order. should be the latest one
	searchOrdersResp, err := squareClient.Orders.Search(context.TODO(), &square.SearchOrdersRequest{
		LocationIDs: []string{locationID},
		Query: &square.SearchOrdersQuery{
			Filter: &square.SearchOrdersFilter{
				StateFilter: &square.SearchOrdersStateFilter{
					States: []square.OrderState{square.OrderStateCompleted},
				},
			},
			Sort: &square.SearchOrdersSort{
				SortField: square.SearchOrdersSortFieldClosedAt,
				SortOrder: square.SortOrderDesc.Ptr(),
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to search orders: %v", err)
	}

	if len(searchOrdersResp.Orders) == 0 {
		t.Fatalf("no completed orders found")
	}

	fmt.Printf("Found completed order: %v\n", searchOrdersResp.Orders[0])

	// validate id matches
	if *searchOrdersResp.Orders[0].ID != *createOrderResp.Order.ID {
		t.Fatalf("expected order id %s, got %v", *createOrderResp.Order.ID, *searchOrdersResp.Orders[0].ID)
	}
}

func TestOrderCustomAttributesLifecycle(t *testing.T) {
	squareClient := initSquareClient()

	// Grab necessary information
	// Location ID
	listLocationsResp, err := squareClient.Locations.List(context.TODO())
	if err != nil {
		t.Fatalf("failed to list locations: %v", err)
	}
	locationID := *listLocationsResp.Locations[0].ID

	// Step 1: Create an order custom attribute definition
	createOrderCustomAttributeDefinitionResp, err := squareClient.OrderCustomAttributes.CreateOrderCustomAttributeDefinition(context.TODO(), &square.CreateOrderCustomAttributeDefinitionRequest{
		IdempotencyKey: square.String(uuid.New().String()),
		CustomAttributeDefinition: &square.CustomAttributeDefinition{
			Key:  square.String("test-address"),
			Name: square.String("Test address attribute"),
			Schema: map[string]interface{}{
				"$ref": "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.Address",
			},
		},
	})

	if err != nil {
		t.Fatalf("failed to create order custom attribute definition: %v", err)
	}

	fmt.Printf("Created order custom attribute definition: %v\n", createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition)

	// Step 2: Retrieve the custom attribute definition
	getOrderCustomAttributeDefinitionResp, err := squareClient.OrderCustomAttributes.RetrieveOrderCustomAttributeDefinition(context.TODO(), *createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key, nil)
	if err != nil {
		t.Fatalf("failed to get order custom attribute definition: %v", err)
	}

	fmt.Printf("Retrieved order custom attribute definition: %v\n", getOrderCustomAttributeDefinitionResp.CustomAttributeDefinition)

	// Step 3: Update the custom attribute definition
	customAttributeDefinitionToUpdate := getOrderCustomAttributeDefinitionResp.CustomAttributeDefinition
	customAttributeDefinitionToUpdate.Description = square.String("Updated description")
	updateOrderCustomAttributeDefinitionResp, err := squareClient.OrderCustomAttributes.UpdateOrderCustomAttributeDefinition(context.TODO(), *createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key, &square.UpdateOrderCustomAttributeDefinitionRequest{
		CustomAttributeDefinition: customAttributeDefinitionToUpdate,
		IdempotencyKey:            square.String(uuid.New().String()),
	})
	if err != nil {
		t.Fatalf("failed to update order custom attribute definition: %v", err)
	}

	fmt.Printf("Updated order custom attribute definition: %v\n", updateOrderCustomAttributeDefinitionResp.CustomAttributeDefinition)

	// Step 4: Create an order with the custom attribute
	createOrderResp, err := squareClient.Orders.Create(context.TODO(), &square.CreateOrderRequest{
		IdempotencyKey: square.String(uuid.New().String()),
		Order: &square.Order{
			LocationID: locationID,
			LineItems: []*square.OrderLineItem{
				{
					Quantity: "1",
					BasePriceMoney: &square.Money{
						Amount:   square.Int64(1000),
						Currency: square.CurrencyUsd.Ptr(),
					},
					CatalogObjectID: nil, // not attached to a catalog item
					Name:            square.String("test-line-item-name"),
				},
			},
		},
	})

	// Step 5: Add the custom attribute to the order
	createOrderCustomAttributeResp, err := squareClient.OrderCustomAttributes.UpsertOrderCustomAttribute(
		context.TODO(),
		*createOrderResp.Order.ID,
		*createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key,
		&square.UpsertOrderCustomAttributeRequest{
			CustomAttribute: &square.CustomAttribute{
				Value: map[string]string{
					"address_line_1":                  "1455 Market Street",
					"postal_code":                     "94103",
					"country":                         "US",
					"administrative_district_level_1": "California",
					"locality":                        "San Francisco",
				},
				Key: createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key,
			},
		})
	if err != nil {
		t.Fatalf("failed to attach order custom attribute: %v", err)
	}

	fmt.Printf("Created order custom attribute: %v\n", createOrderCustomAttributeResp.CustomAttribute)

	// Step 6: List the custom attribute attached to the order
	listOrderCustomAttributesResp, err := squareClient.OrderCustomAttributes.ListOrderCustomAttributes(context.TODO(), *createOrderResp.Order.ID, nil)
	if err != nil {
		t.Fatalf("failed to get order custom attribute: %v", err)
	}

	fmt.Printf("Retrieved order custom attributes: %v\n", listOrderCustomAttributesResp.CustomAttributes)

	// Validate we have the correct custom attribute
	if len(listOrderCustomAttributesResp.CustomAttributes) != 1 {
		t.Fatalf("expected 1 custom attribute, got %d", len(listOrderCustomAttributesResp.CustomAttributes))
	}
	if *listOrderCustomAttributesResp.CustomAttributes[0].Key != *createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key {
		t.Fatalf("expected custom attribute key %s, got %v", *createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key, *listOrderCustomAttributesResp.CustomAttributes[0].Key)
	}

	// Delete the custom attribute definition
	deleteOrderCustomAttributeDefinitionResp, err := squareClient.OrderCustomAttributes.DeleteOrderCustomAttributeDefinition(context.TODO(), *createOrderCustomAttributeDefinitionResp.CustomAttributeDefinition.Key)
	if err != nil {
		t.Fatalf("failed to delete order custom attribute definition: %v", err)
	}

	fmt.Printf("Deleted order custom attribute definition: %v\n", deleteOrderCustomAttributeDefinitionResp)
}
