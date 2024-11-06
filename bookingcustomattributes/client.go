// This file was auto-generated by Fern from our API Definition.

package bookingcustomattributes

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

// Get all bookings custom attribute definitions.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) ListBookingCustomAttributeDefinitions(
	ctx context.Context,
	request *squaregosdk.ListBookingCustomAttributeDefinitionsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListBookingCustomAttributeDefinitionsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/custom-attribute-definitions"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListBookingCustomAttributeDefinitionsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Creates a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) CreateBookingCustomAttributeDefinition(
	ctx context.Context,
	request *squaregosdk.CreateBookingCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/custom-attribute-definitions"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreateBookingCustomAttributeDefinitionResponse
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

// Retrieves a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) RetrieveBookingCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to retrieve. If the requesting application
	// is not the definition owner, you must use the qualified key.
	key string,
	request *squaregosdk.RetrieveBookingCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/custom-attribute-definitions/%v", key)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RetrieveBookingCustomAttributeDefinitionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) UpdateBookingCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to update.
	key string,
	request *squaregosdk.UpdateBookingCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/custom-attribute-definitions/%v", key)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdateBookingCustomAttributeDefinitionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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

// Deletes a bookings custom attribute definition.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) DeleteBookingCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to delete.
	key string,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteBookingCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/custom-attribute-definitions/%v", key)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.DeleteBookingCustomAttributeDefinitionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Bulk deletes bookings custom attributes.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) BulkDeleteBookingCustomAttributes(
	ctx context.Context,
	request *squaregosdk.BulkDeleteBookingCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkDeleteBookingCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/custom-attributes/bulk-delete"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkDeleteBookingCustomAttributesResponse
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

// Bulk upserts bookings custom attributes.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) BulkUpsertBookingCustomAttributes(
	ctx context.Context,
	request *squaregosdk.BulkUpsertBookingCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkUpsertBookingCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/custom-attributes/bulk-upsert"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkUpsertBookingCustomAttributesResponse
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

// Lists a booking's custom attributes.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) ListBookingCustomAttributes(
	ctx context.Context,
	// The ID of the target [booking](entity:Booking).
	bookingID string,
	request *squaregosdk.ListBookingCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListBookingCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/%v/custom-attributes", bookingID)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListBookingCustomAttributesResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Retrieves a bookings custom attribute.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) RetrieveBookingCustomAttribute(
	ctx context.Context,
	// The ID of the target [booking](entity:Booking).
	bookingID string,
	// The key of the custom attribute to retrieve. This key must match the `key` of a custom
	// attribute definition in the Square seller account. If the requesting application is not the
	// definition owner, you must use the qualified key.
	key string,
	request *squaregosdk.RetrieveBookingCustomAttributeRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveBookingCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/bookings/%v/custom-attributes/%v",
		bookingID,
		key,
	)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RetrieveBookingCustomAttributeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Upserts a bookings custom attribute.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) UpsertBookingCustomAttribute(
	ctx context.Context,
	// The ID of the target [booking](entity:Booking).
	bookingID string,
	// The key of the custom attribute to create or update. This key must match the `key` of a
	// custom attribute definition in the Square seller account. If the requesting application is not
	// the definition owner, you must use the qualified key.
	key string,
	request *squaregosdk.UpsertBookingCustomAttributeRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpsertBookingCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/bookings/%v/custom-attributes/%v",
		bookingID,
		key,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpsertBookingCustomAttributeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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

// Deletes a bookings custom attribute.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) DeleteBookingCustomAttribute(
	ctx context.Context,
	// The ID of the target [booking](entity:Booking).
	bookingID string,
	// The key of the custom attribute to delete. This key must match the `key` of a custom
	// attribute definition in the Square seller account. If the requesting application is not the
	// definition owner, you must use the qualified key.
	key string,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteBookingCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/bookings/%v/custom-attributes/%v",
		bookingID,
		key,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.DeleteBookingCustomAttributeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
