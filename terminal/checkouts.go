// This file was auto-generated by Fern from our API Definition.

package terminal

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CancelTerminalCheckoutRequest struct {
}

type CreateTerminalCheckoutRequest struct {
	// A unique string that identifies this `CreateCheckout` request. Keys can be any valid string but
	// must be unique for every `CreateCheckout` request.
	//
	// See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The checkout to create.
	Checkout *squaregosdk.TerminalCheckout `json:"checkout,omitempty" url:"-"`
}

type SearchTerminalCheckoutsRequest struct {
	// Queries Terminal checkouts based on given conditions and the sort order.
	// Leaving these unset returns all checkouts with the default sort order.
	Query *squaregosdk.TerminalCheckoutQuery `json:"query,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"cursor,omitempty" url:"-"`
	// Limits the number of results returned for a single request.
	Limit *int `json:"limit,omitempty" url:"-"`
}
