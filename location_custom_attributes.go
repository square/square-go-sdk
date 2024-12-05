// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type BulkDeleteLocationCustomAttributesRequest struct {
	// The data used to update the `CustomAttribute` objects.
	// The keys must be unique and are used to map to the corresponding response.
	Values map[string]*BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest `json:"values,omitempty" url:"-"`
}

type BulkUpsertLocationCustomAttributesRequest struct {
	// A map containing 1 to 25 individual upsert requests. For each request, provide an
	// arbitrary ID that is unique for this `BulkUpsertLocationCustomAttributes` request and the
	// information needed to create or update a custom attribute.
	Values map[string]*BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest `json:"values,omitempty" url:"-"`
}

type CreateLocationCustomAttributeDefinitionRequest struct {
	// The custom attribute definition to create. Note the following:
	// - With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
	// definition hosted on the Square CDN. For more information, including supported values and constraints, see
	// [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
	// - `name` is required unless `visibility` is set to `VISIBILITY_HIDDEN`.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type DeleteLocationCustomAttributeRequest struct {
	// The ID of the target [location](entity:Location).
	LocationID string `json:"-" url:"-"`
	// The key of the custom attribute to delete. This key must match the `key` of a custom
	// attribute definition in the Square seller account. If the requesting application is not the
	// definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
}

type DeleteLocationCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to delete.
	Key string `json:"-" url:"-"`
}

type ListLocationCustomAttributesRequest struct {
	// The ID of the target [location](entity:Location).
	LocationID string `json:"-" url:"-"`
	// Filters the `CustomAttributeDefinition` results by their `visibility` values.
	VisibilityFilter *VisibilityFilter `json:"-" url:"visibility_filter,omitempty"`
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

type RetrieveLocationCustomAttributeRequest struct {
	// The ID of the target [location](entity:Location).
	LocationID string `json:"-" url:"-"`
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

type RetrieveLocationCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to retrieve. If the requesting application
	// is not the definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// The current version of the custom attribute definition, which is used for strongly consistent
	// reads to guarantee that you receive the most up-to-date data. When included in the request,
	// Square returns the specified version or a higher version if one exists. If the specified version
	// is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type UpdateLocationCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to update.
	Key string `json:"-" url:"-"`
	// The custom attribute definition that contains the fields to update. This endpoint
	// supports sparse updates, so only new or changed fields need to be included in the request.
	// Only the following fields can be updated:
	// - `name`
	// - `description`
	// - `visibility`
	// - `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
	// selections are supported.
	//
	// For more information, see
	// [Update a location custom attribute definition](https://developer.squareup.com/docs/location-custom-attributes-api/custom-attribute-definitions#update-custom-attribute-definition).
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, specify the current version of the custom attribute definition.
	// If this is not important for your application, `version` can be set to -1.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type UpsertLocationCustomAttributeRequest struct {
	// The ID of the target [location](entity:Location).
	LocationID string `json:"-" url:"-"`
	// The key of the custom attribute to create or update. This key must match the `key` of a
	// custom attribute definition in the Square seller account. If the requesting application is not
	// the definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// The custom attribute to create or update, with the following fields:
	// - `value`. This value must conform to the `schema` specified by the definition.
	// For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
	// - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control for an update operation, include the current version of the custom attribute.
	// If this is not important for your application, version can be set to -1.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

// Represents an individual delete request in a [BulkDeleteLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkDeleteLocationCustomAttributes)
// request. An individual request contains an optional ID of the associated custom attribute definition
// and optional key of the associated custom attribute definition.
type BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest struct {
	// The key of the associated custom attribute definition.
	// Represented as a qualified key if the requesting app is not the definition owner.
	Key *string `json:"key,omitempty" url:"key,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest) GetKey() *string {
	if b == nil {
		return nil
	}
	return b.Key
}

func (b *BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [BulkDeleteLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkDeleteLocationCustomAttributes) response,
// which contains a map of responses that each corresponds to an individual delete request.
type BulkDeleteLocationCustomAttributesResponse struct {
	// A map of responses that correspond to individual delete requests. Each response has the
	// same key as the corresponding request.
	Values map[string]*BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse `json:"values,omitempty" url:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDeleteLocationCustomAttributesResponse) GetValues() map[string]*BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkDeleteLocationCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkDeleteLocationCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteLocationCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteLocationCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteLocationCustomAttributesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteLocationCustomAttributesResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents an individual delete response in a [BulkDeleteLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkDeleteLocationCustomAttributes)
// request.
type BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse struct {
	// The ID of the location associated with the custom attribute.
	LocationID *string `json:"location_id,omitempty" url:"location_id,omitempty"`
	// Errors that occurred while processing the individual LocationCustomAttributeDeleteRequest request
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse) GetLocationID() *string {
	if b == nil {
		return nil
	}
	return b.LocationID
}

func (b *BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents an individual upsert request in a [BulkUpsertLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkUpsertLocationCustomAttributes)
// request. An individual request contains a location ID, the custom attribute to create or update,
// and an optional idempotency key.
type BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest struct {
	// The ID of the target [location](entity:Location).
	LocationID string `json:"location_id" url:"location_id"`
	// The custom attribute to create or update, with following fields:
	//
	//   - `key`. This key must match the `key` of a custom attribute definition in the Square seller
	//     account. If the requesting application is not the definition owner, you must provide the qualified key.
	//   - `value`. This value must conform to the `schema` specified by the definition.
	//     For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types)..
	//   - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	//     control, specify the current version of the custom attribute.
	//     If this is not important for your application, `version` can be set to -1.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// A unique identifier for this individual upsert request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"idempotency_key,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) GetLocationID() string {
	if b == nil {
		return ""
	}
	return b.LocationID
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) GetCustomAttribute() *CustomAttribute {
	if b == nil {
		return nil
	}
	return b.CustomAttribute
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) GetIdempotencyKey() *string {
	if b == nil {
		return nil
	}
	return b.IdempotencyKey
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [BulkUpsertLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkUpsertLocationCustomAttributes) response,
// which contains a map of responses that each corresponds to an individual upsert request.
type BulkUpsertLocationCustomAttributesResponse struct {
	// A map of responses that correspond to individual upsert requests. Each response has the
	// same ID as the corresponding request and contains either a `location_id` and `custom_attribute` or an `errors` field.
	Values map[string]*BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse `json:"values,omitempty" url:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkUpsertLocationCustomAttributesResponse) GetValues() map[string]*BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkUpsertLocationCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkUpsertLocationCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertLocationCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertLocationCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertLocationCustomAttributesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertLocationCustomAttributesResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a response for an individual upsert request in a [BulkUpsertLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkUpsertLocationCustomAttributes) operation.
type BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse struct {
	// The ID of the location associated with the custom attribute.
	LocationID *string `json:"location_id,omitempty" url:"location_id,omitempty"`
	// The new or updated custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) GetLocationID() *string {
	if b == nil {
		return nil
	}
	return b.LocationID
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) GetCustomAttribute() *CustomAttribute {
	if b == nil {
		return nil
	}
	return b.CustomAttribute
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [CreateLocationCustomAttributeDefinition](api-endpoint:LocationCustomAttributes-CreateLocationCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type CreateLocationCustomAttributeDefinitionResponse struct {
	// The new custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateLocationCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if c == nil {
		return nil
	}
	return c.CustomAttributeDefinition
}

func (c *CreateLocationCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateLocationCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateLocationCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateLocationCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateLocationCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateLocationCustomAttributeDefinitionResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Represents a response from a delete request containing error messages if there are any.
type DeleteLocationCustomAttributeDefinitionResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteLocationCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteLocationCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteLocationCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteLocationCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteLocationCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteLocationCustomAttributeDefinitionResponse) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Represents a [DeleteLocationCustomAttribute](api-endpoint:LocationCustomAttributes-DeleteLocationCustomAttribute) response.
// Either an empty object `{}` (for a successful deletion) or `errors` is present in the response.
type DeleteLocationCustomAttributeResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteLocationCustomAttributeResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteLocationCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteLocationCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteLocationCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteLocationCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteLocationCustomAttributeResponse) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Represents a [ListLocationCustomAttributes](api-endpoint:LocationCustomAttributes-ListLocationCustomAttributes) response.
// Either `custom_attributes`, an empty object, or `errors` is present in the response. If additional
// results are available, the `cursor` field is also present along with `custom_attributes`.
type ListLocationCustomAttributesResponse struct {
	// The retrieved custom attributes. If `with_definitions` was set to `true` in the request,
	// the custom attribute definition is returned in the `definition` field of each custom attribute.
	// If no custom attributes are found, Square returns an empty object (`{}`).
	CustomAttributes []*CustomAttribute `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
	// The cursor to use in your next call to this endpoint to retrieve the next page of results
	// for your original request. This field is present only if the request succeeded and additional
	// results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListLocationCustomAttributesResponse) GetCustomAttributes() []*CustomAttribute {
	if l == nil {
		return nil
	}
	return l.CustomAttributes
}

func (l *ListLocationCustomAttributesResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListLocationCustomAttributesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListLocationCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListLocationCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListLocationCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListLocationCustomAttributesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListLocationCustomAttributesResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Represents a [RetrieveLocationCustomAttributeDefinition](api-endpoint:LocationCustomAttributes-RetrieveLocationCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveLocationCustomAttributeDefinitionResponse struct {
	// The retrieved custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RetrieveLocationCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if r == nil {
		return nil
	}
	return r.CustomAttributeDefinition
}

func (r *RetrieveLocationCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveLocationCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveLocationCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveLocationCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveLocationCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveLocationCustomAttributeDefinitionResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Represents a [RetrieveLocationCustomAttribute](api-endpoint:LocationCustomAttributes-RetrieveLocationCustomAttribute) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveLocationCustomAttributeResponse struct {
	// The retrieved custom attribute. If `with_definition` was set to `true` in the request,
	// the custom attribute definition is returned in the `definition` field.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RetrieveLocationCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if r == nil {
		return nil
	}
	return r.CustomAttribute
}

func (r *RetrieveLocationCustomAttributeResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveLocationCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveLocationCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveLocationCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveLocationCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveLocationCustomAttributeResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Represents an [UpdateLocationCustomAttributeDefinition](api-endpoint:LocationCustomAttributes-UpdateLocationCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpdateLocationCustomAttributeDefinitionResponse struct {
	// The updated custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateLocationCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if u == nil {
		return nil
	}
	return u.CustomAttributeDefinition
}

func (u *UpdateLocationCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateLocationCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateLocationCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateLocationCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateLocationCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateLocationCustomAttributeDefinitionResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

// Represents an [UpsertLocationCustomAttribute](api-endpoint:LocationCustomAttributes-UpsertLocationCustomAttribute) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpsertLocationCustomAttributeResponse struct {
	// The new or updated custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpsertLocationCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if u == nil {
		return nil
	}
	return u.CustomAttribute
}

func (u *UpsertLocationCustomAttributeResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpsertLocationCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpsertLocationCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertLocationCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertLocationCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertLocationCustomAttributeResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
