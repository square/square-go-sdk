//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/square/square-go-sdk/v2/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	square "github.com/square/square-go-sdk/v2"
)

// Orders API integration tests.
func TestOrdersAPI(t *testing.T) {
	t.Run("create order and attach payment", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		locationId, err := getDefaultLocationID(squareClient)
		require.NoError(t, err)

		testOrder := newTestOrder([]TestOrderItem{
			{
				Name:     "New York Strip Steak",
				Quantity: "1",
				Price:    1599,
			},
		}, locationId)

		// 1. Create an order and verify that it's `OPEN`.
		createOrderResp, err := squareClient.Orders.Create(
			context.Background(),
			&square.CreateOrderRequest{
				Order: testOrder,
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, createOrderResp.Order)
		assert.Len(t, createOrderResp.Order.LineItems, 1)
		assert.Equal(t, *createOrderResp.Order.State, square.OrderStateOpen)

		// Attach a payment to the order (equal to the total amount of the order).
		createPaymentResp, err := squareClient.Payments.Create(
			context.Background(),
			&square.CreatePaymentRequest{
				LocationID:     square.String(locationId),
				IdempotencyKey: newTestUUID(),
				OrderID:        createOrderResp.Order.ID,
				SourceID:       SourceID,
				AmountMoney:    newTestMoney(1599),
				Autocomplete:   square.Bool(true),
			},
		)
		require.NoError(t, err)
		assert.Equal(t, createPaymentResp.Payment.OrderID, createOrderResp.Order.ID)

		// Retrieve the order and verify it's `COMPLETED`.
		getOrderResp, err := squareClient.Orders.Get(
			context.Background(),
			&square.OrdersGetRequest{
				OrderID: *createOrderResp.Order.ID,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, *getOrderResp.Order.State, square.OrderStateCompleted)
	})

	t.Run("create and use a custom attribute for an order", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		locationId, err := getDefaultLocationID(squareClient)
		require.NoError(t, err)

		// Setup: Create an order with the custom attribute.
		testOrder := newTestOrder([]TestOrderItem{
			{
				Name:     "New York Strip Steak",
				Quantity: "1",
				Price:    1599,
			},
		}, locationId)
		createOrderResp, err := squareClient.Orders.Create(
			context.Background(),
			&square.CreateOrderRequest{
				Order: testOrder,
			},
		)
		require.NoError(t, err)

		// 1. Create a custom attribute definition for an email address schema.
		emailSchemaUri := "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.Email"
		attrKey := fmt.Sprintf("custom-attr-%s", newTestUUID())
		customAttrDefinition := &square.CustomAttributeDefinition{
			Key:  square.String(attrKey),
			Name: square.String("My Custom Attribute"),
			Schema: map[string]interface{}{
				"$ref": square.String(emailSchemaUri),
			},
		}
		createAttrResp, err := squareClient.Orders.CustomAttributeDefinitions.Create(
			context.Background(),
			&orders.CreateOrderCustomAttributeDefinitionRequest{
				CustomAttributeDefinition: customAttrDefinition,
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, createAttrResp.CustomAttributeDefinition)

		// 2. For the given order, upsert a valid value for the custom attribute.
		upsertOrderResp, err := squareClient.Orders.CustomAttributes.Upsert(
			context.Background(),
			&orders.UpsertOrderCustomAttributeRequest{
				OrderID:            *createOrderResp.Order.ID,
				CustomAttributeKey: attrKey,
				CustomAttribute: &square.CustomAttribute{
					Value: square.String("foo@example.com"),
				},
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, upsertOrderResp.CustomAttribute)

		// 3. For the given order, upsert an invalid value for the custom attribute.
		_, err = squareClient.Orders.CustomAttributes.Upsert(
			context.Background(),
			&orders.UpsertOrderCustomAttributeRequest{
				OrderID:            *createOrderResp.Order.ID,
				CustomAttributeKey: attrKey,
				CustomAttribute: &square.CustomAttribute{
					Value: square.String("invalid-email"),
				},
			},
		)
		require.Error(t, err)
	})
}
