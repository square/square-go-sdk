// This file was auto-generated by Fern from our API Definition.

package promotions

import (
	context "context"
	fmt "fmt"
	http "net/http"
	os "os"

	squaregosdk "github.com/fern-demo/square-go-sdk"
	core "github.com/fern-demo/square-go-sdk/core"
	programs "github.com/fern-demo/square-go-sdk/loyalty/programs"
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

// Lists the loyalty promotions associated with a [loyalty program]($m/LoyaltyProgram).
// Results are sorted by the `created_at` date in descending order (newest to oldest).
func (c *Client) List(
	ctx context.Context,
	// The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID,
	// call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the `main` keyword.
	programID string,
	request *programs.PromotionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.LoyaltyPromotion], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/loyalty/programs/%v/promotions", programID)

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
	readPageResponse := func(response *squaregosdk.ListLoyaltyPromotionsResponse) *core.PageResponse[*string, *squaregosdk.LoyaltyPromotion] {
		next := response.Cursor
		results := response.LoyaltyPromotions
		return &core.PageResponse[*string, *squaregosdk.LoyaltyPromotion]{
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

// Creates a loyalty promotion for a [loyalty program]($m/LoyaltyProgram). A loyalty promotion
// enables buyers to earn points in addition to those earned from the base loyalty program.
//
// This endpoint sets the loyalty promotion to the `ACTIVE` or `SCHEDULED` status, depending on the
// `available_time` setting. A loyalty program can have a maximum of 10 loyalty promotions with an
// `ACTIVE` or `SCHEDULED` status.
func (c *Client) Create(
	ctx context.Context,
	// The ID of the [loyalty program](entity:LoyaltyProgram) to associate with the promotion.
	// To get the program ID, call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram)
	// using the `main` keyword.
	programID string,
	request *programs.CreateLoyaltyPromotionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateLoyaltyPromotionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/loyalty/programs/%v/promotions", programID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreateLoyaltyPromotionResponse
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

// Retrieves a loyalty promotion.
func (c *Client) Get(
	ctx context.Context,
	// The ID of the [loyalty promotion](entity:LoyaltyPromotion) to retrieve.
	promotionID string,
	// The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID,
	// call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the `main` keyword.
	programID string,
	opts ...option.RequestOption,
) (*squaregosdk.GetLoyaltyPromotionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/loyalty/programs/%v/promotions/%v",
		programID,
		promotionID,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.GetLoyaltyPromotionResponse
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

// Cancels a loyalty promotion. Use this endpoint to cancel an `ACTIVE` promotion earlier than the
// end date, cancel an `ACTIVE` promotion when an end date is not specified, or cancel a `SCHEDULED` promotion.
// Because updating a promotion is not supported, you can also use this endpoint to cancel a promotion before
// you create a new one.
//
// This endpoint sets the loyalty promotion to the `CANCELED` state
func (c *Client) Cancel(
	ctx context.Context,
	// The ID of the [loyalty promotion](entity:LoyaltyPromotion) to cancel. You can cancel a
	// promotion that has an `ACTIVE` or `SCHEDULED` status.
	promotionID string,
	// The ID of the base [loyalty program](entity:LoyaltyProgram).
	programID string,
	opts ...option.RequestOption,
) (*squaregosdk.CancelLoyaltyPromotionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/loyalty/programs/%v/promotions/%v/cancel",
		programID,
		promotionID,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CancelLoyaltyPromotionResponse
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
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
