// This file was auto-generated by Fern from our API Definition.

package shifts

import (
	context "context"
	fmt "fmt"
	http "net/http"
	os "os"

	squaregosdk "github.com/fern-demo/square-go-sdk"
	cashdrawers "github.com/fern-demo/square-go-sdk/cashdrawers"
	core "github.com/fern-demo/square-go-sdk/core"
	option "github.com/fern-demo/square-go-sdk/option"
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

// Provides the details for all of the cash drawer shifts for a location
// in a date range.
func (c *Client) List(
	ctx context.Context,
	request *cashdrawers.ShiftsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CashDrawerShiftSummary], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/cash-drawers/shifts"

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
	readPageResponse := func(response *squaregosdk.ListCashDrawerShiftsResponse) *core.PageResponse[*string, *squaregosdk.CashDrawerShiftSummary] {
		next := response.Cursor
		results := response.CashDrawerShifts
		return &core.PageResponse[*string, *squaregosdk.CashDrawerShiftSummary]{
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

// Provides the summary details for a single cash drawer shift. See
// [ListCashDrawerShiftEvents]($e/CashDrawers/ListCashDrawerShiftEvents) for a list of cash drawer shift events.
func (c *Client) Get(
	ctx context.Context,
	// The shift ID.
	shiftID string,
	request *cashdrawers.ShiftsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetCashDrawerShiftResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/cash-drawers/shifts/%v", shiftID)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetCashDrawerShiftResponse
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

// Provides a paginated list of events for a single cash drawer shift.
func (c *Client) ListEvents(
	ctx context.Context,
	// The shift ID.
	shiftID string,
	request *cashdrawers.ShiftsListEventsRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CashDrawerShiftEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/cash-drawers/shifts/%v/events", shiftID)

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
	readPageResponse := func(response *squaregosdk.ListCashDrawerShiftEventsResponse) *core.PageResponse[*string, *squaregosdk.CashDrawerShiftEvent] {
		next := response.Cursor
		results := response.CashDrawerShiftEvents
		return &core.PageResponse[*string, *squaregosdk.CashDrawerShiftEvent]{
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
