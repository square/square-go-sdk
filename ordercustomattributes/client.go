// This file was auto-generated by Fern from our API Definition.

package ordercustomattributes

import (
	context "context"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	option "github.com/square/square-go-sdk/option"
	http "net/http"
	os "os"
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

// Lists the order-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.
//
// When all response pages are retrieved, the results include all custom attribute definitions
// that are visible to the requesting application, including those that are created by other
// applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that
// seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) ListOrderCustomAttributeDefinitions(
	ctx context.Context,
	request *squaregosdk.ListOrderCustomAttributeDefinitionsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListOrderCustomAttributeDefinitionsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders/custom-attribute-definitions"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListOrderCustomAttributeDefinitionsResponse
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

// Creates an order-related custom attribute definition. Use this endpoint to
// define a custom attribute that can be associated with orders.
//
// After creating a custom attribute definition, you can set the custom attribute for orders
// in the Square seller account.
func (c *Client) CreateOrderCustomAttributeDefinition(
	ctx context.Context,
	request *squaregosdk.CreateOrderCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders/custom-attribute-definitions"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CreateOrderCustomAttributeDefinitionResponse
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

// Retrieves an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// To retrieve a custom attribute definition created by another application, the `visibility`
// setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) RetrieveOrderCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to retrieve.
	key string,
	request *squaregosdk.RetrieveOrderCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/custom-attribute-definitions/%v", key)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RetrieveOrderCustomAttributeDefinitionResponse
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

// Updates an order-related custom attribute definition for a Square seller account.
//
// Only the definition owner can update a custom attribute definition. Note that sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.
func (c *Client) UpdateOrderCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to update.
	key string,
	request *squaregosdk.UpdateOrderCustomAttributeDefinitionRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/custom-attribute-definitions/%v", key)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdateOrderCustomAttributeDefinitionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
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

// Deletes an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
//
// Only the definition owner can delete a custom attribute definition.
func (c *Client) DeleteOrderCustomAttributeDefinition(
	ctx context.Context,
	// The key of the custom attribute definition to delete.
	key string,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteOrderCustomAttributeDefinitionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/custom-attribute-definitions/%v", key)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.DeleteOrderCustomAttributeDefinitionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
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

// Deletes order [custom attributes](entity:CustomAttribute) as a bulk operation.
//
// Use this endpoint to delete one or more custom attributes from one or more orders.
// A custom attribute is based on a custom attribute definition in a Square seller account. (To create a
// custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)
//
// This `BulkDeleteOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual delete
// requests and returns a map of individual delete responses. Each delete request has a unique ID
// and provides an order ID and custom attribute. Each delete response is returned with the ID
// of the corresponding request.
//
// To delete a custom attribute owned by another application, the `visibility` setting
// must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) BulkDeleteOrderCustomAttributes(
	ctx context.Context,
	request *squaregosdk.BulkDeleteOrderCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkDeleteOrderCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders/custom-attributes/bulk-delete"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkDeleteOrderCustomAttributesResponse
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

// Creates or updates order [custom attributes](entity:CustomAttribute) as a bulk operation.
//
// Use this endpoint to delete one or more custom attributes from one or more orders.
// A custom attribute is based on a custom attribute definition in a Square seller account. (To create a
// custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)
//
// This `BulkUpsertOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
// requests and returns a map of individual upsert responses. Each upsert request has a unique ID
// and provides an order ID and custom attribute. Each upsert response is returned with the ID
// of the corresponding request.
//
// To create or update a custom attribute owned by another application, the `visibility` setting
// must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) BulkUpsertOrderCustomAttributes(
	ctx context.Context,
	request *squaregosdk.BulkUpsertOrderCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BulkUpsertOrderCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders/custom-attributes/bulk-upsert"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BulkUpsertOrderCustomAttributesResponse
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

// Lists the [custom attributes](entity:CustomAttribute) associated with an order.
//
// You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
// in the same call.
//
// When all response pages are retrieved, the results include all custom attributes that are
// visible to the requesting application, including those that are owned by other applications
// and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) ListOrderCustomAttributes(
	ctx context.Context,
	// The ID of the target [order](entity:Order).
	orderID string,
	request *squaregosdk.ListOrderCustomAttributesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.ListOrderCustomAttributesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/%v/custom-attributes", orderID)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.ListOrderCustomAttributesResponse
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

// Retrieves a [custom attribute](entity:CustomAttribute) associated with an order.
//
// You can use the `with_definition` query parameter to also retrieve the custom attribute definition
// in the same call.
//
// To retrieve a custom attribute owned by another application, the `visibility` setting must be
// `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) RetrieveOrderCustomAttribute(
	ctx context.Context,
	// The ID of the target [order](entity:Order).
	orderID string,
	// The key of the custom attribute to retrieve. This key must match the key of an
	// existing custom attribute definition.
	customAttributeKey string,
	request *squaregosdk.RetrieveOrderCustomAttributeRequest,
	opts ...option.RequestOption,
) (*squaregosdk.RetrieveOrderCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/orders/%v/custom-attributes/%v",
		orderID,
		customAttributeKey,
	)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.RetrieveOrderCustomAttributeResponse
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

// Creates or updates a [custom attribute](entity:CustomAttribute) for an order.
//
// Use this endpoint to set the value of a custom attribute for a specific order.
// A custom attribute is based on a custom attribute definition in a Square seller account. (To create a
// custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)
//
// To create or update a custom attribute owned by another application, the `visibility` setting
// must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) UpsertOrderCustomAttribute(
	ctx context.Context,
	// The ID of the target [order](entity:Order).
	orderID string,
	// The key of the custom attribute to create or update. This key must match the key
	// of an existing custom attribute definition.
	customAttributeKey string,
	request *squaregosdk.UpsertOrderCustomAttributeRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpsertOrderCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/orders/%v/custom-attributes/%v",
		orderID,
		customAttributeKey,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpsertOrderCustomAttributeResponse
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

// Deletes a [custom attribute](entity:CustomAttribute) associated with a customer profile.
//
// To delete a custom attribute owned by another application, the `visibility` setting must be
// `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
// (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
func (c *Client) DeleteOrderCustomAttribute(
	ctx context.Context,
	// The ID of the target [order](entity:Order).
	orderID string,
	// The key of the custom attribute to delete. This key must match the key of an
	// existing custom attribute definition.
	customAttributeKey string,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteOrderCustomAttributeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v2/orders/%v/custom-attributes/%v",
		orderID,
		customAttributeKey,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.DeleteOrderCustomAttributeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
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
