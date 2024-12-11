// This file was auto-generated by Fern from our API Definition.

package bookings

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CreateBookingCustomAttributeDefinitionRequest struct {
	// The custom attribute definition to create, with the following fields:
	//
	// - `key`
	//
	// - `name`. If provided, `name` must be unique (case-sensitive) across all visible booking-related custom attribute
	// definitions for the seller.
	//
	// - `description`
	//
	// - `visibility`. Note that all custom attributes are visible in exported booking data, including those set to
	// `VISIBILITY_HIDDEN`.
	//
	// - `schema`. With the exception of the `Selection` data type, the `schema` is specified as a
	// simple URL to the JSON schema definition hosted on the Square CDN. For more information, see
	// [Specifying the schema](https://developer.squareup.com/docs/booking-custom-attributes-api/custom-attribute-definitions#specify-schema).
	CustomAttributeDefinition *squaregosdk.CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type CustomAttributeDefinitionsDeleteRequest struct {
	// The key of the custom attribute definition to delete.
	Key string `json:"-" url:"-"`
}

type CustomAttributeDefinitionsGetRequest struct {
	// The key of the custom attribute definition to retrieve. If the requesting application
	// is not the definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// The current version of the custom attribute definition, which is used for strongly consistent
	// reads to guarantee that you receive the most up-to-date data. When included in the request,
	// Square returns the specified version or a higher version if one exists. If the specified version
	// is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type CustomAttributeDefinitionsListRequest struct {
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type UpdateBookingCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to update.
	Key string `json:"-" url:"-"`
	// The custom attribute definition that contains the fields to update. Only the following fields can be updated:
	// - `name`
	// - `description`
	// - `visibility`
	// - `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
	// selections are supported.
	//
	// For more information, see
	// [Updatable definition fields](https://developer.squareup.com/docs/booking-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).
	//
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include the optional `version` field and specify the current version of the custom attribute definition.
	CustomAttributeDefinition *squaregosdk.CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}