// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type ListEventTypesRequest struct {
	// The API version for which to list event types. Setting this field overrides the default version used by the application.
	APIVersion *string `json:"-" url:"api_version,omitempty"`
}

type SearchEventsRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of events for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
	// The maximum number of events to return in a single page. The response might contain fewer events. The default value is 100, which is also the maximum allowed value.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	//
	// Default: 100
	Limit *int `json:"limit,omitempty" url:"-"`
	// The filtering and sorting criteria for the search request. To retrieve additional pages using a cursor, you must use the original query.
	Query *SearchEventsQuery `json:"query,omitempty" url:"-"`
}

// Defines the fields that are included in the response body of
// a request to the [DisableEvents](api-endpoint:Events-DisableEvents) endpoint.
//
// Note: if there are errors processing the request, the events field will not be
// present.
type DisableEventsResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DisableEventsResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DisableEventsResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DisableEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DisableEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DisableEventsResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DisableEventsResponse) String() string {
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

// Defines the fields that are included in the response body of
// a request to the [EnableEvents](api-endpoint:Events-EnableEvents) endpoint.
//
// Note: if there are errors processing the request, the events field will not be
// present.
type EnableEventsResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EnableEventsResponse) GetErrors() []*Error {
	if e == nil {
		return nil
	}
	return e.Errors
}

func (e *EnableEventsResponse) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EnableEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EnableEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnableEventsResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnableEventsResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type Event struct {
	// The ID of the target merchant associated with the event.
	MerchantID *string `json:"merchant_id,omitempty" url:"merchant_id,omitempty"`
	// The ID of the target location associated with the event.
	LocationID *string `json:"location_id,omitempty" url:"location_id,omitempty"`
	// The type of event this represents.
	Type *string `json:"type,omitempty" url:"type,omitempty"`
	// A unique ID for the event.
	EventID *string `json:"event_id,omitempty" url:"event_id,omitempty"`
	// Timestamp of when the event was created, in RFC 3339 format.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The data associated with the event.
	Data *EventData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *Event) GetMerchantID() *string {
	if e == nil {
		return nil
	}
	return e.MerchantID
}

func (e *Event) GetLocationID() *string {
	if e == nil {
		return nil
	}
	return e.LocationID
}

func (e *Event) GetType() *string {
	if e == nil {
		return nil
	}
	return e.Type
}

func (e *Event) GetEventID() *string {
	if e == nil {
		return nil
	}
	return e.EventID
}

func (e *Event) GetCreatedAt() *string {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *Event) GetData() *EventData {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *Event) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type unmarshaler Event
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Event(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *Event) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EventData struct {
	// The name of the affected object’s type.
	Type *string `json:"type,omitempty" url:"type,omitempty"`
	// The ID of the affected object.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// This is true if the affected object has been deleted; otherwise, it's absent.
	Deleted *bool `json:"deleted,omitempty" url:"deleted,omitempty"`
	// An object containing fields and values relevant to the event. It is absent if the affected object has been deleted.
	Object map[string]interface{} `json:"object,omitempty" url:"object,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EventData) GetType() *string {
	if e == nil {
		return nil
	}
	return e.Type
}

func (e *EventData) GetID() *string {
	if e == nil {
		return nil
	}
	return e.ID
}

func (e *EventData) GetDeleted() *bool {
	if e == nil {
		return nil
	}
	return e.Deleted
}

func (e *EventData) GetObject() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.Object
}

func (e *EventData) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EventData) UnmarshalJSON(data []byte) error {
	type unmarshaler EventData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EventData(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventData) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Contains metadata about a particular [Event](entity:Event).
type EventMetadata struct {
	// A unique ID for the event.
	EventID *string `json:"event_id,omitempty" url:"event_id,omitempty"`
	// The API version of the event. This corresponds to the default API version of the developer application at the time when the event was created.
	APIVersion *string `json:"api_version,omitempty" url:"api_version,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EventMetadata) GetEventID() *string {
	if e == nil {
		return nil
	}
	return e.EventID
}

func (e *EventMetadata) GetAPIVersion() *string {
	if e == nil {
		return nil
	}
	return e.APIVersion
}

func (e *EventMetadata) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EventMetadata) UnmarshalJSON(data []byte) error {
	type unmarshaler EventMetadata
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EventMetadata(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventMetadata) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Defines the fields that are included in the response body of
// a request to the [ListEventTypes](api-endpoint:Events-ListEventTypes) endpoint.
//
// Note: if there are errors processing the request, the event types field will not be
// present.
type ListEventTypesResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The list of event types.
	EventTypes []string `json:"event_types,omitempty" url:"event_types,omitempty"`
	// Contains the metadata of an event type. For more information, see [EventTypeMetadata](entity:EventTypeMetadata).
	Metadata []*EventTypeMetadata `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListEventTypesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListEventTypesResponse) GetEventTypes() []string {
	if l == nil {
		return nil
	}
	return l.EventTypes
}

func (l *ListEventTypesResponse) GetMetadata() []*EventTypeMetadata {
	if l == nil {
		return nil
	}
	return l.Metadata
}

func (l *ListEventTypesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListEventTypesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListEventTypesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListEventTypesResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListEventTypesResponse) String() string {
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

// Criteria to filter events by.
type SearchEventsFilter struct {
	// Filter events by event types.
	EventTypes []string `json:"event_types,omitempty" url:"event_types,omitempty"`
	// Filter events by merchant.
	MerchantIDs []string `json:"merchant_ids,omitempty" url:"merchant_ids,omitempty"`
	// Filter events by location.
	LocationIDs []string `json:"location_ids,omitempty" url:"location_ids,omitempty"`
	// Filter events by when they were created.
	CreatedAt *TimeRange `json:"created_at,omitempty" url:"created_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchEventsFilter) GetEventTypes() []string {
	if s == nil {
		return nil
	}
	return s.EventTypes
}

func (s *SearchEventsFilter) GetMerchantIDs() []string {
	if s == nil {
		return nil
	}
	return s.MerchantIDs
}

func (s *SearchEventsFilter) GetLocationIDs() []string {
	if s == nil {
		return nil
	}
	return s.LocationIDs
}

func (s *SearchEventsFilter) GetCreatedAt() *TimeRange {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *SearchEventsFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchEventsFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchEventsFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchEventsFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchEventsFilter) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Contains query criteria for the search.
type SearchEventsQuery struct {
	// Criteria to filter events by.
	Filter *SearchEventsFilter `json:"filter,omitempty" url:"filter,omitempty"`
	// Criteria to sort events by.
	Sort *SearchEventsSort `json:"sort,omitempty" url:"sort,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchEventsQuery) GetFilter() *SearchEventsFilter {
	if s == nil {
		return nil
	}
	return s.Filter
}

func (s *SearchEventsQuery) GetSort() *SearchEventsSort {
	if s == nil {
		return nil
	}
	return s.Sort
}

func (s *SearchEventsQuery) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchEventsQuery) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchEventsQuery
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchEventsQuery(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchEventsQuery) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Defines the fields that are included in the response body of
// a request to the [SearchEvents](api-endpoint:Events-SearchEvents) endpoint.
//
// Note: if there are errors processing the request, the events field will not be
// present.
type SearchEventsResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The list of [Event](entity:Event)s returned by the search.
	Events []*Event `json:"events,omitempty" url:"events,omitempty"`
	// Contains the metadata of an event. For more information, see [Event](entity:Event).
	Metadata []*EventMetadata `json:"metadata,omitempty" url:"metadata,omitempty"`
	// When a response is truncated, it includes a cursor that you can use in a subsequent request to fetch the next set of events. If empty, this is the final response.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchEventsResponse) GetErrors() []*Error {
	if s == nil {
		return nil
	}
	return s.Errors
}

func (s *SearchEventsResponse) GetEvents() []*Event {
	if s == nil {
		return nil
	}
	return s.Events
}

func (s *SearchEventsResponse) GetMetadata() []*EventMetadata {
	if s == nil {
		return nil
	}
	return s.Metadata
}

func (s *SearchEventsResponse) GetCursor() *string {
	if s == nil {
		return nil
	}
	return s.Cursor
}

func (s *SearchEventsResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchEventsResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchEventsResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Criteria to sort events by.
type SearchEventsSort struct {
	// Sort events by event types.
	// See [SearchEventsSortField](#type-searcheventssortfield) for possible values
	Field *SearchEventsSortField `json:"field,omitempty" url:"field,omitempty"`
	// The order to use for sorting the events.
	// See [SortOrder](#type-sortorder) for possible values
	Order *SortOrder `json:"order,omitempty" url:"order,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchEventsSort) GetOrder() *SortOrder {
	if s == nil {
		return nil
	}
	return s.Order
}

func (s *SearchEventsSort) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchEventsSort) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchEventsSort
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchEventsSort(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchEventsSort) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Specifies the sort key for events returned from a search.
type SearchEventsSortField = string
