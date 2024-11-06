// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	images "github.com/square/square-go-sdk/catalog/images"
	object "github.com/square/square-go-sdk/catalog/object"
	core "github.com/square/square-go-sdk/core"
	option "github.com/square/square-go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Images *images.Client
	Object *object.Client
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
		Images: images.NewClient(opts...),
		Object: object.NewClient(opts...),
	}
}

// Deletes a set of [CatalogItem](entity:CatalogItem)s based on the
// provided list of target IDs and returns a set of successfully deleted IDs in
// the response. Deletion is a cascading event such that all children of the
// targeted object are also deleted. For example, deleting a CatalogItem will
// also delete all of its [CatalogItemVariation](entity:CatalogItemVariation)
// children.
//
// `BatchDeleteCatalogObjects` succeeds even if only a portion of the targeted
// IDs can be deleted. The response will only include IDs that were
// actually deleted.
//
// To ensure consistency, only one delete request is processed at a time per seller account.
// While one (batch or non-batch) delete request is being processed, other (batched and non-batched)
// delete requests are rejected with the `429` error code.
func (c *Client) BatchDelete(
	ctx context.Context,
	request *squaregosdk.BatchDeleteCatalogObjectsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BatchDeleteCatalogObjectsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/batch-delete"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BatchDeleteCatalogObjectsResponse
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

// Returns a set of objects based on the provided ID.
// Each [CatalogItem](entity:CatalogItem) returned in the set includes all of its
// child information including: all of its
// [CatalogItemVariation](entity:CatalogItemVariation) objects, references to
// its [CatalogModifierList](entity:CatalogModifierList) objects, and the ids of
// any [CatalogTax](entity:CatalogTax) objects that apply to it.
func (c *Client) BatchGet(
	ctx context.Context,
	request *squaregosdk.BatchGetCatalogObjectsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BatchGetCatalogObjectsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/batch-retrieve"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BatchGetCatalogObjectsResponse
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

// Creates or updates up to 10,000 target objects based on the provided
// list of objects. The target objects are grouped into batches and each batch is
// inserted/updated in an all-or-nothing manner. If an object within a batch is
// malformed in some way, or violates a database constraint, the entire batch
// containing that item will be disregarded. However, other batches in the same
// request may still succeed. Each batch may contain up to 1,000 objects, and
// batches will be processed in order as long as the total object count for the
// request (items, variations, modifier lists, discounts, and taxes) is no more
// than 10,000.
//
// To ensure consistency, only one update request is processed at a time per seller account.
// While one (batch or non-batch) update request is being processed, other (batched and non-batched)
// update requests are rejected with the `429` error code.
func (c *Client) BatchUpsert(
	ctx context.Context,
	request *squaregosdk.BatchUpsertCatalogObjectsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.BatchUpsertCatalogObjectsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/batch-upsert"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.BatchUpsertCatalogObjectsResponse
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

// Retrieves information about the Square Catalog API, such as batch size
// limits that can be used by the `BatchUpsertCatalogObjects` endpoint.
func (c *Client) Info(
	ctx context.Context,
	opts ...option.RequestOption,
) (*squaregosdk.CatalogInfoResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/info"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.CatalogInfoResponse
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

// Returns a list of all [CatalogObject](entity:CatalogObject)s of the specified types in the catalog.
//
// The `types` parameter is specified as a comma-separated list of the [CatalogObjectType](entity:CatalogObjectType) values,
// for example, "`ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, `CATEGORY`, `DISCOUNT`, `TAX`, `IMAGE`".
//
// **Important:** ListCatalog does not return deleted catalog items. To retrieve
// deleted catalog items, use [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects)
// and set the `include_deleted_objects` attribute value to `true`.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.CatalogListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.CatalogObject], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/list"

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
	readPageResponse := func(response *squaregosdk.ListCatalogResponse) *core.PageResponse[*string, *squaregosdk.CatalogObject] {
		next := response.Cursor
		results := response.Objects
		return &core.PageResponse[*string, *squaregosdk.CatalogObject]{
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

// Searches for [CatalogObject](entity:CatalogObject) of any type by matching supported search attribute values,
// excluding custom attribute values on items or item variations, against one or more of the specified query filters.
//
// This (`SearchCatalogObjects`) endpoint differs from the [SearchCatalogItems](api-endpoint:Catalog-SearchCatalogItems)
// endpoint in the following aspects:
//
// - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
// - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
// - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
// - The both endpoints have different call conventions, including the query filter formats.
func (c *Client) Search(
	ctx context.Context,
	request *squaregosdk.SearchCatalogObjectsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchCatalogObjectsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.SearchCatalogObjectsResponse
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

// Searches for catalog items or item variations by matching supported search attribute values, including
// custom attribute values, against one or more of the specified query filters.
//
// This (`SearchCatalogItems`) endpoint differs from the [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects)
// endpoint in the following aspects:
//
// - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
// - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
// - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
// - The both endpoints use different call conventions, including the query filter formats.
func (c *Client) SearchItems(
	ctx context.Context,
	request *squaregosdk.SearchCatalogItemsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchCatalogItemsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/search-catalog-items"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.SearchCatalogItemsResponse
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

// Updates the [CatalogModifierList](entity:CatalogModifierList) objects
// that apply to the targeted [CatalogItem](entity:CatalogItem) without having
// to perform an upsert on the entire item.
func (c *Client) UpdateItemModifierLists(
	ctx context.Context,
	request *squaregosdk.UpdateItemModifierListsRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateItemModifierListsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/update-item-modifier-lists"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdateItemModifierListsResponse
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

// Updates the [CatalogTax](entity:CatalogTax) objects that apply to the
// targeted [CatalogItem](entity:CatalogItem) without having to perform an
// upsert on the entire item.
func (c *Client) UpdateItemTaxes(
	ctx context.Context,
	request *squaregosdk.UpdateItemTaxesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateItemTaxesResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/catalog/update-item-taxes"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *squaregosdk.UpdateItemTaxesResponse
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
