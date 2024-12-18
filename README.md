# Square Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/square/square-go-sdk)

The Square Go library provides convenient access to the Square API from Go.

## Requirements

This module requires Go version >= 1.18.

# Installation

Run the following command to use the Square Go library in your module:

```sh
go get github.com/square/square-go-sdk
```

## Usage

```go
package main

import (
    "github.com/square/square-go-sdk"
    squareclient "github.com/square/square-go-sdk/client"
    "github.com/square/square-go-sdk/option"

    "context"
    "fmt"
)


func main() {
	client := squareclient.NewClient(
		option.WithToken("<YOUR_ACCESS_TOKEN>"),
	)
	
	response, err := client.Payments.Create(
		context.TODO(),
		&square.CreatePaymentRequest{
			IdempotencyKey: "4935a656-a929-4792-b97c-8848be85c27c",
			SourceID:       "CASH",
			AmountMoney: &square.Money{
				Amount:   square.Int64(100),
				Currency: square.CurrencyUsd.Ptr(),
			},
			TipMoney: &square.Money{
				Amount:   square.Int64(50),
				Currency: square.CurrencyUsd.Ptr(),
			},
			CashDetails: &square.CashPaymentDetails{
				BuyerSuppliedMoney: &square.Money{
					Amount:   square.Int64(200),
					Currency: square.CurrencyUsd.Ptr(),
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.Payment)
}
```

## Optional Parameters

This library models optional primitives and enum types as pointers. This is primarily meant to distinguish
default zero values from explicit values (e.g. `false` for `bool` and `""` for `string`). A collection of
helper functions are provided to easily map a primitive or enum to its pointer-equivalent (e.g. `square.String`).

For example, consider the `client.Payments.List` endpoint usage below:

```go
response, err := client.Payments.List(
    context.TODO(),
    &square.PaymentsListRequest{
        Total: square.Int64(100),
    },
)
```

## Automatic Pagination

List endpoints are paginated. The SDK provides an iterator so that you can simply loop over the items:

```go
ctx := context.TODO()
page, err := client.Payments.List(
    ctx,
    &square.PaymentsListRequest{
        Total: square.Int64(100),
    },
)
if err != nil {
    return nil, err
}
iter := page.Iterator()
for iter.Next(ctx) {
    payment := iter.Current()
    fmt.Printf("Got payment: %v\n", *payment.ID)
}
if err := iter.Err(); err != nil {
    // Handle the error!
}
```

You can also iterate page-by-page:

```go
for page != nil {
    for _, payment := range page.Results {
        fmt.Printf("Got payment: %v\n", *payment.ID)
    }
    page, err = page.GetNextPage(ctx)
    if errors.Is(err, core.ErrNoPages) {
        break
    }
    if err != nil {
        // Handle the error!
    }
}
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.Payments.List(
    ctx,
    &square.PaymentsListRequest{
        Total: square.Int64(100),
    },
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to an unauthorized request (i.e. status code 401) with the following:

```go
response, err := client.Payments.Create(...)
if err != nil {
    if apiError, ok := err.(*core.APIError); ok {
        switch (apiError.StatusCode) {
            case http.StatusUnauthorized:
                // Do something with the unauthorized request ...
        }
    }
    return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.Payments.Create(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, apiError) {
        // Do something with the API error ...
    }
    return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability
to access the type with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.Payments.Create(...)
if err != nil {
    return fmt.Errorf("failed to create payment: %w", err)
}
```

## Webhook Signature Verification

The SDK provides a utility method that allow you to verify webhook signatures and ensure that all webhook events
originate from Square. The `client.Webhooks.VerifySignature` method will verify the signature:

```go
err := client.Webhooks.VerifySignature(
    context.TODO(),
    &square.VerifySignatureRequest{
        RequestBody: requestBody,
        SignatureHeader: header.Get("x-square-hmacsha256-signature"),
        SignatureKey: "YOUR_SIGNATURE_KEY",
        NotificationURL: "https://example.com/webhook", // The URL where event notifications are sent.
    },
);
if err != nil {
    return nil, err
}
```

## Advanced

### Request Options

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := squareclient.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)
```

These request options can either be specified on the client so that they're applied on _every_
request (shown above), or for an individual request like so:

```go
response, err := client.Payments.List(
    ctx,
    &square.PaymentsListRequest{
        Total: square.Int64(100),
    },
    option.WithToken("<YOUR_API_KEY>"),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

### Send Extra Properties

All endpoints support sending additional request body properties and query parameters that are not already
supported by the SDK. This is useful whenever you need to interact with an unreleased or hidden feature.

For example, suppose that a new feature was rolled out that allowed users to list all deactivated team members.
You could the relevant query parameters like so:

```go
response, err := client.TeamMembers.Search(
    context.TODO(),
    &square.SearchTeamMembersRequest{
        Limit: square.Int(100),
    },
    option.WithQueryParameters(
        url.Values{
            "status": []string{"DEACTIVATED"},
        },
    ),
)
```

### Receive Extra Properties

Every response type includes the `GetExtraProperties` method, which returns a map that contains any properties in the JSON
response that were not specified in the struct. Similar to the use case for sending additional parameters, this can be useful for
API features not present in the SDK yet.

You can receive and interact with the extra properties like so:

```go
response, err := client.Payments.Create(...)
if err != nil {
    return nil, err
}
extraProperties := response.GetExtraProperties()
```

### Retries

The Square Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := squareclient.NewClient(
    option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
response, err := client.Payments.List(
    context.TODO(),
    &square.PaymentsListRequest{
        Total: square.Int64(100),
    },
    option.WithMaxAttempts(1),
)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
