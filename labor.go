// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type CreateBreakTypeRequest struct {
	// A unique string value to ensure the idempotency of the operation.
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The `BreakType` to be created.
	BreakType *BreakType `json:"break_type,omitempty" url:"-"`
}

type LaborListRequest struct {
	// Filter the returned `BreakType` results to only those that are associated with the
	// specified location.
	LocationID *string `json:"-" url:"location_id,omitempty"`
	// The maximum number of `BreakType` results to return per page. The number can range between 1
	// and 200. The default is 200.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pointer to the next page of `BreakType` results to fetch.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

// The response to the request to create a `BreakType`. The response contains
// the created `BreakType` object and might contain a set of `Error` objects if
// the request resulted in errors.
type CreateBreakTypeResponse struct {
	// The `BreakType` that was created by the request.
	BreakType *BreakType `json:"break_type,omitempty" url:"break_type,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateBreakTypeResponse) GetBreakType() *BreakType {
	if c == nil {
		return nil
	}
	return c.BreakType
}

func (c *CreateBreakTypeResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateBreakTypeResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateBreakTypeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateBreakTypeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateBreakTypeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateBreakTypeResponse) String() string {
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

// The response to a request for a set of `BreakType` objects. The response contains
// the requested `BreakType` objects and might contain a set of `Error` objects if
// the request resulted in errors.
type ListBreakTypesResponse struct {
	// A page of `BreakType` results.
	BreakTypes []*BreakType `json:"break_types,omitempty" url:"break_types,omitempty"`
	// The value supplied in the subsequent request to fetch the next page
	// of `BreakType` results.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListBreakTypesResponse) GetBreakTypes() []*BreakType {
	if l == nil {
		return nil
	}
	return l.BreakTypes
}

func (l *ListBreakTypesResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListBreakTypesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListBreakTypesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListBreakTypesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListBreakTypesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListBreakTypesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListBreakTypesResponse) String() string {
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
