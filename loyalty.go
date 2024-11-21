// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type SearchLoyaltyEventsRequest struct {
	// A set of one or more predefined query filters to apply when
	// searching for loyalty events. The endpoint performs a logical AND to
	// evaluate multiple filters and performs a logical OR on arrays
	// that specifies multiple field values.
	Query *LoyaltyEventQuery `json:"query,omitempty" url:"-"`
	// The maximum number of results to include in the response.
	// The last page might contain fewer events.
	// The default is 30 events.
	Limit *int `json:"limit,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for your original query.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
}

// Filter events by date time range.
type LoyaltyEventDateTimeFilter struct {
	// The `created_at` date time range used to filter the result.
	CreatedAt *TimeRange `json:"created_at,omitempty" url:"created_at,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventDateTimeFilter) GetCreatedAt() *TimeRange {
	if l == nil {
		return nil
	}
	return l.CreatedAt
}

func (l *LoyaltyEventDateTimeFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventDateTimeFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventDateTimeFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventDateTimeFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventDateTimeFilter) String() string {
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

// The filtering criteria. If the request specifies multiple filters,
// the endpoint uses a logical AND to evaluate them.
type LoyaltyEventFilter struct {
	// Filter events by loyalty account.
	LoyaltyAccountFilter *LoyaltyEventLoyaltyAccountFilter `json:"loyalty_account_filter,omitempty" url:"loyalty_account_filter,omitempty"`
	// Filter events by event type.
	TypeFilter *LoyaltyEventTypeFilter `json:"type_filter,omitempty" url:"type_filter,omitempty"`
	// Filter events by date time range.
	// For each range, the start time is inclusive and the end time
	// is exclusive.
	DateTimeFilter *LoyaltyEventDateTimeFilter `json:"date_time_filter,omitempty" url:"date_time_filter,omitempty"`
	// Filter events by location.
	LocationFilter *LoyaltyEventLocationFilter `json:"location_filter,omitempty" url:"location_filter,omitempty"`
	// Filter events by the order associated with the event.
	OrderFilter *LoyaltyEventOrderFilter `json:"order_filter,omitempty" url:"order_filter,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventFilter) GetLoyaltyAccountFilter() *LoyaltyEventLoyaltyAccountFilter {
	if l == nil {
		return nil
	}
	return l.LoyaltyAccountFilter
}

func (l *LoyaltyEventFilter) GetTypeFilter() *LoyaltyEventTypeFilter {
	if l == nil {
		return nil
	}
	return l.TypeFilter
}

func (l *LoyaltyEventFilter) GetDateTimeFilter() *LoyaltyEventDateTimeFilter {
	if l == nil {
		return nil
	}
	return l.DateTimeFilter
}

func (l *LoyaltyEventFilter) GetLocationFilter() *LoyaltyEventLocationFilter {
	if l == nil {
		return nil
	}
	return l.LocationFilter
}

func (l *LoyaltyEventFilter) GetOrderFilter() *LoyaltyEventOrderFilter {
	if l == nil {
		return nil
	}
	return l.OrderFilter
}

func (l *LoyaltyEventFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventFilter) String() string {
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

// Filter events by location.
type LoyaltyEventLocationFilter struct {
	// The [location](entity:Location) IDs for loyalty events to query.
	// If multiple values are specified, the endpoint uses
	// a logical OR to combine them.
	LocationIDs []string `json:"location_ids,omitempty" url:"location_ids,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventLocationFilter) GetLocationIDs() []string {
	if l == nil {
		return nil
	}
	return l.LocationIDs
}

func (l *LoyaltyEventLocationFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventLocationFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventLocationFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventLocationFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventLocationFilter) String() string {
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

// Filter events by loyalty account.
type LoyaltyEventLoyaltyAccountFilter struct {
	// The ID of the [loyalty account](entity:LoyaltyAccount) associated with loyalty events.
	LoyaltyAccountID string `json:"loyalty_account_id" url:"loyalty_account_id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventLoyaltyAccountFilter) GetLoyaltyAccountID() string {
	if l == nil {
		return ""
	}
	return l.LoyaltyAccountID
}

func (l *LoyaltyEventLoyaltyAccountFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventLoyaltyAccountFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventLoyaltyAccountFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventLoyaltyAccountFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventLoyaltyAccountFilter) String() string {
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

// Filter events by the order associated with the event.
type LoyaltyEventOrderFilter struct {
	// The ID of the [order](entity:Order) associated with the event.
	OrderID string `json:"order_id" url:"order_id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventOrderFilter) GetOrderID() string {
	if l == nil {
		return ""
	}
	return l.OrderID
}

func (l *LoyaltyEventOrderFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventOrderFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventOrderFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventOrderFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventOrderFilter) String() string {
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

// Represents a query used to search for loyalty events.
type LoyaltyEventQuery struct {
	// The query filter criteria.
	Filter *LoyaltyEventFilter `json:"filter,omitempty" url:"filter,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventQuery) GetFilter() *LoyaltyEventFilter {
	if l == nil {
		return nil
	}
	return l.Filter
}

func (l *LoyaltyEventQuery) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventQuery) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventQuery
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventQuery(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventQuery) String() string {
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

// Filter events by event type.
type LoyaltyEventTypeFilter struct {
	// The loyalty event types used to filter the result.
	// If multiple values are specified, the endpoint uses a
	// logical OR to combine them.
	// See [LoyaltyEventType](#type-loyaltyeventtype) for possible values
	Types []LoyaltyEventType `json:"types,omitempty" url:"types,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LoyaltyEventTypeFilter) GetTypes() []LoyaltyEventType {
	if l == nil {
		return nil
	}
	return l.Types
}

func (l *LoyaltyEventTypeFilter) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LoyaltyEventTypeFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler LoyaltyEventTypeFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LoyaltyEventTypeFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LoyaltyEventTypeFilter) String() string {
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

// A response that contains loyalty events that satisfy the search
// criteria, in order by the `created_at` date.
type SearchLoyaltyEventsResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The loyalty events that satisfy the search criteria.
	Events []*LoyaltyEvent `json:"events,omitempty" url:"events,omitempty"`
	// The pagination cursor to be used in a subsequent
	// request. If empty, this is the final response.
	// For more information,
	// see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SearchLoyaltyEventsResponse) GetErrors() []*Error {
	if s == nil {
		return nil
	}
	return s.Errors
}

func (s *SearchLoyaltyEventsResponse) GetEvents() []*LoyaltyEvent {
	if s == nil {
		return nil
	}
	return s.Events
}

func (s *SearchLoyaltyEventsResponse) GetCursor() *string {
	if s == nil {
		return nil
	}
	return s.Cursor
}

func (s *SearchLoyaltyEventsResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchLoyaltyEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchLoyaltyEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchLoyaltyEventsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchLoyaltyEventsResponse) String() string {
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
