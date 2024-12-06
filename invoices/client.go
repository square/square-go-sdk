// This file was auto-generated by Fern from our API Definition.

package invoices

import (
	context "context"
	fmt "fmt"
	squaregosdk "github.com/square/square-go-sdk"
	core "github.com/square/square-go-sdk/core"
	internal "github.com/square/square-go-sdk/internal"
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

// Returns a list of invoices for a given location. The response
// is paginated. If truncated, the response includes a `cursor` that you
// use in a subsequent request to retrieve the next set of invoices.
func (c *Client) List(
	ctx context.Context,
	request *squaregosdk.InvoicesListRequest,
	opts ...option.RequestOption,
) (*core.Page[*squaregosdk.Invoice], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := baseURL + "/v2/invoices"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	prepareCall := func(pageRequest *internal.PageRequest[*string]) *internal.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &internal.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *squaregosdk.ListInvoicesResponse) *internal.PageResponse[*string, *squaregosdk.Invoice] {
		next := response.Cursor
		results := response.Invoices
		return &internal.PageResponse[*string, *squaregosdk.Invoice]{
			Next:    next,
			Results: results,
		}
	}
	pager := internal.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Creates a draft [invoice]($m/Invoice)
// for an order created using the Orders API.
//
// A draft invoice remains in your account and no action is taken.
// You must publish the invoice before Square can process it (send it to the customer's email address or charge the customer’s card on file).
func (c *Client) Create(
	ctx context.Context,
	request *squaregosdk.CreateInvoiceRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := baseURL + "/v2/invoices"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CreateInvoiceResponse
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

// Searches for invoices from a location specified in
// the filter. You can optionally specify customers in the filter for whom to
// retrieve invoices. In the current implementation, you can only specify one location and
// optionally one customer.
//
// The response is paginated. If truncated, the response includes a `cursor`
// that you use in a subsequent request to retrieve the next set of invoices.
func (c *Client) Search(
	ctx context.Context,
	request *squaregosdk.SearchInvoicesRequest,
	opts ...option.RequestOption,
) (*squaregosdk.SearchInvoicesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := baseURL + "/v2/invoices/search"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.SearchInvoicesResponse
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

// Retrieves an invoice by invoice ID.
func (c *Client) Get(
	ctx context.Context,
	request *squaregosdk.InvoicesGetRequest,
	opts ...option.RequestOption,
) (*squaregosdk.GetInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v",
		request.InvoiceID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.GetInvoiceResponse
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

// Updates an invoice. This endpoint supports sparse updates, so you only need
// to specify the fields you want to change along with the required `version` field.
// Some restrictions apply to updating invoices. For example, you cannot change the
// `order_id` or `location_id` field.
func (c *Client) Update(
	ctx context.Context,
	request *squaregosdk.UpdateInvoiceRequest,
	opts ...option.RequestOption,
) (*squaregosdk.UpdateInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v",
		request.InvoiceID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.UpdateInvoiceResponse
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

// Deletes the specified invoice. When an invoice is deleted, the
// associated order status changes to CANCELED. You can only delete a draft
// invoice (you cannot delete a published invoice, including one that is scheduled for processing).
func (c *Client) Delete(
	ctx context.Context,
	request *squaregosdk.InvoicesDeleteRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v",
		request.InvoiceID,
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

	var response *squaregosdk.DeleteInvoiceResponse
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

// Uploads a file and attaches it to an invoice. This endpoint accepts HTTP multipart/form-data file uploads
// with a JSON `request` part and a `file` part. The `file` part must be a `readable stream` that contains a file
// in a supported format: GIF, JPEG, PNG, TIFF, BMP, or PDF.
//
// Invoices can have up to 10 attachments with a total file size of 25 MB. Attachments can be added only to invoices
// in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.
func (c *Client) CreateInvoiceAttachment(
	ctx context.Context,
	request *squaregosdk.CreateInvoiceAttachmentRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CreateInvoiceAttachmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v/attachments",
		request.InvoiceID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	writer := internal.NewMultipartWriter()
	if request.ImageFile != nil {
		if err := writer.WriteFile("image_file", request.ImageFile, internal.WithDefaultContentType("image/jpeg")); err != nil {
			return nil, err
		}
	}
	if request.Request != nil {
		if err := writer.WriteJSON("request", request.Request, internal.WithDefaultContentType("application/json; charset=utf-8")); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers.Set("Content-Type", writer.ContentType())

	var response *squaregosdk.CreateInvoiceAttachmentResponse
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
			Request:         writer.Buffer(),
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Removes an attachment from an invoice and permanently deletes the file. Attachments can be removed only
// from invoices in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.
func (c *Client) DeleteInvoiceAttachment(
	ctx context.Context,
	request *squaregosdk.DeleteInvoiceAttachmentRequest,
	opts ...option.RequestOption,
) (*squaregosdk.DeleteInvoiceAttachmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v/attachments/%v",
		request.InvoiceID,
		request.AttachmentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *squaregosdk.DeleteInvoiceAttachmentResponse
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

// Cancels an invoice. The seller cannot collect payments for
// the canceled invoice.
//
// You cannot cancel an invoice in the `DRAFT` state or in a terminal state: `PAID`, `REFUNDED`, `CANCELED`, or `FAILED`.
func (c *Client) Cancel(
	ctx context.Context,
	request *squaregosdk.CancelInvoiceRequest,
	opts ...option.RequestOption,
) (*squaregosdk.CancelInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v/cancel",
		request.InvoiceID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.CancelInvoiceResponse
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

// Publishes the specified draft invoice.
//
// After an invoice is published, Square
// follows up based on the invoice configuration. For example, Square
// sends the invoice to the customer's email address, charges the customer's card on file, or does
// nothing. Square also makes the invoice available on a Square-hosted invoice page.
//
// The invoice `status` also changes from `DRAFT` to a status
// based on the invoice configuration. For example, the status changes to `UNPAID` if
// Square emails the invoice or `PARTIALLY_PAID` if Square charges a card on file for a portion of the
// invoice amount.
//
// In addition to the required `ORDERS_WRITE` and `INVOICES_WRITE` permissions, `CUSTOMERS_READ`
// and `PAYMENTS_WRITE` are required when publishing invoices configured for card-on-file payments.
func (c *Client) Publisb(
	ctx context.Context,
	request *squaregosdk.PublishInvoiceRequest,
	opts ...option.RequestOption,
) (*squaregosdk.PublishInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://connect.squareupsandbox.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v2/invoices/%v/publish",
		request.InvoiceID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *squaregosdk.PublishInvoiceResponse
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
