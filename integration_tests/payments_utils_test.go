//go:build integration

package integration

import (
	square "github.com/square/square-go-sdk/v2"
)

type TestPaymentRequest struct {
	Amount         *int64
	IdempotencyKey *string
}

// Create a [square.CreatePaymentRequest] object with the given amount and
// idempotency key.
func newTestPaymentRequest(
	request TestPaymentRequest,
) *square.CreatePaymentRequest {
	defaultAmount := int64(1001)
	if request.Amount == nil {
		request.Amount = &defaultAmount
	}

	defaultIdempotencyKey := newTestUUID()
	if request.IdempotencyKey == nil {
		request.IdempotencyKey = &defaultIdempotencyKey
	}

	return &square.CreatePaymentRequest{
		IdempotencyKey: *request.IdempotencyKey,
		SourceID:       SourceID,
		AmountMoney:    newTestMoney(*request.Amount),
		Autocomplete:   square.Bool(true),
	}
}
