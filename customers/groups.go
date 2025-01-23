// This file was auto-generated by Fern from our API Definition.

package customers

import (
	v2 "github.com/square/square-go-sdk/v2"
)

type GroupsAddRequest struct {
	// The ID of the customer to add to a group.
	CustomerID string `json:"-" url:"-"`
	// The ID of the customer group to add the customer to.
	GroupID string `json:"-" url:"-"`
}

type CreateCustomerGroupRequest struct {
	// The idempotency key for the request. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The customer group to create.
	Group *v2.CustomerGroup `json:"group,omitempty" url:"-"`
}

type GroupsDeleteRequest struct {
	// The ID of the customer group to delete.
	GroupID string `json:"-" url:"-"`
}

type GroupsGetRequest struct {
	// The ID of the customer group to retrieve.
	GroupID string `json:"-" url:"-"`
}

type GroupsListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results.
	// If the limit is less than 1 or greater than 50, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 50.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
}

type GroupsRemoveRequest struct {
	// The ID of the customer to remove from the group.
	CustomerID string `json:"-" url:"-"`
	// The ID of the customer group to remove the customer from.
	GroupID string `json:"-" url:"-"`
}

type UpdateCustomerGroupRequest struct {
	// The ID of the customer group to update.
	GroupID string `json:"-" url:"-"`
	// The `CustomerGroup` object including all the updates you want to make.
	Group *v2.CustomerGroup `json:"group,omitempty" url:"-"`
}
