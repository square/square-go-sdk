// This file was auto-generated by Fern from our API Definition.

package rewards

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	loyalty "github.com/square/square-go-sdk/loyalty"
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

// Creates a loyalty reward. In the process, the endpoint does following:
//
//   - Uses the `reward_tier_id` in the request to determine the number of points
//     to lock for this reward.
//   - If the request includes `order_id`, it adds the reward and related discount to the order.
//
// After a reward is created, the points are locked and
// not available for the buyer to redeem another reward.
func (c *Client) Create(
	ctx context.Context,
	request *loyalty.CreateLoyaltyRewardRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateLoyaltyRewardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/loyalty/rewards"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateLoyaltyRewardResponse
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

// Searches for loyalty rewards. This endpoint accepts a request with no query filters and returns results for all loyalty accounts.
// If you include a `query` object, `loyalty_account_id` is required and `status` is optional.
//
// If you know a reward ID, use the
// [RetrieveLoyaltyReward](api-endpoint:Loyalty-RetrieveLoyaltyReward) endpoint.
//
// Search results are sorted by `updated_at` in descending order.
func (c *Client) Search(
	ctx context.Context,
	request *loyalty.SearchLoyaltyRewardsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchLoyaltyRewardsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/loyalty/rewards/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SearchLoyaltyRewardsResponse
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

// Retrieves a loyalty reward.
func (c *Client) Get(
	ctx context.Context,
	request *loyalty.RewardsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetLoyaltyRewardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/rewards/%v",
		request.RewardID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetLoyaltyRewardResponse
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

// Deletes a loyalty reward by doing the following:
//
//   - Returns the loyalty points back to the loyalty account.
//   - If an order ID was specified when the reward was created
//     (see [CreateLoyaltyReward](api-endpoint:Loyalty-CreateLoyaltyReward)),
//     it updates the order by removing the reward and related
//     discounts.
//
// You cannot delete a reward that has reached the terminal state (REDEEMED).
func (c *Client) Delete(
	ctx context.Context,
	request *loyalty.RewardsDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteLoyaltyRewardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/rewards/%v",
		request.RewardID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.DeleteLoyaltyRewardResponse
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

// Redeems a loyalty reward.
//
// The endpoint sets the reward to the `REDEEMED` terminal state.
//
// If you are using your own order processing system (not using the
// Orders API), you call this endpoint after the buyer paid for the
// purchase.
//
// After the reward reaches the terminal state, it cannot be deleted.
// In other words, points used for the reward cannot be returned
// to the account.
func (c *Client) Redeem(
	ctx context.Context,
	request *loyalty.RedeemLoyaltyRewardRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RedeemLoyaltyRewardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/loyalty/rewards/%v/redeem",
		request.RewardID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.RedeemLoyaltyRewardResponse
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
