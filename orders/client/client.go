// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	option "github.com/square/square-go-sdk/option"
	customattributedefinitions "github.com/square/square-go-sdk/orders/customattributedefinitions"
	customattributes "github.com/square/square-go-sdk/orders/customattributes"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	CustomAttributeDefinitions *customattributedefinitions.Client
	CustomAttributes           *customattributes.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.Token == "" {
		options.Token = os.Getenv("SQUARE_TOKEN")
	}
	if options.Version == "" {
		options.Version = os.Getenv("VERSION")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:                     options.ToHeader(),
		CustomAttributeDefinitions: customattributedefinitions.NewClient(opts...),
		CustomAttributes:           customattributes.NewClient(opts...),
	}
}

// Creates a new [order](entity:Order) that can include information about products for
// purchase and settings to apply to the purchase.
//
// To pay for a created order, see
// [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).
//
// You can modify open orders using the [UpdateOrder](api-endpoint:Orders-UpdateOrder) endpoint.
func (c *Client) Create(
	ctx context.Context,
	request *squaregosdk.CreateOrderRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieves a set of [orders](entity:Order) by their IDs.
//
// If a given order ID does not exist, the ID is ignored instead of generating an error.
func (c *Client) BatchGet(
	ctx context.Context,
	request *squaregosdk.BatchGetOrdersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BatchGetOrdersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/batch-retrieve"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BatchGetOrdersResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Enables applications to preview order pricing without creating an order.
func (c *Client) Calculate(
	ctx context.Context,
	request *squaregosdk.CalculateOrderRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CalculateOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/calculate"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CalculateOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Creates a new order, in the `DRAFT` state, by duplicating an existing order. The newly created order has
// only the core fields (such as line items, taxes, and discounts) copied from the original order.
func (c *Client) Clone(
	ctx context.Context,
	request *squaregosdk.CloneOrderRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CloneOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/clone"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CloneOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Search all orders for one or more locations. Orders include all sales,
// returns, and exchanges regardless of how or when they entered the Square
// ecosystem (such as Point of Sale, Invoices, and Connect APIs).
//
// `SearchOrders` requests need to specify which locations to search and define a
// [SearchOrdersQuery](entity:SearchOrdersQuery) object that controls
// how to sort or filter the results. Your `SearchOrdersQuery` can:
//
//	Set filter criteria.
//	Set the sort order.
//	Determine whether to return results as complete `Order` objects or as
//
// [OrderEntry](entity:OrderEntry) objects.
//
// Note that details for orders processed with Square Point of Sale while in
// offline mode might not be transmitted to Square for up to 72 hours. Offline
// orders have a `created_at` value that reflects the time the order was created,
// not the time it was subsequently transmitted to Square.
func (c *Client) Search(
	ctx context.Context,
	request *squaregosdk.SearchOrdersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchOrdersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SearchOrdersResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieves an [Order](entity:Order) by ID.
func (c *Client) Get(
	ctx context.Context,
	request *squaregosdk.OrdersGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/%v",
		request.OrderID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates an open [order](entity:Order) by adding, replacing, or deleting
// fields. Orders with a `COMPLETED` or `CANCELED` state cannot be updated.
//
// An `UpdateOrder` request requires the following:
//
// - The `order_id` in the endpoint path, identifying the order to update.
// - The latest `version` of the order to update.
// - The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#sparse-order-objects)
// containing only the fields to update and the version to which the update is
// being applied.
// - If deleting fields, the [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#identifying-fields-to-delete)
// identifying the fields to clear.
//
// To pay for an order, see
// [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).
func (c *Client) Update(
	ctx context.Context,
	request *squaregosdk.UpdateOrderRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/%v",
		request.OrderID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Pay for an [order](entity:Order) using one or more approved [payments](entity:Payment)
// or settle an order with a total of `0`.
//
// The total of the `payment_ids` listed in the request must be equal to the order
// total. Orders with a total amount of `0` can be marked as paid by specifying an empty
// array of `payment_ids` in the request.
//
// To be used with `PayOrder`, a payment must:
//
// - Reference the order by specifying the `order_id` when [creating the payment](api-endpoint:Payments-CreatePayment).
// Any approved payments that reference the same `order_id` not specified in the
// `payment_ids` is canceled.
// - Be approved with [delayed capture](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture).
// Using a delayed capture payment with `PayOrder` completes the approved payment.
func (c *Client) Pay(
	ctx context.Context,
	request *squaregosdk.PayOrderRequest,
	opts ...option.RequestOption,
) (*squaregosdk.PayOrderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/%v/pay",
		request.OrderID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.PayOrderResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
