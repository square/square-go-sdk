// This file was auto-generated by Fern from our API Definition.

package object

import (
	context "context"
	v2 "github.com/square/square-go-sdk/v2"
	catalog "github.com/square/square-go-sdk/v2/catalog"
	core "github.com/square/square-go-sdk/v2/core"
	internal "github.com/square/square-go-sdk/v2/internal"
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

// Creates a new or updates the specified [CatalogObject](entity:CatalogObject).
//
// To ensure consistency, only one update request is processed at a time per seller account.
// While one (batch or non-batch) update request is being processed, other (batched and non-batched)
// update requests are rejected with the `429` error code.
func (c *Client) Upsert(
	ctx context.Context,
	request *catalog.UpsertCatalogObjectRequest,
	opts ...option.RequestOption,
) (*v2.UpsertCatalogObjectResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := baseURL + "/v2/catalog/object"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *v2.UpsertCatalogObjectResponse
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

// Returns a single [CatalogItem](entity:CatalogItem) as a
// [CatalogObject](entity:CatalogObject) based on the provided ID. The returned
// object includes all of the relevant [CatalogItem](entity:CatalogItem)
// information including: [CatalogItemVariation](entity:CatalogItemVariation)
// children, references to its
// [CatalogModifierList](entity:CatalogModifierList) objects, and the ids of
// any [CatalogTax](entity:CatalogTax) objects that apply to it.
func (c *Client) Get(
	ctx context.Context,
	request *catalog.GetObjectRequest,
	opts ...option.RequestOption,
) (*v2.GetCatalogObjectResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/catalog/object/%v",
		request.ObjectID,
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

	var response *v2.GetCatalogObjectResponse
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

// Deletes a single [CatalogObject](entity:CatalogObject) based on the
// provided ID and returns the set of successfully deleted IDs in the response.
// Deletion is a cascading event such that all children of the targeted object
// are also deleted. For example, deleting a [CatalogItem](entity:CatalogItem)
// will also delete all of its
// [CatalogItemVariation](entity:CatalogItemVariation) children.
//
// To ensure consistency, only one delete request is processed at a time per seller account.
// While one (batch or non-batch) delete request is being processed, other (batched and non-batched)
// delete requests are rejected with the `429` error code.
func (c *Client) Delete(
	ctx context.Context,
	request *catalog.DeleteObjectRequest,
	opts ...option.RequestOption,
) (*v2.DeleteCatalogObjectResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareup.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/catalog/object/%v",
		request.ObjectID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *v2.DeleteCatalogObjectResponse
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
