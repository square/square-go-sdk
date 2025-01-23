// This file was auto-generated by Fern from our API Definition.

package payments

import (
	context "context"
	fmt "fmt"
	v2 "github.com/square/square-go-sdk/v2"
	core "github.com/square/square-go-sdk/v2/core"
	internal "github.com/square/square-go-sdk/v2/internal"
	option "github.com/square/square-go-sdk/v2/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Retrieves a list of payments taken by the account making the request.
//
// Results are eventually consistent, and new payments or changes to payments might take several
// seconds to appear.
//
// The maximum results per page is 100.
func (c *Client) List(
	ctx context.Context,
	request *v2.PaymentsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*v2.Payment], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/payments"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	prepareCall := func(pageRequest *internal.PageRequest[*string]) *internal.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &internal.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *v2.ListPaymentsResponse) *internal.PageResponse[*string, *v2.Payment] {
		next := response.Cursor
		results := response.Payments
		return &internal.PageResponse[*string, *v2.Payment]{
			Next:    next,
			Results: results,
		}
	}
	pager := internal.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Creates a payment using the provided source. You can use this endpoint
// to charge a card (credit/debit card or
// Square gift card) or record a payment that the seller received outside of Square
// (cash payment from a buyer or a payment that an external entity
// processed on behalf of the seller).
//
// The endpoint creates a
// `Payment` object and returns it in the response.
func (c *Client) Create(
	ctx context.Context,
	request *v2.CreatePaymentRequest,
	opts ...option.RequestOption,
) (*v2.CreatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/payments"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CreatePaymentResponse
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

// Cancels (voids) a payment identified by the idempotency key that is specified in the
// request.
//
// Use this method when the status of a `CreatePayment` request is unknown (for example, after you send a
// `CreatePayment` request, a network error occurs and you do not get a response). In this case, you can
// direct Square to cancel the payment using this endpoint. In the request, you provide the same
// idempotency key that you provided in your `CreatePayment` request that you want to cancel. After
// canceling the payment, you can submit your `CreatePayment` request again.
//
// Note that if no payment with the specified idempotency key is found, no action is taken and the endpoint
// returns successfully.
func (c *Client) CancelByIdempotencyKey(
	ctx context.Context,
	request *v2.CancelPaymentByIdempotencyKeyRequest,
	opts ...option.RequestOption,
) (*v2.CancelPaymentByIdempotencyKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/payments/cancel"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CancelPaymentByIdempotencyKeyResponse
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

// Retrieves details for a specific payment.
func (c *Client) Get(
	ctx context.Context,
	request *v2.PaymentsGetRequest,
	opts ...option.RequestOption,
) (*v2.GetPaymentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/payments/%v",
		request.PaymentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.GetPaymentResponse
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

// Updates a payment with the APPROVED status.
// You can update the `amount_money` and `tip_money` using this endpoint.
func (c *Client) Update(
	ctx context.Context,
	request *v2.UpdatePaymentRequest,
	opts ...option.RequestOption,
) (*v2.UpdatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/payments/%v",
		request.PaymentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.UpdatePaymentResponse
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

// Cancels (voids) a payment. You can use this endpoint to cancel a payment with
// the APPROVED `status`.
func (c *Client) Cancel(
	ctx context.Context,
	request *v2.PaymentsCancelRequest,
	opts ...option.RequestOption,
) (*v2.CancelPaymentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/payments/%v/cancel",
		request.PaymentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.CancelPaymentResponse
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
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Completes (captures) a payment.
// By default, payments are set to complete immediately after they are created.
//
// You can use this endpoint to complete a payment with the APPROVED `status`.
func (c *Client) Complete(
	ctx context.Context,
	request *v2.CompletePaymentRequest,
	opts ...option.RequestOption,
) (*v2.CompletePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/payments/%v/complete",
		request.PaymentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CompletePaymentResponse
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
