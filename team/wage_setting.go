// This file was auto-generated by Fern from our API Definition.

package team

import (
	squaregosdk "github.com/square/square-go-sdk"
)

type UpdateWageSettingRequest struct {
	// The new `WageSetting` object that completely replaces the existing one.
	WageSetting *squaregosdk.WageSetting `json:"wage_setting,omitempty" url:"-"`
}
