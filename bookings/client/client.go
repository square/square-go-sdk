// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	teammemberprofiles "github.com/square/square-go-sdk/bookings/teammemberprofiles"
	core "github.com/square/square-go-sdk/core"
	option "github.com/square/square-go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	TeamMemberProfiles *teammemberprofiles.Client
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
		header:             options.ToHeader(),
		TeamMemberProfiles: teammemberprofiles.NewClient(opts...),
	}
}

// Retrieve a collection of bookings.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.BookingsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.Booking], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*string]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &core.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *squaregosdk.ListBookingsResponse) *core.PageResponse[*string, *squaregosdk.Booking] {
		next := response.Cursor
		results := response.Bookings
		return &core.PageResponse[*string, *squaregosdk.Booking]{
			Next:    next,
			Results: results,
		}
	}
	pager := core.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Creates a booking.
//
// The required input must include the following:
//
// - `Booking.location_id`
// - `Booking.start_at`
// - `Booking.AppointmentSegment.team_member_id`
// - `Booking.AppointmentSegment.service_variation_id`
// - `Booking.AppointmentSegment.service_variation_version`
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) Create(
	ctx context.Context,
	request *squaregosdk.CreateBookingRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateBookingResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreateBookingResponse
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

// Searches for availabilities for booking.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) SearchAvailability(
	ctx context.Context,
	request *squaregosdk.SearchAvailabilityRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchAvailabilityResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/availability/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.SearchAvailabilityResponse
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

// Bulk-Retrieves a list of bookings by booking IDs.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) BulkRetrieveBookings(
	ctx context.Context,
	request *squaregosdk.BulkRetrieveBookingsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkRetrieveBookingsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/bulk-retrieve"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkRetrieveBookingsResponse
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

// Retrieves a seller's booking profile.
func (c *Client) GetBusinessProfile(
	ctx context.Context,
	opts ...option.RequestOption,
) (*squaregosdk.GetBusinessBookingProfileResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/business-booking-profile"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetBusinessBookingProfileResponse
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

// Lists location booking profiles of a seller.
func (c *Client) ListLocationBookingProfiles(
	ctx context.Context,
	request *squaregosdk.ListLocationBookingProfilesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListLocationBookingProfilesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/location-booking-profiles"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListLocationBookingProfilesResponse
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

// Retrieves a seller's location booking profile.
func (c *Client) RetrieveLocationBookingProfile(
	ctx context.Context,
	// The ID of the location to retrieve the booking profile.
	locationID string,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveLocationBookingProfileResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/location-booking-profiles/%v", locationID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RetrieveLocationBookingProfileResponse
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

// Retrieves one or more team members' booking profiles.
func (c *Client) BulkRetrieveTeamMemberBookingProfiles(
	ctx context.Context,
	request *squaregosdk.BulkRetrieveTeamMemberBookingProfilesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkRetrieveTeamMemberBookingProfilesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/bookings/team-member-booking-profiles/bulk-retrieve"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkRetrieveTeamMemberBookingProfilesResponse
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

// Retrieves a booking.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
func (c *Client) Get(
	ctx context.Context,
	// The ID of the [Booking](entity:Booking) object representing the to-be-retrieved booking.
	bookingID string,
	opts ...option.RequestOption,
) (*squaregosdk.GetBookingResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/%v", bookingID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetBookingResponse
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

// Updates a booking.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) Update(
	ctx context.Context,
	// The ID of the [Booking](entity:Booking) object representing the to-be-updated booking.
	bookingID string,
	request *squaregosdk.UpdateBookingRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateBookingResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/%v", bookingID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdateBookingResponse
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

// Cancels an existing booking.
//
// To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
// To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.
//
// For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to _Appointments Plus_
// or _Appointments Premium_.
func (c *Client) Cancel(
	ctx context.Context,
	// The ID of the [Booking](entity:Booking) object representing the to-be-cancelled booking.
	bookingID string,
	request *squaregosdk.CancelBookingRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CancelBookingResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/bookings/%v/cancel", bookingID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CancelBookingResponse
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
