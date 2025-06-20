// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/v2/internal"
)

type GetDevicesRequest struct {
	// The unique ID for the desired `Device`.
	DeviceID string `json:"-" url:"-"`
}

type ListDevicesRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The order in which results are listed.
	// - `ASC` - Oldest to newest.
	// - `DESC` - Newest to oldest (default).
	SortOrder *SortOrder `json:"-" url:"sort_order,omitempty"`
	// The number of results to return in a single page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// If present, only returns devices at the target location.
	LocationID *string `json:"-" url:"location_id,omitempty"`
}

type ApplicationType = string

// The wrapper object for the component entries of a given component type.
type Component struct {
	// The type of this component. Each component type has expected properties expressed
	// in a structured format within its corresponding `*_details` field.
	// See [ComponentType](#type-componenttype) for possible values
	Type ComponentComponentType `json:"type" url:"type"`
	// Structured data for an `Application`, set for Components of type `APPLICATION`.
	ApplicationDetails *DeviceComponentDetailsApplicationDetails `json:"application_details,omitempty" url:"application_details,omitempty"`
	// Structured data for a `CardReader`, set for Components of type `CARD_READER`.
	CardReaderDetails *DeviceComponentDetailsCardReaderDetails `json:"card_reader_details,omitempty" url:"card_reader_details,omitempty"`
	// Structured data for a `Battery`, set for Components of type `BATTERY`.
	BatteryDetails *DeviceComponentDetailsBatteryDetails `json:"battery_details,omitempty" url:"battery_details,omitempty"`
	// Structured data for a `WiFi` interface, set for Components of type `WIFI`.
	WifiDetails *DeviceComponentDetailsWiFiDetails `json:"wifi_details,omitempty" url:"wifi_details,omitempty"`
	// Structured data for an `Ethernet` interface, set for Components of type `ETHERNET`.
	EthernetDetails *DeviceComponentDetailsEthernetDetails `json:"ethernet_details,omitempty" url:"ethernet_details,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Component) GetType() ComponentComponentType {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *Component) GetApplicationDetails() *DeviceComponentDetailsApplicationDetails {
	if c == nil {
		return nil
	}
	return c.ApplicationDetails
}

func (c *Component) GetCardReaderDetails() *DeviceComponentDetailsCardReaderDetails {
	if c == nil {
		return nil
	}
	return c.CardReaderDetails
}

func (c *Component) GetBatteryDetails() *DeviceComponentDetailsBatteryDetails {
	if c == nil {
		return nil
	}
	return c.BatteryDetails
}

func (c *Component) GetWifiDetails() *DeviceComponentDetailsWiFiDetails {
	if c == nil {
		return nil
	}
	return c.WifiDetails
}

func (c *Component) GetEthernetDetails() *DeviceComponentDetailsEthernetDetails {
	if c == nil {
		return nil
	}
	return c.EthernetDetails
}

func (c *Component) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Component) UnmarshalJSON(data []byte) error {
	type unmarshaler Component
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Component(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Component) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// An enum for ComponentType.
type ComponentComponentType string

const (
	ComponentComponentTypeApplication ComponentComponentType = "APPLICATION"
	ComponentComponentTypeCardReader  ComponentComponentType = "CARD_READER"
	ComponentComponentTypeBattery     ComponentComponentType = "BATTERY"
	ComponentComponentTypeWifi        ComponentComponentType = "WIFI"
	ComponentComponentTypeEthernet    ComponentComponentType = "ETHERNET"
	ComponentComponentTypePrinter     ComponentComponentType = "PRINTER"
)

func NewComponentComponentTypeFromString(s string) (ComponentComponentType, error) {
	switch s {
	case "APPLICATION":
		return ComponentComponentTypeApplication, nil
	case "CARD_READER":
		return ComponentComponentTypeCardReader, nil
	case "BATTERY":
		return ComponentComponentTypeBattery, nil
	case "WIFI":
		return ComponentComponentTypeWifi, nil
	case "ETHERNET":
		return ComponentComponentTypeEthernet, nil
	case "PRINTER":
		return ComponentComponentTypePrinter, nil
	}
	var t ComponentComponentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ComponentComponentType) Ptr() *ComponentComponentType {
	return &c
}

type Device struct {
	// A synthetic identifier for the device. The identifier includes a standardized prefix and
	// is otherwise an opaque id generated from key device fields.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// A collection of DeviceAttributes representing the device.
	Attributes *DeviceAttributes `json:"attributes,omitempty" url:"attributes,omitempty"`
	// A list of components applicable to the device.
	Components []*Component `json:"components,omitempty" url:"components,omitempty"`
	// The current status of the device.
	Status *DeviceStatus `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *Device) GetID() *string {
	if d == nil {
		return nil
	}
	return d.ID
}

func (d *Device) GetAttributes() *DeviceAttributes {
	if d == nil {
		return nil
	}
	return d.Attributes
}

func (d *Device) GetComponents() []*Component {
	if d == nil {
		return nil
	}
	return d.Components
}

func (d *Device) GetStatus() *DeviceStatus {
	if d == nil {
		return nil
	}
	return d.Status
}

func (d *Device) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *Device) UnmarshalJSON(data []byte) error {
	type unmarshaler Device
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = Device(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *Device) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceAttributes struct {
	// The device type.
	// See [DeviceType](#type-devicetype) for possible values
	Type DeviceAttributesDeviceType `json:"type,omitempty" url:"type,omitempty"`
	// The maker of the device.
	Manufacturer string `json:"manufacturer" url:"manufacturer"`
	// The specific model of the device.
	Model *string `json:"model,omitempty" url:"model,omitempty"`
	// A seller-specified name for the device.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The manufacturer-supplied identifier for the device (where available). In many cases,
	// this identifier will be a serial number.
	ManufacturersID *string `json:"manufacturers_id,omitempty" url:"manufacturers_id,omitempty"`
	// The RFC 3339-formatted value of the most recent update to the device information.
	// (Could represent any field update on the device.)
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The current version of software installed on the device.
	Version *string `json:"version,omitempty" url:"version,omitempty"`
	// The merchant_token identifying the merchant controlling the device.
	MerchantToken *string `json:"merchant_token,omitempty" url:"merchant_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceAttributes) GetManufacturer() string {
	if d == nil {
		return ""
	}
	return d.Manufacturer
}

func (d *DeviceAttributes) GetModel() *string {
	if d == nil {
		return nil
	}
	return d.Model
}

func (d *DeviceAttributes) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DeviceAttributes) GetManufacturersID() *string {
	if d == nil {
		return nil
	}
	return d.ManufacturersID
}

func (d *DeviceAttributes) GetUpdatedAt() *string {
	if d == nil {
		return nil
	}
	return d.UpdatedAt
}

func (d *DeviceAttributes) GetVersion() *string {
	if d == nil {
		return nil
	}
	return d.Version
}

func (d *DeviceAttributes) GetMerchantToken() *string {
	if d == nil {
		return nil
	}
	return d.MerchantToken
}

func (d *DeviceAttributes) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceAttributes) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceAttributes
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceAttributes(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceAttributes) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// An enum identifier of the device type.
type DeviceAttributesDeviceType = string

type DeviceComponentDetailsApplicationDetails struct {
	// The type of application.
	// See [ApplicationType](#type-applicationtype) for possible values
	ApplicationType *ApplicationType `json:"application_type,omitempty" url:"application_type,omitempty"`
	// The version of the application.
	Version *string `json:"version,omitempty" url:"version,omitempty"`
	// The location_id of the session for the application.
	SessionLocation *string `json:"session_location,omitempty" url:"session_location,omitempty"`
	// The id of the device code that was used to log in to the device.
	DeviceCodeID *string `json:"device_code_id,omitempty" url:"device_code_id,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsApplicationDetails) GetVersion() *string {
	if d == nil {
		return nil
	}
	return d.Version
}

func (d *DeviceComponentDetailsApplicationDetails) GetSessionLocation() *string {
	if d == nil {
		return nil
	}
	return d.SessionLocation
}

func (d *DeviceComponentDetailsApplicationDetails) GetDeviceCodeID() *string {
	if d == nil {
		return nil
	}
	return d.DeviceCodeID
}

func (d *DeviceComponentDetailsApplicationDetails) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsApplicationDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsApplicationDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsApplicationDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsApplicationDetails) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceComponentDetailsBatteryDetails struct {
	// The battery charge percentage as displayed on the device.
	VisiblePercent *int `json:"visible_percent,omitempty" url:"visible_percent,omitempty"`
	// The status of external_power.
	// See [ExternalPower](#type-externalpower) for possible values
	ExternalPower *DeviceComponentDetailsExternalPower `json:"external_power,omitempty" url:"external_power,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsBatteryDetails) GetVisiblePercent() *int {
	if d == nil {
		return nil
	}
	return d.VisiblePercent
}

func (d *DeviceComponentDetailsBatteryDetails) GetExternalPower() *DeviceComponentDetailsExternalPower {
	if d == nil {
		return nil
	}
	return d.ExternalPower
}

func (d *DeviceComponentDetailsBatteryDetails) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsBatteryDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsBatteryDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsBatteryDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsBatteryDetails) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceComponentDetailsCardReaderDetails struct {
	// The version of the card reader.
	Version *string `json:"version,omitempty" url:"version,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsCardReaderDetails) GetVersion() *string {
	if d == nil {
		return nil
	}
	return d.Version
}

func (d *DeviceComponentDetailsCardReaderDetails) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsCardReaderDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsCardReaderDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsCardReaderDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsCardReaderDetails) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceComponentDetailsEthernetDetails struct {
	// A boolean to represent whether the Ethernet interface is currently active.
	Active *bool `json:"active,omitempty" url:"active,omitempty"`
	// The string representation of the device’s IPv4 address.
	IPAddressV4 *string `json:"ip_address_v4,omitempty" url:"ip_address_v4,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsEthernetDetails) GetActive() *bool {
	if d == nil {
		return nil
	}
	return d.Active
}

func (d *DeviceComponentDetailsEthernetDetails) GetIPAddressV4() *string {
	if d == nil {
		return nil
	}
	return d.IPAddressV4
}

func (d *DeviceComponentDetailsEthernetDetails) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsEthernetDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsEthernetDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsEthernetDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsEthernetDetails) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// An enum for ExternalPower.
type DeviceComponentDetailsExternalPower string

const (
	DeviceComponentDetailsExternalPowerAvailableCharging     DeviceComponentDetailsExternalPower = "AVAILABLE_CHARGING"
	DeviceComponentDetailsExternalPowerAvailableNotInUse     DeviceComponentDetailsExternalPower = "AVAILABLE_NOT_IN_USE"
	DeviceComponentDetailsExternalPowerUnavailable           DeviceComponentDetailsExternalPower = "UNAVAILABLE"
	DeviceComponentDetailsExternalPowerAvailableInsufficient DeviceComponentDetailsExternalPower = "AVAILABLE_INSUFFICIENT"
)

func NewDeviceComponentDetailsExternalPowerFromString(s string) (DeviceComponentDetailsExternalPower, error) {
	switch s {
	case "AVAILABLE_CHARGING":
		return DeviceComponentDetailsExternalPowerAvailableCharging, nil
	case "AVAILABLE_NOT_IN_USE":
		return DeviceComponentDetailsExternalPowerAvailableNotInUse, nil
	case "UNAVAILABLE":
		return DeviceComponentDetailsExternalPowerUnavailable, nil
	case "AVAILABLE_INSUFFICIENT":
		return DeviceComponentDetailsExternalPowerAvailableInsufficient, nil
	}
	var t DeviceComponentDetailsExternalPower
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DeviceComponentDetailsExternalPower) Ptr() *DeviceComponentDetailsExternalPower {
	return &d
}

// A value qualified by unit of measure.
type DeviceComponentDetailsMeasurement struct {
	Value *int `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsMeasurement) GetValue() *int {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DeviceComponentDetailsMeasurement) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsMeasurement) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsMeasurement
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsMeasurement(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsMeasurement) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceComponentDetailsWiFiDetails struct {
	// A boolean to represent whether the WiFI interface is currently active.
	Active *bool `json:"active,omitempty" url:"active,omitempty"`
	// The name of the connected WIFI network.
	Ssid *string `json:"ssid,omitempty" url:"ssid,omitempty"`
	// The string representation of the device’s IPv4 address.
	IPAddressV4 *string `json:"ip_address_v4,omitempty" url:"ip_address_v4,omitempty"`
	// The security protocol for a secure connection (e.g. WPA2). None provided if the connection
	// is unsecured.
	SecureConnection *string `json:"secure_connection,omitempty" url:"secure_connection,omitempty"`
	// A representation of signal strength of the WIFI network connection.
	SignalStrength *DeviceComponentDetailsMeasurement `json:"signal_strength,omitempty" url:"signal_strength,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceComponentDetailsWiFiDetails) GetActive() *bool {
	if d == nil {
		return nil
	}
	return d.Active
}

func (d *DeviceComponentDetailsWiFiDetails) GetSsid() *string {
	if d == nil {
		return nil
	}
	return d.Ssid
}

func (d *DeviceComponentDetailsWiFiDetails) GetIPAddressV4() *string {
	if d == nil {
		return nil
	}
	return d.IPAddressV4
}

func (d *DeviceComponentDetailsWiFiDetails) GetSecureConnection() *string {
	if d == nil {
		return nil
	}
	return d.SecureConnection
}

func (d *DeviceComponentDetailsWiFiDetails) GetSignalStrength() *DeviceComponentDetailsMeasurement {
	if d == nil {
		return nil
	}
	return d.SignalStrength
}

func (d *DeviceComponentDetailsWiFiDetails) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceComponentDetailsWiFiDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceComponentDetailsWiFiDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceComponentDetailsWiFiDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceComponentDetailsWiFiDetails) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceStatus struct {
	// See [Category](#type-category) for possible values
	Category *DeviceStatusCategory `json:"category,omitempty" url:"category,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeviceStatus) GetCategory() *DeviceStatusCategory {
	if d == nil {
		return nil
	}
	return d.Category
}

func (d *DeviceStatus) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeviceStatus) UnmarshalJSON(data []byte) error {
	type unmarshaler DeviceStatus
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeviceStatus(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeviceStatus) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeviceStatusCategory string

const (
	DeviceStatusCategoryAvailable      DeviceStatusCategory = "AVAILABLE"
	DeviceStatusCategoryNeedsAttention DeviceStatusCategory = "NEEDS_ATTENTION"
	DeviceStatusCategoryOffline        DeviceStatusCategory = "OFFLINE"
)

func NewDeviceStatusCategoryFromString(s string) (DeviceStatusCategory, error) {
	switch s {
	case "AVAILABLE":
		return DeviceStatusCategoryAvailable, nil
	case "NEEDS_ATTENTION":
		return DeviceStatusCategoryNeedsAttention, nil
	case "OFFLINE":
		return DeviceStatusCategoryOffline, nil
	}
	var t DeviceStatusCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DeviceStatusCategory) Ptr() *DeviceStatusCategory {
	return &d
}

type GetDeviceResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The requested `Device`.
	Device *Device `json:"device,omitempty" url:"device,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetDeviceResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetDeviceResponse) GetDevice() *Device {
	if g == nil {
		return nil
	}
	return g.Device
}

func (g *GetDeviceResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetDeviceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetDeviceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetDeviceResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetDeviceResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type ListDevicesResponse struct {
	// Information about errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The requested list of `Device` objects.
	Devices []*Device `json:"devices,omitempty" url:"devices,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty,
	// this is the final response.
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListDevicesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListDevicesResponse) GetDevices() []*Device {
	if l == nil {
		return nil
	}
	return l.Devices
}

func (l *ListDevicesResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListDevicesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListDevicesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListDevicesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListDevicesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListDevicesResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}
