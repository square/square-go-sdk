// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	http "net/http"
	os "os"

	squaregosdk "github.com/fern-demo/square-go-sdk"
	core "github.com/fern-demo/square-go-sdk/core"
	codes "github.com/fern-demo/square-go-sdk/devices/codes"
	option "github.com/fern-demo/square-go-sdk/option"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Codes *codes.Client
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
		Codes:  codes.NewClient(opts...),
	}
}

// List devices associated with the merchant. Currently, only Terminal API
// devices are supported.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.DevicesListRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListDevicesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/devices"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListDevicesResponse
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

// Retrieves Device with the associated `device_id`.
func (c *Client) Get(
	ctx context.Context,
	// The unique ID for the desired `Device`.
	deviceID string,
	opts ...option.RequestOption,
) (*squaregosdk.GetDeviceResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/devices/%v", deviceID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetDeviceResponse
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
