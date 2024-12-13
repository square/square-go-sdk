// This file was auto-generated by Fern from our API Definition.

package customattributedefinitions

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	bookings "github.com/square/square-go-sdk/bookings"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	option "github.com/square/square-go-sdk/option"
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

// Get all bookings custom attribute definitions.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) List(
	ctx context.Context,
	request *bookings.CustomAttributeDefinitionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CustomAttributeDefinition], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/bookings/custom-attribute-definitions"
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
	readPageResponse := func(response *squaregosdk.ListBookingCustomAttributeDefinitionsResponse) *internal.PageResponse[*string, *squaregosdk.CustomAttributeDefinition] {
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

// Creates a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
// or *Appointments Premium*.
func (c *Client) Create(
	ctx context.Context,
	request *bookings.CreateBookingCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/bookings/custom-attribute-definitions"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateBookingCustomAttributeDefinitionResponse
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

// Retrieves a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) Get(
	ctx context.Context,
	request *bookings.CustomAttributeDefinitionsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/bookings/custom-attribute-definitions/%v",
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

	var response *squaregosdk.RetrieveBookingCustomAttributeDefinitionResponse
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

// Updates a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
// or *Appointments Premium*.
func (c *Client) Update(
	ctx context.Context,
	request *bookings.UpdateBookingCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/bookings/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateBookingCustomAttributeDefinitionResponse
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

// Deletes a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
// or *Appointments Premium*.
func (c *Client) Delete(
	ctx context.Context,
	request *bookings.CustomAttributeDefinitionsDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/bookings/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.DeleteBookingCustomAttributeDefinitionResponse
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
