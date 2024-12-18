// This file was auto-generated by Fern from our API Definition.

package wagesetting

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	option "github.com/square/square-go-sdk/option"
	teammembers "github.com/square/square-go-sdk/teammembers"
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

// Retrieves a `WageSetting` object for a team member specified
// by `TeamMember.id`. For more information, see
// [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrievewagesetting).
//
// Square recommends using [RetrieveTeamMember]($e/Team/RetrieveTeamMember) or [SearchTeamMembers]($e/Team/SearchTeamMembers)
// to get this information directly from the `TeamMember.wage_setting` field.
func (c *Client) Get(
	ctx context.Context,
	request *teammembers.WageSettingGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetWageSettingResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/team-members/%v/wage-setting",
		request.TeamMemberID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetWageSettingResponse
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

// Creates or updates a `WageSetting` object. The object is created if a
// `WageSetting` with the specified `team_member_id` doesn't exist. Otherwise,
// it fully replaces the `WageSetting` object for the team member.
// The `WageSetting` is returned on a successful update. For more information, see
// [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#create-or-update-a-wage-setting).
//
// Square recommends using [CreateTeamMember]($e/Team/CreateTeamMember) or [UpdateTeamMember]($e/Team/UpdateTeamMember)
// to manage the `TeamMember.wage_setting` field directly.
func (c *Client) Update(
	ctx context.Context,
	request *teammembers.UpdateWageSettingRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateWageSettingResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/team-members/%v/wage-setting",
		request.TeamMemberID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateWageSettingResponse
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
