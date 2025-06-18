package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"hash"

	v2 "github.com/square/square-go-sdk/v2"
)

// VerifySignature verifies and validates an event notification.
func (c *Client) VerifySignature(_ context.Context, request *v2.VerifySignatureRequest) error {
	if request.RequestBody == "" {
		return nil
	}
	if err := validateVerifySignatureRequest(request); err != nil {
		return err
	}
	expectedSignature, err := createHmac(request.NotificationURL+request.RequestBody, request.SignatureKey)
	if err != nil {
		return err
	}
	if !timingSafeEqual(expectedSignature, request.SignatureHeader) {
		return errors.New("signature hash does not match the expected signature hash for payload")
	}
	return nil
}

func validateVerifySignatureRequest(request *v2.VerifySignatureRequest) error {
	if request.SignatureHeader == "" {
		return errors.New("signature header is required")
	}
	if request.SignatureKey == "" {
		return errors.New("signature key is required")
	}
	if request.NotificationURL == "" {
		return errors.New("notification URL is required")
	}
	return nil
}

func createHmac(data, key string) (string, error) {
	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func timingSafeEqual(a, b string) bool {
	return hmac.Equal([]byte(a), []byte(b))
}
