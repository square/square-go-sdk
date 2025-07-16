package square

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPtr(s string) *string { return &s }

func TestGetPayoutResponse(t *testing.T) {
	for _, test := range []struct {
		description string
		jsonStr     string
		expected    *GetPayoutResponse
		expectedErr string
	}{
		{
			description: "valid payout input",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
					RawJSON:    json.RawMessage(`{"id": "test_id", "location_id": "loc_123"}`),
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
						RawJSON:  json.RawMessage(`{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}`),
					},
				},
				RawJSON: json.RawMessage(`{"payout": {"id": "test_id", "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`),
			},
		},
		{
			description: "invalid payout input (payout id is a number)",
			jsonStr:     `{"payout": {"id": 123, "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					LocationID: "loc_123",
					RawJSON:    json.RawMessage(`{"id": 123, "location_id": "loc_123"}`),
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
						RawJSON:  json.RawMessage(`{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}`),
					},
				},
				RawJSON: json.RawMessage(`{"payout": {"id": 123, "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`),
			},
		},
		{
			description: "extra property at top level",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "foo": "bar"}`,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
					RawJSON:    json.RawMessage(`{"id": "test_id", "location_id": "loc_123"}`),
				},
				ExtraProperties: map[string]interface{}{
					"foo": "bar",
				},
				RawJSON: json.RawMessage(`{"payout": {"id": "test_id", "location_id": "loc_123"}, "foo": "bar"}`),
			},
		},
		{
			description: "extra property inside payout",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123", "baz": 42}}`,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
					ExtraProperties: map[string]interface{}{
						"baz": float64(42),
					},
					RawJSON: json.RawMessage(`{"id": "test_id", "location_id": "loc_123", "baz": 42}`),
				},
				RawJSON: json.RawMessage(`{"payout": {"id": "test_id", "location_id": "loc_123", "baz": 42}}`),
			},
		},
	} {
		var resp GetPayoutResponse
		err := json.NewDecoder(strings.NewReader(test.jsonStr)).Decode(&resp)

		// Check for errors.
		if test.expectedErr != "" {
			assert.EqualError(t, err, test.expectedErr, test.description)
		} else {
			assert.NoError(t, err, test.description)
		}

		// Check for expected values.
		if test.expected != nil {
			assert.EqualValues(t, test.expected, &resp, test.description)
		}
	}
}
