// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	customattributes "github.com/square/square-go-sdk/merchants/customattributes"
	option "github.com/square/square-go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	CustomAttributes *customattributes.Client
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
		header:           options.ToHeader(),
		CustomAttributes: customattributes.NewClient(opts...),
	}
}

// Provides details about the merchant associated with a given access token.
//
// The access token used to connect your application to a Square seller is associated
// with a single merchant. That means that `ListMerchants` returns a list
// with a single `Merchant` object. You can specify your personal access token
// to get your own merchant information or specify an OAuth token to get the
// information for the merchant that granted your application access.
//
// If you know the merchant ID, you can also use the [RetrieveMerchant]($e/Merchants/RetrieveMerchant)
// endpoint to retrieve the merchant information.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.MerchantsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.Merchant], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := baseURL + "/v2/merchants"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	prepareCall := func(pageRequest *internal.PageRequest[*int]) *internal.CallParams {
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
	readPageResponse := func(response *squaregosdk.ListMerchantsResponse) *internal.PageResponse[*int, *squaregosdk.Merchant] {
		next := response.Cursor
		results := response.Merchant
		return &internal.PageResponse[*int, *squaregosdk.Merchant]{
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

// Lists the merchant-related [custom attribute definitions]($m/CustomAttributeDefinition) that belong to a Square seller account.
// When all response pages are retrieved, the results include all custom attribute definitions
// that are visible to the requesting application, including those that are created by other
// applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) GetCustomAttributeDefinitions(
	ctx context.Context,
	request *squaregosdk.MerchantsGetCustomAttributeDefinitionsRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CustomAttributeDefinition], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := baseURL + "/v2/merchants/custom-attribute-definitions"
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
	readPageResponse := func(response *squaregosdk.ListMerchantCustomAttributeDefinitionsResponse) *internal.PageResponse[*string, *squaregosdk.CustomAttributeDefinition] {
		next := response.Cursor
		results := response.CustomAttributeDefinitions
		return &internal.PageResponse[*string, *squaregosdk.CustomAttributeDefinition]{
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

// Retrieves the `Merchant` object for the given `merchant_id`.
func (c *Client) Get(
	ctx context.Context,
	request *squaregosdk.MerchantsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetMerchantResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/merchants/%v",
		request.MerchantID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetMerchantResponse
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