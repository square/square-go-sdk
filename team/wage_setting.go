// This file was auto-generated by Fern from our API Definition.

package team

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type WageSettingGetRequest struct {
	// The ID of the team member for which to retrieve the wage setting.
	TeamMemberID string `json:"-" url:"-"`
}

type UpdateWageSettingRequest struct {
	// The ID of the team member for which to update the `WageSetting` object.
	TeamMemberID string `json:"-" url:"-"`
	// The new `WageSetting` object that completely replaces the existing one.
	WageSetting *squaregosdk.WageSetting `json:"wage_setting,omitempty" url:"-"`
}
