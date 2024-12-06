//go:build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	square "github.com/square/square-go-sdk"
	client "github.com/square/square-go-sdk/client"
	option "github.com/square/square-go-sdk/option"
)

// SourceID represents the payment source ID for the sandbox environment.
const SourceID string = "cnon:card-nonce-ok"

const MaxCatalogPageSize int = 100

func newTestSquareClient(t *testing.T) *client.Client {
	token := os.Getenv("TEST_SQUARE_TOKEN")
	require.NotEmpty(t, token, "The TEST_SQUARE_TOKEN environment variable is required to run this test")
	return client.NewClient(
		option.WithToken(token),
		option.WithBaseURL("https://connect.squareupsandbox.com"),
	)
}

// Generates a new UUID for use as an idempotency key.
//
// For info on Square idempotency keys, see:
// * https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency
func newTestUUID() string {
	return uuid.New().String()
}

// Creates a Square temporary ID (for use in e.g. creation on CatalogObjects),
// which is simply a string prefixed with '#'.
//
// See: https://developer.squareup.com/docs/catalog-api/what-it-does#reference-other-objects-by-their-ids
func newTestSquareTempID() string {
	return fmt.Sprintf("#%d", uuid.New().ID())
}

func newTestMoney(amount int64) *square.Money {
	return &square.Money{
		Amount:   square.Int64(amount),
		Currency: square.CurrencyUsd.Ptr(),
	}
}

func getDefaultLocationID(c *client.Client) (string, error) {
	locationsResp, err := c.Locations.List(context.Background())
	if err != nil {
		return "", err
	}
	return *locationsResp.Locations[0].GetID(), nil
}
