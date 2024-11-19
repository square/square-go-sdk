// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type BulkDeleteMerchantCustomAttributesRequest struct {
	// The data used to update the `CustomAttribute` objects.
	// The keys must be unique and are used to map to the corresponding response.
	Values map[string]*BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest `json:"values,omitempty" url:"-"`
}

type BulkUpsertMerchantCustomAttributesRequest struct {
	// A map containing 1 to 25 individual upsert requests. For each request, provide an
	// arbitrary ID that is unique for this `BulkUpsertMerchantCustomAttributes` request and the
	// information needed to create or update a custom attribute.
	Values map[string]*BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest `json:"values,omitempty" url:"-"`
}

type CreateMerchantCustomAttributeDefinitionRequest struct {
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

type ListMerchantCustomAttributeDefinitionsRequest struct {
	// Filters the `CustomAttributeDefinition` results by their `visibility` values.
	VisibilityFilter *VisibilityFilter `json:"-" url:"visibility_filter,omitempty"`
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type ListMerchantCustomAttributesRequest struct {
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

type RetrieveMerchantCustomAttributeRequest struct {
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

type RetrieveMerchantCustomAttributeDefinitionRequest struct {
	// The current version of the custom attribute definition, which is used for strongly consistent
	// reads to guarantee that you receive the most up-to-date data. When included in the request,
	// Square returns the specified version or a higher version if one exists. If the specified version
	// is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type UpdateMerchantCustomAttributeDefinitionRequest struct {
	// The custom attribute definition that contains the fields to update. This endpoint
	// supports sparse updates, so only new or changed fields need to be included in the request.
	// Only the following fields can be updated:
	// - `name`
	// - `description`
	// - `visibility`
	// - `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
	// selections are supported.
	// For more information, see
	// [Update a merchant custom attribute definition](https://developer.squareup.com/docs/merchant-custom-attributes-api/custom-attribute-definitions#update-custom-attribute-definition).
	// The version field must match the current version of the custom attribute definition to enable
	// [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// If this is not important for your application, version can be set to -1. For any other values, the request fails with a BAD_REQUEST error.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type UpsertMerchantCustomAttributeRequest struct {
	// The custom attribute to create or update, with the following fields:
	// - `value`. This value must conform to the `schema` specified by the definition.
	// For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
	// - The version field must match the current version of the custom attribute definition to enable
	// [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// If this is not important for your application, version can be set to -1. For any other values, the request fails with a BAD_REQUEST error.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

// Represents an individual delete request in a [BulkDeleteMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkDeleteMerchantCustomAttributes)
// request. An individual request contains an optional ID of the associated custom attribute definition
// and optional key of the associated custom attribute definition.
type BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest struct {
	// The key of the associated custom attribute definition.
	// Represented as a qualified key if the requesting app is not the definition owner.
	Key *string `json:"key,omitempty" url:"key,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest) GetKey() *string {
	if b == nil {
		return nil
	}
	return b.Key
}

func (b *BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [BulkDeleteMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkDeleteMerchantCustomAttributes) response,
// which contains a map of responses that each corresponds to an individual delete request.
type BulkDeleteMerchantCustomAttributesResponse struct {
	// A map of responses that correspond to individual delete requests. Each response has the
	// same key as the corresponding request.
	Values map[string]*BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse `json:"values,omitempty" url:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkDeleteMerchantCustomAttributesResponse) GetValues() map[string]*BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkDeleteMerchantCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkDeleteMerchantCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteMerchantCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteMerchantCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteMerchantCustomAttributesResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteMerchantCustomAttributesResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents an individual delete response in a [BulkDeleteMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkDeleteMerchantCustomAttributes)
// request.
type BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse struct {
	// Errors that occurred while processing the individual MerchantCustomAttributeDeleteRequest request
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteMerchantCustomAttributesResponseMerchantCustomAttributeDeleteResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents an individual upsert request in a [BulkUpsertMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkUpsertMerchantCustomAttributes)
// request. An individual request contains a merchant ID, the custom attribute to create or update,
// and an optional idempotency key.
type BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest struct {
	// The ID of the target [merchant](entity:Merchant).
	MerchantID string `json:"merchant_id" url:"merchant_id"`
	// The custom attribute to create or update, with following fields:
	//
	//   - `key`. This key must match the `key` of a custom attribute definition in the Square seller
	//     account. If the requesting application is not the definition owner, you must provide the qualified key.
	//   - `value`. This value must conform to the `schema` specified by the definition.
	//     For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
	//   - The version field must match the current version of the custom attribute definition to enable
	//     [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	//     If this is not important for your application, version can be set to -1. For any other values, the request fails with a BAD_REQUEST error.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// A unique identifier for this individual upsert request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"idempotency_key,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) GetMerchantID() string {
	if b == nil {
		return ""
	}
	return b.MerchantID
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) GetCustomAttribute() *CustomAttribute {
	if b == nil {
		return nil
	}
	return b.CustomAttribute
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) GetIdempotencyKey() *string {
	if b == nil {
		return nil
	}
	return b.IdempotencyKey
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [BulkUpsertMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkUpsertMerchantCustomAttributes) response,
// which contains a map of responses that each corresponds to an individual upsert request.
type BulkUpsertMerchantCustomAttributesResponse struct {
	// A map of responses that correspond to individual upsert requests. Each response has the
	// same ID as the corresponding request and contains either a `merchant_id` and `custom_attribute` or an `errors` field.
	Values map[string]*BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse `json:"values,omitempty" url:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkUpsertMerchantCustomAttributesResponse) GetValues() map[string]*BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkUpsertMerchantCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkUpsertMerchantCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertMerchantCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertMerchantCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertMerchantCustomAttributesResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertMerchantCustomAttributesResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a response for an individual upsert request in a [BulkUpsertMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkUpsertMerchantCustomAttributes) operation.
type BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse struct {
	// The ID of the merchant associated with the custom attribute.
	MerchantID *string `json:"merchant_id,omitempty" url:"merchant_id,omitempty"`
	// The new or updated custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) GetMerchantID() *string {
	if b == nil {
		return nil
	}
	return b.MerchantID
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) GetCustomAttribute() *CustomAttribute {
	if b == nil {
		return nil
	}
	return b.CustomAttribute
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a [CreateMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-CreateMerchantCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type CreateMerchantCustomAttributeDefinitionResponse struct {
	// The new custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateMerchantCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if c == nil {
		return nil
	}
	return c.CustomAttributeDefinition
}

func (c *CreateMerchantCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateMerchantCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateMerchantCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateMerchantCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateMerchantCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateMerchantCustomAttributeDefinitionResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Represents a response from a delete request containing error messages if there are any.
type DeleteMerchantCustomAttributeDefinitionResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteMerchantCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteMerchantCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteMerchantCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteMerchantCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteMerchantCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteMerchantCustomAttributeDefinitionResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Represents a [DeleteMerchantCustomAttribute](api-endpoint:MerchantCustomAttributes-DeleteMerchantCustomAttribute) response.
// Either an empty object `{}` (for a successful deletion) or `errors` is present in the response.
type DeleteMerchantCustomAttributeResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteMerchantCustomAttributeResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteMerchantCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteMerchantCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteMerchantCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteMerchantCustomAttributeResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteMerchantCustomAttributeResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Represents a [ListMerchantCustomAttributeDefinitions](api-endpoint:MerchantCustomAttributes-ListMerchantCustomAttributeDefinitions) response.
// Either `custom_attribute_definitions`, an empty object, or `errors` is present in the response.
// If additional results are available, the `cursor` field is also present along with `custom_attribute_definitions`.
type ListMerchantCustomAttributeDefinitionsResponse struct {
	// The retrieved custom attribute definitions. If no custom attribute definitions are found,
	// Square returns an empty object (`{}`).
	CustomAttributeDefinitions []*CustomAttributeDefinition `json:"custom_attribute_definitions,omitempty" url:"custom_attribute_definitions,omitempty"`
	// The cursor to provide in your next call to this endpoint to retrieve the next page of
	// results for your original request. This field is present only if the request succeeded and
	// additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) GetCustomAttributeDefinitions() []*CustomAttributeDefinition {
	if l == nil {
		return nil
	}
	return l.CustomAttributeDefinitions
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListMerchantCustomAttributeDefinitionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListMerchantCustomAttributeDefinitionsResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListMerchantCustomAttributeDefinitionsResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Represents a [ListMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-ListMerchantCustomAttributes) response.
// Either `custom_attributes`, an empty object, or `errors` is present in the response. If additional
// results are available, the `cursor` field is also present along with `custom_attributes`.
type ListMerchantCustomAttributesResponse struct {
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
	_rawJSON        json.RawMessage
}

func (l *ListMerchantCustomAttributesResponse) GetCustomAttributes() []*CustomAttribute {
	if l == nil {
		return nil
	}
	return l.CustomAttributes
}

func (l *ListMerchantCustomAttributesResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListMerchantCustomAttributesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListMerchantCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListMerchantCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListMerchantCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListMerchantCustomAttributesResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListMerchantCustomAttributesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Represents a [RetrieveMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-RetrieveMerchantCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveMerchantCustomAttributeDefinitionResponse struct {
	// The retrieved custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if r == nil {
		return nil
	}
	return r.CustomAttributeDefinition
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveMerchantCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveMerchantCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Represents a [RetrieveMerchantCustomAttribute](api-endpoint:MerchantCustomAttributes-RetrieveMerchantCustomAttribute) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveMerchantCustomAttributeResponse struct {
	// The retrieved custom attribute. If `with_definition` was set to `true` in the request,
	// the custom attribute definition is returned in the `definition` field.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RetrieveMerchantCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if r == nil {
		return nil
	}
	return r.CustomAttribute
}

func (r *RetrieveMerchantCustomAttributeResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveMerchantCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveMerchantCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveMerchantCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveMerchantCustomAttributeResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveMerchantCustomAttributeResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Represents an [UpdateMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-UpdateMerchantCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpdateMerchantCustomAttributeDefinitionResponse struct {
	// The updated custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if u == nil {
		return nil
	}
	return u.CustomAttributeDefinition
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateMerchantCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateMerchantCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

// Represents an [UpsertMerchantCustomAttribute](api-endpoint:MerchantCustomAttributes-UpsertMerchantCustomAttribute) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpsertMerchantCustomAttributeResponse struct {
	// The new or updated custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpsertMerchantCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if u == nil {
		return nil
	}
	return u.CustomAttribute
}

func (u *UpsertMerchantCustomAttributeResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpsertMerchantCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpsertMerchantCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertMerchantCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertMerchantCustomAttributeResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertMerchantCustomAttributeResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
