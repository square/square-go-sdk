// This file was auto-generated by Fern from our API Definition.

package labor

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CreateBreakTypeRequest struct {
	// A unique string value to ensure the idempotency of the operation.
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The `BreakType` to be created.
	BreakType *squaregosdk.BreakType `json:"break_type,omitempty" url:"-"`
}

type DeleteBreakTypesRequest struct {
	// The UUID for the `BreakType` being deleted.
	ID string `json:"-" url:"-"`
}

type GetBreakTypesRequest struct {
	// The UUID for the `BreakType` being retrieved.
	ID string `json:"-" url:"-"`
}

type ListBreakTypesRequest struct {
	// Filter the returned `BreakType` results to only those that are associated with the
	// specified location.
	LocationID *string `json:"-" url:"location_id,omitempty"`
	// The maximum number of `BreakType` results to return per page. The number can range between 1
	// and 200. The default is 200.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pointer to the next page of `BreakType` results to fetch.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type UpdateBreakTypeRequest struct {
	// The UUID for the `BreakType` being updated.
	ID string `json:"-" url:"-"`
	// The updated `BreakType`.
	BreakType *squaregosdk.BreakType `json:"break_type,omitempty" url:"-"`
}
