//go:build integration

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	square "github.com/square/square-go-sdk/v3"
	client "github.com/square/square-go-sdk/v3/client"
	option "github.com/square/square-go-sdk/v3/option"
	reporting "github.com/square/square-go-sdk/v3/reporting"
)

// Reporting API live integration tests.
//
// The Reporting API is a beta, bespoke offering served ONLY from production
// (connect.squareup.com/reporting) — it is not routed on sandbox (returns 404
// there). Validating it live therefore needs a production, reporting-provisioned
// token. The repo's other integration tests use TEST_SQUARE_TOKEN against
// sandbox (which 401s against prod), so this suite is gated behind — and
// authenticates with — TEST_SQUARE_REPORTING, the production,
// reporting-provisioned access token. It skips by default when that is unset,
// keeping CI green. The endpoints are read-only (schema discovery + queries).
// The polling *logic* is covered without a live account in
// reporting/load_and_wait_test.go.
//
// Run it against a real prod account:
//
//	TEST_SQUARE_REPORTING=<prod-reporting-token> \
//	  go test ./integration_tests/... -tags=integration -run TestReportingAPI -v
//	# override the host with TEST_SQUARE_BASE_URL=<url> if reporting moves.
func TestReportingAPI(t *testing.T) {
	token := os.Getenv("TEST_SQUARE_REPORTING")
	if token == "" {
		t.Skip("set TEST_SQUARE_REPORTING=<prod-reporting-token> to run the live reporting suite")
	}

	// Reporting only exists on production; allow overriding the host via TEST_SQUARE_BASE_URL.
	baseURL := os.Getenv("TEST_SQUARE_BASE_URL")
	if baseURL == "" {
		baseURL = square.Environments.Production
	}
	// Make the live target unambiguous in the test output (useful when triaging CI).
	t.Logf("[reporting] base URL: %s  ->  %s/reporting/v1/{meta,load}", baseURL, baseURL)

	squareClient := client.NewClient(
		option.WithToken(token),
		option.WithBaseURL(baseURL),
	)
	ctx := context.Background()

	// firstMeasureName resolves the first queryable measure from the live schema,
	// e.g. "Orders.count".
	firstMeasureName := func(t *testing.T) string {
		metadata, err := squareClient.Reporting.GetMetadata(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, metadata.GetCubes(), "no cubes are available on the reporting schema for this account")
		measures := metadata.GetCubes()[0].GetMeasures()
		require.NotEmpty(t, measures, "the first cube exposes no measures")
		return measures[0].GetName()
	}

	t.Run("GetMetadata returns the queryable schema (cubes + measures)", func(t *testing.T) {
		metadata, err := squareClient.Reporting.GetMetadata(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, metadata.GetCubes())
	})

	t.Run("Load returns either results or the 'Continue wait' sentinel for an in-flight query", func(t *testing.T) {
		measure := firstMeasureName(t)
		response, err := squareClient.Reporting.Load(ctx, &square.LoadRequest{
			Query: &square.Query{Measures: []string{measure}},
		})
		require.NoError(t, err)

		if sentinel, ok := response.GetExtraProperties()["error"].(string); ok {
			// Documented async behavior: a still-processing query comes back as HTTP 200
			// with { "error": "Continue wait" } instead of results.
			assert.Equal(t, "Continue wait", sentinel)
		} else {
			// A resolved query carries its rows in the flat top-level `data`
			// field (e.g. [{"Orders.count": 0}]); there is no `results` wrapper.
			assert.NotNil(t, response.GetData())
		}
	})

	t.Run("LoadAndWait resolves a query end-to-end without surfacing 'Continue wait'", func(t *testing.T) {
		measure := firstMeasureName(t)

		// Polling can take minutes; bound it with a context deadline.
		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()

		response, err := squareClient.Reporting.LoadAndWait(
			ctx,
			&square.LoadRequest{Query: &square.Query{Measures: []string{measure}}},
			&reporting.LoadAndWaitOptions{MaxAttempts: 20, InitialDelay: 2 * time.Second, MaxDelay: 20 * time.Second},
		)
		require.NoError(t, err)

		// The polling helper must never hand back the raw "Continue wait" sentinel.
		_, hasError := response.GetExtraProperties()["error"]
		assert.False(t, hasError)
		// The resolved payload lives in the flat top-level `data` field.
		assert.NotNil(t, response.GetData())
	})
}
