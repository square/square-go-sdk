// This file was auto-generated by Fern from our API Definition.

package customattributedefinitions

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	option "github.com/square/square-go-sdk/option"
	orders "github.com/square/square-go-sdk/orders"
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

// Lists the order-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.
//
// When all response pages are retrieved, the results include all custom attribute definitions
// that are visible to the requesting application, including those that are created by other
// applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that
// seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) List(
	ctx context.Context,
	request *orders.CustomAttributeDefinitionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CustomAttributeDefinition], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/custom-attribute-definitions"
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
	readPageResponse := func(response *squaregosdk.ListOrderCustomAttributeDefinitionsResponse) *internal.PageResponse[*string, *squaregosdk.CustomAttributeDefinition] {
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

// Creates an order-related custom attribute definition.  Use this endpoint to
// define a custom attribute that can be associated with orders.
//
// After creating a custom attribute definition, you can set the custom attribute for orders
// in the Square seller account.
func (c *Client) Create(
	ctx context.Context,
	request *orders.CreateOrderCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/orders/custom-attribute-definitions"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateOrderCustomAttributeDefinitionResponse
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

// Retrieves an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// To retrieve a custom attribute definition created by another application, the `visibility`
// setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Get(
	ctx context.Context,
	request *orders.CustomAttributeDefinitionsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/custom-attribute-definitions/%v",
		request.Key,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.RetrieveOrderCustomAttributeDefinitionResponse
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

// Updates an order-related custom attribute definition for a Square seller account.
//
// Only the definition owner can update a custom attribute definition. Note that sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.
func (c *Client) Update(
	ctx context.Context,
	request *orders.UpdateOrderCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateOrderCustomAttributeDefinitionResponse
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

// Deletes an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// Only the definition owner can delete a custom attribute definition.
func (c *Client) Delete(
	ctx context.Context,
	request *orders.CustomAttributeDefinitionsDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/orders/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.DeleteOrderCustomAttributeDefinitionResponse
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
