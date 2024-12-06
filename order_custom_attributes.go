// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type BulkDeleteOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual delete operations for custom attributes.
	Values map[string]*BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute `json:"values,omitempty" url:"-"`
}

type BulkUpsertOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual upsert operations for custom attributes.
	Values map[string]*BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute `json:"values,omitempty" url:"-"`
}

type CreateOrderCustomAttributeDefinitionRequest struct {
	// The custom attribute definition to create. Note the following:
	// - With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
	// definition hosted on the Square CDN. For more information, including supported values and constraints, see
	// [Specifying the schema](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#specify-schema).
	// - If provided, `name` must be unique (case-sensitive) across all visible customer-related custom attribute definitions for the seller.
	// - All custom attributes are visible in exported customer data, including those set to `VISIBILITY_HIDDEN`.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type DeleteOrderCustomAttributeRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to delete. This key must match the key of an
	// existing custom attribute definition.
	CustomAttributeKey string `json:"-" url:"-"`
}

type DeleteOrderCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to delete.
	Key string `json:"-" url:"-"`
}

type RetrieveOrderCustomAttributeRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to retrieve. This key must match the key of an
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

type RetrieveOrderCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to retrieve.
	Key string `json:"-" url:"-"`
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include this optional field and specify the current version of the custom attribute.
	Version *int `json:"-" url:"version,omitempty"`
}

type UpdateOrderCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to update.
	Key string `json:"-" url:"-"`
	// The custom attribute definition that contains the fields to update. This endpoint supports sparse updates,
	// so only new or changed fields need to be included in the request.  For more information, see
	// [Updatable definition fields](https://developer.squareup.com/docs/orders-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).
	//
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control, include the optional `version` field and specify the current version of the custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type UpsertOrderCustomAttributeRequest struct {
	// The ID of the target [order](entity:Order).
	OrderID string `json:"-" url:"-"`
	// The key of the custom attribute to create or update. This key must match the key
	// of an existing custom attribute definition.
	CustomAttributeKey string `json:"-" url:"-"`
	// The custom attribute to create or update, with the following fields:
	//
	// - `value`. This value must conform to the `schema` specified by the definition.
	// For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).
	//
	// - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include this optional field and specify the current version of the custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

// Represents one delete within the bulk operation.
type BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute struct {
	// The key of the custom attribute to delete. This key must match the key
	// of an existing custom attribute definition.
	Key *string `json:"key,omitempty" url:"key,omitempty"`
	// The ID of the target [order](entity:Order).
	OrderID string `json:"order_id" url:"order_id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute) GetKey() *string {
	if b == nil {
		return nil
	}
	return b.Key
}

func (b *BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute) GetOrderID() string {
	if b == nil {
		return ""
	}
	return b.OrderID
}

func (b *BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute) String() string {
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

// Represents a response from deleting one or more order custom attributes.
type BulkDeleteOrderCustomAttributesResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// A map of responses that correspond to individual delete requests. Each response has the same ID
	// as the corresponding request and contains either a `custom_attribute` or an `errors` field.
	Values map[string]*DeleteOrderCustomAttributeResponse `json:"values,omitempty" url:"values,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDeleteOrderCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkDeleteOrderCustomAttributesResponse) GetValues() map[string]*DeleteOrderCustomAttributeResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkDeleteOrderCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDeleteOrderCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDeleteOrderCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDeleteOrderCustomAttributesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDeleteOrderCustomAttributesResponse) String() string {
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

// Represents one upsert within the bulk operation.
type BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute struct {
	// The custom attribute to create or update, with the following fields:
	//
	//   - `value`. This value must conform to the `schema` specified by the definition.
	//     For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).
	//
	//   - `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	//     control, include this optional field and specify the current version of the custom attribute.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// A unique identifier for this request, used to ensure idempotency.
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"idempotency_key,omitempty"`
	// The ID of the target [order](entity:Order).
	OrderID string `json:"order_id" url:"order_id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) GetCustomAttribute() *CustomAttribute {
	if b == nil {
		return nil
	}
	return b.CustomAttribute
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) GetIdempotencyKey() *string {
	if b == nil {
		return nil
	}
	return b.IdempotencyKey
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) GetOrderID() string {
	if b == nil {
		return ""
	}
	return b.OrderID
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute) String() string {
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

// Represents a response from a bulk upsert of order custom attributes.
type BulkUpsertOrderCustomAttributesResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// A map of responses that correspond to individual upsert operations for custom attributes.
	Values map[string]*UpsertOrderCustomAttributeResponse `json:"values,omitempty" url:"values,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkUpsertOrderCustomAttributesResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BulkUpsertOrderCustomAttributesResponse) GetValues() map[string]*UpsertOrderCustomAttributeResponse {
	if b == nil {
		return nil
	}
	return b.Values
}

func (b *BulkUpsertOrderCustomAttributesResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkUpsertOrderCustomAttributesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkUpsertOrderCustomAttributesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkUpsertOrderCustomAttributesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkUpsertOrderCustomAttributesResponse) String() string {
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

// Represents a response from creating an order custom attribute definition.
type CreateOrderCustomAttributeDefinitionResponse struct {
	// The new custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateOrderCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if c == nil {
		return nil
	}
	return c.CustomAttributeDefinition
}

func (c *CreateOrderCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateOrderCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateOrderCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateOrderCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateOrderCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateOrderCustomAttributeDefinitionResponse) String() string {
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

// Represents a response from deleting an order custom attribute definition.
type DeleteOrderCustomAttributeDefinitionResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteOrderCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteOrderCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteOrderCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteOrderCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteOrderCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteOrderCustomAttributeDefinitionResponse) String() string {
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

// Represents a response from deleting an order custom attribute.
type DeleteOrderCustomAttributeResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteOrderCustomAttributeResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteOrderCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteOrderCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteOrderCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteOrderCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteOrderCustomAttributeResponse) String() string {
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

// Represents a response from getting an order custom attribute definition.
type RetrieveOrderCustomAttributeDefinitionResponse struct {
	// The retrieved custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RetrieveOrderCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if r == nil {
		return nil
	}
	return r.CustomAttributeDefinition
}

func (r *RetrieveOrderCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveOrderCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveOrderCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveOrderCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveOrderCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveOrderCustomAttributeDefinitionResponse) String() string {
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

// Represents a response from getting an order custom attribute.
type RetrieveOrderCustomAttributeResponse struct {
	// The retrieved custom attribute. If `with_definition` was set to `true` in the request, the custom attribute definition is returned in the `definition field.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RetrieveOrderCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if r == nil {
		return nil
	}
	return r.CustomAttribute
}

func (r *RetrieveOrderCustomAttributeResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RetrieveOrderCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RetrieveOrderCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RetrieveOrderCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetrieveOrderCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveOrderCustomAttributeResponse) String() string {
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

// Represents a response from updating an order custom attribute definition.
type UpdateOrderCustomAttributeDefinitionResponse struct {
	// The updated order custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateOrderCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if u == nil {
		return nil
	}
	return u.CustomAttributeDefinition
}

func (u *UpdateOrderCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateOrderCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateOrderCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateOrderCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateOrderCustomAttributeDefinitionResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateOrderCustomAttributeDefinitionResponse) String() string {
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

// Represents a response from upserting order custom attribute definitions.
type UpsertOrderCustomAttributeResponse struct {
	// The order custom attribute that was created or modified.
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty" url:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpsertOrderCustomAttributeResponse) GetCustomAttribute() *CustomAttribute {
	if u == nil {
		return nil
	}
	return u.CustomAttribute
}

func (u *UpsertOrderCustomAttributeResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpsertOrderCustomAttributeResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpsertOrderCustomAttributeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertOrderCustomAttributeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertOrderCustomAttributeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertOrderCustomAttributeResponse) String() string {
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
