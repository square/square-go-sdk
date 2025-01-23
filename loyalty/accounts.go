// This file was auto-generated by Fern from our API Definition.

package loyalty

import (
	v2 "github.com/square/square-go-sdk/v2"
)

type AccumulateLoyaltyPointsRequest struct {
	// The ID of the target [loyalty account](entity:LoyaltyAccount).
	AccountID string `json:"-" url:"-"`
	// The points to add to the account.
	// If you are using the Orders API to manage orders, specify the order ID.
	// Otherwise, specify the points to add.
	AccumulatePoints *v2.LoyaltyEventAccumulatePoints `json:"accumulate_points,omitempty" url:"-"`
	// A unique string that identifies the `AccumulateLoyaltyPoints` request.
	// Keys can be any valid string but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The [location](entity:Location) where the purchase was made.
	LocationID string `json:"location_id" url:"-"`
}

type AdjustLoyaltyPointsRequest struct {
	// The ID of the target [loyalty account](entity:LoyaltyAccount).
	AccountID string `json:"-" url:"-"`
	// A unique string that identifies this `AdjustLoyaltyPoints` request.
	// Keys can be any valid string, but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The points to add or subtract and the reason for the adjustment. To add points, specify a positive integer.
	// To subtract points, specify a negative integer.
	AdjustPoints *v2.LoyaltyEventAdjustPoints `json:"adjust_points,omitempty" url:"-"`
	// Indicates whether to allow a negative adjustment to result in a negative balance. If `true`, a negative
	// balance is allowed when subtracting points. If `false`, Square returns a `BAD_REQUEST` error when subtracting
	// the specified number of points would result in a negative balance. The default value is `false`.
	AllowNegativeBalance *bool `json:"allow_negative_balance,omitempty" url:"-"`
}

type CreateLoyaltyAccountRequest struct {
	// The loyalty account to create.
	LoyaltyAccount *v2.LoyaltyAccount `json:"loyalty_account,omitempty" url:"-"`
	// A unique string that identifies this `CreateLoyaltyAccount` request.
	// Keys can be any valid string, but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
}

type AccountsGetRequest struct {
	// The ID of the [loyalty account](entity:LoyaltyAccount) to retrieve.
	AccountID string `json:"-" url:"-"`
}

type SearchLoyaltyAccountsRequest struct {
	// The search criteria for the request.
	Query *v2.SearchLoyaltyAccountsRequestLoyaltyAccountQuery `json:"query,omitempty" url:"-"`
	// The maximum number of results to include in the response. The default value is 30.
	Limit *int `json:"limit,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to
	// this endpoint. Provide this to retrieve the next set of
	// results for the original query.
	//
	// For more information,
	// see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
}
