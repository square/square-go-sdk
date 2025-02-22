// This file was auto-generated by Fern from our API Definition.

package client

import (
	applepay "github.com/square/square-go-sdk/applepay"
	bankaccounts "github.com/square/square-go-sdk/bankaccounts"
	bookingsclient "github.com/square/square-go-sdk/bookings/client"
	cards "github.com/square/square-go-sdk/cards"
	cashdrawersclient "github.com/square/square-go-sdk/cashdrawers/client"
	catalogclient "github.com/square/square-go-sdk/catalog/client"
	checkoutclient "github.com/square/square-go-sdk/checkout/client"
	core "github.com/square/square-go-sdk/core"
	customersclient "github.com/square/square-go-sdk/customers/client"
	devicesclient "github.com/square/square-go-sdk/devices/client"
	disputesclient "github.com/square/square-go-sdk/disputes/client"
	employees "github.com/square/square-go-sdk/employees"
	events "github.com/square/square-go-sdk/events"
	giftcardsclient "github.com/square/square-go-sdk/giftcards/client"
	internal "github.com/square/square-go-sdk/internal"
	inventory "github.com/square/square-go-sdk/inventory"
	invoices "github.com/square/square-go-sdk/invoices"
	laborclient "github.com/square/square-go-sdk/labor/client"
	locationsclient "github.com/square/square-go-sdk/locations/client"
	loyaltyclient "github.com/square/square-go-sdk/loyalty/client"
	merchantsclient "github.com/square/square-go-sdk/merchants/client"
	mobile "github.com/square/square-go-sdk/mobile"
	oauth "github.com/square/square-go-sdk/oauth"
	option "github.com/square/square-go-sdk/option"
	ordersclient "github.com/square/square-go-sdk/orders/client"
	payments "github.com/square/square-go-sdk/payments"
	payouts "github.com/square/square-go-sdk/payouts"
	refunds "github.com/square/square-go-sdk/refunds"
	sites "github.com/square/square-go-sdk/sites"
	snippets "github.com/square/square-go-sdk/snippets"
	subscriptions "github.com/square/square-go-sdk/subscriptions"
	team "github.com/square/square-go-sdk/team"
	teammembersclient "github.com/square/square-go-sdk/teammembers/client"
	terminalclient "github.com/square/square-go-sdk/terminal/client"
	v1transactions "github.com/square/square-go-sdk/v1transactions"
	vendors "github.com/square/square-go-sdk/vendors"
	webhooksclient "github.com/square/square-go-sdk/webhooks/client"
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
