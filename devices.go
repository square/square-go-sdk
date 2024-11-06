// This file was auto-generated by Fern from our API Definition.

package square

type DevicesListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The order in which results are listed.
	//
	// - `ASC` - Oldest to newest.
	// - `DESC` - Newest to oldest (default).
	SortOrder *SortOrder `json:"-" url:"sort_order,omitempty"`
	// The number of results to return in a single page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// If present, only returns devices at the target location.
	LocationID *string `json:"-" url:"location_id,omitempty"`
}
