// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/v2/internal"
)

type BatchCreateVendorsRequest struct {
	// Specifies a set of new [Vendor](entity:Vendor) objects as represented by a collection of idempotency-key/`Vendor`-object pairs.
	Vendors map[string]*Vendor `json:"vendors,omitempty" url:"-"`
}

type BatchGetVendorsRequest struct {
	// IDs of the [Vendor](entity:Vendor) objects to retrieve.
	VendorIDs []string `json:"vendor_ids,omitempty" url:"-"`
}

type BatchUpdateVendorsRequest struct {
	// A set of [UpdateVendorRequest](entity:UpdateVendorRequest) objects encapsulating to-be-updated [Vendor](entity:Vendor)
	// objects. The set is represented by  a collection of `Vendor`-ID/`UpdateVendorRequest`-object pairs.
	Vendors map[string]*UpdateVendorRequest `json:"vendors,omitempty" url:"-"`
}

type CreateVendorRequest struct {
	// A client-supplied, universally unique identifier (UUID) to make this [CreateVendor](api-endpoint:Vendors-CreateVendor) call idempotent.
	//
	// See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) in the
	// [API Development 101](https://developer.squareup.com/docs/buildbasics) section for more
	// information.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The requested [Vendor](entity:Vendor) to be created.
	Vendor *Vendor `json:"vendor,omitempty" url:"-"`
}

type GetVendorsRequest struct {
	// ID of the [Vendor](entity:Vendor) to retrieve.
	VendorID string `json:"-" url:"-"`
}

type SearchVendorsRequest struct {
	// Specifies a filter used to search for vendors.
	Filter *SearchVendorsRequestFilter `json:"filter,omitempty" url:"-"`
	// Specifies a sorter used to sort the returned vendors.
	Sort *SearchVendorsRequestSort `json:"sort,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	//
	// See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor *string `json:"cursor,omitempty" url:"-"`
}

// Represents an output from a call to [BulkCreateVendors](api-endpoint:Vendors-BulkCreateVendors).
type BatchCreateVendorsResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// A set of [CreateVendorResponse](entity:CreateVendorResponse) objects encapsulating successfully created [Vendor](entity:Vendor)
	// objects or error responses for failed attempts. The set is represented by
	// a collection of idempotency-key/`Vendor`-object or idempotency-key/error-object pairs. The idempotency keys correspond to those specified
	// in the input.
	Responses map[string]*CreateVendorResponse `json:"responses,omitempty" url:"responses,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BatchCreateVendorsResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BatchCreateVendorsResponse) GetResponses() map[string]*CreateVendorResponse {
	if b == nil {
		return nil
	}
	return b.Responses
}

func (b *BatchCreateVendorsResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchCreateVendorsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchCreateVendorsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchCreateVendorsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchCreateVendorsResponse) String() string {
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

// Represents an output from a call to [BulkRetrieveVendors](api-endpoint:Vendors-BulkRetrieveVendors).
type BatchGetVendorsResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The set of [RetrieveVendorResponse](entity:RetrieveVendorResponse) objects encapsulating successfully retrieved [Vendor](entity:Vendor)
	// objects or error responses for failed attempts. The set is represented by
	// a collection of `Vendor`-ID/`Vendor`-object or `Vendor`-ID/error-object pairs.
	Responses map[string]*GetVendorResponse `json:"responses,omitempty" url:"responses,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BatchGetVendorsResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BatchGetVendorsResponse) GetResponses() map[string]*GetVendorResponse {
	if b == nil {
		return nil
	}
	return b.Responses
}

func (b *BatchGetVendorsResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchGetVendorsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchGetVendorsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchGetVendorsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchGetVendorsResponse) String() string {
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

// Represents an output from a call to [BulkUpdateVendors](api-endpoint:Vendors-BulkUpdateVendors).
type BatchUpdateVendorsResponse struct {
	// Errors encountered when the request fails.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// A set of [UpdateVendorResponse](entity:UpdateVendorResponse) objects encapsulating successfully created [Vendor](entity:Vendor)
	// objects or error responses for failed attempts. The set is represented by a collection of `Vendor`-ID/`UpdateVendorResponse`-object or
	// `Vendor`-ID/error-object pairs.
	Responses map[string]*UpdateVendorResponse `json:"responses,omitempty" url:"responses,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BatchUpdateVendorsResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BatchUpdateVendorsResponse) GetResponses() map[string]*UpdateVendorResponse {
	if b == nil {
		return nil
	}
	return b.Responses
}

func (b *BatchUpdateVendorsResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchUpdateVendorsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchUpdateVendorsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchUpdateVendorsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchUpdateVendorsResponse) String() string {
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

// Represents an output from a call to [CreateVendor](api-endpoint:Vendors-CreateVendor).
type CreateVendorResponse struct {
	// Errors encountered when the request fails.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The successfully created [Vendor](entity:Vendor) object.
	Vendor *Vendor `json:"vendor,omitempty" url:"vendor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateVendorResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateVendorResponse) GetVendor() *Vendor {
	if c == nil {
		return nil
	}
	return c.Vendor
}

func (c *CreateVendorResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateVendorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateVendorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateVendorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateVendorResponse) String() string {
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

// Represents an output from a call to [RetrieveVendor](api-endpoint:Vendors-RetrieveVendor).
type GetVendorResponse struct {
	// Errors encountered when the request fails.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The successfully retrieved [Vendor](entity:Vendor) object.
	Vendor *Vendor `json:"vendor,omitempty" url:"vendor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetVendorResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetVendorResponse) GetVendor() *Vendor {
	if g == nil {
		return nil
	}
	return g.Vendor
}

func (g *GetVendorResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetVendorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetVendorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetVendorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetVendorResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Defines supported query expressions to search for vendors by.
type SearchVendorsRequestFilter struct {
	// The names of the [Vendor](entity:Vendor) objects to retrieve.
	Name []string `json:"name,omitempty" url:"name,omitempty"`
	// The statuses of the [Vendor](entity:Vendor) objects to retrieve.
	// See [VendorStatus](#type-vendorstatus) for possible values
	Status []VendorStatus `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SearchVendorsRequestFilter) GetName() []string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SearchVendorsRequestFilter) GetStatus() []VendorStatus {
	if s == nil {
		return nil
	}
	return s.Status
}

func (s *SearchVendorsRequestFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchVendorsRequestFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchVendorsRequestFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchVendorsRequestFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchVendorsRequestFilter) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Defines a sorter used to sort results from [SearchVendors](api-endpoint:Vendors-SearchVendors).
type SearchVendorsRequestSort struct {
	// Specifies the sort key to sort the returned vendors.
	// See [Field](#type-field) for possible values
	Field *SearchVendorsRequestSortField `json:"field,omitempty" url:"field,omitempty"`
	// Specifies the sort order for the returned vendors.
	// See [SortOrder](#type-sortorder) for possible values
	Order *SortOrder `json:"order,omitempty" url:"order,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SearchVendorsRequestSort) GetField() *SearchVendorsRequestSortField {
	if s == nil {
		return nil
	}
	return s.Field
}

func (s *SearchVendorsRequestSort) GetOrder() *SortOrder {
	if s == nil {
		return nil
	}
	return s.Order
}

func (s *SearchVendorsRequestSort) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchVendorsRequestSort) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchVendorsRequestSort
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchVendorsRequestSort(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchVendorsRequestSort) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// The field to sort the returned [Vendor](entity:Vendor) objects by.
type SearchVendorsRequestSortField string

const (
	SearchVendorsRequestSortFieldName      SearchVendorsRequestSortField = "NAME"
	SearchVendorsRequestSortFieldCreatedAt SearchVendorsRequestSortField = "CREATED_AT"
)

func NewSearchVendorsRequestSortFieldFromString(s string) (SearchVendorsRequestSortField, error) {
	switch s {
	case "NAME":
		return SearchVendorsRequestSortFieldName, nil
	case "CREATED_AT":
		return SearchVendorsRequestSortFieldCreatedAt, nil
	}
	var t SearchVendorsRequestSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SearchVendorsRequestSortField) Ptr() *SearchVendorsRequestSortField {
	return &s
}

// Represents an output from a call to [SearchVendors](api-endpoint:Vendors-SearchVendors).
type SearchVendorsResponse struct {
	// Errors encountered when the request fails.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The [Vendor](entity:Vendor) objects matching the specified search filter.
	Vendors []*Vendor `json:"vendors,omitempty" url:"vendors,omitempty"`
	// The pagination cursor to be used in a subsequent request. If unset,
	// this is the final response.
	//
	// See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SearchVendorsResponse) GetErrors() []*Error {
	if s == nil {
		return nil
	}
	return s.Errors
}

func (s *SearchVendorsResponse) GetVendors() []*Vendor {
	if s == nil {
		return nil
	}
	return s.Vendors
}

func (s *SearchVendorsResponse) GetCursor() *string {
	if s == nil {
		return nil
	}
	return s.Cursor
}

func (s *SearchVendorsResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchVendorsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchVendorsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchVendorsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchVendorsResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Represents an input to a call to [UpdateVendor](api-endpoint:Vendors-UpdateVendor).
type UpdateVendorRequest struct {
	// A client-supplied, universally unique identifier (UUID) for the
	// request.
	//
	// See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) in the
	// [API Development 101](https://developer.squareup.com/docs/buildbasics) section for more
	// information.
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"idempotency_key,omitempty"`
	// The specified [Vendor](entity:Vendor) to be updated.
	Vendor *Vendor `json:"vendor,omitempty" url:"vendor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateVendorRequest) GetIdempotencyKey() *string {
	if u == nil {
		return nil
	}
	return u.IdempotencyKey
}

func (u *UpdateVendorRequest) GetVendor() *Vendor {
	if u == nil {
		return nil
	}
	return u.Vendor
}

func (u *UpdateVendorRequest) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateVendorRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateVendorRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateVendorRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateVendorRequest) String() string {
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

// Represents an output from a call to [UpdateVendor](api-endpoint:Vendors-UpdateVendor).
type UpdateVendorResponse struct {
	// Errors occurred when the request fails.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The [Vendor](entity:Vendor) that has been updated.
	Vendor *Vendor `json:"vendor,omitempty" url:"vendor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateVendorResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateVendorResponse) GetVendor() *Vendor {
	if u == nil {
		return nil
	}
	return u.Vendor
}

func (u *UpdateVendorResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateVendorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateVendorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateVendorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateVendorResponse) String() string {
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

// Represents a supplier to a seller.
type Vendor struct {
	// A unique Square-generated ID for the [Vendor](entity:Vendor).
	// This field is required when attempting to update a [Vendor](entity:Vendor).
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when the
	// [Vendor](entity:Vendor) was created.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when the
	// [Vendor](entity:Vendor) was last updated.
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The name of the [Vendor](entity:Vendor).
	// This field is required when attempting to create or update a [Vendor](entity:Vendor).
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The address of the [Vendor](entity:Vendor).
	Address *Address `json:"address,omitempty" url:"address,omitempty"`
	// The contacts of the [Vendor](entity:Vendor).
	Contacts []*VendorContact `json:"contacts,omitempty" url:"contacts,omitempty"`
	// The account number of the [Vendor](entity:Vendor).
	AccountNumber *string `json:"account_number,omitempty" url:"account_number,omitempty"`
	// A note detailing information about the [Vendor](entity:Vendor).
	Note *string `json:"note,omitempty" url:"note,omitempty"`
	// The version of the [Vendor](entity:Vendor).
	Version *int `json:"version,omitempty" url:"version,omitempty"`
	// The status of the [Vendor](entity:Vendor).
	// See [Status](#type-status) for possible values
	Status *VendorStatus `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (v *Vendor) GetID() *string {
	if v == nil {
		return nil
	}
	return v.ID
}

func (v *Vendor) GetCreatedAt() *string {
	if v == nil {
		return nil
	}
	return v.CreatedAt
}

func (v *Vendor) GetUpdatedAt() *string {
	if v == nil {
		return nil
	}
	return v.UpdatedAt
}

func (v *Vendor) GetName() *string {
	if v == nil {
		return nil
	}
	return v.Name
}

func (v *Vendor) GetAddress() *Address {
	if v == nil {
		return nil
	}
	return v.Address
}

func (v *Vendor) GetContacts() []*VendorContact {
	if v == nil {
		return nil
	}
	return v.Contacts
}

func (v *Vendor) GetAccountNumber() *string {
	if v == nil {
		return nil
	}
	return v.AccountNumber
}

func (v *Vendor) GetNote() *string {
	if v == nil {
		return nil
	}
	return v.Note
}

func (v *Vendor) GetVersion() *int {
	if v == nil {
		return nil
	}
	return v.Version
}

func (v *Vendor) GetStatus() *VendorStatus {
	if v == nil {
		return nil
	}
	return v.Status
}

func (v *Vendor) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *Vendor) UnmarshalJSON(data []byte) error {
	type unmarshaler Vendor
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = Vendor(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	v.rawJSON = json.RawMessage(data)
	return nil
}

func (v *Vendor) String() string {
	if len(v.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(v.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

// Represents a contact of a [Vendor](entity:Vendor).
type VendorContact struct {
	// A unique Square-generated ID for the [VendorContact](entity:VendorContact).
	// This field is required when attempting to update a [VendorContact](entity:VendorContact).
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The name of the [VendorContact](entity:VendorContact).
	// This field is required when attempting to create a [Vendor](entity:Vendor).
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The email address of the [VendorContact](entity:VendorContact).
	EmailAddress *string `json:"email_address,omitempty" url:"email_address,omitempty"`
	// The phone number of the [VendorContact](entity:VendorContact).
	PhoneNumber *string `json:"phone_number,omitempty" url:"phone_number,omitempty"`
	// The state of the [VendorContact](entity:VendorContact).
	Removed *bool `json:"removed,omitempty" url:"removed,omitempty"`
	// The ordinal of the [VendorContact](entity:VendorContact).
	Ordinal int `json:"ordinal" url:"ordinal"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (v *VendorContact) GetID() *string {
	if v == nil {
		return nil
	}
	return v.ID
}

func (v *VendorContact) GetName() *string {
	if v == nil {
		return nil
	}
	return v.Name
}

func (v *VendorContact) GetEmailAddress() *string {
	if v == nil {
		return nil
	}
	return v.EmailAddress
}

func (v *VendorContact) GetPhoneNumber() *string {
	if v == nil {
		return nil
	}
	return v.PhoneNumber
}

func (v *VendorContact) GetRemoved() *bool {
	if v == nil {
		return nil
	}
	return v.Removed
}

func (v *VendorContact) GetOrdinal() int {
	if v == nil {
		return 0
	}
	return v.Ordinal
}

func (v *VendorContact) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorContact) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorContact
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorContact(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	v.rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorContact) String() string {
	if len(v.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(v.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

// The status of the [Vendor](entity:Vendor),
// whether a [Vendor](entity:Vendor) is active or inactive.
type VendorStatus string

const (
	VendorStatusActive   VendorStatus = "ACTIVE"
	VendorStatusInactive VendorStatus = "INACTIVE"
)

func NewVendorStatusFromString(s string) (VendorStatus, error) {
	switch s {
	case "ACTIVE":
		return VendorStatusActive, nil
	case "INACTIVE":
		return VendorStatusInactive, nil
	}
	var t VendorStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (v VendorStatus) Ptr() *VendorStatus {
	return &v
}

type UpdateVendorsRequest struct {
	// ID of the [Vendor](entity:Vendor) to retrieve.
	VendorID string               `json:"-" url:"-"`
	Body     *UpdateVendorRequest `json:"-" url:"-"`
}

func (u *UpdateVendorsRequest) UnmarshalJSON(data []byte) error {
	body := new(UpdateVendorRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.Body = body
	return nil
}

func (u *UpdateVendorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}
