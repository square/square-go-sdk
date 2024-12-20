// This file was auto-generated by Fern from our API Definition.

package locations

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CreateLocationCustomAttributeDefinitionRequest struct {
	// The custom attribute definition to create. Note the following:
	// - With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
	// definition hosted on the Square CDN. For more information, including supported values and constraints, see
	// [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
	// - `name` is required unless `visibility` is set to `VISIBILITY_HIDDEN`.
	CustomAttributeDefinition *squaregosdk.CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
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
	// Filters the `CustomAttributeDefinition` results by their `visibility` values.
	VisibilityFilter *squaregosdk.VisibilityFilter `json:"-" url:"visibility_filter,omitempty"`
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
