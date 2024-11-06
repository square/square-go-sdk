// This file was auto-generated by Fern from our API Definition.

package square

type V1ListOrdersRequest struct {
	// The order in which payments are listed in the response.
	Order *SortOrder `json:"-" url:"order,omitempty"`
	// The maximum number of payments to return in a single response. This value cannot exceed 200.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken *string `json:"-" url:"batch_token,omitempty"`
}

type V1UpdateOrderRequest struct {
	// The action to perform on the order (COMPLETE, CANCEL, or REFUND).
	// See [V1UpdateOrderRequestAction](#type-v1updateorderrequestaction) for possible values
	Action V1UpdateOrderRequestAction `json:"action" url:"-"`
	// The tracking number of the shipment associated with the order. Only valid if action is COMPLETE.
	ShippedTrackingNumber *string `json:"shipped_tracking_number,omitempty" url:"-"`
	// A merchant-specified note about the completion of the order. Only valid if action is COMPLETE.
	CompletedNote *string `json:"completed_note,omitempty" url:"-"`
	// A merchant-specified note about the refunding of the order. Only valid if action is REFUND.
	RefundedNote *string `json:"refunded_note,omitempty" url:"-"`
	// A merchant-specified note about the canceling of the order. Only valid if action is CANCEL.
	CanceledNote *string `json:"canceled_note,omitempty" url:"-"`
}
