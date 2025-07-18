// This file was auto-generated by Fern from our API Definition.

package customers

import (
	v2 "github.com/square/square-go-sdk/v2"
)

type DeleteCustomAttributesRequest struct {
	// The ID of the target [customer profile](entity:Customer).
	CustomerID string `json:"-" url:"-"`
	// The key of the custom attribute to delete. This key must match the `key` of a custom
	// attribute definition in the Square seller account. If the requesting application is not the
	// definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
}

type GetCustomAttributesRequest struct {
	// The ID of the target [customer profile](entity:Customer).
	CustomerID string `json:"-" url:"-"`
	// The key of the custom attribute to retrieve. This key must match the `key` of a custom
	// attribute definition in the Square seller account. If the requesting application is not the
	// definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of
	// the custom attribute. Set this parameter to `true` to get the name and description of the custom
	// attribute, information about the data type, or other definition details. The default value is `false`.
	WithDefinition *bool `json:"-" url:"with_definition,omitempty"`
	// The current version of the custom attribute, which is used for strongly consistent reads to
	// guarantee that you receive the most up-to-date data. When included in the request, Square
	// returns the specified version or a higher version if one exists. If the specified version is
	// higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type ListCustomAttributesRequest struct {
	// The ID of the target [customer profile](entity:Customer).
	CustomerID string `json:"-" url:"-"`
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request. For more
	// information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
	// custom attribute. Set this parameter to `true` to get the name and description of each custom
	// attribute, information about the data type, or other definition details. The default value is `false`.
	WithDefinitions *bool `json:"-" url:"with_definitions,omitempty"`
}

type UpsertCustomerCustomAttributeRequest struct {
	// The ID of the target [customer profile](entity:Customer).
	CustomerID string `json:"-" url:"-"`
	// The key of the custom attribute to create or update. This key must match the `key` of a
	// custom attribute definition in the Square seller account. If the requesting application is not
	// the definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// The custom attribute to create or update, with the following fields:
	//
	// - `value`. This value must conform to the `schema` specified by the definition.
	// For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).
	//
	// - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control for an update operation, include this optional field and specify the current version
	// of the custom attribute.
	CustomAttribute *v2.CustomAttribute `json:"custom_attribute,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}
