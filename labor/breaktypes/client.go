// This file was auto-generated by Fern from our API Definition.

package breaktypes

import (
	context "context"
	fmt "fmt"
	v40 "github.com/square/square-go-sdk/v40"
	core "github.com/square/square-go-sdk/v40/core"
	internal "github.com/square/square-go-sdk/v40/internal"
	labor "github.com/square/square-go-sdk/v40/labor"
	option "github.com/square/square-go-sdk/v40/option"
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

// Returns a paginated list of `BreakType` instances for a business.
func (c *Client) List(
	ctx context.Context,
	request *labor.BreakTypesListRequest,
	opts ...option.RequestOption,
) (*core.Page[*v40.BreakType], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/labor/break-types"
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
	readPageResponse := func(response *v40.ListBreakTypesResponse) *internal.PageResponse[*string, *v40.BreakType] {
		next := response.Cursor
		results := response.BreakTypes
		return &internal.PageResponse[*string, *v40.BreakType]{
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

// Creates a new `BreakType`.
//
// A `BreakType` is a template for creating `Break` objects.
// You must provide the following values in your request to this
// endpoint:
//
// - `location_id`
// - `break_name`
// - `expected_duration`
// - `is_paid`
//
// You can only have three `BreakType` instances per location. If you attempt to add a fourth
// `BreakType` for a location, an `INVALID_REQUEST_ERROR` "Exceeded limit of 3 breaks per location."
// is returned.
func (c *Client) Create(
	ctx context.Context,
	request *labor.CreateBreakTypeRequest,
	opts ...option.RequestOption,
) (*v40.CreateBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/labor/break-types"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v40.CreateBreakTypeResponse
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

// Returns a single `BreakType` specified by `id`.
func (c *Client) Get(
	ctx context.Context,
	request *labor.BreakTypesGetRequest,
	opts ...option.RequestOption,
) (*v40.GetBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/labor/break-types/%v",
		request.ID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v40.GetBreakTypeResponse
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

// Updates an existing `BreakType`.
func (c *Client) Update(
	ctx context.Context,
	request *labor.UpdateBreakTypeRequest,
	opts ...option.RequestOption,
) (*v40.UpdateBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/labor/break-types/%v",
		request.ID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v40.UpdateBreakTypeResponse
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

// Deletes an existing `BreakType`.
//
// A `BreakType` can be deleted even if it is referenced from a `Shift`.
func (c *Client) Delete(
	ctx context.Context,
	request *labor.BreakTypesDeleteRequest,
	opts ...option.RequestOption,
) (*v40.DeleteBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/labor/break-types/%v",
		request.ID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v40.DeleteBreakTypeResponse
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
