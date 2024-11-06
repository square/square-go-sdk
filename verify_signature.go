package square

type VerifySignatureRequest struct {
	RequestBody     string
	SignatureHeader string
	SignatureKey    string
	NotificationURL string
}
