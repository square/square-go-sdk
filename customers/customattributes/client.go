// This file was auto-generated by Fern from our API Definition.

package customattributes

import (
	context "context"
	fmt "fmt"
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

// Lists the [custom attributes](entity:CustomAttribute) associated with a customer profile.
//
// You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
// in the same call.
//
// When all response pages are retrieved, the results include all custom attributes that are
// visible to the requesting application, including those that are owned by other applications
// and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) List(
	ctx context.Context,
	request *customers.ListCustomAttributesRequest,
	opts ...option.RequestOption,
) (*core.Page[*v2.CustomAttribute], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/custom-attributes",
		request.CustomerID,
	)
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
	readPageResponse := func(response *v2.ListCustomerCustomAttributesResponse) *internal.PageResponse[*string, *v2.CustomAttribute] {
		var zeroValue *string
		next := response.Cursor
		results := response.CustomAttributes
		return &internal.PageResponse[*string, *v2.CustomAttribute]{
			Next:    next,
			Results: results,
			Done:    next == zeroValue,
		}
	}
	pager := internal.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Retrieves a [custom attribute](entity:CustomAttribute) associated with a customer profile.
//
// You can use the `with_definition` query parameter to also retrieve the custom attribute definition
// in the same call.
//
// To retrieve a custom attribute owned by another application, the `visibility` setting must be
// `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Get(
	ctx context.Context,
	request *customers.GetCustomAttributesRequest,
	opts ...option.RequestOption,
) (*v2.GetCustomerCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/custom-attributes/%v",
		request.CustomerID,
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

	var response *v2.GetCustomerCustomAttributeResponse
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

// Creates or updates a [custom attribute](entity:CustomAttribute) for a customer profile.
//
// Use this endpoint to set the value of a custom attribute for a specified customer profile.
// A custom attribute is based on a custom attribute definition in a Square seller account, which
// is created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.
//
// To create or update a custom attribute owned by another application, the `visibility` setting
// must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Upsert(
	ctx context.Context,
	request *customers.UpsertCustomerCustomAttributeRequest,
	opts ...option.RequestOption,
) (*v2.UpsertCustomerCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/custom-attributes/%v",
		request.CustomerID,
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.UpsertCustomerCustomAttributeResponse
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

// Deletes a [custom attribute](entity:CustomAttribute) associated with a customer profile.
//
// To delete a custom attribute owned by another application, the `visibility` setting must be
// `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Delete(
	ctx context.Context,
	request *customers.DeleteCustomAttributesRequest,
	opts ...option.RequestOption,
) (*v2.DeleteCustomerCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v/custom-attributes/%v",
		request.CustomerID,
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.DeleteCustomerCustomAttributeResponse
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
