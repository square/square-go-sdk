// This file was auto-generated by Fern from our API Definition.

package customers

import (
	squaregosdk "github.com/fern-demo/square-go-sdk"
)

type CreateCustomerCardRequest struct {
	// A card nonce representing the credit card to link to the customer.
	//
	// Card nonces are generated by the Square payment form when customers enter
	// their card information. For more information, see
	// [Walkthrough: Integrate Square Payments in a Website](https://developer.squareup.com/docs/web-payments/take-card-payment).
	//
	// __NOTE:__ Card nonces generated by digital wallets (such as Apple Pay)
	// cannot be used to create a customer card.
	CardNonce string `json:"card_nonce" url:"-"`
	// Address information for the card on file.
	//
	// __NOTE:__ If a billing address is provided in the request, the
	// `CreateCustomerCardRequest.billing_address.postal_code` must match
	// the postal code encoded in the card nonce.
	BillingAddress *squaregosdk.Address `json:"billing_address,omitempty" url:"-"`
	// The full name printed on the credit card.
	CardholderName *string `json:"cardholder_name,omitempty" url:"-"`
	// An identifying token generated by [Payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer).
	// Verification tokens encapsulate customer device information and 3-D Secure
	// challenge results to indicate that Square has verified the buyer identity.
	VerificationToken *string `json:"verification_token,omitempty" url:"-"`
}
