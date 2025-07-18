// This file was auto-generated by Fern from our API Definition.

package cards

import (
	context "context"
	v2 "github.com/square/square-go-sdk/v2"
	core "github.com/square/square-go-sdk/v2/core"
	customers "github.com/square/square-go-sdk/v2/customers"
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

// Adds a card on file to an existing customer.
//
// As with charges, calls to `CreateCustomerCard` are idempotent. Multiple
// calls with the same card nonce return the same card record that was created
// with the provided nonce during the _first_ call.
func (c *Client) Create(
	ctx context.Context,
	request *customers.CreateCustomerCardRequest,
	opts ...option.RequestOption,
) (*v2.CreateCustomerCardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/cards",
		request.CustomerID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CreateCustomerCardResponse
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

// Removes a card on file from a customer.
func (c *Client) Delete(
	ctx context.Context,
	request *customers.DeleteCardsRequest,
	opts ...option.RequestOption,
) (*v2.DeleteCustomerCardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/cards/%v",
		request.CustomerID,
		request.CardID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.DeleteCustomerCardResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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
