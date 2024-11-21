// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type RegisterDomainRequest struct {
	// A domain name as described in RFC-1034 that will be registered with ApplePay.
	DomainName string `json:"domain_name" url:"-"`
}

// Defines the fields that are included in the response body of
// a request to the [RegisterDomain](api-endpoint:ApplePay-RegisterDomain) endpoint.
//
// Either `errors` or `status` are present in a given response (never both).
type RegisterDomainResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The status of the domain registration.
	//
	// See [RegisterDomainResponseStatus](entity:RegisterDomainResponseStatus) for possible values.
	// See [RegisterDomainResponseStatus](#type-registerdomainresponsestatus) for possible values
	Status *RegisterDomainResponseStatus `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RegisterDomainResponse) GetErrors() []*Error {
	if r == nil {
		return nil
	}
	return r.Errors
}

func (r *RegisterDomainResponse) GetStatus() *RegisterDomainResponseStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *RegisterDomainResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RegisterDomainResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RegisterDomainResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RegisterDomainResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RegisterDomainResponse) String() string {
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

// The status of the domain registration.
type RegisterDomainResponseStatus string

const (
	RegisterDomainResponseStatusPending  RegisterDomainResponseStatus = "PENDING"
	RegisterDomainResponseStatusVerified RegisterDomainResponseStatus = "VERIFIED"
	RegisterDomainResponseStatusDoNotUse RegisterDomainResponseStatus = "DO_NOT_USE"
)

func NewRegisterDomainResponseStatusFromString(s string) (RegisterDomainResponseStatus, error) {
	switch s {
	case "PENDING":
		return RegisterDomainResponseStatusPending, nil
	case "VERIFIED":
		return RegisterDomainResponseStatusVerified, nil
	case "DO_NOT_USE":
		return RegisterDomainResponseStatusDoNotUse, nil
	}
	var t RegisterDomainResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RegisterDomainResponseStatus) Ptr() *RegisterDomainResponseStatus {
	return &r
}
