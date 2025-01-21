// This file was auto-generated by Fern from our API Definition.

package client

import (
	applepay "github.com/square/square-go-sdk/v40/applepay"
	bankaccounts "github.com/square/square-go-sdk/v40/bankaccounts"
	bookingsclient "github.com/square/square-go-sdk/v40/bookings/client"
	cards "github.com/square/square-go-sdk/v40/cards"
	cashdrawersclient "github.com/square/square-go-sdk/v40/cashdrawers/client"
	catalogclient "github.com/square/square-go-sdk/v40/catalog/client"
	checkoutclient "github.com/square/square-go-sdk/v40/checkout/client"
	core "github.com/square/square-go-sdk/v40/core"
	customersclient "github.com/square/square-go-sdk/v40/customers/client"
	devicesclient "github.com/square/square-go-sdk/v40/devices/client"
	disputesclient "github.com/square/square-go-sdk/v40/disputes/client"
	employees "github.com/square/square-go-sdk/v40/employees"
	events "github.com/square/square-go-sdk/v40/events"
	giftcardsclient "github.com/square/square-go-sdk/v40/giftcards/client"
	internal "github.com/square/square-go-sdk/v40/internal"
	inventory "github.com/square/square-go-sdk/v40/inventory"
	invoices "github.com/square/square-go-sdk/v40/invoices"
	laborclient "github.com/square/square-go-sdk/v40/labor/client"
	locationsclient "github.com/square/square-go-sdk/v40/locations/client"
	loyaltyclient "github.com/square/square-go-sdk/v40/loyalty/client"
	merchantsclient "github.com/square/square-go-sdk/v40/merchants/client"
	mobile "github.com/square/square-go-sdk/v40/mobile"
	oauth "github.com/square/square-go-sdk/v40/oauth"
	option "github.com/square/square-go-sdk/v40/option"
	ordersclient "github.com/square/square-go-sdk/v40/orders/client"
	payments "github.com/square/square-go-sdk/v40/payments"
	payouts "github.com/square/square-go-sdk/v40/payouts"
	refunds "github.com/square/square-go-sdk/v40/refunds"
	sites "github.com/square/square-go-sdk/v40/sites"
	snippets "github.com/square/square-go-sdk/v40/snippets"
	subscriptions "github.com/square/square-go-sdk/v40/subscriptions"
	team "github.com/square/square-go-sdk/v40/team"
	teammembersclient "github.com/square/square-go-sdk/v40/teammembers/client"
	terminalclient "github.com/square/square-go-sdk/v40/terminal/client"
	v1transactions "github.com/square/square-go-sdk/v40/v1transactions"
	vendors "github.com/square/square-go-sdk/v40/vendors"
	webhooksclient "github.com/square/square-go-sdk/v40/webhooks/client"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Mobile         *mobile.Client
	OAuth          *oauth.Client
	V1Transactions *v1transactions.Client
	ApplePay       *applepay.Client
	BankAccounts   *bankaccounts.Client
	Bookings       *bookingsclient.Client
	Cards          *cards.Client
	Catalog        *catalogclient.Client
	Customers      *customersclient.Client
	Devices        *devicesclient.Client
	Disputes       *disputesclient.Client
	Employees      *employees.Client
	Events         *events.Client
	GiftCards      *giftcardsclient.Client
	Inventory      *inventory.Client
	Invoices       *invoices.Client
	Locations      *locationsclient.Client
	Loyalty        *loyaltyclient.Client
	Merchants      *merchantsclient.Client
	Checkout       *checkoutclient.Client
	Orders         *ordersclient.Client
	Payments       *payments.Client
	Payouts        *payouts.Client
	Refunds        *refunds.Client
	Sites          *sites.Client
	Snippets       *snippets.Client
	Subscriptions  *subscriptions.Client
	TeamMembers    *teammembersclient.Client
	Team           *team.Client
	Terminal       *terminalclient.Client
	Vendors        *vendors.Client
	CashDrawers    *cashdrawersclient.Client
	Labor          *laborclient.Client
	Webhooks       *webhooksclient.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.Token == "" {
		options.Token = os.Getenv("SQUARE_TOKEN")
	}
	if options.Version == "" {
		options.Version = os.Getenv("VERSION")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:         options.ToHeader(),
		Mobile:         mobile.NewClient(opts...),
		OAuth:          oauth.NewClient(opts...),
		V1Transactions: v1transactions.NewClient(opts...),
		ApplePay:       applepay.NewClient(opts...),
		BankAccounts:   bankaccounts.NewClient(opts...),
		Bookings:       bookingsclient.NewClient(opts...),
		Cards:          cards.NewClient(opts...),
		Catalog:        catalogclient.NewClient(opts...),
		Customers:      customersclient.NewClient(opts...),
		Devices:        devicesclient.NewClient(opts...),
		Disputes:       disputesclient.NewClient(opts...),
		Employees:      employees.NewClient(opts...),
		Events:         events.NewClient(opts...),
		GiftCards:      giftcardsclient.NewClient(opts...),
		Inventory:      inventory.NewClient(opts...),
		Invoices:       invoices.NewClient(opts...),
		Locations:      locationsclient.NewClient(opts...),
		Loyalty:        loyaltyclient.NewClient(opts...),
		Merchants:      merchantsclient.NewClient(opts...),
		Checkout:       checkoutclient.NewClient(opts...),
		Orders:         ordersclient.NewClient(opts...),
		Payments:       payments.NewClient(opts...),
		Payouts:        payouts.NewClient(opts...),
		Refunds:        refunds.NewClient(opts...),
		Sites:          sites.NewClient(opts...),
		Snippets:       snippets.NewClient(opts...),
		Subscriptions:  subscriptions.NewClient(opts...),
		TeamMembers:    teammembersclient.NewClient(opts...),
		Team:           team.NewClient(opts...),
		Terminal:       terminalclient.NewClient(opts...),
		Vendors:        vendors.NewClient(opts...),
		CashDrawers:    cashdrawersclient.NewClient(opts...),
		Labor:          laborclient.NewClient(opts...),
		Webhooks:       webhooksclient.NewClient(opts...),
	}
}
