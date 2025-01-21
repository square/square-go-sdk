// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/v40/internal"
)

type CreateMobileAuthorizationCodeRequest struct {
	// The Square location ID that the authorization code should be tied to.
	LocationID *string `json:"location_id,omitempty" url:"-"`
}

// Defines the fields that are included in the response body of
// a request to the `CreateMobileAuthorizationCode` endpoint.
type CreateMobileAuthorizationCodeResponse struct {
	// The generated authorization code that connects a mobile application instance
	// to a Square account.
	AuthorizationCode *string `json:"authorization_code,omitempty" url:"authorization_code,omitempty"`
	// The timestamp when `authorization_code` expires, in
	// [RFC 3339](https://tools.ietf.org/html/rfc3339) format (for example, "2016-09-04T23:59:33.123Z").
	ExpiresAt *string `json:"expires_at,omitempty" url:"expires_at,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateMobileAuthorizationCodeResponse) GetAuthorizationCode() *string {
	if c == nil {
		return nil
	}
	return c.AuthorizationCode
}

func (c *CreateMobileAuthorizationCodeResponse) GetExpiresAt() *string {
	if c == nil {
		return nil
	}
	return c.ExpiresAt
}

func (c *CreateMobileAuthorizationCodeResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateMobileAuthorizationCodeResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateMobileAuthorizationCodeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateMobileAuthorizationCodeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateMobileAuthorizationCodeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateMobileAuthorizationCodeResponse) String() string {
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
