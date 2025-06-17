package client

import (
	"context"
	"testing"

	square "github.com/square/square-go-sdk/v2"
	"github.com/stretchr/testify/assert"
)

func TestWebhookSignatureValidation(t *testing.T) {
	validRequest := &square.VerifySignatureRequest{
		RequestBody:     `{"merchant_id":"MLEFBHHSJGVHD","type":"webhooks.test_notification","event_id":"ac3ac95b-f97d-458c-a6e6-18981597e05f","created_at":"2022-07-13T20:30:59.037339943Z","data":{"type":"webhooks","id":"bc368e64-01aa-407e-b46e-3231809b1129"}}`,
		SignatureHeader: "GF4YkrJgGBDZ9NIYbNXBnMzqb2HoL4RW/S6vkZ9/2N4=",
		SignatureKey:    "Ibxx_5AKakO-3qeNVR61Dw",
		NotificationURL: "https://webhook.site/679a4f3a-dcfa-49ee-bac5-9d0edad886b9",
	}
	tests := []struct {
		description   string
		request       *square.VerifySignatureRequest
		expectedError string
	}{
		{
			description: "valid signature should pass",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: validRequest.NotificationURL,
			},
		},
		{
			description: "unescaped chars validation should pass",
			request: &square.VerifySignatureRequest{
				RequestBody:     `{"data":{"type":"webhooks","id":"fake_id"}}`,
				SignatureHeader: `W3FlCNk5IA3ZQ2LHTWoajvzfaDu/OwY2tNHIHC3IUOA=`,
				SignatureKey:    "signature-key",
				NotificationURL: "https://webhook.site/webhooks",
			},
		},
		{
			description: "escaped characters should pass",
			request: &square.VerifySignatureRequest{
				RequestBody:     `{"data":{"type":"webhooks","id":">id<"}}`,
				SignatureHeader: `Cxt7+aTi4rKgcA0bC4g9EHdVtLSDWdqccmL5MvihU4U=`,
				SignatureKey:    "signature-key",
				NotificationURL: "https://webhook.site/webhooks",
			},
		},
		{
			description: "notification URL mismatch should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: "https://webhook.site/79a4f3a-dcfa-49ee-bac5-9d0edad886b9",
			},
			expectedError: "signature hash does not match the expected signature hash for payload",
		},
		{
			description: "invalid signature key should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    "MCmhFRxGX82xMwg5FsaoQA",
				NotificationURL: validRequest.NotificationURL,
			},
			expectedError: "signature hash does not match the expected signature hash for payload",
		},
		{
			description: "invalid signature header should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: "1z2n3DEJJiUPKcPzQo1ftvbxGdw=",
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: validRequest.NotificationURL,
			},
			expectedError: "signature hash does not match the expected signature hash for payload",
		},
		{
			description: "request body mismatch should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     `{"merchant_id":"MLEFBHHSJGVHD","type":"webhooks.test_notification","event_id":"ac3ac95b-f97d-458c-a6e6-18981597e05f","created_at":"2022-06-13T20:30:59.037339943Z","data":{"type":"webhooks","id":"bc368e64-01aa-407e-b46e-3231809b1129"}}`,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: validRequest.NotificationURL,
			},
			expectedError: "signature hash does not match the expected signature hash for payload",
		}, {
			description: "formatted request body should fail",
			request: &square.VerifySignatureRequest{
				RequestBody: `{
	"merchant_id": "MLEFBHHSJGVHD",
	"type": "webhooks.test_notification",
	"event_id": "ac3ac95b-f97d-458c-a6e6-18981597e05f",
	"created_at": "2022-07-13T20:30:59.037339943Z",
	"data": {
		"type": "webhooks",
		"id": "bc368e64-01aa-407e-b46e-3231809b1129"
	},
}`,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: validRequest.NotificationURL,
			},
			expectedError: "signature hash does not match the expected signature hash for payload",
		},
		{
			description: "empty signature key should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    "",
				NotificationURL: validRequest.NotificationURL,
			},
			expectedError: "signature key is required",
		},
		{
			description: "empty notification URL should fail",
			request: &square.VerifySignatureRequest{
				RequestBody:     validRequest.RequestBody,
				SignatureHeader: validRequest.SignatureHeader,
				SignatureKey:    validRequest.SignatureKey,
				NotificationURL: "",
			},
			expectedError: "notification URL is required",
		},
	}

	client := &Client{}
	ctx := context.Background()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := client.VerifySignature(ctx, test.request)
			if test.expectedError == "" {
				assert.NoError(t, err, "expected no error, got %v", err)
			} else {
				assert.Error(t, err, "expected error %q, got nil", test.expectedError)
			}
		})
	}
}
