// This file was auto-generated by Fern from our API Definition.

package loyalty

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CalculateLoyaltyPointsRequest struct {
	// The ID of the [loyalty program](entity:LoyaltyProgram), which defines the rules for accruing points.
	ProgramID string `json:"-" url:"-"`
	// The [order](entity:Order) ID for which to calculate the points.
	// Specify this field if your application uses the Orders API to process orders.
	// Otherwise, specify the `transaction_amount_money`.
	OrderID *string `json:"order_id,omitempty" url:"-"`
	// The purchase amount for which to calculate the points.
	// Specify this field if your application does not use the Orders API to process orders.
	// Otherwise, specify the `order_id`.
	TransactionAmountMoney *squaregosdk.Money `json:"transaction_amount_money,omitempty" url:"-"`
	// The ID of the target [loyalty account](entity:LoyaltyAccount). Optionally specify this field
	// if your application uses the Orders API to process orders.
	//
	// If specified, the `promotion_points` field in the response shows the number of points the buyer would
	// earn from the purchase. In this case, Square uses the account ID to determine whether the promotion's
	// `trigger_limit` (the maximum number of times that a buyer can trigger the promotion) has been reached.
	// If not specified, the `promotion_points` field shows the number of points the purchase qualifies
	// for regardless of the trigger limit.
	LoyaltyAccountID *string `json:"loyalty_account_id,omitempty" url:"-"`
}

type ProgramsGetRequest struct {
	// The ID of the loyalty program or the keyword `main`. Either value can be used to retrieve the single loyalty program that belongs to the seller.
	ProgramID string `json:"-" url:"-"`
}
