//go:build integration

package integration

import (
	square "github.com/square/square-go-sdk/v2"
)

// Simplified data structure for [square.OrderLineItem].
type TestOrderItem struct {
	Name     string
	Quantity string
	Price    int64
}

// Create a test order with the given order items.
func newTestOrder(orderItems []TestOrderItem, locationId string) *square.Order {
	squareItems := make([]*square.OrderLineItem, 0, len(orderItems))

	for _, item := range orderItems {
		squareItems = append(squareItems, &square.OrderLineItem{
			Name:           square.String(item.Name),
			Quantity:       item.Quantity,
			BasePriceMoney: newTestMoney(item.Price),
		})
	}

	return &square.Order{
		LocationID: locationId,
		LineItems:  squareItems,
	}
}
