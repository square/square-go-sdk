// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
	io "io"
)

type DisputesCreateEvidenceFileRequest struct {
	ImageFile io.Reader                         `json:"-" url:"-"`
	Request   *CreateDisputeEvidenceFileRequest `json:"request,omitempty" url:"-"`
}

type CreateDisputeEvidenceTextRequest struct {
	// A unique key identifying the request. For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The type of evidence you are uploading.
	// See [DisputeEvidenceType](#type-disputeevidencetype) for possible values
	EvidenceType *DisputeEvidenceType `json:"evidence_type,omitempty" url:"-"`
	// The evidence string.
	EvidenceText string `json:"evidence_text" url:"-"`
}

type DisputesListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The dispute states used to filter the result. If not specified, the endpoint returns all disputes.
	States *DisputeState `json:"-" url:"states,omitempty"`
	// The ID of the location for which to return a list of disputes.
	// If not specified, the endpoint returns disputes associated with all locations.
	LocationID *string `json:"-" url:"location_id,omitempty"`
}

// Defines the fields in an `AcceptDispute` response.
type AcceptDisputeResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// Details about the accepted dispute.
	Dispute *Dispute `json:"dispute,omitempty" url:"dispute,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AcceptDisputeResponse) GetErrors() []*Error {
	if a == nil {
		return nil
	}
	return a.Errors
}

func (a *AcceptDisputeResponse) GetDispute() *Dispute {
	if a == nil {
		return nil
	}
	return a.Dispute
}

func (a *AcceptDisputeResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AcceptDisputeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AcceptDisputeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AcceptDisputeResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AcceptDisputeResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// Defines the parameters for a `CreateDisputeEvidenceFile` request.
type CreateDisputeEvidenceFileRequest struct {
	// A unique key identifying the request. For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"idempotency_key"`
	// The type of evidence you are uploading.
	// See [DisputeEvidenceType](#type-disputeevidencetype) for possible values
	EvidenceType *DisputeEvidenceType `json:"evidence_type,omitempty" url:"evidence_type,omitempty"`
	// The MIME type of the uploaded file.
	// The type can be image/heic, image/heif, image/jpeg, application/pdf, image/png, or image/tiff.
	ContentType *string `json:"content_type,omitempty" url:"content_type,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateDisputeEvidenceFileRequest) GetIdempotencyKey() string {
	if c == nil {
		return ""
	}
	return c.IdempotencyKey
}

func (c *CreateDisputeEvidenceFileRequest) GetEvidenceType() *DisputeEvidenceType {
	if c == nil {
		return nil
	}
	return c.EvidenceType
}

func (c *CreateDisputeEvidenceFileRequest) GetContentType() *string {
	if c == nil {
		return nil
	}
	return c.ContentType
}

func (c *CreateDisputeEvidenceFileRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateDisputeEvidenceFileRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateDisputeEvidenceFileRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateDisputeEvidenceFileRequest(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateDisputeEvidenceFileRequest) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Defines the fields in a `CreateDisputeEvidenceFile` response.
type CreateDisputeEvidenceFileResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The metadata of the newly uploaded dispute evidence.
	Evidence *DisputeEvidence `json:"evidence,omitempty" url:"evidence,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateDisputeEvidenceFileResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateDisputeEvidenceFileResponse) GetEvidence() *DisputeEvidence {
	if c == nil {
		return nil
	}
	return c.Evidence
}

func (c *CreateDisputeEvidenceFileResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateDisputeEvidenceFileResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateDisputeEvidenceFileResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateDisputeEvidenceFileResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateDisputeEvidenceFileResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Defines the fields in a `CreateDisputeEvidenceText` response.
type CreateDisputeEvidenceTextResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The newly uploaded dispute evidence metadata.
	Evidence *DisputeEvidence `json:"evidence,omitempty" url:"evidence,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateDisputeEvidenceTextResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateDisputeEvidenceTextResponse) GetEvidence() *DisputeEvidence {
	if c == nil {
		return nil
	}
	return c.Evidence
}

func (c *CreateDisputeEvidenceTextResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateDisputeEvidenceTextResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateDisputeEvidenceTextResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateDisputeEvidenceTextResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateDisputeEvidenceTextResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Represents a [dispute](https://developer.squareup.com/docs/disputes-api/overview) a cardholder initiated with their bank.
type Dispute struct {
	// The unique ID for this `Dispute`, generated by Square.
	DisputeID *string `json:"dispute_id,omitempty" url:"dispute_id,omitempty"`
	// The unique ID for this `Dispute`, generated by Square.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The disputed amount, which can be less than the total transaction amount.
	// For instance, if multiple items were purchased but the cardholder only initiates a dispute over some of the items.
	AmountMoney *Money `json:"amount_money,omitempty" url:"amount_money,omitempty"`
	// The reason why the cardholder initiated the dispute.
	// See [DisputeReason](#type-disputereason) for possible values
	Reason *DisputeReason `json:"reason,omitempty" url:"reason,omitempty"`
	// The current state of this dispute.
	// See [DisputeState](#type-disputestate) for possible values
	State *DisputeState `json:"state,omitempty" url:"state,omitempty"`
	// The deadline by which the seller must respond to the dispute, in [RFC 3339 format](https://developer.squareup.com/docs/build-basics/common-data-types/working-with-dates).
	DueAt *string `json:"due_at,omitempty" url:"due_at,omitempty"`
	// The payment challenged in this dispute.
	DisputedPayment *DisputedPayment `json:"disputed_payment,omitempty" url:"disputed_payment,omitempty"`
	// The IDs of the evidence associated with the dispute.
	EvidenceIDs []string `json:"evidence_ids,omitempty" url:"evidence_ids,omitempty"`
	// The card brand used in the disputed payment.
	// See [CardBrand](#type-cardbrand) for possible values
	CardBrand *CardBrand `json:"card_brand,omitempty" url:"card_brand,omitempty"`
	// The timestamp when the dispute was created, in RFC 3339 format.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The timestamp when the dispute was last updated, in RFC 3339 format.
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The ID of the dispute in the card brand system, generated by the card brand.
	BrandDisputeID *string `json:"brand_dispute_id,omitempty" url:"brand_dispute_id,omitempty"`
	// The timestamp when the dispute was reported, in RFC 3339 format.
	ReportedDate *string `json:"reported_date,omitempty" url:"reported_date,omitempty"`
	// The timestamp when the dispute was reported, in RFC 3339 format.
	ReportedAt *string `json:"reported_at,omitempty" url:"reported_at,omitempty"`
	// The current version of the `Dispute`.
	Version *int `json:"version,omitempty" url:"version,omitempty"`
	// The ID of the location where the dispute originated.
	LocationID *string `json:"location_id,omitempty" url:"location_id,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *Dispute) GetDisputeID() *string {
	if d == nil {
		return nil
	}
	return d.DisputeID
}

func (d *Dispute) GetID() *string {
	if d == nil {
		return nil
	}
	return d.ID
}

func (d *Dispute) GetAmountMoney() *Money {
	if d == nil {
		return nil
	}
	return d.AmountMoney
}

func (d *Dispute) GetReason() *DisputeReason {
	if d == nil {
		return nil
	}
	return d.Reason
}

func (d *Dispute) GetState() *DisputeState {
	if d == nil {
		return nil
	}
	return d.State
}

func (d *Dispute) GetDueAt() *string {
	if d == nil {
		return nil
	}
	return d.DueAt
}

func (d *Dispute) GetDisputedPayment() *DisputedPayment {
	if d == nil {
		return nil
	}
	return d.DisputedPayment
}

func (d *Dispute) GetEvidenceIDs() []string {
	if d == nil {
		return nil
	}
	return d.EvidenceIDs
}

func (d *Dispute) GetCardBrand() *CardBrand {
	if d == nil {
		return nil
	}
	return d.CardBrand
}

func (d *Dispute) GetCreatedAt() *string {
	if d == nil {
		return nil
	}
	return d.CreatedAt
}

func (d *Dispute) GetUpdatedAt() *string {
	if d == nil {
		return nil
	}
	return d.UpdatedAt
}

func (d *Dispute) GetBrandDisputeID() *string {
	if d == nil {
		return nil
	}
	return d.BrandDisputeID
}

func (d *Dispute) GetReportedDate() *string {
	if d == nil {
		return nil
	}
	return d.ReportedDate
}

func (d *Dispute) GetReportedAt() *string {
	if d == nil {
		return nil
	}
	return d.ReportedAt
}

func (d *Dispute) GetVersion() *int {
	if d == nil {
		return nil
	}
	return d.Version
}

func (d *Dispute) GetLocationID() *string {
	if d == nil {
		return nil
	}
	return d.LocationID
}

func (d *Dispute) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *Dispute) UnmarshalJSON(data []byte) error {
	type unmarshaler Dispute
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = Dispute(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *Dispute) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// The list of possible reasons why a cardholder might initiate a
// dispute with their bank.
type DisputeReason string

const (
	DisputeReasonUnknownReason          DisputeReason = "UNKNOWN_REASON"
	DisputeReasonAmountDiffers          DisputeReason = "AMOUNT_DIFFERS"
	DisputeReasonCancelled              DisputeReason = "CANCELLED"
	DisputeReasonCompliance             DisputeReason = "COMPLIANCE"
	DisputeReasonDissatisfied           DisputeReason = "DISSATISFIED"
	DisputeReasonDuplicate              DisputeReason = "DUPLICATE"
	DisputeReasonFraud                  DisputeReason = "FRAUD"
	DisputeReasonNoKnowledge            DisputeReason = "NO_KNOWLEDGE"
	DisputeReasonNotAsDescribed         DisputeReason = "NOT_AS_DESCRIBED"
	DisputeReasonNotReceived            DisputeReason = "NOT_RECEIVED"
	DisputeReasonPaidByOtherMeans       DisputeReason = "PAID_BY_OTHER_MEANS"
	DisputeReasonCustomerRequestsCredit DisputeReason = "CUSTOMER_REQUESTS_CREDIT"
	DisputeReasonUnauthorized           DisputeReason = "UNAUTHORIZED"
	DisputeReasonReturned               DisputeReason = "RETURNED"
	DisputeReasonInsufficientResponse   DisputeReason = "INSUFFICIENT_RESPONSE"
	DisputeReasonRequestDocumentation   DisputeReason = "REQUEST_DOCUMENTATION"
	DisputeReasonEmvLiabilityShift      DisputeReason = "EMV_LIABILITY_SHIFT"
)

func NewDisputeReasonFromString(s string) (DisputeReason, error) {
	switch s {
	case "UNKNOWN_REASON":
		return DisputeReasonUnknownReason, nil
	case "AMOUNT_DIFFERS":
		return DisputeReasonAmountDiffers, nil
	case "CANCELLED":
		return DisputeReasonCancelled, nil
	case "COMPLIANCE":
		return DisputeReasonCompliance, nil
	case "DISSATISFIED":
		return DisputeReasonDissatisfied, nil
	case "DUPLICATE":
		return DisputeReasonDuplicate, nil
	case "FRAUD":
		return DisputeReasonFraud, nil
	case "NO_KNOWLEDGE":
		return DisputeReasonNoKnowledge, nil
	case "NOT_AS_DESCRIBED":
		return DisputeReasonNotAsDescribed, nil
	case "NOT_RECEIVED":
		return DisputeReasonNotReceived, nil
	case "PAID_BY_OTHER_MEANS":
		return DisputeReasonPaidByOtherMeans, nil
	case "CUSTOMER_REQUESTS_CREDIT":
		return DisputeReasonCustomerRequestsCredit, nil
	case "UNAUTHORIZED":
		return DisputeReasonUnauthorized, nil
	case "RETURNED":
		return DisputeReasonReturned, nil
	case "INSUFFICIENT_RESPONSE":
		return DisputeReasonInsufficientResponse, nil
	case "REQUEST_DOCUMENTATION":
		return DisputeReasonRequestDocumentation, nil
	case "EMV_LIABILITY_SHIFT":
		return DisputeReasonEmvLiabilityShift, nil
	}
	var t DisputeReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DisputeReason) Ptr() *DisputeReason {
	return &d
}

// The list of possible dispute states.
type DisputeState string

const (
	DisputeStateUnknownState            DisputeState = "UNKNOWN_STATE"
	DisputeStateInquiryEvidenceRequired DisputeState = "INQUIRY_EVIDENCE_REQUIRED"
	DisputeStateInquiryProcessing       DisputeState = "INQUIRY_PROCESSING"
	DisputeStateInquiryClosed           DisputeState = "INQUIRY_CLOSED"
	DisputeStateEvidenceRequired        DisputeState = "EVIDENCE_REQUIRED"
	DisputeStateProcessing              DisputeState = "PROCESSING"
	DisputeStateWon                     DisputeState = "WON"
	DisputeStateLost                    DisputeState = "LOST"
	DisputeStateAccepted                DisputeState = "ACCEPTED"
	DisputeStateWaitingThirdParty       DisputeState = "WAITING_THIRD_PARTY"
)

func NewDisputeStateFromString(s string) (DisputeState, error) {
	switch s {
	case "UNKNOWN_STATE":
		return DisputeStateUnknownState, nil
	case "INQUIRY_EVIDENCE_REQUIRED":
		return DisputeStateInquiryEvidenceRequired, nil
	case "INQUIRY_PROCESSING":
		return DisputeStateInquiryProcessing, nil
	case "INQUIRY_CLOSED":
		return DisputeStateInquiryClosed, nil
	case "EVIDENCE_REQUIRED":
		return DisputeStateEvidenceRequired, nil
	case "PROCESSING":
		return DisputeStateProcessing, nil
	case "WON":
		return DisputeStateWon, nil
	case "LOST":
		return DisputeStateLost, nil
	case "ACCEPTED":
		return DisputeStateAccepted, nil
	case "WAITING_THIRD_PARTY":
		return DisputeStateWaitingThirdParty, nil
	}
	var t DisputeState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DisputeState) Ptr() *DisputeState {
	return &d
}

// The payment the cardholder disputed.
type DisputedPayment struct {
	// Square-generated unique ID of the payment being disputed.
	PaymentID *string `json:"payment_id,omitempty" url:"payment_id,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DisputedPayment) GetPaymentID() *string {
	if d == nil {
		return nil
	}
	return d.PaymentID
}

func (d *DisputedPayment) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DisputedPayment) UnmarshalJSON(data []byte) error {
	type unmarshaler DisputedPayment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DisputedPayment(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DisputedPayment) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Defines fields in a `RetrieveDispute` response.
type GetDisputeResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// Details about the requested `Dispute`.
	Dispute *Dispute `json:"dispute,omitempty" url:"dispute,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetDisputeResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetDisputeResponse) GetDispute() *Dispute {
	if g == nil {
		return nil
	}
	return g.Dispute
}

func (g *GetDisputeResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetDisputeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetDisputeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetDisputeResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetDisputeResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Defines fields in a `ListDisputes` response.
type ListDisputesResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The list of disputes.
	Disputes []*Dispute `json:"disputes,omitempty" url:"disputes,omitempty"`
	// The pagination cursor to be used in a subsequent request.
	// If unset, this is the final response. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListDisputesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListDisputesResponse) GetDisputes() []*Dispute {
	if l == nil {
		return nil
	}
	return l.Disputes
}

func (l *ListDisputesResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListDisputesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListDisputesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListDisputesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListDisputesResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListDisputesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Defines the fields in a `SubmitEvidence` response.
type SubmitEvidenceResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The `Dispute` for which evidence was submitted.
	Dispute *Dispute `json:"dispute,omitempty" url:"dispute,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SubmitEvidenceResponse) GetErrors() []*Error {
	if s == nil {
		return nil
	}
	return s.Errors
}

func (s *SubmitEvidenceResponse) GetDispute() *Dispute {
	if s == nil {
		return nil
	}
	return s.Dispute
}

func (s *SubmitEvidenceResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SubmitEvidenceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SubmitEvidenceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SubmitEvidenceResponse(value)

	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SubmitEvidenceResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
