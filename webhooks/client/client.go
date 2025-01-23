// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/square/square-go-sdk/v2/core"
	internal "github.com/square/square-go-sdk/v2/internal"
	option "github.com/square/square-go-sdk/v2/option"
	eventtypes "github.com/square/square-go-sdk/v2/webhooks/eventtypes"
	subscriptions "github.com/square/square-go-sdk/v2/webhooks/subscriptions"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	EventTypes    *eventtypes.Client
	Subscriptions *subscriptions.Client
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
		header:        options.ToHeader(),
		EventTypes:    eventtypes.NewClient(opts...),
		Subscriptions: subscriptions.NewClient(opts...),
	}
}
