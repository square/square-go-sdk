// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type DeleteMerchantCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to delete.
	Key string `json:"-" url:"-"`
}

type RetrieveMerchantCustomAttributeDefinitionRequest struct {
	// The key of the custom attribute definition to retrieve. If the requesting application
	// is not the definition owner, you must use the qualified key.
	Key string `json:"-" url:"-"`
	// The current version of the custom attribute definition, which is used for strongly consistent
	// reads to guarantee that you receive the most up-to-date data. When included in the request,
	// Square returns the specified version or a higher version if one exists. If the specified version
	// is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version *int `json:"-" url:"version,omitempty"`
}

type UpdateMerchantCustomAttributeDefinitionRequest struct {
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

// Represents a response from a delete request containing error messages if there are any.
type DeleteMerchantCustomAttributeDefinitionResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
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
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteMerchantCustomAttributeDefinitionResponse) String() string {
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

// Represents a [RetrieveMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-RetrieveMerchantCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveMerchantCustomAttributeDefinitionResponse struct {
	// The retrieved custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
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
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RetrieveMerchantCustomAttributeDefinitionResponse) String() string {
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

// Represents an [UpdateMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-UpdateMerchantCustomAttributeDefinition) response.
// Either `custom_attribute_definition` or `errors` is present in the response.
type UpdateMerchantCustomAttributeDefinitionResponse struct {
	// The updated custom attribute definition.
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty" url:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
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
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateMerchantCustomAttributeDefinitionResponse) String() string {
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
