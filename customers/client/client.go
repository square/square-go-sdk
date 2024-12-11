// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	cards "github.com/square/square-go-sdk/customers/cards"
	customattributedefinitions "github.com/square/square-go-sdk/customers/customattributedefinitions"
	customattributes "github.com/square/square-go-sdk/customers/customattributes"
	groups "github.com/square/square-go-sdk/customers/groups"
	segments "github.com/square/square-go-sdk/customers/segments"
	internal "github.com/square/square-go-sdk/internal"
	option "github.com/square/square-go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	CustomAttributeDefinitions *customattributedefinitions.Client
	Groups                     *groups.Client
	Segments                   *segments.Client
	Cards                      *cards.Client
	CustomAttributes           *customattributes.Client
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
		header:                     options.ToHeader(),
		CustomAttributeDefinitions: customattributedefinitions.NewClient(opts...),
		Groups:                     groups.NewClient(opts...),
		Segments:                   segments.NewClient(opts...),
		Cards:                      cards.NewClient(opts...),
		CustomAttributes:           customattributes.NewClient(opts...),
	}
}

// Lists customer profiles associated with a Square account.
//
// Under normal operating conditions, newly created or updated customer profiles become available
// for the listing operation in well under 30 seconds. Occasionally, propagation of the new or updated
// profiles can take closer to one minute or longer, especially during network incidents and outages.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.CustomersListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.Customer], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers"
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
	readPageResponse := func(response *squaregosdk.ListCustomersResponse) *internal.PageResponse[*string, *squaregosdk.Customer] {
		next := response.Cursor
		results := response.Customers
		return &internal.PageResponse[*string, *squaregosdk.Customer]{
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

// Creates a new customer for a business.
//
// You must provide at least one of the following values in your request to this
// endpoint:
//
// - `given_name`
// - `family_name`
// - `company_name`
// - `email_address`
// - `phone_number`
func (c *Client) Create(
	ctx context.Context,
	request *squaregosdk.CreateCustomerRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateCustomerResponse
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

// Creates multiple [customer profiles](entity:Customer) for a business.
//
// This endpoint takes a map of individual create requests and returns a map of responses.
//
// You must provide at least one of the following values in each create request:
//
// - `given_name`
// - `family_name`
// - `company_name`
// - `email_address`
// - `phone_number`
func (c *Client) BulkCreateCustomers(
	ctx context.Context,
	request *squaregosdk.BulkCreateCustomersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkCreateCustomersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/bulk-create"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BulkCreateCustomersResponse
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

// Deletes multiple customer profiles.
//
// The endpoint takes a list of customer IDs and returns a map of responses.
func (c *Client) BulkDeleteCustomers(
	ctx context.Context,
	request *squaregosdk.BulkDeleteCustomersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkDeleteCustomersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/bulk-delete"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BulkDeleteCustomersResponse
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

// Retrieves multiple customer profiles.
//
// This endpoint takes a list of customer IDs and returns a map of responses.
func (c *Client) BulkRetrieveCustomers(
	ctx context.Context,
	request *squaregosdk.BulkRetrieveCustomersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkRetrieveCustomersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/bulk-retrieve"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BulkRetrieveCustomersResponse
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

// Updates multiple customer profiles.
//
// This endpoint takes a map of individual update requests and returns a map of responses.
//
// You cannot use this endpoint to change cards on file. To make changes, use the [Cards API](api:Cards) or [Gift Cards API](api:GiftCards).
func (c *Client) BulkUpdateCustomers(
	ctx context.Context,
	request *squaregosdk.BulkUpdateCustomersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkUpdateCustomersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/bulk-update"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BulkUpdateCustomersResponse
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

// Searches the customer profiles associated with a Square account using one or more supported query filters.
//
// Calling `SearchCustomers` without any explicit query filter returns all
// customer profiles ordered alphabetically based on `given_name` and
// `family_name`.
//
// Under normal operating conditions, newly created or updated customer profiles become available
// for the search operation in well under 30 seconds. Occasionally, propagation of the new or updated
// profiles can take closer to one minute or longer, especially during network incidents and outages.
func (c *Client) Search(
	ctx context.Context,
	request *squaregosdk.SearchCustomersRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchCustomersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/customers/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SearchCustomersResponse
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

// Returns details for a single customer.
func (c *Client) Get(
	ctx context.Context,
	request *squaregosdk.CustomersGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v",
		request.CustomerID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetCustomerResponse
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

// Updates a customer profile. This endpoint supports sparse updates, so only new or changed fields are required in the request.
// To add or update a field, specify the new value. To remove a field, specify `null`.
//
// To update a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
//
// You cannot use this endpoint to change cards on file. To make changes, use the [Cards API](api:Cards) or [Gift Cards API](api:GiftCards).
func (c *Client) Update(
	ctx context.Context,
	request *squaregosdk.UpdateCustomerRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v",
		request.CustomerID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateCustomerResponse
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

// Deletes a customer profile from a business. This operation also unlinks any associated cards on file.
//
// To delete a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
func (c *Client) Delete(
	ctx context.Context,
	request *squaregosdk.CustomersDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/customers/%v",
		request.CustomerID,
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

	var response *squaregosdk.DeleteCustomerResponse
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
