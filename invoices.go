// This file was auto-generated by Fern from our API Definition.

package square

type CreateInvoiceAttachmentRequest struct {
	Request interface{} `json:"request,omitempty" url:"-"`
}

type CancelInvoiceRequest struct {
	// The version of the [invoice](entity:Invoice) to cancel.
	// If you do not know the version, you can call
	// [GetInvoice](api-endpoint:Invoices-GetInvoice) or [ListInvoices](api-endpoint:Invoices-ListInvoices).
	Version int `json:"version" url:"-"`
}

type CreateInvoiceRequest struct {
	// The invoice to create.
	Invoice *Invoice `json:"invoice,omitempty" url:"-"`
	// A unique string that identifies the `CreateInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type InvoicesDeleteRequest struct {
	// The version of the [invoice](entity:Invoice) to delete.
	// If you do not know the version, you can call [GetInvoice](api-endpoint:Invoices-GetInvoice) or
	// [ListInvoices](api-endpoint:Invoices-ListInvoices).
	Version *int `json:"-" url:"version,omitempty"`
}

type InvoicesListRequest struct {
	// The ID of the location for which to list invoices.
	LocationID string `json:"-" url:"location_id"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The maximum number of invoices to return (200 is the maximum `limit`).
	// If not provided, the server uses a default limit of 100 invoices.
	Limit *int `json:"-" url:"limit,omitempty"`
}

type PublishInvoiceRequest struct {
	// The version of the [invoice](entity:Invoice) to publish.
	// This must match the current version of the invoice; otherwise, the request is rejected.
	Version int `json:"version" url:"-"`
	// A unique string that identifies the `PublishInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}

type SearchInvoicesRequest struct {
	// Describes the query criteria for searching invoices.
	Query *InvoiceQuery `json:"query,omitempty" url:"-"`
	// The maximum number of invoices to return (200 is the maximum `limit`).
	// If not provided, the server uses a default limit of 100 invoices.
	Limit *int `json:"limit,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
}

type UpdateInvoiceRequest struct {
	// The invoice fields to add, change, or clear. Fields can be cleared using
	// null values or the `remove` field (for individual payment requests or reminders).
	// The current invoice `version` is also required. For more information, including requirements,
	// limitations, and more examples, see [Update an Invoice](https://developer.squareup.com/docs/invoices-api/update-invoices).
	Invoice *Invoice `json:"invoice,omitempty" url:"-"`
	// A unique string that identifies the `UpdateInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The list of fields to clear. Although this field is currently supported, we
	// recommend using null values or the `remove` field when possible. For examples, see
	// [Update an Invoice](https://developer.squareup.com/docs/invoices-api/update-invoices).
	FieldsToClear []string `json:"fields_to_clear,omitempty" url:"-"`
}
