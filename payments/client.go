// This file was auto-generated by Fern from our API Definition.

package payments

import (
	context "context"
	fmt "fmt"
	http "net/http"
	os "os"

	squaregosdk "github.com/fern-demo/square-go-sdk"
	core "github.com/fern-demo/square-go-sdk/core"
	option "github.com/fern-demo/square-go-sdk/option"
)

type Client struct {
	baseURL string
	caller  *core.Caller
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
		caller: core.NewCaller(
			&core.CallerParams{
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
	request *squaregosdk.PaymentsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.Payment], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/payments"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*string]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &core.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *squaregosdk.ListPaymentsResponse) *core.PageResponse[*string, *squaregosdk.Payment] {
		next := response.Cursor
		results := response.Payments
		return &core.PageResponse[*string, *squaregosdk.Payment]{
			Next:    next,
			Results: results,
		}
	}
	pager := core.NewCursorPager(
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
	request *squaregosdk.CreatePaymentRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/payments"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreatePaymentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
	// The ID of the payment to cancel.
	paymentID string,
	opts ...option.RequestOption,
) (*squaregosdk.CancelPaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v/cancel", paymentID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CancelPaymentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Retrieves details for a specific payment.
func (c *Client) Get(
	ctx context.Context,
	// A unique ID for the desired payment.
	paymentID string,
	opts ...option.RequestOption,
) (*squaregosdk.GetPaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v", paymentID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetPaymentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
	// The ID of the payment to update.
	paymentID string,
	request *squaregosdk.UpdatePaymentRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v", paymentID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdatePaymentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Completes (captures) a payment.
// By default, payments are set to complete immediately after they are created.
//
// You can use this endpoint to complete a payment with the APPROVED `status`.
func (c *Client) Complete(
	ctx context.Context,
	// The unique ID identifying the payment to be completed.
	paymentID string,
	request *squaregosdk.CompletePaymentRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CompletePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v/complete", paymentID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CompletePaymentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
