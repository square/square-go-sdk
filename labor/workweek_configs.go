// This file was auto-generated by Fern from our API Definition.

package labor

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type UpdateWorkweekConfigRequest struct {
	// The UUID for the `WorkweekConfig` object being updated.
	ID string `json:"-" url:"-"`
	// The updated `WorkweekConfig` object.
	WorkweekConfig *squaregosdk.WorkweekConfig `json:"workweek_config,omitempty" url:"-"`
}

type WorkweekConfigsListRequest struct {
	// The maximum number of `WorkweekConfigs` results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pointer to the next page of `WorkweekConfig` results to fetch.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
