// This file was auto-generated by Fern from our API Definition.

package transactions

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type PaymentsListRequest struct {
	// The order in which payments are listed in the response.
	Order *squaregosdk.SortOrder `json:"-" url:"order,omitempty"`
	// The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year.
	BeginTime *string `json:"-" url:"begin_time,omitempty"`
	// The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time.
	EndTime *string `json:"-" url:"end_time,omitempty"`
	// The maximum number of payments to return in a single response. This value cannot exceed 200.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken *string `json:"-" url:"batch_token,omitempty"`
	// Indicates whether or not to include partial payments in the response. Partial payments will have the tenders collected so far, but the itemizations will be empty until the payment is completed.
	IncludePartial *bool `json:"-" url:"include_partial,omitempty"`
}
