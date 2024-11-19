//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/square/square-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Payments API integration tests.
func TestPaymentsAPI(t *testing.T) {
	t.Run("create and retrieve a payment", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		// Create a valid payment.
		createPaymentResp, err := squareClient.Payments.Create(
			context.Background(),
			newTestPaymentRequest(
				TestPaymentRequest{},
			),
		)
		require.NoError(t, err)
		require.NotNil(t, createPaymentResp.Payment)

		// Fetch the payment to verify it was created.
		getPaymentResp, err := squareClient.Payments.Get(
			context.Background(),
			*createPaymentResp.Payment.ID,
		)

		require.NoError(t, err)
		assert.NotNil(t, getPaymentResp.Payment)
		assert.Equal(t, *createPaymentResp.Payment.ID, *getPaymentResp.Payment.ID)
		assert.Equal(t, getPaymentResp.Payment.AmountMoney.Amount, square.Int64(1001))
	})

	t.Run("payment fails when reusing an idempotency key", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		idempotencyKey := newTestUUID()

		// Create a valid payment with a unique idempotency key.
		_, err := squareClient.Payments.Create(
			context.Background(),
			newTestPaymentRequest(
				TestPaymentRequest{
					Amount:         square.Int64(1001),
					IdempotencyKey: &idempotencyKey,
				},
			),
		)
		require.NoError(t, err)

		// Attempt to create a payment with the same idempotency key, but
		// a different monetary amount (which should fail).
		_, err = squareClient.Payments.Create(
			context.Background(),
			newTestPaymentRequest(
				TestPaymentRequest{
					Amount:         square.Int64(1002),
					IdempotencyKey: &idempotencyKey,
				},
			),
		)
		require.Error(t, err)
	})

	t.Run("refund a payment", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		// Create a valid payment.
		createPaymentResp, err := squareClient.Payments.Create(
			context.Background(),
			newTestPaymentRequest(
				TestPaymentRequest{},
			),
		)
		require.NoError(t, err)

		// Retrieve the payment
		getPaymentResp, err := squareClient.Payments.Get(
			context.Background(),
			*createPaymentResp.Payment.ID,
		)

		require.NoError(t, err)
		assert.NotNil(t, getPaymentResp.Payment)
		assert.Equal(t, *createPaymentResp.Payment.ID, *getPaymentResp.Payment.ID)
		assert.Equal(t, *getPaymentResp.Payment.Status, "COMPLETED")
		assert.Nil(t, getPaymentResp.Payment.RefundIDs)

		// Refund the payment
		refundResp, err := squareClient.Refunds.RefundPayment(
			context.Background(),
			&square.RefundPaymentRequest{
				IdempotencyKey: newTestUUID(),
				PaymentID:      createPaymentResp.Payment.ID,
				AmountMoney:    newTestMoney(1001),
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, refundResp.Refund)

		// Retrieve the payment and verify the refund
		getPaymentResp, err = squareClient.Payments.Get(
			context.Background(),
			*createPaymentResp.Payment.ID,
		)
		require.NoError(t, err)
		assert.NotNil(t, getPaymentResp.Payment)
		assert.Equal(t, len(getPaymentResp.Payment.RefundIDs), 1)
	})
}
