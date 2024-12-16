// This file was auto-generated by Fern from our API Definition.

package checkout

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type CreatePaymentLinkRequest struct {
	// A unique string that identifies this `CreatePaymentLinkRequest` request.
	// If you do not provide a unique string (or provide an empty string as the value),
	// the endpoint treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// A description of the payment link. You provide this optional description that is useful in your
	// application context. It is not used anywhere.
	Description *string `json:"description,omitempty" url:"-"`
	// Describes an ad hoc item and price for which to generate a quick pay checkout link.
	// For more information,
	// see [Quick Pay Checkout](https://developer.squareup.com/docs/checkout-api/quick-pay-checkout).
	QuickPay *squaregosdk.QuickPay `json:"quick_pay,omitempty" url:"-"`
	// Describes the `Order` for which to create a checkout link.
	// For more information,
	// see [Square Order Checkout](https://developer.squareup.com/docs/checkout-api/square-order-checkout).
	Order *squaregosdk.Order `json:"order,omitempty" url:"-"`
	// Describes optional fields to add to the resulting checkout page.
	// For more information,
	// see [Optional Checkout Configurations](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations).
	CheckoutOptions *squaregosdk.CheckoutOptions `json:"checkout_options,omitempty" url:"-"`
	// Describes fields to prepopulate in the resulting checkout page.
	// For more information, see [Prepopulate the shipping address](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations#prepopulate-the-shipping-address).
	PrePopulatedData *squaregosdk.PrePopulatedData `json:"pre_populated_data,omitempty" url:"-"`
	// A note for the payment. After processing the payment, Square adds this note to the resulting `Payment`.
	PaymentNote *string `json:"payment_note,omitempty" url:"-"`
}

type PaymentLinksDeleteRequest struct {
	// The ID of the payment link to delete.
	ID string `json:"-" url:"-"`
}

type PaymentLinksGetRequest struct {
	// The ID of link to retrieve.
	ID string `json:"-" url:"-"`
}

type PaymentLinksListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// If a cursor is not provided, the endpoint returns the first page of the results.
	// For more  information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// A limit on the number of results to return per page. The limit is advisory and
	// the implementation might return more or less results. If the supplied limit is negative, zero, or
	// greater than the maximum limit of 1000, it is ignored.
	//
	// Default value: `100`
	Limit *int `json:"-" url:"limit,omitempty"`
}

type UpdatePaymentLinkRequest struct {
	// The ID of the payment link to update.
	ID string `json:"-" url:"-"`
	// The `payment_link` object describing the updates to apply.
	// For more information, see [Update a payment link](https://developer.squareup.com/docs/checkout-api/manage-checkout#update-a-payment-link).
	PaymentLink *squaregosdk.PaymentLink `json:"payment_link,omitempty" url:"-"`
}
