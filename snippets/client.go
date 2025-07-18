// This file was auto-generated by Fern from our API Definition.

package snippets

import (
	context "context"
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

// Retrieves your snippet from a Square Online site. A site can contain snippets from multiple snippet applications, but you can retrieve only the snippet that was added by your application.
//
// You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.
//
// __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
func (c *Client) Get(
	ctx context.Context,
	request *v2.GetSnippetsRequest,
	opts ...option.RequestOption,
) (*v2.GetSnippetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/sites/%v/snippet",
		request.SiteID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.GetSnippetResponse
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

// Adds a snippet to a Square Online site or updates the existing snippet on the site.
// The snippet code is appended to the end of the `head` element on every page of the site, except checkout pages. A snippet application can add one snippet to a given site.
//
// You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.
//
// __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
func (c *Client) Upsert(
	ctx context.Context,
	request *v2.UpsertSnippetRequest,
	opts ...option.RequestOption,
) (*v2.UpsertSnippetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/sites/%v/snippet",
		request.SiteID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.UpsertSnippetResponse
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

// Removes your snippet from a Square Online site.
//
// You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.
//
// __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
func (c *Client) Delete(
	ctx context.Context,
	request *v2.DeleteSnippetsRequest,
	opts ...option.RequestOption,
) (*v2.DeleteSnippetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/sites/%v/snippet",
		request.SiteID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.DeleteSnippetResponse
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
