// This file was auto-generated by Fern from our API Definition.

package customattributedefinitions

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
	merchants "github.com/square/square-go-sdk/merchants"
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

// Retrieves a merchant-related [custom attribute definition]($m/CustomAttributeDefinition) from a Square seller account.
// To retrieve a custom attribute definition created by another application, the `visibility`
// setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) Get(
	ctx context.Context,
	request *merchants.CustomAttributeDefinitionsGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveMerchantCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/merchants/custom-attribute-definitions/%v",
		request.Key,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.RetrieveMerchantCustomAttributeDefinitionResponse
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

// Updates a merchant-related [custom attribute definition]($m/CustomAttributeDefinition) for a Square seller account.
// Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the
// `schema` for a `Selection` data type.
// Only the definition owner can update a custom attribute definition.
func (c *Client) Update(
	ctx context.Context,
	request *merchants.UpdateMerchantCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateMerchantCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/merchants/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateMerchantCustomAttributeDefinitionResponse
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

// Deletes a merchant-related [custom attribute definition]($m/CustomAttributeDefinition) from a Square seller account.
// Deleting a custom attribute definition also deletes the corresponding custom attribute from
// the merchant.
// Only the definition owner can delete a custom attribute definition.
func (c *Client) Delete(
	ctx context.Context,
	request *merchants.CustomAttributeDefinitionsDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteMerchantCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/merchants/custom-attribute-definitions/%v",
		request.Key,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.DeleteMerchantCustomAttributeDefinitionResponse
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
