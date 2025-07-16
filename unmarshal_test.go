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
		settings    GlobalUnmarshalType
		expected    *GetPayoutResponse
		expectedErr string
	}{
		/*
			Original unmarshaler. Existing logic.
		*/
		{
			description: "valid input with original settings",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeOriginal,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
					},
				},
			},
		},
		{
			description: "extra property at top level",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "foo": "bar"}`,
			settings:    GlobalUnmarshalTypeOriginal,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				extraProperties: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		{
			description: "extra property inside payout with original settings",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123", "baz": 42}}`,
			settings:    GlobalUnmarshalTypeOriginal,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
					extraProperties: map[string]interface{}{
						"baz": float64(42),
					},
				},
			},
		},
		{
			description: "payout input with invalid type for payout id",
			jsonStr:     `{"payout": {"id": 123, "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeOriginal,
			expected:    nil,
			expectedErr: "json: cannot unmarshal number into Go struct field unmarshaler.payout of type string",
		},
		/*
			Mapstructure for json unmarshalling. Attempts weak type matching.
		*/
		{
			description: "valid payout input with mapstructure settings",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeMapstructure,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
					},
				},
			},
		},
		{
			description: "invalid payout input with mapstructure settings",
			jsonStr:     `{"payout": {"id": 123, "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeMapstructure,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "123",
					LocationID: "loc_123",
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
					},
				},
			},
		},
		{
			description: "extra property at top level with mapstructure settings",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "foo": "bar"}`,
			settings:    GlobalUnmarshalTypeMapstructure,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				extraProperties: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		// TODO: Implement custom unmarshaler for Payout struct.
		// {
		// 	description: "extra property inside payout with mapstructure settings",
		// 	jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123", "baz": 42}}`,
		// 	settings:    GlobalUnmarshalTypeMapstructure,
		// 	expected: &GetPayoutResponse{
		// 		Payout: &Payout{
		// 			ID:         "test_id",
		// 			LocationID: "loc_123",
		// 			extraProperties: map[string]interface{}{
		// 				"baz": float64(42),
		// 			},
		// 		},
		// 		extraProperties: map[string]interface{}{},
		// 	},
		// },
		/*
			Custom unmarshaler.
		*/
		{
			description: "valid payout input with no third party library",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeNative,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
					},
				},
			},
		},
		{
			description: "invalid payout input with no third party library",
			jsonStr:     `{"payout": {"id": 123, "location_id": "loc_123"}, "errors": [{"category": "API_ERROR", "code": "INTERNAL_SERVER_ERROR", "detail": "detail", "field": "field"}]}`,
			settings:    GlobalUnmarshalTypeNative,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					LocationID: "loc_123",
				},
				Errors: []*Error{
					{
						Category: "API_ERROR",
						Code:     "INTERNAL_SERVER_ERROR",
						Detail:   stringPtr("detail"),
						Field:    stringPtr("field"),
					},
				},
			},
		},
		{
			description: "extra property at top level with no third party library",
			jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123"}, "foo": "bar"}`,
			settings:    GlobalUnmarshalTypeNative,
			expected: &GetPayoutResponse{
				Payout: &Payout{
					ID:         "test_id",
					LocationID: "loc_123",
				},
				extraProperties: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		// TODO: Implement custom unmarshaler for Payout struct.
		// {
		// 	description: "extra property inside payout with no third party library",
		// 	jsonStr:     `{"payout": {"id": "test_id", "location_id": "loc_123", "baz": 42}}`,
		// 	settings:    GlobalUnmarshalTypeNative,
		// 	expected: &GetPayoutResponse{
		// 		Payout: &Payout{
		// 			ID:         "test_id",
		// 			LocationID: "loc_123",
		// 			extraProperties: map[string]interface{}{
		// 				"baz": float64(42),
		// 			},
		// 		},
		// 		extraProperties: map[string]interface{}{},
		// 	},
		// },
	} {
		GlobalUnmarshalSettings = test.settings
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
			// Check that the raw JSON is the same.
			assert.EqualValues(t, json.RawMessage(test.jsonStr), resp.rawJSON, test.description)
			// An alternative to setting nested rawJSON fields. Ignores validation of untagged JSON fields.
			assert.Equal(t, prettyJSON(test.expected), prettyJSON(&resp), test.description)
			// Check that the extra properties are the same.
			assert.Equal(t, test.expected.extraProperties, resp.extraProperties, test.description)
		}
	}
}

func prettyJSON(v interface{}) string {
	jsonBytes, _ := json.MarshalIndent(v, "", "  ")
	return string(jsonBytes)
}
