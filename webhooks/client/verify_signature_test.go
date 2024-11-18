package client

import (
	"context"
	"testing"

	"github.com/square/square-go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestWebhookSignatureValidation(t *testing.T) {
	validRequest := &square.VerifySignatureRequest{
		RequestBody:     `{"merchant_id":"MLEFBHHSJGVHD","type":"webhooks.test_notification","event_id":"ac3ac95b-f97d-458c-a6e6-18981597e05f","created_at":"2022-07-13T20:30:59.037339943Z","data":{"type":"webhooks","id":"bc368e64-01aa-407e-b46e-3231809b1129"}}`,
		SignatureHeader: "185e1892b2601810d9f4d2186cd5c19cccea6f61e82f8456fd2eaf919f7fd8de",
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
				SignatureHeader: `5b716508d939200dd94362c74d6a1a8efcdf683bbf3b0636b4d1c81c2dc850e0`,
				SignatureKey:    "signature-key",
				NotificationURL: "https://webhook.site/webhooks",
			},
		},
		{
			description: "escaped characters should pass",
			request: &square.VerifySignatureRequest{
				RequestBody:     `{"data":{"type":"webhooks","id":">id<"}}`,
				SignatureHeader: `0b1b7bf9a4e2e2b2a0700d1b0b883d107755b4b48359da9c7262f932f8a15385`,
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
