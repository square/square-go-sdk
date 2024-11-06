// This file was auto-generated by Fern from our API Definition.

package batch

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	option "github.com/square/square-go-sdk/option"
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

// Bundle multiple requests to Connect V1 API endpoints as a single
// request.
//
// The **V1SubmitBatch** endpoint does not require an access token
// in the request header. Instead, provide an `access_token` parameter for
// each request included in the batch.
//
// **V1SubmitBatch** responds with an array that contains response objects for
// each of the batched requests. There is no guarantee of the order in
// which batched requests are performed.
//
// **IMPORTANT**
//
// You cannot include more than 30 requests in a single batch and
// recursive requests to **V1SubmitBatch** are not allowed. In
// other words, none of the requests included in a batch can itself be a
// request to the **V1SubmitBatch** endpoint.
func (c *Client) Submit(
	ctx context.Context,
	request *squaregosdk.BatchRequest,
	opts ...option.RequestOption,
) ([]*squaregosdk.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/batch"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response []*squaregosdk.BatchResponse
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
