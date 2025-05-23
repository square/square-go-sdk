// This file was auto-generated by Fern from our API Definition.

package bookings

type ListLocationProfilesRequest struct {
	// The maximum number of results to return in a paged response.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
