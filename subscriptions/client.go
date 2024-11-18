// This file was auto-generated by Fern from our API Definition.

package subscriptions

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
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

// Enrolls a customer in a subscription.
//
// If you provide a card on file in the request, Square charges the card for
// the subscription. Otherwise, Square sends an invoice to the customer's email
// address. The subscription starts immediately, unless the request includes
// the optional `start_date`. Each individual subscription is associated with a particular location.
//
// For more information, see [Create a subscription](https://developer.squareup.com/docs/subscriptions-api/manage-subscriptions#create-a-subscription).
func (c *Client) Create(
	ctx context.Context,
	request *squaregosdk.CreateSubscriptionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/subscriptions"

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Schedules a plan variation change for all active subscriptions under a given plan
// variation. For more information, see [Swap Subscription Plan Variations](https://developer.squareup.com/docs/subscriptions-api/swap-plan-variations).
func (c *Client) BulkSwapPlan(
	ctx context.Context,
	request *squaregosdk.BulkSwapPlanRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkSwapPlanResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/subscriptions/bulk-swap-plan"

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.BulkSwapPlanResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Searches for subscriptions.
//
// Results are ordered chronologically by subscription creation date. If
// the request specifies more than one location ID,
// the endpoint orders the result
// by location ID, and then by creation date within each location. If no locations are given
// in the query, all locations are searched.
//
// You can also optionally specify `customer_ids` to search by customer.
// If left unset, all customers
// associated with the specified locations are returned.
// If the request specifies customer IDs, the endpoint orders results
// first by location, within location by customer ID, and within
// customer by subscription creation date.
func (c *Client) Search(
	ctx context.Context,
	request *squaregosdk.SearchSubscriptionsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchSubscriptionsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/subscriptions/search"

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SearchSubscriptionsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Retrieves a specific subscription.
func (c *Client) Get(
	ctx context.Context,
	// The ID of the subscription to retrieve.
	subscriptionID string,
	request *squaregosdk.SubscriptionsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v", subscriptionID)

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Updates a subscription by modifying or clearing `subscription` field values.
// To clear a field, set its value to `null`.
func (c *Client) Update(
	ctx context.Context,
	// The ID of the subscription to update.
	subscriptionID string,
	request *squaregosdk.UpdateSubscriptionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Deletes a scheduled action for a subscription.
func (c *Client) DeleteAction(
	ctx context.Context,
	// The ID of the subscription the targeted action is to act upon.
	subscriptionID string,
	// The ID of the targeted action to be deleted.
	actionID string,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteSubscriptionActionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/subscriptions/%v/actions/%v",
		subscriptionID,
		actionID,
	)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.DeleteSubscriptionActionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Changes the [billing anchor date](https://developer.squareup.com/docs/subscriptions-api/subscription-billing#billing-dates)
// for a subscription.
func (c *Client) ChangeBillingAnchorDate(
	ctx context.Context,
	// The ID of the subscription to update the billing anchor date.
	subscriptionID string,
	request *squaregosdk.ChangeBillingAnchorDateRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ChangeBillingAnchorDateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/billing-anchor", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.ChangeBillingAnchorDateResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Schedules a `CANCEL` action to cancel an active subscription. This
// sets the `canceled_date` field to the end of the active billing period. After this date,
// the subscription status changes from ACTIVE to CANCELED.
func (c *Client) Cancel(
	ctx context.Context,
	// The ID of the subscription to cancel.
	subscriptionID string,
	opts ...option.RequestOption,
) (*squaregosdk.CancelSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/cancel", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CancelSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
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

// Lists all [events](https://developer.squareup.com/docs/subscriptions-api/actions-events) for a specific subscription.
func (c *Client) ListEvents(
	ctx context.Context,
	// The ID of the subscription to retrieve the events for.
	subscriptionID string,
	request *squaregosdk.SubscriptionsListEventsRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.SubscriptionEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/events", subscriptionID)

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())

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
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *squaregosdk.ListSubscriptionEventsResponse) *internal.PageResponse[*string, *squaregosdk.SubscriptionEvent] {
		next := response.Cursor
		results := response.SubscriptionEvents
		return &internal.PageResponse[*string, *squaregosdk.SubscriptionEvent]{
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

// Schedules a `PAUSE` action to pause an active subscription.
func (c *Client) Pause(
	ctx context.Context,
	// The ID of the subscription to pause.
	subscriptionID string,
	request *squaregosdk.PauseSubscriptionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.PauseSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/pause", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.PauseSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Schedules a `RESUME` action to resume a paused or a deactivated subscription.
func (c *Client) Resume(
	ctx context.Context,
	// The ID of the subscription to resume.
	subscriptionID string,
	request *squaregosdk.ResumeSubscriptionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ResumeSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/resume", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.ResumeSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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

// Schedules a `SWAP_PLAN` action to swap a subscription plan variation in an existing subscription.
// For more information, see [Swap Subscription Plan Variations](https://developer.squareup.com/docs/subscriptions-api/swap-plan-variations).
func (c *Client) SwapPlan(
	ctx context.Context,
	// The ID of the subscription to swap the subscription plan for.
	subscriptionID string,
	request *squaregosdk.SwapPlanRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SwapPlanResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := internal.EncodeURL(baseURL+"/v2/subscriptions/%v/swap-plan", subscriptionID)

	headers := internal.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SwapPlanResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
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
