// This file was auto-generated by Fern from our API Definition.

package refunds

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	option "github.com/square/square-go-sdk/option"
	terminal "github.com/square/square-go-sdk/terminal"
	http "net/http"
	os "os"
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

// Creates a request to refund an Interac payment completed on a Square Terminal. Refunds for Interac payments on a Square Terminal are supported only for Interac debit cards in Canada. Other refunds for Terminal payments should use the Refunds API. For more information, see [Refunds API](api:Refunds).
func (c *Client) Create(
	ctx context.Context,
	request *terminal.CreateTerminalRefundRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateTerminalRefundResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/terminals/refunds"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreateTerminalRefundResponse
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

// Retrieves a filtered list of Interac Terminal refund requests created by the seller making the request. Terminal refund requests are available for 30 days.
func (c *Client) Search(
	ctx context.Context,
	request *terminal.SearchTerminalRefundsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchTerminalRefundsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/terminals/refunds/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.SearchTerminalRefundsResponse
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

// Retrieves an Interac Terminal refund object by ID. Terminal refund objects are available for 30 days.
func (c *Client) Get(
	ctx context.Context,
	// The unique ID for the desired `TerminalRefund`.
	terminalRefundID string,
	opts ...option.RequestOption,
) (*squaregosdk.GetTerminalRefundResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/terminals/refunds/%v", terminalRefundID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetTerminalRefundResponse
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

// Cancels an Interac Terminal refund request by refund request ID if the status of the request permits it.
func (c *Client) Cancel(
	ctx context.Context,
	// The unique ID for the desired `TerminalRefund`.
	terminalRefundID string,
	request *terminal.CancelTerminalRefundRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CancelTerminalRefundResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/terminals/refunds/%v/cancel", terminalRefundID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CancelTerminalRefundResponse
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
