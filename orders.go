// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type BatchGetOrdersRequest struct {
	// The ID of the location for these orders. This field is optional: omit it to retrieve
	// orders within the scope of the current authorization's merchant ID.
	LocationID *string `json:"location_id,omitempty" url:"-"`
	// The IDs of the orders to retrieve. A maximum of 100 orders can be retrieved per request.
	OrderIDs []string `json:"order_ids,omitempty" url:"-"`
}

type CalculateOrderRequest struct {
	// The order to be calculated. Expects the entire order, not a sparse update.
	Order *Order `json:"order,omitempty" url:"-"`
	// Identifies one or more loyalty reward tiers to apply during the order calculation.
	// The discounts defined by the reward tiers are added to the order only to preview the
	// effect of applying the specified rewards. The rewards do not correspond to actual
	// redemptions; that is, no `reward`s are created. Therefore, the reward `id`s are
	// random strings used only to reference the reward tier.
	ProposedRewards []*OrderReward `json:"proposed_rewards,omitempty" url:"-"`
}

type CloneOrderRequest struct {
	// The ID of the order to clone.
	OrderID string `json:"order_id" url:"-"`
	// An optional order version for concurrency protection.
	//
	// If a version is provided, it must match the latest stored version of the order to clone.
	// If a version is not provided, the API clones the latest version.
	Version *int `json:"version,omitempty" url:"-"`
	// A value you specify that uniquely identifies this clone request.
	//
	// If you are unsure whether a particular order was cloned successfully,
	// you can reattempt the call with the same idempotency key without
	// worrying about creating duplicate cloned orders.
	// The originally cloned order is returned.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type PayOrderRequest struct {
	// A value you specify that uniquely identifies this request among requests you have sent. If
	// you are unsure whether a particular payment request was completed successfully, you can reattempt
	// it with the same idempotency key without worrying about duplicate payments.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The version of the order being paid. If not supplied, the latest version will be paid.
	OrderVersion *int `json:"order_version,omitempty" url:"-"`
	// The IDs of the [payments](entity:Payment) to collect.
	// The payment total must match the order total.
	PaymentIDs []string `json:"payment_ids,omitempty" url:"-"`
}

type SearchOrdersRequest struct {
	// The location IDs for the orders to query. All locations must belong to
	// the same merchant.
	//
	// Max: 10 location IDs.
	LocationIDs []string `json:"location_ids,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
	// Query conditions used to filter or sort the results. Note that when
	// retrieving additional pages using a cursor, you must use the original query.
	Query *SearchOrdersQuery `json:"query,omitempty" url:"-"`
	// The maximum number of results to be returned in a single page.
	//
	// Default: `500`
	// Max: `1000`
	Limit *int `json:"limit,omitempty" url:"-"`
	// A Boolean that controls the format of the search results. If `true`,
	// `SearchOrders` returns [OrderEntry](entity:OrderEntry) objects. If `false`, `SearchOrders`
	// returns complete order objects.
	//
	// Default: `false`.
	ReturnEntries *bool `json:"return_entries,omitempty" url:"-"`
}

// Defines the fields that are included in the response body of
// a request to the `BatchRetrieveOrders` endpoint.
type BatchGetOrdersResponse struct {
	// The requested orders. This will omit any requested orders that do not exist.
	Orders []*Order `json:"orders,omitempty" url:"orders,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BatchGetOrdersResponse) GetOrders() []*Order {
	if b == nil {
		return nil
	}
	return b.Orders
}

func (b *BatchGetOrdersResponse) GetErrors() []*Error {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BatchGetOrdersResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchGetOrdersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchGetOrdersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchGetOrdersResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchGetOrdersResponse) String() string {
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

type CalculateOrderResponse struct {
	// The calculated version of the order provided in the request.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CalculateOrderResponse) GetOrder() *Order {
	if c == nil {
		return nil
	}
	return c.Order
}

func (c *CalculateOrderResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CalculateOrderResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CalculateOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CalculateOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CalculateOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CalculateOrderResponse) String() string {
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

// Defines the fields that are included in the response body of
// a request to the [CloneOrder](api-endpoint:Orders-CloneOrder) endpoint.
type CloneOrderResponse struct {
	// The cloned order.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CloneOrderResponse) GetOrder() *Order {
	if c == nil {
		return nil
	}
	return c.Order
}

func (c *CloneOrderResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CloneOrderResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CloneOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CloneOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CloneOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CloneOrderResponse) String() string {
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

// Defines the fields that are included in the response body of
// a request to the `CreateOrder` endpoint.
//
// Either `errors` or `order` is present in a given response, but never both.
type CreateOrderResponse struct {
	// The newly created order.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateOrderResponse) GetOrder() *Order {
	if c == nil {
		return nil
	}
	return c.Order
}

func (c *CreateOrderResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateOrderResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateOrderResponse) String() string {
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

type GetOrderResponse struct {
	// The requested order.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetOrderResponse) GetOrder() *Order {
	if g == nil {
		return nil
	}
	return g.Order
}

func (g *GetOrderResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetOrderResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetOrderResponse) String() string {
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

// A lightweight description of an [order](entity:Order) that is returned when
// `returned_entries` is `true` on a [SearchOrdersRequest](api-endpoint:Orders-SearchOrders).
type OrderEntry struct {
	// The ID of the order.
	OrderID *string `json:"order_id,omitempty" url:"order_id,omitempty"`
	// The version number, which is incremented each time an update is committed to the order.
	// Orders that were not created through the API do not include a version number and
	// therefore cannot be updated.
	//
	// [Read more about working with versions.](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders)
	Version *int `json:"version,omitempty" url:"version,omitempty"`
	// The location ID the order belongs to.
	LocationID *string `json:"location_id,omitempty" url:"location_id,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (o *OrderEntry) GetOrderID() *string {
	if o == nil {
		return nil
	}
	return o.OrderID
}

func (o *OrderEntry) GetVersion() *int {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *OrderEntry) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *OrderEntry) GetExtraProperties() map[string]interface{} {
	return o.extraProperties
}

func (o *OrderEntry) UnmarshalJSON(data []byte) error {
	type unmarshaler OrderEntry
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrderEntry(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties

	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrderEntry) String() string {
	if len(o._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

// Defines the fields that are included in the response body of a request to the
// [PayOrder](api-endpoint:Orders-PayOrder) endpoint.
type PayOrderResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The paid, updated [order](entity:Order).
	Order *Order `json:"order,omitempty" url:"order,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PayOrderResponse) GetErrors() []*Error {
	if p == nil {
		return nil
	}
	return p.Errors
}

func (p *PayOrderResponse) GetOrder() *Order {
	if p == nil {
		return nil
	}
	return p.Order
}

func (p *PayOrderResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PayOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PayOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PayOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PayOrderResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// A filter based on the order `customer_id` and any tender `customer_id`
// associated with the order. It does not filter based on the
// [FulfillmentRecipient](entity:FulfillmentRecipient) `customer_id`.
type SearchOrdersCustomerFilter struct {
	// A list of customer IDs to filter by.
	//
	// Max: 10 customer ids.
	CustomerIDs []string `json:"customer_ids,omitempty" url:"customer_ids,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersCustomerFilter) GetCustomerIDs() []string {
	if s == nil {
		return nil
	}
	return s.CustomerIDs
}

func (s *SearchOrdersCustomerFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersCustomerFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersCustomerFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersCustomerFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersCustomerFilter) String() string {
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

// Filter for `Order` objects based on whether their `CREATED_AT`,
// `CLOSED_AT`, or `UPDATED_AT` timestamps fall within a specified time range.
// You can specify the time range and which timestamp to filter for. You can filter
// for only one time range at a time.
//
// For each time range, the start time and end time are inclusive. If the end time
// is absent, it defaults to the time of the first request for the cursor.
//
// **Important:** If you use the `DateTimeFilter` in a `SearchOrders` query,
// you must set the `sort_field` in [OrdersSort](entity:SearchOrdersSort)
// to the same field you filter for. For example, if you set the `CLOSED_AT` field
// in `DateTimeFilter`, you must set the `sort_field` in `SearchOrdersSort` to
// `CLOSED_AT`. Otherwise, `SearchOrders` throws an error.
// [Learn more about filtering orders by time range.](https://developer.squareup.com/docs/orders-api/manage-orders/search-orders#important-note-about-filtering-orders-by-time-range)
type SearchOrdersDateTimeFilter struct {
	// The time range for filtering on the `created_at` timestamp. If you use this
	// value, you must set the `sort_field` in the `OrdersSearchSort` object to
	// `CREATED_AT`.
	CreatedAt *TimeRange `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The time range for filtering on the `updated_at` timestamp. If you use this
	// value, you must set the `sort_field` in the `OrdersSearchSort` object to
	// `UPDATED_AT`.
	UpdatedAt *TimeRange `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The time range for filtering on the `closed_at` timestamp. If you use this
	// value, you must set the `sort_field` in the `OrdersSearchSort` object to
	// `CLOSED_AT`.
	ClosedAt *TimeRange `json:"closed_at,omitempty" url:"closed_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersDateTimeFilter) GetCreatedAt() *TimeRange {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *SearchOrdersDateTimeFilter) GetUpdatedAt() *TimeRange {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SearchOrdersDateTimeFilter) GetClosedAt() *TimeRange {
	if s == nil {
		return nil
	}
	return s.ClosedAt
}

func (s *SearchOrdersDateTimeFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersDateTimeFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersDateTimeFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersDateTimeFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersDateTimeFilter) String() string {
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

// Filtering criteria to use for a `SearchOrders` request. Multiple filters
// are ANDed together.
type SearchOrdersFilter struct {
	// Filter by [OrderState](entity:OrderState).
	StateFilter *SearchOrdersStateFilter `json:"state_filter,omitempty" url:"state_filter,omitempty"`
	// Filter for results within a time range.
	//
	// **Important:** If you filter for orders by time range, you must set `SearchOrdersSort`
	// to sort by the same field.
	// [Learn more about filtering orders by time range.](https://developer.squareup.com/docs/orders-api/manage-orders/search-orders#important-note-about-filtering-orders-by-time-range)
	DateTimeFilter *SearchOrdersDateTimeFilter `json:"date_time_filter,omitempty" url:"date_time_filter,omitempty"`
	// Filter by the fulfillment type or state.
	FulfillmentFilter *SearchOrdersFulfillmentFilter `json:"fulfillment_filter,omitempty" url:"fulfillment_filter,omitempty"`
	// Filter by the source of the order.
	SourceFilter *SearchOrdersSourceFilter `json:"source_filter,omitempty" url:"source_filter,omitempty"`
	// Filter by customers associated with the order.
	CustomerFilter *SearchOrdersCustomerFilter `json:"customer_filter,omitempty" url:"customer_filter,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersFilter) GetStateFilter() *SearchOrdersStateFilter {
	if s == nil {
		return nil
	}
	return s.StateFilter
}

func (s *SearchOrdersFilter) GetDateTimeFilter() *SearchOrdersDateTimeFilter {
	if s == nil {
		return nil
	}
	return s.DateTimeFilter
}

func (s *SearchOrdersFilter) GetFulfillmentFilter() *SearchOrdersFulfillmentFilter {
	if s == nil {
		return nil
	}
	return s.FulfillmentFilter
}

func (s *SearchOrdersFilter) GetSourceFilter() *SearchOrdersSourceFilter {
	if s == nil {
		return nil
	}
	return s.SourceFilter
}

func (s *SearchOrdersFilter) GetCustomerFilter() *SearchOrdersCustomerFilter {
	if s == nil {
		return nil
	}
	return s.CustomerFilter
}

func (s *SearchOrdersFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersFilter) String() string {
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

// Filter based on [order fulfillment](entity:Fulfillment) information.
type SearchOrdersFulfillmentFilter struct {
	// A list of [fulfillment types](entity:FulfillmentType) to filter
	// for. The list returns orders if any of its fulfillments match any of the fulfillment types
	// listed in this field.
	// See [FulfillmentType](#type-fulfillmenttype) for possible values
	FulfillmentTypes []FulfillmentType `json:"fulfillment_types,omitempty" url:"fulfillment_types,omitempty"`
	// A list of [fulfillment states](entity:FulfillmentState) to filter
	// for. The list returns orders if any of its fulfillments match any of the
	// fulfillment states listed in this field.
	// See [FulfillmentState](#type-fulfillmentstate) for possible values
	FulfillmentStates []FulfillmentState `json:"fulfillment_states,omitempty" url:"fulfillment_states,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersFulfillmentFilter) GetFulfillmentTypes() []FulfillmentType {
	if s == nil {
		return nil
	}
	return s.FulfillmentTypes
}

func (s *SearchOrdersFulfillmentFilter) GetFulfillmentStates() []FulfillmentState {
	if s == nil {
		return nil
	}
	return s.FulfillmentStates
}

func (s *SearchOrdersFulfillmentFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersFulfillmentFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersFulfillmentFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersFulfillmentFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersFulfillmentFilter) String() string {
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
type SearchOrdersQuery struct {
	// Criteria to filter results by.
	Filter *SearchOrdersFilter `json:"filter,omitempty" url:"filter,omitempty"`
	// Criteria to sort results by.
	Sort *SearchOrdersSort `json:"sort,omitempty" url:"sort,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersQuery) GetFilter() *SearchOrdersFilter {
	if s == nil {
		return nil
	}
	return s.Filter
}

func (s *SearchOrdersQuery) GetSort() *SearchOrdersSort {
	if s == nil {
		return nil
	}
	return s.Sort
}

func (s *SearchOrdersQuery) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersQuery) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersQuery
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersQuery(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersQuery) String() string {
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

// Either the `order_entries` or `orders` field is set, depending on whether
// `return_entries` is set on the [SearchOrdersRequest](api-endpoint:Orders-SearchOrders).
type SearchOrdersResponse struct {
	// A list of [OrderEntries](entity:OrderEntry) that fit the query
	// conditions. The list is populated only if `return_entries` is set to `true` in the request.
	OrderEntries []*OrderEntry `json:"order_entries,omitempty" url:"order_entries,omitempty"`
	// A list of
	// [Order](entity:Order) objects that match the query conditions. The list is populated only if
	// `return_entries` is set to `false` in the request.
	Orders []*Order `json:"orders,omitempty" url:"orders,omitempty"`
	// The pagination cursor to be used in a subsequent request. If unset,
	// this is the final response.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`
	// [Errors](entity:Error) encountered during the search.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersResponse) GetOrderEntries() []*OrderEntry {
	if s == nil {
		return nil
	}
	return s.OrderEntries
}

func (s *SearchOrdersResponse) GetOrders() []*Order {
	if s == nil {
		return nil
	}
	return s.Orders
}

func (s *SearchOrdersResponse) GetCursor() *string {
	if s == nil {
		return nil
	}
	return s.Cursor
}

func (s *SearchOrdersResponse) GetErrors() []*Error {
	if s == nil {
		return nil
	}
	return s.Errors
}

func (s *SearchOrdersResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersResponse) String() string {
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

// Sorting criteria for a `SearchOrders` request. Results can only be sorted
// by a timestamp field.
type SearchOrdersSort struct {
	// The field to sort by.
	//
	// **Important:** When using a [DateTimeFilter](entity:SearchOrdersFilter),
	// `sort_field` must match the timestamp field that the `DateTimeFilter` uses to
	// filter. For example, if you set your `sort_field` to `CLOSED_AT` and you use a
	// `DateTimeFilter`, your `DateTimeFilter` must filter for orders by their `CLOSED_AT` date.
	// If this field does not match the timestamp field in `DateTimeFilter`,
	// `SearchOrders` returns an error.
	//
	// Default: `CREATED_AT`.
	// See [SearchOrdersSortField](#type-searchorderssortfield) for possible values
	SortField SearchOrdersSortField `json:"sort_field" url:"sort_field"`
	// The chronological order in which results are returned. Defaults to `DESC`.
	// See [SortOrder](#type-sortorder) for possible values
	SortOrder *SortOrder `json:"sort_order,omitempty" url:"sort_order,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersSort) GetSortField() SearchOrdersSortField {
	if s == nil {
		return ""
	}
	return s.SortField
}

func (s *SearchOrdersSort) GetSortOrder() *SortOrder {
	if s == nil {
		return nil
	}
	return s.SortOrder
}

func (s *SearchOrdersSort) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersSort) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersSort
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersSort(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersSort) String() string {
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

// Specifies which timestamp to use to sort `SearchOrder` results.
type SearchOrdersSortField string

const (
	SearchOrdersSortFieldDoNotUse  SearchOrdersSortField = "DO_NOT_USE"
	SearchOrdersSortFieldCreatedAt SearchOrdersSortField = "CREATED_AT"
	SearchOrdersSortFieldUpdatedAt SearchOrdersSortField = "UPDATED_AT"
	SearchOrdersSortFieldClosedAt  SearchOrdersSortField = "CLOSED_AT"
	SearchOrdersSortFieldPaidAt    SearchOrdersSortField = "PAID_AT"
	SearchOrdersSortFieldScore     SearchOrdersSortField = "SCORE"
	SearchOrdersSortFieldDueAt     SearchOrdersSortField = "DUE_AT"
)

func NewSearchOrdersSortFieldFromString(s string) (SearchOrdersSortField, error) {
	switch s {
	case "DO_NOT_USE":
		return SearchOrdersSortFieldDoNotUse, nil
	case "CREATED_AT":
		return SearchOrdersSortFieldCreatedAt, nil
	case "UPDATED_AT":
		return SearchOrdersSortFieldUpdatedAt, nil
	case "CLOSED_AT":
		return SearchOrdersSortFieldClosedAt, nil
	case "PAID_AT":
		return SearchOrdersSortFieldPaidAt, nil
	case "SCORE":
		return SearchOrdersSortFieldScore, nil
	case "DUE_AT":
		return SearchOrdersSortFieldDueAt, nil
	}
	var t SearchOrdersSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SearchOrdersSortField) Ptr() *SearchOrdersSortField {
	return &s
}

// A filter based on order `source` information.
type SearchOrdersSourceFilter struct {
	// Filters by the [Source](entity:OrderSource) `name`. The filter returns any orders
	// with a `source.name` that matches any of the listed source names.
	//
	// Max: 10 source names.
	SourceNames []string `json:"source_names,omitempty" url:"source_names,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersSourceFilter) GetSourceNames() []string {
	if s == nil {
		return nil
	}
	return s.SourceNames
}

func (s *SearchOrdersSourceFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersSourceFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersSourceFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersSourceFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersSourceFilter) String() string {
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

// Filter by the current order `state`.
type SearchOrdersStateFilter struct {
	// States to filter for.
	// See [OrderState](#type-orderstate) for possible values
	States []OrderState `json:"states,omitempty" url:"states,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchOrdersStateFilter) GetStates() []OrderState {
	if s == nil {
		return nil
	}
	return s.States
}

func (s *SearchOrdersStateFilter) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchOrdersStateFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchOrdersStateFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchOrdersStateFilter(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchOrdersStateFilter) String() string {
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
// a request to the [UpdateOrder](api-endpoint:Orders-UpdateOrder) endpoint.
type UpdateOrderResponse struct {
	// The updated order.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateOrderResponse) GetOrder() *Order {
	if u == nil {
		return nil
	}
	return u.Order
}

func (u *UpdateOrderResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateOrderResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateOrderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateOrderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateOrderResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateOrderResponse) String() string {
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

type UpdateOrderRequest struct {
	// The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#sparse-order-objects)
	// containing only the fields to update and the version to which the update is
	// being applied.
	Order *Order `json:"order,omitempty" url:"-"`
	// The [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#identifying-fields-to-delete)
	// fields to clear. For example, `line_items[uid].note`.
	// For more information, see [Deleting fields](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#deleting-fields).
	FieldsToClear []string `json:"fields_to_clear,omitempty" url:"-"`
	// A value you specify that uniquely identifies this update request.
	//
	// If you are unsure whether a particular update was applied to an order successfully,
	// you can reattempt it with the same idempotency key without
	// worrying about creating duplicate updates to the order.
	// The latest order version is returned.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}
