// This file was auto-generated by Fern from our API Definition.

package customattributedefinitions

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

// Lists the customer-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.
//
// When all response pages are retrieved, the results include all custom attribute definitions
// that are visible to the requesting application, including those that are created by other
// applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that
// seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) List(
	ctx context.Context,
	request *customers.CustomAttributeDefinitionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*v2.CustomAttributeDefinition], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/custom-attribute-definitions"
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
	readPageResponse := func(response *v2.ListCustomerCustomAttributeDefinitionsResponse) *internal.PageResponse[*string, *v2.CustomAttributeDefinition] {
		next := response.Cursor
		results := response.CustomAttributeDefinitions
		return &internal.PageResponse[*string, *v2.CustomAttributeDefinition]{
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

// Creates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
// Use this endpoint to define a custom attribute that can be associated with customer profiles.
//
// A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties
// for a custom attribute. After the definition is created, you can call
// [UpsertCustomerCustomAttribute](api-endpoint:CustomerCustomAttributes-UpsertCustomerCustomAttribute) or
// [BulkUpsertCustomerCustomAttributes](api-endpoint:CustomerCustomAttributes-BulkUpsertCustomerCustomAttributes)
// to set the custom attribute for customer profiles in the seller's Customer Directory.
//
// Sellers can view all custom attributes in exported customer data, including those set to
// `VISIBILITY_HIDDEN`.
func (c *Client) Create(
	ctx context.Context,
	request *customers.CreateCustomerCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*v2.CreateCustomerCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/custom-attribute-definitions"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CreateCustomerCustomAttributeDefinitionResponse
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

// Retrieves a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// To retrieve a custom attribute definition created by another application, the `visibility`
// setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Get(
	ctx context.Context,
	request *customers.CustomAttributeDefinitionsGetRequest,
	opts ...option.RequestOption,
) (*v2.GetCustomerCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/custom-attribute-definitions/%v",
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

	var response *v2.GetCustomerCustomAttributeDefinitionResponse
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

// Updates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
//
// Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the
// `schema` for a `Selection` data type.
//
// Only the definition owner can update a custom attribute definition. Note that sellers can view
// all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.
func (c *Client) Update(
	ctx context.Context,
	request *customers.UpdateCustomerCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*v2.UpdateCustomerCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.UpdateCustomerCustomAttributeDefinitionResponse
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

// Deletes a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// Deleting a custom attribute definition also deletes the corresponding custom attribute from
// all customer profiles in the seller's Customer Directory.
//
// Only the definition owner can delete a custom attribute definition.
func (c *Client) Delete(
	ctx context.Context,
	request *customers.CustomAttributeDefinitionsDeleteRequest,
	opts ...option.RequestOption,
) (*v2.DeleteCustomerCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.DeleteCustomerCustomAttributeDefinitionResponse
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

// Creates or updates [custom attributes](entity:CustomAttribute) for customer profiles as a bulk operation.
//
// Use this endpoint to set the value of one or more custom attributes for one or more customer profiles.
// A custom attribute is based on a custom attribute definition in a Square seller account, which is
// created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.
//
// This `BulkUpsertCustomerCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
// requests and returns a map of individual upsert responses. Each upsert request has a unique ID
// and provides a customer ID and custom attribute. Each upsert response is returned with the ID
// of the corresponding request.
//
// To create or update a custom attribute owned by another application, the `visibility` setting
// must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) BatchUpsert(
	ctx context.Context,
	request *customers.BatchUpsertCustomerCustomAttributesRequest,
	opts ...option.RequestOption,
) (*v2.BatchUpsertCustomerCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/custom-attributes/bulk-upsert"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.BatchUpsertCustomerCustomAttributesResponse
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
