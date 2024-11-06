// This file was auto-generated by Fern from our API Definition.

package cashdrawers

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type ShiftsGetRequest struct {
	// The ID of the location to retrieve cash drawer shifts from.
	LocationID string `json:"-" url:"location_id"`
}

type ShiftsListRequest struct {
	// The ID of the location to query for a list of cash drawer shifts.
	LocationID string `json:"-" url:"location_id"`
	// The order in which cash drawer shifts are listed in the response,
	// based on their opened_at field. Default value: ASC
	SortOrder *squaregosdk.SortOrder `json:"-" url:"sort_order,omitempty"`
	// The inclusive start time of the query on opened_at, in ISO 8601 format.
	BeginTime *string `json:"-" url:"begin_time,omitempty"`
	// The exclusive end date of the query on opened_at, in ISO 8601 format.
	EndTime *string `json:"-" url:"end_time,omitempty"`
	// Number of cash drawer shift events in a page of results (200 by
	// default, 1000 max).
	Limit *int `json:"-" url:"limit,omitempty"`
	// Opaque cursor for fetching the next page of results.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type ShiftsListEventsRequest struct {
	// The ID of the location to list cash drawer shifts for.
	LocationID string `json:"-" url:"location_id"`
	// Number of resources to be returned in a page of results (200 by
	// default, 1000 max).
	Limit *int `json:"-" url:"limit,omitempty"`
	// Opaque cursor for fetching the next page of results.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
