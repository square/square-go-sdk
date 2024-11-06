// This file was auto-generated by Fern from our API Definition.

package applepay

import (
	context "context"
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

// Activates a domain for use with Apple Pay on the Web and Square. A validation
// is performed on this domain by Apple to ensure that it is properly set up as
// an Apple Pay enabled domain.
//
// This endpoint provides an easy way for platform developers to bulk activate
// Apple Pay on the Web with Square for merchants using their platform.
//
// Note: You will need to host a valid domain verification file on your domain to support Apple Pay. The
// current version of this file is always available at https://app.squareup.com/digital-wallets/apple-pay/apple-developer-merchantid-domain-association,
// and should be hosted at `.well_known/apple-developer-merchantid-domain-association` on your
// domain. This file is subject to change; we strongly recommend checking for updates regularly and avoiding
// long-lived caches that might not keep in sync with the correct file version.
//
// To learn more about the Web Payments SDK and how to add Apple Pay, see [Take an Apple Pay Payment](https://developer.squareup.com/docs/web-payments/apple-pay).
func (c *Client) RegisterDomain(
	ctx context.Context,
	request *squaregosdk.RegisterDomainRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RegisterDomainResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/apple-pay/domains"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RegisterDomainResponse
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
