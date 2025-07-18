// This file was auto-generated by Fern from our API Definition.

package accounts

import (
	context "context"
	v2 "github.com/square/square-go-sdk/v2"
	core "github.com/square/square-go-sdk/v2/core"
	internal "github.com/square/square-go-sdk/v2/internal"
	loyalty "github.com/square/square-go-sdk/v2/loyalty"
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

// Creates a loyalty account. To create a loyalty account, you must provide the `program_id` and a `mapping` with the `phone_number` of the buyer.
func (c *Client) Create(
	ctx context.Context,
	request *loyalty.CreateLoyaltyAccountRequest,
	opts ...option.RequestOption,
) (*v2.CreateLoyaltyAccountResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/loyalty/accounts"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.CreateLoyaltyAccountResponse
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

// Searches for loyalty accounts in a loyalty program.
//
// You can search for a loyalty account using the phone number or customer ID associated with the account. To return all loyalty accounts, specify an empty `query` object or omit it entirely.
//
// Search results are sorted by `created_at` in ascending order.
func (c *Client) Search(
	ctx context.Context,
	request *loyalty.SearchLoyaltyAccountsRequest,
	opts ...option.RequestOption,
) (*v2.SearchLoyaltyAccountsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/loyalty/accounts/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.SearchLoyaltyAccountsResponse
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

// Retrieves a loyalty account.
func (c *Client) Get(
	ctx context.Context,
	request *loyalty.GetAccountsRequest,
	opts ...option.RequestOption,
) (*v2.GetLoyaltyAccountResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/accounts/%v",
		request.AccountID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.GetLoyaltyAccountResponse
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

// Adds points earned from a purchase to a [loyalty account](entity:LoyaltyAccount).
//
// - If you are using the Orders API to manage orders, provide the `order_id`. Square reads the order
// to compute the points earned from both the base loyalty program and an associated
// [loyalty promotion](entity:LoyaltyPromotion). For purchases that qualify for multiple accrual
// rules, Square computes points based on the accrual rule that grants the most points.
// For purchases that qualify for multiple promotions, Square computes points based on the most
// recently created promotion. A purchase must first qualify for program points to be eligible for promotion points.
//
// - If you are not using the Orders API to manage orders, provide `points` with the number of points to add.
// You must first perform a client-side computation of the points earned from the loyalty program and
// loyalty promotion. For spend-based and visit-based programs, you can call [CalculateLoyaltyPoints](api-endpoint:Loyalty-CalculateLoyaltyPoints)
// to compute the points earned from the base loyalty program. For information about computing points earned from a loyalty promotion, see
// [Calculating promotion points](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#calculate-promotion-points).
func (c *Client) AccumulatePoints(
	ctx context.Context,
	request *loyalty.AccumulateLoyaltyPointsRequest,
	opts ...option.RequestOption,
) (*v2.AccumulateLoyaltyPointsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/accounts/%v/accumulate",
		request.AccountID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.AccumulateLoyaltyPointsResponse
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

// Adds points to or subtracts points from a buyer's account.
//
// Use this endpoint only when you need to manually adjust points. Otherwise, in your application flow, you call
// [AccumulateLoyaltyPoints](api-endpoint:Loyalty-AccumulateLoyaltyPoints)
// to add points when a buyer pays for the purchase.
func (c *Client) Adjust(
	ctx context.Context,
	request *loyalty.AdjustLoyaltyPointsRequest,
	opts ...option.RequestOption,
) (*v2.AdjustLoyaltyPointsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/accounts/%v/adjust",
		request.AccountID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.AdjustLoyaltyPointsResponse
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
