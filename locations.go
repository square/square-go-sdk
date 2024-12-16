// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type CreateCheckoutRequest struct {
	// The ID of the business location to associate the checkout with.
	LocationID string `json:"-" url:"-"`
	// A unique string that identifies this checkout among others you have created. It can be
	// any valid string but must be unique for every order sent to Square Checkout for a given location ID.
	//
	// The idempotency key is used to avoid processing the same order more than once. If you are
	// unsure whether a particular checkout was created successfully, you can attempt it again with
	// the same idempotency key and all the same other parameters without worrying about creating duplicates.
	//
	// You should use a random number/string generator native to the language
	// you are working in to generate strings for your idempotency keys.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The order including line items to be checked out.
	Order *CreateOrderRequest `json:"order,omitempty" url:"-"`
	// If `true`, Square Checkout collects shipping information on your behalf and stores
	// that information with the transaction information in the Square Seller Dashboard.
	//
	// Default: `false`.
	AskForShippingAddress *bool `json:"ask_for_shipping_address,omitempty" url:"-"`
	// The email address to display on the Square Checkout confirmation page
	// and confirmation email that the buyer can use to contact the seller.
	//
	// If this value is not set, the confirmation page and email display the
	// primary email address associated with the seller's Square account.
	//
	// Default: none; only exists if explicitly set.
	MerchantSupportEmail *string `json:"merchant_support_email,omitempty" url:"-"`
	// If provided, the buyer's email is prepopulated on the checkout page
	// as an editable text field.
	//
	// Default: none; only exists if explicitly set.
	PrePopulateBuyerEmail *string `json:"pre_populate_buyer_email,omitempty" url:"-"`
	// If provided, the buyer's shipping information is prepopulated on the
	// checkout page as editable text fields.
	//
	// Default: none; only exists if explicitly set.
	PrePopulateShippingAddress *Address `json:"pre_populate_shipping_address,omitempty" url:"-"`
	// The URL to redirect to after the checkout is completed with `checkoutId`,
	// `transactionId`, and `referenceId` appended as URL parameters. For example,
	// if the provided redirect URL is `http://www.example.com/order-complete`, a
	// successful transaction redirects the customer to:
	//
	// `http://www.example.com/order-complete?checkoutId=xxxxxx&amp;referenceId=xxxxxx&amp;transactionId=xxxxxx`
	//
	// If you do not provide a redirect URL, Square Checkout displays an order
	// confirmation page on your behalf; however, it is strongly recommended that
	// you provide a redirect URL so you can verify the transaction results and
	// finalize the order through your existing/normal confirmation workflow.
	//
	// Default: none; only exists if explicitly set.
	RedirectURL *string `json:"redirect_url,omitempty" url:"-"`
	// The basic primitive of a multi-party transaction. The value is optional.
	// The transaction facilitated by you can be split from here.
	//
	// If you provide this value, the `amount_money` value in your `additional_recipients` field
	// cannot be more than 90% of the `total_money` calculated by Square for your order.
	// The `location_id` must be a valid seller location where the checkout is occurring.
	//
	// This field requires `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission.
	//
	// This field is currently not supported in the Square Sandbox.
	AdditionalRecipients []*ChargeRequestAdditionalRecipient `json:"additional_recipients,omitempty" url:"-"`
	// An optional note to associate with the `checkout` object.
	//
	// This value cannot exceed 60 characters.
	Note *string `json:"note,omitempty" url:"-"`
}

type CreateLocationRequest struct {
	// The initial values of the location being created. The `name` field is required and must be unique within a seller account.
	// All other fields are optional, but any information you care about for the location should be included.
	// The remaining fields are automatically added based on the data from the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location).
	Location *Location `json:"location,omitempty" url:"-"`
}

type LocationsGetRequest struct {
	// The ID of the location to retrieve. Specify the string
	// "main" to return the main location.
	LocationID string `json:"-" url:"-"`
}

// The hours of operation for a location.
type BusinessHours struct {
	// The list of time periods during which the business is open. There can be at most 10 periods per day.
	Periods []*BusinessHoursPeriod `json:"periods,omitempty" url:"periods,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BusinessHours) GetPeriods() []*BusinessHoursPeriod {
	if b == nil {
		return nil
	}
	return b.Periods
}

func (b *BusinessHours) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BusinessHours) UnmarshalJSON(data []byte) error {
	type unmarshaler BusinessHours
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BusinessHours(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BusinessHours) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents a period of time during which a business location is open.
type BusinessHoursPeriod struct {
	// The day of the week for this time period.
	// See [DayOfWeek](#type-dayofweek) for possible values
	DayOfWeek *DayOfWeek `json:"day_of_week,omitempty" url:"day_of_week,omitempty"`
	// The start time of a business hours period, specified in local time using partial-time
	// RFC 3339 format. For example, `8:30:00` for a period starting at 8:30 in the morning.
	// Note that the seconds value is always :00, but it is appended for conformance to the RFC.
	StartLocalTime *string `json:"start_local_time,omitempty" url:"start_local_time,omitempty"`
	// The end time of a business hours period, specified in local time using partial-time
	// RFC 3339 format. For example, `21:00:00` for a period ending at 9:00 in the evening.
	// Note that the seconds value is always :00, but it is appended for conformance to the RFC.
	EndLocalTime *string `json:"end_local_time,omitempty" url:"end_local_time,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BusinessHoursPeriod) GetDayOfWeek() *DayOfWeek {
	if b == nil {
		return nil
	}
	return b.DayOfWeek
}

func (b *BusinessHoursPeriod) GetStartLocalTime() *string {
	if b == nil {
		return nil
	}
	return b.StartLocalTime
}

func (b *BusinessHoursPeriod) GetEndLocalTime() *string {
	if b == nil {
		return nil
	}
	return b.EndLocalTime
}

func (b *BusinessHoursPeriod) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BusinessHoursPeriod) UnmarshalJSON(data []byte) error {
	type unmarshaler BusinessHoursPeriod
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BusinessHoursPeriod(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BusinessHoursPeriod) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Represents an additional recipient (other than the merchant) entitled to a portion of the tender.
// Support is currently limited to USD, CAD and GBP currencies
type ChargeRequestAdditionalRecipient struct {
	// The location ID for a recipient (other than the merchant) receiving a portion of the tender.
	LocationID string `json:"location_id" url:"location_id"`
	// The description of the additional recipient.
	Description string `json:"description" url:"description"`
	// The amount of money distributed to the recipient.
	AmountMoney *Money `json:"amount_money,omitempty" url:"amount_money,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ChargeRequestAdditionalRecipient) GetLocationID() string {
	if c == nil {
		return ""
	}
	return c.LocationID
}

func (c *ChargeRequestAdditionalRecipient) GetDescription() string {
	if c == nil {
		return ""
	}
	return c.Description
}

func (c *ChargeRequestAdditionalRecipient) GetAmountMoney() *Money {
	if c == nil {
		return nil
	}
	return c.AmountMoney
}

func (c *ChargeRequestAdditionalRecipient) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ChargeRequestAdditionalRecipient) UnmarshalJSON(data []byte) error {
	type unmarshaler ChargeRequestAdditionalRecipient
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ChargeRequestAdditionalRecipient(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ChargeRequestAdditionalRecipient) String() string {
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

// Square Checkout lets merchants accept online payments for supported
// payment types using a checkout workflow hosted on squareup.com.
type Checkout struct {
	// ID generated by Square Checkout when a new checkout is requested.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The URL that the buyer's browser should be redirected to after the
	// checkout is completed.
	CheckoutPageURL *string `json:"checkout_page_url,omitempty" url:"checkout_page_url,omitempty"`
	// If `true`, Square Checkout will collect shipping information on your
	// behalf and store that information with the transaction information in your
	// Square Dashboard.
	//
	// Default: `false`.
	AskForShippingAddress *bool `json:"ask_for_shipping_address,omitempty" url:"ask_for_shipping_address,omitempty"`
	// The email address to display on the Square Checkout confirmation page
	// and confirmation email that the buyer can use to contact the merchant.
	//
	// If this value is not set, the confirmation page and email will display the
	// primary email address associated with the merchant's Square account.
	//
	// Default: none; only exists if explicitly set.
	MerchantSupportEmail *string `json:"merchant_support_email,omitempty" url:"merchant_support_email,omitempty"`
	// If provided, the buyer's email is pre-populated on the checkout page
	// as an editable text field.
	//
	// Default: none; only exists if explicitly set.
	PrePopulateBuyerEmail *string `json:"pre_populate_buyer_email,omitempty" url:"pre_populate_buyer_email,omitempty"`
	// If provided, the buyer's shipping info is pre-populated on the
	// checkout page as editable text fields.
	//
	// Default: none; only exists if explicitly set.
	PrePopulateShippingAddress *Address `json:"pre_populate_shipping_address,omitempty" url:"pre_populate_shipping_address,omitempty"`
	// The URL to redirect to after checkout is completed with `checkoutId`,
	// Square's `orderId`, `transactionId`, and `referenceId` appended as URL
	// parameters. For example, if the provided redirect_url is
	// `http://www.example.com/order-complete`, a successful transaction redirects
	// the customer to:
	//
	// <pre><code>http://www.example.com/order-complete?checkoutId=xxxxxx&amp;orderId=xxxxxx&amp;referenceId=xxxxxx&amp;transactionId=xxxxxx</code></pre>
	//
	// If you do not provide a redirect URL, Square Checkout will display an order
	// confirmation page on your behalf; however Square strongly recommends that
	// you provide a redirect URL so you can verify the transaction results and
	// finalize the order through your existing/normal confirmation workflow.
	RedirectURL *string `json:"redirect_url,omitempty" url:"redirect_url,omitempty"`
	// Order to be checked out.
	Order *Order `json:"order,omitempty" url:"order,omitempty"`
	// The time when the checkout was created, in RFC 3339 format.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// Additional recipients (other than the merchant) receiving a portion of this checkout.
	// For example, fees assessed on the purchase by a third party integration.
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients,omitempty" url:"additional_recipients,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Checkout) GetID() *string {
	if c == nil {
		return nil
	}
	return c.ID
}

func (c *Checkout) GetCheckoutPageURL() *string {
	if c == nil {
		return nil
	}
	return c.CheckoutPageURL
}

func (c *Checkout) GetAskForShippingAddress() *bool {
	if c == nil {
		return nil
	}
	return c.AskForShippingAddress
}

func (c *Checkout) GetMerchantSupportEmail() *string {
	if c == nil {
		return nil
	}
	return c.MerchantSupportEmail
}

func (c *Checkout) GetPrePopulateBuyerEmail() *string {
	if c == nil {
		return nil
	}
	return c.PrePopulateBuyerEmail
}

func (c *Checkout) GetPrePopulateShippingAddress() *Address {
	if c == nil {
		return nil
	}
	return c.PrePopulateShippingAddress
}

func (c *Checkout) GetRedirectURL() *string {
	if c == nil {
		return nil
	}
	return c.RedirectURL
}

func (c *Checkout) GetOrder() *Order {
	if c == nil {
		return nil
	}
	return c.Order
}

func (c *Checkout) GetCreatedAt() *string {
	if c == nil {
		return nil
	}
	return c.CreatedAt
}

func (c *Checkout) GetAdditionalRecipients() []*AdditionalRecipient {
	if c == nil {
		return nil
	}
	return c.AdditionalRecipients
}

func (c *Checkout) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Checkout) UnmarshalJSON(data []byte) error {
	type unmarshaler Checkout
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Checkout(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Checkout) String() string {
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

// Latitude and longitude coordinates.
type Coordinates struct {
	// The latitude of the coordinate expressed in degrees.
	Latitude *float64 `json:"latitude,omitempty" url:"latitude,omitempty"`
	// The longitude of the coordinate expressed in degrees.
	Longitude *float64 `json:"longitude,omitempty" url:"longitude,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Coordinates) GetLatitude() *float64 {
	if c == nil {
		return nil
	}
	return c.Latitude
}

func (c *Coordinates) GetLongitude() *float64 {
	if c == nil {
		return nil
	}
	return c.Longitude
}

func (c *Coordinates) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Coordinates) UnmarshalJSON(data []byte) error {
	type unmarshaler Coordinates
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Coordinates(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Coordinates) String() string {
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

// Defines the fields that are included in the response body of
// a request to the `CreateCheckout` endpoint.
type CreateCheckoutResponse struct {
	// The newly created `checkout` object associated with the provided idempotency key.
	Checkout *Checkout `json:"checkout,omitempty" url:"checkout,omitempty"`
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateCheckoutResponse) GetCheckout() *Checkout {
	if c == nil {
		return nil
	}
	return c.Checkout
}

func (c *CreateCheckoutResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateCheckoutResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateCheckoutResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCheckoutResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateCheckoutResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateCheckoutResponse) String() string {
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

// The response object returned by the [CreateLocation](api-endpoint:Locations-CreateLocation) endpoint.
type CreateLocationResponse struct {
	// Information about [errors](https://developer.squareup.com/docs/build-basics/handling-errors) encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The newly created `Location` object.
	Location *Location `json:"location,omitempty" url:"location,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateLocationResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateLocationResponse) GetLocation() *Location {
	if c == nil {
		return nil
	}
	return c.Location
}

func (c *CreateLocationResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateLocationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateLocationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateLocationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateLocationResponse) String() string {
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

// Indicates the specific day  of the week.
type DayOfWeek string

const (
	DayOfWeekSun DayOfWeek = "SUN"
	DayOfWeekMon DayOfWeek = "MON"
	DayOfWeekTue DayOfWeek = "TUE"
	DayOfWeekWed DayOfWeek = "WED"
	DayOfWeekThu DayOfWeek = "THU"
	DayOfWeekFri DayOfWeek = "FRI"
	DayOfWeekSat DayOfWeek = "SAT"
)

func NewDayOfWeekFromString(s string) (DayOfWeek, error) {
	switch s {
	case "SUN":
		return DayOfWeekSun, nil
	case "MON":
		return DayOfWeekMon, nil
	case "TUE":
		return DayOfWeekTue, nil
	case "WED":
		return DayOfWeekWed, nil
	case "THU":
		return DayOfWeekThu, nil
	case "FRI":
		return DayOfWeekFri, nil
	case "SAT":
		return DayOfWeekSat, nil
	}
	var t DayOfWeek
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DayOfWeek) Ptr() *DayOfWeek {
	return &d
}

// Defines the fields that the [RetrieveLocation](api-endpoint:Locations-RetrieveLocation)
// endpoint returns in a response.
type GetLocationResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The requested location.
	Location *Location `json:"location,omitempty" url:"location,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetLocationResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetLocationResponse) GetLocation() *Location {
	if g == nil {
		return nil
	}
	return g.Location
}

func (g *GetLocationResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetLocationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetLocationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetLocationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetLocationResponse) String() string {
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

// Defines the fields that are included in the response body of a request
// to the [ListLocations](api-endpoint:Locations-ListLocations) endpoint.
//
// Either `errors` or `locations` is present in a given response (never both).
type ListLocationsResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The business locations.
	Locations []*Location `json:"locations,omitempty" url:"locations,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListLocationsResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListLocationsResponse) GetLocations() []*Location {
	if l == nil {
		return nil
	}
	return l.Locations
}

func (l *ListLocationsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListLocationsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListLocationsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListLocationsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListLocationsResponse) String() string {
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

// Represents one of a business' [locations](https://developer.squareup.com/docs/locations-api).
type Location struct {
	// A short generated string of letters and numbers that uniquely identifies this location instance.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The name of the location.
	// This information appears in the Seller Dashboard as the nickname.
	// A location name must be unique within a seller account.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The physical address of the location.
	Address *Address `json:"address,omitempty" url:"address,omitempty"`
	// The [IANA time zone](https://www.iana.org/time-zones) identifier for
	// the time zone of the location. For example, `America/Los_Angeles`.
	Timezone *string `json:"timezone,omitempty" url:"timezone,omitempty"`
	// The Square features that are enabled for the location.
	// See [LocationCapability](entity:LocationCapability) for possible values.
	// See [LocationCapability](#type-locationcapability) for possible values
	Capabilities []LocationCapability `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	// The status of the location.
	// See [LocationStatus](#type-locationstatus) for possible values
	Status *LocationStatus `json:"status,omitempty" url:"status,omitempty"`
	// The time when the location was created, in RFC 3339 format.
	// For more information, see [Working with Dates](https://developer.squareup.com/docs/build-basics/working-with-dates).
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The ID of the merchant that owns the location.
	MerchantID *string `json:"merchant_id,omitempty" url:"merchant_id,omitempty"`
	// The country of the location, in the two-letter format of ISO 3166. For example, `US` or `JP`.
	//
	// See [Country](entity:Country) for possible values.
	// See [Country](#type-country) for possible values
	Country *Country `json:"country,omitempty" url:"country,omitempty"`
	// The language associated with the location, in
	// [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A).
	// For more information, see [Language Preferences](https://developer.squareup.com/docs/build-basics/general-considerations/language-preferences).
	LanguageCode *string `json:"language_code,omitempty" url:"language_code,omitempty"`
	// The currency used for all transactions at this location,
	// in ISO 4217 format. For example, the currency code for US dollars is `USD`.
	// See [Currency](entity:Currency) for possible values.
	// See [Currency](#type-currency) for possible values
	Currency *Currency `json:"currency,omitempty" url:"currency,omitempty"`
	// The phone number of the location. For example, `+1 855-700-6000`.
	PhoneNumber *string `json:"phone_number,omitempty" url:"phone_number,omitempty"`
	// The name of the location's overall business. This name is present on receipts and other customer-facing branding, and can be changed no more than three times in a twelve-month period.
	BusinessName *string `json:"business_name,omitempty" url:"business_name,omitempty"`
	// The type of the location.
	// See [LocationType](#type-locationtype) for possible values
	Type *LocationType `json:"type,omitempty" url:"type,omitempty"`
	// The website URL of the location.  For example, `https://squareup.com`.
	WebsiteURL *string `json:"website_url,omitempty" url:"website_url,omitempty"`
	// The hours of operation for the location.
	BusinessHours *BusinessHours `json:"business_hours,omitempty" url:"business_hours,omitempty"`
	// The email address of the location. This can be unique to the location and is not always the email address for the business owner or administrator.
	BusinessEmail *string `json:"business_email,omitempty" url:"business_email,omitempty"`
	// The description of the location. For example, `Main Street location`.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// The Twitter username of the location without the '@' symbol. For example, `Square`.
	TwitterUsername *string `json:"twitter_username,omitempty" url:"twitter_username,omitempty"`
	// The Instagram username of the location without the '@' symbol. For example, `square`.
	InstagramUsername *string `json:"instagram_username,omitempty" url:"instagram_username,omitempty"`
	// The Facebook profile URL of the location. The URL should begin with 'facebook.com/'. For example, `https://www.facebook.com/square`.
	FacebookURL *string `json:"facebook_url,omitempty" url:"facebook_url,omitempty"`
	// The physical coordinates (latitude and longitude) of the location.
	Coordinates *Coordinates `json:"coordinates,omitempty" url:"coordinates,omitempty"`
	// The URL of the logo image for the location. When configured in the Seller
	// Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller.
	// This image should have a roughly square (1:1) aspect ratio and should be at least 200x200 pixels.
	LogoURL *string `json:"logo_url,omitempty" url:"logo_url,omitempty"`
	// The URL of the Point of Sale background image for the location.
	PosBackgroundURL *string `json:"pos_background_url,omitempty" url:"pos_background_url,omitempty"`
	// A four-digit number that describes the kind of goods or services sold at the location.
	// The [merchant category code (MCC)](https://developer.squareup.com/docs/locations-api#initialize-a-merchant-category-code) of the location as standardized by ISO 18245.
	// For example, `5045`, for a location that sells computer goods and software.
	Mcc *string `json:"mcc,omitempty" url:"mcc,omitempty"`
	// The URL of a full-format logo image for the location. When configured in the Seller
	// Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller.
	// This image can be wider than it is tall and should be at least 1280x648 pixels.
	FullFormatLogoURL *string `json:"full_format_logo_url,omitempty" url:"full_format_logo_url,omitempty"`
	// The tax IDs for this location.
	TaxIDs *TaxIDs `json:"tax_ids,omitempty" url:"tax_ids,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *Location) GetID() *string {
	if l == nil {
		return nil
	}
	return l.ID
}

func (l *Location) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *Location) GetAddress() *Address {
	if l == nil {
		return nil
	}
	return l.Address
}

func (l *Location) GetTimezone() *string {
	if l == nil {
		return nil
	}
	return l.Timezone
}

func (l *Location) GetCapabilities() []LocationCapability {
	if l == nil {
		return nil
	}
	return l.Capabilities
}

func (l *Location) GetStatus() *LocationStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *Location) GetCreatedAt() *string {
	if l == nil {
		return nil
	}
	return l.CreatedAt
}

func (l *Location) GetMerchantID() *string {
	if l == nil {
		return nil
	}
	return l.MerchantID
}

func (l *Location) GetCountry() *Country {
	if l == nil {
		return nil
	}
	return l.Country
}

func (l *Location) GetLanguageCode() *string {
	if l == nil {
		return nil
	}
	return l.LanguageCode
}

func (l *Location) GetCurrency() *Currency {
	if l == nil {
		return nil
	}
	return l.Currency
}

func (l *Location) GetPhoneNumber() *string {
	if l == nil {
		return nil
	}
	return l.PhoneNumber
}

func (l *Location) GetBusinessName() *string {
	if l == nil {
		return nil
	}
	return l.BusinessName
}

func (l *Location) GetType() *LocationType {
	if l == nil {
		return nil
	}
	return l.Type
}

func (l *Location) GetWebsiteURL() *string {
	if l == nil {
		return nil
	}
	return l.WebsiteURL
}

func (l *Location) GetBusinessHours() *BusinessHours {
	if l == nil {
		return nil
	}
	return l.BusinessHours
}

func (l *Location) GetBusinessEmail() *string {
	if l == nil {
		return nil
	}
	return l.BusinessEmail
}

func (l *Location) GetDescription() *string {
	if l == nil {
		return nil
	}
	return l.Description
}

func (l *Location) GetTwitterUsername() *string {
	if l == nil {
		return nil
	}
	return l.TwitterUsername
}

func (l *Location) GetInstagramUsername() *string {
	if l == nil {
		return nil
	}
	return l.InstagramUsername
}

func (l *Location) GetFacebookURL() *string {
	if l == nil {
		return nil
	}
	return l.FacebookURL
}

func (l *Location) GetCoordinates() *Coordinates {
	if l == nil {
		return nil
	}
	return l.Coordinates
}

func (l *Location) GetLogoURL() *string {
	if l == nil {
		return nil
	}
	return l.LogoURL
}

func (l *Location) GetPosBackgroundURL() *string {
	if l == nil {
		return nil
	}
	return l.PosBackgroundURL
}

func (l *Location) GetMcc() *string {
	if l == nil {
		return nil
	}
	return l.Mcc
}

func (l *Location) GetFullFormatLogoURL() *string {
	if l == nil {
		return nil
	}
	return l.FullFormatLogoURL
}

func (l *Location) GetTaxIDs() *TaxIDs {
	if l == nil {
		return nil
	}
	return l.TaxIDs
}

func (l *Location) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *Location) UnmarshalJSON(data []byte) error {
	type unmarshaler Location
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = Location(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *Location) String() string {
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

// The capabilities a location might have.
type LocationCapability string

const (
	LocationCapabilityCreditCardProcessing LocationCapability = "CREDIT_CARD_PROCESSING"
	LocationCapabilityAutomaticTransfers   LocationCapability = "AUTOMATIC_TRANSFERS"
	LocationCapabilityUnlinkedRefunds      LocationCapability = "UNLINKED_REFUNDS"
)

func NewLocationCapabilityFromString(s string) (LocationCapability, error) {
	switch s {
	case "CREDIT_CARD_PROCESSING":
		return LocationCapabilityCreditCardProcessing, nil
	case "AUTOMATIC_TRANSFERS":
		return LocationCapabilityAutomaticTransfers, nil
	case "UNLINKED_REFUNDS":
		return LocationCapabilityUnlinkedRefunds, nil
	}
	var t LocationCapability
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationCapability) Ptr() *LocationCapability {
	return &l
}

// A location's status.
type LocationStatus string

const (
	LocationStatusActive   LocationStatus = "ACTIVE"
	LocationStatusInactive LocationStatus = "INACTIVE"
)

func NewLocationStatusFromString(s string) (LocationStatus, error) {
	switch s {
	case "ACTIVE":
		return LocationStatusActive, nil
	case "INACTIVE":
		return LocationStatusInactive, nil
	}
	var t LocationStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationStatus) Ptr() *LocationStatus {
	return &l
}

// A location's type.
type LocationType string

const (
	LocationTypePhysical LocationType = "PHYSICAL"
	LocationTypeMobile   LocationType = "MOBILE"
)

func NewLocationTypeFromString(s string) (LocationType, error) {
	switch s {
	case "PHYSICAL":
		return LocationTypePhysical, nil
	case "MOBILE":
		return LocationTypeMobile, nil
	}
	var t LocationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationType) Ptr() *LocationType {
	return &l
}

// Identifiers for the location used by various governments for tax purposes.
type TaxIDs struct {
	// The EU VAT number for this location. For example, `IE3426675K`.
	// If the EU VAT number is present, it is well-formed and has been
	// validated with VIES, the VAT Information Exchange System.
	EuVat *string `json:"eu_vat,omitempty" url:"eu_vat,omitempty"`
	// The SIRET (Système d'Identification du Répertoire des Entreprises et de leurs Etablissements)
	// number is a 14-digit code issued by the French INSEE. For example, `39922799000021`.
	FrSiret *string `json:"fr_siret,omitempty" url:"fr_siret,omitempty"`
	// The French government uses the NAF (Nomenclature des Activités Françaises) to display and
	// track economic statistical data. This is also called the APE (Activite Principale de l’Entreprise) code.
	// For example, `6910Z`.
	FrNaf *string `json:"fr_naf,omitempty" url:"fr_naf,omitempty"`
	// The NIF (Numero de Identificacion Fiscal) number is a nine-character tax identifier used in Spain.
	// If it is present, it has been validated. For example, `73628495A`.
	EsNif *string `json:"es_nif,omitempty" url:"es_nif,omitempty"`
	// The QII (Qualified Invoice Issuer) number is a 14-character tax identifier used in Japan.
	// For example, `T1234567890123`.
	JpQii *string `json:"jp_qii,omitempty" url:"jp_qii,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaxIDs) GetEuVat() *string {
	if t == nil {
		return nil
	}
	return t.EuVat
}

func (t *TaxIDs) GetFrSiret() *string {
	if t == nil {
		return nil
	}
	return t.FrSiret
}

func (t *TaxIDs) GetFrNaf() *string {
	if t == nil {
		return nil
	}
	return t.FrNaf
}

func (t *TaxIDs) GetEsNif() *string {
	if t == nil {
		return nil
	}
	return t.EsNif
}

func (t *TaxIDs) GetJpQii() *string {
	if t == nil {
		return nil
	}
	return t.JpQii
}

func (t *TaxIDs) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TaxIDs) UnmarshalJSON(data []byte) error {
	type unmarshaler TaxIDs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaxIDs(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaxIDs) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

// The response object returned by the [UpdateLocation](api-endpoint:Locations-UpdateLocation) endpoint.
type UpdateLocationResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The updated `Location` object.
	Location *Location `json:"location,omitempty" url:"location,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateLocationResponse) GetErrors() []*Error {
	if u == nil {
		return nil
	}
	return u.Errors
}

func (u *UpdateLocationResponse) GetLocation() *Location {
	if u == nil {
		return nil
	}
	return u.Location
}

func (u *UpdateLocationResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateLocationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateLocationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateLocationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateLocationResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateLocationRequest struct {
	// The ID of the location to update.
	LocationID string `json:"-" url:"-"`
	// The `Location` object with only the fields to update.
	Location *Location `json:"location,omitempty" url:"-"`
}
