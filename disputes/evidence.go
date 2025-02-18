// This file was auto-generated by Fern from our API Definition.

package disputes

type DeleteEvidenceRequest struct {
	// The ID of the dispute from which you want to remove evidence.
	DisputeID string `json:"-" url:"-"`
	// The ID of the evidence you want to remove.
	EvidenceID string `json:"-" url:"-"`
}

type GetEvidenceRequest struct {
	// The ID of the dispute from which you want to retrieve evidence metadata.
	DisputeID string `json:"-" url:"-"`
	// The ID of the evidence to retrieve.
	EvidenceID string `json:"-" url:"-"`
}

type ListEvidenceRequest struct {
	// The ID of the dispute.
	DisputeID string `json:"-" url:"-"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
