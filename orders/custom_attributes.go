// This file was auto-generated by Fern from our API Definition.

package orders

import (
	v2 "github.com/square/square-go-sdk/v2"
)

type BulkDeleteOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual delete operations for custom attributes.
	Values map[string]*v2.BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute `json:"values,omitempty" url:"-"`
}

type BulkUpsertOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual upsert operations for custom attributes.
	Values map[string]*v2.BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute `json:"values,omitempty" url:"-"`
}

type CustomAttributesDeleteRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to delete.  This key must match the key of an
	// existing custom attribute definition.
	CustomAttributeKey string `json:"-" url:"-"`
}

type CustomAttributesGetRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to retrieve.  This key must match the key of an
	// existing custom attribute definition.
	CustomAttributeKey string `json:"-" url:"-"`
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include this optional field and specify the current version of the custom attribute.
	Version *int `json:"-" url:"version,omitempty"`
	// Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
	// custom attribute. Set this parameter to `true` to get the name and description of each custom attribute,
	// information about the data type, or other definition details. The default value is `false`.
	WithDefinition *bool `json:"-" url:"with_definition,omitempty"`
}

type CustomAttributesListRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// Requests that all of the custom attributes be returned, or only those that are read-only or read-write.
	VisibilityFilter *v2.VisibilityFilter `json:"-" url:"visibility_filter,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request.
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20.
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
	// custom attribute. Set this parameter to `true` to get the name and description of each custom attribute,
	// information about the data type, or other definition details. The default value is `false`.
	WithDefinitions *bool `json:"-" url:"with_definitions,omitempty"`
}

type UpsertOrderCustomAttributeRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to create or update.  This key must match the key
	// of an existing custom attribute definition.
	CustomAttributeKey string `json:"-" url:"-"`
	// The custom attribute to create or update, with the following fields:
	//
	// - `value`. This value must conform to the `schema` specified by the definition.
	// For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).
	//
	// - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include this optional field and specify the current version of the custom attribute.
	CustomAttribute *v2.CustomAttribute `json:"custom_attribute,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}
