// This file was auto-generated by Fern from our API Definition.

package square

type LinkCustomerToGiftCardRequest struct {
	// The ID of the customer to link to the gift card.
	CustomerID string `json:"customer_id" url:"-"`
}

type UnlinkCustomerFromGiftCardRequest struct {
	// The ID of the customer to unlink from the gift card.
	CustomerID string `json:"customer_id" url:"-"`
}

type CreateGiftCardRequest struct {
	// A unique identifier for this request, used to ensure idempotency. For more information,
	// see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The ID of the [location](entity:Location) where the gift card should be registered for
	// reporting purposes. Gift cards can be redeemed at any of the seller's locations.
	LocationID string `json:"location_id" url:"-"`
	// The gift card to create. The `type` field is required for this request. The `gan_source`
	// and `gan` fields are included as follows:
	//
	// To direct Square to generate a 16-digit GAN, omit `gan_source` and `gan`.
	//
	// To provide a custom GAN, include `gan_source` and `gan`.
	// - For `gan_source`, specify `OTHER`.
	// - For `gan`, provide a custom GAN containing 8 to 20 alphanumeric characters. The GAN must be
	// unique for the seller and cannot start with the same bank identification number (BIN) as major
	// credit cards. Do not use GANs that are easy to guess (such as 12345678) because they greatly
	// increase the risk of fraud. It is the responsibility of the developer to ensure the security
	// of their custom GANs. For more information, see
	// [Custom GANs](https://developer.squareup.com/docs/gift-cards/using-gift-cards-api#custom-gans).
	//
	// To register an unused, physical gift card that the seller previously ordered from Square,
	// include `gan` and provide the GAN that is printed on the gift card.
	GiftCard *GiftCard `json:"gift_card,omitempty" url:"-"`
}

type GetGiftCardFromGanRequest struct {
	// The gift card account number (GAN) of the gift card to retrieve.
	// The maximum length of a GAN is 255 digits to account for third-party GANs that have been imported.
	// Square-issued gift cards have 16-digit GANs.
	Gan string `json:"gan" url:"-"`
}

type GetGiftCardFromNonceRequest struct {
	// The payment token of the gift card to retrieve. Payment tokens are generated by the
	// Web Payments SDK or In-App Payments SDK.
	Nonce string `json:"nonce" url:"-"`
}

type GiftCardsListRequest struct {
	// If a [type](entity:GiftCardType) is provided, the endpoint returns gift cards of the specified type.
	// Otherwise, the endpoint returns gift cards of all types.
	Type *string `json:"-" url:"type,omitempty"`
	// If a [state](entity:GiftCardStatus) is provided, the endpoint returns the gift cards in the specified state.
	// Otherwise, the endpoint returns the gift cards of all states.
	State *string `json:"-" url:"state,omitempty"`
	// If a limit is provided, the endpoint returns only the specified number of results per page.
	// The maximum value is 200. The default value is 30.
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// If a cursor is not provided, the endpoint returns the first page of the results.
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// If a customer ID is provided, the endpoint returns only the gift cards linked to the specified customer.
	CustomerID *string `json:"-" url:"customer_id,omitempty"`
}
