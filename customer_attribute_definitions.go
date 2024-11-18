// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type CreateCustomerCustomAttributeDefinitionRequest struct {
	// The custom attribute definition to create. Note the following:
	// - With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
	// definition hosted on the Square CDN. For more information, including supported values and constraints, see
	// [Specifying the schema](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#specify-schema).
	// - If provided, `name` must be unique (case-sensitive) across all visible customer-related custom attribute definitions for the seller.
	// - All custom attributes are visible in exported customer data, including those set to `VISIBILITY_HIDDEN`.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type CustomerAttributeDefinitionsGetRequest struct {
	// The current version of the custom attribute definition, which is used for strongly consistent
	// reads to guarantee that you receive the most up-to-date data. When included in the request,
	// Square returns the specified version or a higher version if one exists. If the specified version
	// is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type CustomerAttributeDefinitionsListRequest struct {
	// The maximum number of results to return in a single paged response. This limit is advisory.
	// The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
	// The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.
	// Provide this cursor to retrieve the next page of results for your original request.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

// Represents a [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type CreateCustomerCustomAttributeDefinitionResponse struct {
	// The new custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateCustomerCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if c == nil {
		return nil
	}
	return c.CustomAttributeDefinition
}

func (c *CreateCustomerCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateCustomerCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateCustomerCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCustomerCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateCustomerCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateCustomerCustomAttributeDefinitionResponse) String() string {
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
type DeleteCustomerCustomAttributeDefinitionResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteCustomerCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DeleteCustomerCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteCustomerCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteCustomerCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteCustomerCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteCustomerCustomAttributeDefinitionResponse) String() string {
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

// Represents a [RetrieveCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-RetrieveCustomerCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type GetCustomerCustomAttributeDefinitionResponse struct {
	// The retrieved custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetCustomerCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if g == nil {
		return nil
	}
	return g.CustomAttributeDefinition
}

func (g *GetCustomerCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetCustomerCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetCustomerCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCustomerCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCustomerCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCustomerCustomAttributeDefinitionResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Represents a [ListCustomerCustomAttributeDefinitions](api-endpoint:CustomerCustomAttributes-ListCustomerCustomAttributeDefinitions) response.
// Either `custom_attribute_definitions`, an empty object, or `errors` is present in the response.
// If additional results are available, the `cursor` field is also present along with `custom_attribute_definitions`.
type ListCustomerCustomAttributeDefinitionsResponse struct {
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

func (l *ListCustomerCustomAttributeDefinitionsResponse) GetCustomAttributeDefinitions() []*CustomAttributeDefinition {
	if l == nil {
		return nil
	}
	return l.CustomAttributeDefinitions
}

func (l *ListCustomerCustomAttributeDefinitionsResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListCustomerCustomAttributeDefinitionsResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListCustomerCustomAttributeDefinitionsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListCustomerCustomAttributeDefinitionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListCustomerCustomAttributeDefinitionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListCustomerCustomAttributeDefinitionsResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListCustomerCustomAttributeDefinitionsResponse) String() string {
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

// Represents an [UpdateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-UpdateCustomerCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpdateCustomerCustomAttributeDefinitionResponse struct {
	// The updated custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateCustomerCustomAttributeDefinitionResponse) GetCustomAttributeDefinition() *CustomAttributeDefinition {
	if u == nil {
		return nil
	}
	return u.CustomAttributeDefinition
}

func (u *UpdateCustomerCustomAttributeDefinitionResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateCustomerCustomAttributeDefinitionResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateCustomerCustomAttributeDefinitionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCustomerCustomAttributeDefinitionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCustomerCustomAttributeDefinitionResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateCustomerCustomAttributeDefinitionResponse) String() string {
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

type UpdateCustomerCustomAttributeDefinitionRequest struct {
	// The custom attribute definition that contains the fields to update. This endpoint
	// supports sparse updates, so only new or changed fields need to be included in the request.
	// Only the following fields can be updated:
	//
	// - `name`
	// - `description`
	// - `visibility`
	// - `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
	// selections are supported.
	//
	// For more information, see
	// [Updatable definition fields](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).
	//
	// To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
	// control, include the optional `version` field and specify the current version of the custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"-"`
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}
