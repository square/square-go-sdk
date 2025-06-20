// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	v2 "github.com/square/square-go-sdk/v2"
	core "github.com/square/square-go-sdk/v2/core"
	internal "github.com/square/square-go-sdk/v2/internal"
	accounts "github.com/square/square-go-sdk/v2/loyalty/accounts"
	programsclient "github.com/square/square-go-sdk/v2/loyalty/programs/client"
	rewards "github.com/square/square-go-sdk/v2/loyalty/rewards"
	option "github.com/square/square-go-sdk/v2/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Accounts *accounts.Client
	Programs *programsclient.Client
	Rewards  *rewards.Client
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
		header:   options.ToHeader(),
		Accounts: accounts.NewClient(opts...),
		Programs: programsclient.NewClient(opts...),
		Rewards:  rewards.NewClient(opts...),
	}
}

// Searches for loyalty events.
//
// A Square loyalty program maintains a ledger of events that occur during the lifetime of a
// buyer's loyalty account. Each change in the point balance
// (for example, points earned, points redeemed, and points expired) is
// recorded in the ledger. Using this endpoint, you can search the ledger for events.
//
// Search results are sorted by `created_at` in descending order.
func (c *Client) SearchEvents(
	ctx context.Context,
	request *v2.SearchLoyaltyEventsRequest,
	opts ...option.RequestOption,
) (*v2.SearchLoyaltyEventsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/loyalty/events/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.SearchLoyaltyEventsResponse
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
