# Reference
## Mobile
<details><summary><code>client.Mobile.AuthorizationCode(request) -> *square.CreateMobileAuthorizationCodeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

__Note:__ This endpoint is used by the deprecated Reader SDK. 
Developers should update their integration to use the [Mobile Payments SDK](https://developer.squareup.com/docs/mobile-payments-sdk), which includes its own authorization methods. 

Generates code to authorize a mobile application to connect to a Square card reader.

Authorization codes are one-time-use codes and expire 60 minutes after being issued.

The `Authorization` header you provide to this endpoint must have the following format:

```
Authorization: Bearer ACCESS_TOKEN
```

Replace `ACCESS_TOKEN` with a
[valid production authorization credential](https://developer.squareup.com/docs/build-basics/access-tokens).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateMobileAuthorizationCodeRequest{
        LocationID: square.String(
            "YOUR_LOCATION_ID",
        ),
    }
client.Mobile.AuthorizationCode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `*string` — The Square location ID that the authorization code should be tied to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## OAuth
<details><summary><code>client.OAuth.RevokeToken(request) -> *square.RevokeTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revokes an access token generated with the OAuth flow.

If an account has more than one OAuth access token for your application, this
endpoint revokes all of them, regardless of which token you specify. 

__Important:__ The `Authorization` header for this endpoint must have the
following format:

```
Authorization: Client APPLICATION_SECRET
```

Replace `APPLICATION_SECRET` with the application secret on the **OAuth**
page for your application in the Developer Dashboard.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RevokeTokenRequest{
        ClientID: square.String(
            "CLIENT_ID",
        ),
        AccessToken: square.String(
            "ACCESS_TOKEN",
        ),
    }
client.OAuth.RevokeToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` 

The Square-issued ID for your application, which is available on the **OAuth** page in the
[Developer Dashboard](https://developer.squareup.com/apps).
    
</dd>
</dl>

<dl>
<dd>

**accessToken:** `*string` 

The access token of the merchant whose token you want to revoke.
Do not provide a value for `merchant_id` if you provide this parameter.
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `*string` 

The ID of the merchant whose token you want to revoke.
Do not provide a value for `access_token` if you provide this parameter.
    
</dd>
</dl>

<dl>
<dd>

**revokeOnlyAccessToken:** `*bool` 

If `true`, terminate the given single access token, but do not
terminate the entire authorization.
Default: `false`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.ObtainToken(request) -> *square.ObtainTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an OAuth access token and refresh token using the `authorization_code`
or `refresh_token` grant type.

When `grant_type` is `authorization_code`:
- With the [code flow](https://developer.squareup.com/docs/oauth-api/overview#code-flow),
provide `code`, `client_id`, and `client_secret`.
- With the [PKCE flow](https://developer.squareup.com/docs/oauth-api/overview#pkce-flow),
provide `code`, `client_id`, and `code_verifier`. 

When `grant_type` is `refresh_token`:
- With the code flow, provide `refresh_token`, `client_id`, and `client_secret`.
The response returns the same refresh token provided in the request.
- With the PKCE flow, provide `refresh_token` and `client_id`. The response returns
a new refresh token.

You can use the `scopes` parameter to limit the set of permissions authorized by the
access token. You can use the `short_lived` parameter to create an access token that
expires in 24 hours.

__Important:__ OAuth tokens should be encrypted and stored on a secure server.
Application clients should never interact directly with OAuth tokens.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ObtainTokenRequest{
        ClientID: "sq0idp-uaPHILoPzWZk3tlJqlML0g",
        ClientSecret: square.String(
            "sq0csp-30a-4C_tVOnTh14Piza2BfTPBXyLafLPWSzY1qAjeBfM",
        ),
        Code: square.String(
            "sq0cgb-l0SBqxs4uwxErTVyYOdemg",
        ),
        GrantType: "authorization_code",
    }
client.OAuth.ObtainToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 

The Square-issued ID of your application, which is available as the **Application ID**
on the **OAuth** page in the [Developer Console](https://developer.squareup.com/apps).

Required for the code flow and PKCE flow for any grant type.
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `*string` 

The secret key for your application, which is available as the **Application secret**
on the **OAuth** page in the [Developer Console](https://developer.squareup.com/apps).

Required for the code flow for any grant type. Don't confuse your client secret with your
personal access token.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` 

The authorization code to exchange for an OAuth access token. This is the `code`
value that Square sent to your redirect URL in the authorization response.

Required for the code flow and PKCE flow if `grant_type` is `authorization_code`.
    
</dd>
</dl>

<dl>
<dd>

**redirectURI:** `*string` 

The redirect URL for your application, which you registered as the **Redirect URL**
on the **OAuth** page in the [Developer Console](https://developer.squareup.com/apps).

Required for the code flow and PKCE flow if `grant_type` is `authorization_code` and
you provided the `redirect_uri` parameter in your authorization URL.
    
</dd>
</dl>

<dl>
<dd>

**grantType:** `string` 

The method used to obtain an OAuth access token. The request must include the
credential that corresponds to the specified grant type. Valid values are:
- `authorization_code` - Requires the `code` field.
- `refresh_token` - Requires the `refresh_token` field.
- `migration_token` - LEGACY for access tokens obtained using a Square API version prior
to 2019-03-13. Requires the `migration_token` field.
    
</dd>
</dl>

<dl>
<dd>

**refreshToken:** `*string` 

A valid refresh token used to generate a new OAuth access token. This is a
refresh token that was returned in a previous `ObtainToken` response.

Required for the code flow and PKCE flow if `grant_type` is `refresh_token`.
    
</dd>
</dl>

<dl>
<dd>

**migrationToken:** `*string` 

__LEGACY__ A valid access token (obtained using a Square API version prior to 2019-03-13)
used to generate a new OAuth access token.

Required if `grant_type` is `migration_token`. For more information, see
[Migrate to Using Refresh Tokens](https://developer.squareup.com/docs/oauth-api/migrate-to-refresh-tokens).
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]string` 

The list of permissions that are explicitly requested for the access token.
For example, ["MERCHANT_PROFILE_READ","PAYMENTS_READ","BANK_ACCOUNTS_READ"].

The returned access token is limited to the permissions that are the intersection
of these requested permissions and those authorized by the provided `refresh_token`.

Optional for the code flow and PKCE flow if `grant_type` is `refresh_token`.
    
</dd>
</dl>

<dl>
<dd>

**shortLived:** `*bool` 

Indicates whether the returned access token should expire in 24 hours.

Optional for the code flow and PKCE flow for any grant type. The default value is `false`.
    
</dd>
</dl>

<dl>
<dd>

**codeVerifier:** `*string` 

The secret your application generated for the authorization request used to
obtain the authorization code. This is the source of the `code_challenge` hash you
provided in your authorization URL.

Required for the PKCE flow if `grant_type` is `authorization_code`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.RetrieveTokenStatus() -> *square.RetrieveTokenStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns information about an [OAuth access token](https://developer.squareup.com/docs/build-basics/access-tokens#get-an-oauth-access-token) or an application’s [personal access token](https://developer.squareup.com/docs/build-basics/access-tokens#get-a-personal-access-token).

Add the access token to the Authorization header of the request.

__Important:__ The `Authorization` header you provide to this endpoint must have the following format:

```
Authorization: Bearer ACCESS_TOKEN
```

where `ACCESS_TOKEN` is a
[valid production authorization credential](https://developer.squareup.com/docs/build-basics/access-tokens).

If the access token is expired or not a valid access token, the endpoint returns an `UNAUTHORIZED` error.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.RetrieveTokenStatus(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.Authorize() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.Authorize(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## V1Transactions
<details><summary><code>client.V1Transactions.V1ListOrders(LocationID) -> []*square.V1Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides summary information for a merchant's online store orders.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.V1ListOrdersRequest{
        LocationID: "location_id",
        Order: square.SortOrderDesc.Ptr(),
        Limit: square.Int(
            1,
        ),
        BatchToken: square.String(
            "batch_token",
        ),
    }
client.V1Transactions.V1ListOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location to list online store orders for.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*square.SortOrder` — The order in which payments are listed in the response.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of payments to return in a single response. This value cannot exceed 200.
    
</dd>
</dl>

<dl>
<dd>

**batchToken:** `*string` 

A pagination cursor to retrieve the next set of results for your
original query to the endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.V1Transactions.V1RetrieveOrder(LocationID, OrderID) -> *square.V1Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides comprehensive information for a single online store order, including the order's history.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.V1RetrieveOrderRequest{
        LocationID: "location_id",
        OrderID: "order_id",
    }
client.V1Transactions.V1RetrieveOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the order's associated location.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The order's Square-issued ID. You obtain this value from Order objects returned by the List Orders endpoint
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.V1Transactions.V1UpdateOrder(LocationID, OrderID, request) -> *square.V1Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the details of an online store order. Every update you perform on an order corresponds to one of three actions:
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.V1UpdateOrderRequest{
        LocationID: "location_id",
        OrderID: "order_id",
        Action: square.V1UpdateOrderRequestActionComplete,
    }
client.V1Transactions.V1UpdateOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the order's associated location.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The order's Square-issued ID. You obtain this value from Order objects returned by the List Orders endpoint
    
</dd>
</dl>

<dl>
<dd>

**action:** `*square.V1UpdateOrderRequestAction` 

The action to perform on the order (COMPLETE, CANCEL, or REFUND).
See [V1UpdateOrderRequestAction](#type-v1updateorderrequestaction) for possible values
    
</dd>
</dl>

<dl>
<dd>

**shippedTrackingNumber:** `*string` — The tracking number of the shipment associated with the order. Only valid if action is COMPLETE.
    
</dd>
</dl>

<dl>
<dd>

**completedNote:** `*string` — A merchant-specified note about the completion of the order. Only valid if action is COMPLETE.
    
</dd>
</dl>

<dl>
<dd>

**refundedNote:** `*string` — A merchant-specified note about the refunding of the order. Only valid if action is REFUND.
    
</dd>
</dl>

<dl>
<dd>

**canceledNote:** `*string` — A merchant-specified note about the canceling of the order. Only valid if action is CANCEL.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay
<details><summary><code>client.ApplePay.RegisterDomain(request) -> *square.RegisterDomainResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Activates a domain for use with Apple Pay on the Web and Square. A validation
is performed on this domain by Apple to ensure that it is properly set up as
an Apple Pay enabled domain.

This endpoint provides an easy way for platform developers to bulk activate
Apple Pay on the Web with Square for merchants using their platform.

Note: You will need to host a valid domain verification file on your domain to support Apple Pay.  The
current version of this file is always available at https://app.squareup.com/digital-wallets/apple-pay/apple-developer-merchantid-domain-association,
and should be hosted at `.well_known/apple-developer-merchantid-domain-association` on your
domain.  This file is subject to change; we strongly recommend checking for updates regularly and avoiding
long-lived caches that might not keep in sync with the correct file version.

To learn more about the Web Payments SDK and how to add Apple Pay, see [Take an Apple Pay Payment](https://developer.squareup.com/docs/web-payments/apple-pay).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RegisterDomainRequest{
        DomainName: "example.com",
    }
client.ApplePay.RegisterDomain(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainName:** `string` — A domain name as described in RFC-1034 that will be registered with ApplePay.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BankAccounts
<details><summary><code>client.BankAccounts.List() -> *square.ListBankAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of [BankAccount](entity:BankAccount) objects linked to a Square account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListBankAccountsRequest{
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
        LocationID: square.String(
            "location_id",
        ),
    }
client.BankAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

The pagination cursor returned by a previous call to this endpoint.
Use it in the next `ListBankAccounts` request to retrieve the next set 
of results.

See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

Upper limit on the number of bank accounts to return in the response. 
Currently, 1000 is the largest supported limit. You can specify a limit 
of up to 1000 bank accounts. This is also the default limit.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

Location ID. You can specify this optional filter 
to retrieve only the linked bank accounts belonging to a specific location.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankAccounts.GetByV1ID(V1BankAccountID) -> *square.GetBankAccountByV1IDResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns details of a [BankAccount](entity:BankAccount) identified by V1 bank account ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetByV1IDBankAccountsRequest{
        V1BankAccountID: "v1_bank_account_id",
    }
client.BankAccounts.GetByV1ID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**v1BankAccountID:** `string` 

Connect V1 ID of the desired `BankAccount`. For more information, see 
[Retrieve a bank account by using an ID issued by V1 Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api#retrieve-a-bank-account-by-using-an-id-issued-by-v1-bank-accounts-api).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankAccounts.Get(BankAccountID) -> *square.GetBankAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns details of a [BankAccount](entity:BankAccount)
linked to a Square account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetBankAccountsRequest{
        BankAccountID: "bank_account_id",
    }
client.BankAccounts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bankAccountID:** `string` — Square-issued ID of the desired `BankAccount`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bookings
<details><summary><code>client.Bookings.List() -> *square.ListBookingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a collection of bookings.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListBookingsRequest{
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        CustomerID: square.String(
            "customer_id",
        ),
        TeamMemberID: square.String(
            "team_member_id",
        ),
        LocationID: square.String(
            "location_id",
        ),
        StartAtMin: square.String(
            "start_at_min",
        ),
        StartAtMax: square.String(
            "start_at_max",
        ),
    }
client.Bookings.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — The maximum number of results per page to return in a paged response.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` — The [customer](entity:Customer) for whom to retrieve bookings. If this is not set, bookings for all customers are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**teamMemberID:** `*string` — The team member for whom to retrieve bookings. If this is not set, bookings of all members are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` — The location for which to retrieve bookings. If this is not set, all locations' bookings are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**startAtMin:** `*string` — The RFC 3339 timestamp specifying the earliest of the start time. If this is not set, the current time is used.
    
</dd>
</dl>

<dl>
<dd>

**startAtMax:** `*string` — The RFC 3339 timestamp specifying the latest of the start time. If this is not set, the time of 31 days after `start_at_min` is used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.Create(request) -> *square.CreateBookingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a booking.

The required input must include the following:
- `Booking.location_id`
- `Booking.start_at`
- `Booking.AppointmentSegment.team_member_id`
- `Booking.AppointmentSegment.service_variation_id`
- `Booking.AppointmentSegment.service_variation_version`

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateBookingRequest{
        Booking: &square.Booking{},
    }
client.Bookings.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique key to make this request an idempotent operation.
    
</dd>
</dl>

<dl>
<dd>

**booking:** `*square.Booking` — The details of the booking to be created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.SearchAvailability(request) -> *square.SearchAvailabilityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for availabilities for booking.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchAvailabilityRequest{
        Query: &square.SearchAvailabilityQuery{
            Filter: &square.SearchAvailabilityFilter{
                StartAtRange: &square.TimeRange{},
            },
        },
    }
client.Bookings.SearchAvailability(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.SearchAvailabilityQuery` — Query conditions used to filter buyer-accessible booking availabilities.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.BulkRetrieveBookings(request) -> *square.BulkRetrieveBookingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Bulk-Retrieves a list of bookings by booking IDs.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkRetrieveBookingsRequest{
        BookingIDs: []string{
            "booking_ids",
        },
    }
client.Bookings.BulkRetrieveBookings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingIDs:** `[]string` — A non-empty list of [Booking](entity:Booking) IDs specifying bookings to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.GetBusinessProfile() -> *square.GetBusinessBookingProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a seller's booking profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bookings.GetBusinessProfile(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.RetrieveLocationBookingProfile(LocationID) -> *square.RetrieveLocationBookingProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a seller's location booking profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RetrieveLocationBookingProfileRequest{
        LocationID: "location_id",
    }
client.Bookings.RetrieveLocationBookingProfile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location to retrieve the booking profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.BulkRetrieveTeamMemberBookingProfiles(request) -> *square.BulkRetrieveTeamMemberBookingProfilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one or more team members' booking profiles.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkRetrieveTeamMemberBookingProfilesRequest{
        TeamMemberIDs: []string{
            "team_member_ids",
        },
    }
client.Bookings.BulkRetrieveTeamMemberBookingProfiles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberIDs:** `[]string` — A non-empty list of IDs of team members whose booking profiles you want to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.Get(BookingID) -> *square.GetBookingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a booking.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetBookingsRequest{
        BookingID: "booking_id",
    }
client.Bookings.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the [Booking](entity:Booking) object representing the to-be-retrieved booking.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.Update(BookingID, request) -> *square.UpdateBookingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a booking.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateBookingRequest{
        BookingID: "booking_id",
        Booking: &square.Booking{},
    }
client.Bookings.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the [Booking](entity:Booking) object representing the to-be-updated booking.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique key to make this request an idempotent operation.
    
</dd>
</dl>

<dl>
<dd>

**booking:** `*square.Booking` — The booking to be updated. Individual attributes explicitly specified here override the corresponding values of the existing booking.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.Cancel(BookingID, request) -> *square.CancelBookingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels an existing booking.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelBookingRequest{
        BookingID: "booking_id",
    }
client.Bookings.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the [Booking](entity:Booking) object representing the to-be-cancelled booking.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique key to make this request an idempotent operation.
    
</dd>
</dl>

<dl>
<dd>

**bookingVersion:** `*int` — The revision number for the booking used for optimistic concurrency.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Cards
<details><summary><code>client.Cards.List() -> *square.ListCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of cards owned by the account making the request.
A max of 25 cards will be returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListCardsRequest{
        Cursor: square.String(
            "cursor",
        ),
        CustomerID: square.String(
            "customer_id",
        ),
        IncludeDisabled: square.Bool(
            true,
        ),
        ReferenceID: square.String(
            "reference_id",
        ),
        SortOrder: square.SortOrderDesc.Ptr(),
    }
client.Cards.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for your original query.

See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` 

Limit results to cards associated with the customer supplied.
By default, all cards owned by the merchant are returned.
    
</dd>
</dl>

<dl>
<dd>

**includeDisabled:** `*bool` 

Includes disabled cards.
By default, all enabled cards owned by the merchant are returned.
    
</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` — Limit results to cards associated with the reference_id supplied.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

Sorts the returned list by when the card was created with the specified order.
This field defaults to ASC.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Create(request) -> *square.CreateCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a card on file to an existing merchant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateCardRequest{
        IdempotencyKey: "4935a656-a929-4792-b97c-8848be85c27c",
        SourceID: "cnon:uIbfJXhXETSP197M3GB",
        Card: &square.Card{
            CardholderName: square.String(
                "Amelia Earhart",
            ),
            BillingAddress: &square.Address{
                AddressLine1: square.String(
                    "500 Electric Ave",
                ),
                AddressLine2: square.String(
                    "Suite 600",
                ),
                Locality: square.String(
                    "New York",
                ),
                AdministrativeDistrictLevel1: square.String(
                    "NY",
                ),
                PostalCode: square.String(
                    "10003",
                ),
                Country: square.CountryUs.Ptr(),
            },
            CustomerID: square.String(
                "VDKXEEKPJN48QDG3BGGFAK05P8",
            ),
            ReferenceID: square.String(
                "user-id-1",
            ),
        },
    }
client.Cards.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this CreateCard request. Keys can be
any valid string and must be unique for every request.

Max: 45 characters

See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**sourceID:** `string` — The ID of the source which represents the card information to be stored. This can be a card nonce or a payment id.
    
</dd>
</dl>

<dl>
<dd>

**verificationToken:** `*string` 

An identifying token generated by [Payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer).
Verification tokens encapsulate customer device information and 3-D Secure
challenge results to indicate that Square has verified the buyer identity.

See the [SCA Overview](https://developer.squareup.com/docs/sca-overview).
    
</dd>
</dl>

<dl>
<dd>

**card:** `*square.Card` — Payment details associated with the card to be stored.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Get(CardID) -> *square.GetCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves details for a specific Card.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetCardsRequest{
        CardID: "card_id",
    }
client.Cards.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cardID:** `string` — Unique ID for the desired Card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Disable(CardID) -> *square.DisableCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disables the card, preventing any further updates or charges.
Disabling an already disabled card is allowed but has no effect.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DisableCardsRequest{
        CardID: "card_id",
    }
client.Cards.Disable(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cardID:** `string` — Unique ID for the desired Card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Catalog
<details><summary><code>client.Catalog.BatchDelete(request) -> *square.BatchDeleteCatalogObjectsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a set of [CatalogItem](entity:CatalogItem)s based on the
provided list of target IDs and returns a set of successfully deleted IDs in
the response. Deletion is a cascading event such that all children of the
targeted object are also deleted. For example, deleting a CatalogItem will
also delete all of its [CatalogItemVariation](entity:CatalogItemVariation)
children.

`BatchDeleteCatalogObjects` succeeds even if only a portion of the targeted
IDs can be deleted. The response will only include IDs that were
actually deleted.

To ensure consistency, only one delete request is processed at a time per seller account.
While one (batch or non-batch) delete request is being processed, other (batched and non-batched)
delete requests are rejected with the `429` error code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchDeleteCatalogObjectsRequest{
        ObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
            "AA27W3M2GGTF3H6AVPNB77CK",
        },
    }
client.Catalog.BatchDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**objectIDs:** `[]string` 

The IDs of the CatalogObjects to be deleted. When an object is deleted, other objects
in the graph that depend on that object will be deleted as well (for example, deleting a
CatalogItem will delete its CatalogItemVariation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.BatchGet(request) -> *square.BatchGetCatalogObjectsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a set of objects based on the provided ID.
Each [CatalogItem](entity:CatalogItem) returned in the set includes all of its
child information including: all of its
[CatalogItemVariation](entity:CatalogItemVariation) objects, references to
its [CatalogModifierList](entity:CatalogModifierList) objects, and the ids of
any [CatalogTax](entity:CatalogTax) objects that apply to it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchGetCatalogObjectsRequest{
        ObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
            "AA27W3M2GGTF3H6AVPNB77CK",
        },
        IncludeRelatedObjects: square.Bool(
            true,
        ),
    }
client.Catalog.BatchGet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**objectIDs:** `[]string` — The IDs of the CatalogObjects to be retrieved.
    
</dd>
</dl>

<dl>
<dd>

**includeRelatedObjects:** `*bool` 

If `true`, the response will include additional objects that are related to the
requested objects. Related objects are defined as any objects referenced by ID by the results in the `objects` field
of the response. These objects are put in the `related_objects` field. Setting this to `true` is
helpful when the objects are needed for immediate display to a user.
This process only goes one level deep. Objects referenced by the related objects will not be included. For example,

if the `objects` field of the response contains a CatalogItem, its associated
CatalogCategory objects, CatalogTax objects, CatalogImage objects and
CatalogModifierLists will be returned in the `related_objects` field of the
response. If the `objects` field of the response contains a CatalogItemVariation,
its parent CatalogItem will be returned in the `related_objects` field of
the response.

Default value: `false`
    
</dd>
</dl>

<dl>
<dd>

**catalogVersion:** `*int64` 

The specific version of the catalog objects to be included in the response. 
This allows you to retrieve historical versions of objects. The specified version value is matched against
the [CatalogObject](entity:CatalogObject)s' `version` attribute. If not included, results will
be from the current version of the catalog.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedObjects:** `*bool` — Indicates whether to include (`true`) or not (`false`) in the response deleted objects, namely, those with the `is_deleted` attribute set to `true`.
    
</dd>
</dl>

<dl>
<dd>

**includeCategoryPathToRoot:** `*bool` 

Specifies whether or not to include the `path_to_root` list for each returned category instance. The `path_to_root` list consists
of `CategoryPathToRootNode` objects and specifies the path that starts with the immediate parent category of the returned category
and ends with its root category. If the returned category is a top-level category, the `path_to_root` list is empty and is not returned
in the response payload.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.BatchUpsert(request) -> *square.BatchUpsertCatalogObjectsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates up to 10,000 target objects based on the provided
list of objects. The target objects are grouped into batches and each batch is
inserted/updated in an all-or-nothing manner. If an object within a batch is
malformed in some way, or violates a database constraint, the entire batch
containing that item will be disregarded. However, other batches in the same
request may still succeed. Each batch may contain up to 1,000 objects, and
batches will be processed in order as long as the total object count for the
request (items, variations, modifier lists, discounts, and taxes) is no more
than 10,000.

To ensure consistency, only one update request is processed at a time per seller account.
While one (batch or non-batch) update request is being processed, other (batched and non-batched)
update requests are rejected with the `429` error code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchUpsertCatalogObjectsRequest{
        IdempotencyKey: "789ff020-f723-43a9-b4b5-43b5dc1fa3dc",
        Batches: []*square.CatalogObjectBatch{
            &square.CatalogObjectBatch{
                Objects: []*square.CatalogObject{
                    &square.CatalogObject{
                        Item: &square.CatalogObjectItem{
                            ID: "id",
                        },
                    },
                    &square.CatalogObject{
                        Item: &square.CatalogObjectItem{
                            ID: "id",
                        },
                    },
                    &square.CatalogObject{
                        Item: &square.CatalogObjectItem{
                            ID: "id",
                        },
                    },
                    &square.CatalogObject{
                        Tax: &square.CatalogObjectTax{
                            ID: "id",
                        },
                    },
                },
            },
        },
    }
client.Catalog.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A value you specify that uniquely identifies this
request among all your requests. A common way to create
a valid idempotency key is to use a Universally unique
identifier (UUID).

If you're unsure whether a particular request was successful,
you can reattempt it with the same idempotency key without
worrying about creating duplicate objects.

See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**batches:** `[]*square.CatalogObjectBatch` 

A batch of CatalogObjects to be inserted/updated atomically.
The objects within a batch will be inserted in an all-or-nothing fashion, i.e., if an error occurs
attempting to insert or update an object within a batch, the entire batch will be rejected. However, an error
in one batch will not affect other batches within the same request.

For each object, its `updated_at` field is ignored and replaced with a current [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates), and its
`is_deleted` field must not be set to `true`.

To modify an existing object, supply its ID. To create a new object, use an ID starting
with `#`. These IDs may be used to create relationships between an object and attributes of
other objects that reference it. For example, you can create a CatalogItem with
ID `#ABC` and a CatalogItemVariation with its `item_id` attribute set to
`#ABC` in order to associate the CatalogItemVariation with its parent
CatalogItem.

Any `#`-prefixed IDs are valid only within a single atomic batch, and will be replaced by server-generated IDs.

Each batch may contain up to 1,000 objects. The total number of objects across all batches for a single request
may not exceed 10,000. If either of these limits is violated, an error will be returned and no objects will
be inserted or updated.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.Info() -> *square.CatalogInfoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves information about the Square Catalog API, such as batch size
limits that can be used by the `BatchUpsertCatalogObjects` endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Catalog.Info(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.List() -> *square.ListCatalogResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all [CatalogObject](entity:CatalogObject)s of the specified types in the catalog.

The `types` parameter is specified as a comma-separated list of the [CatalogObjectType](entity:CatalogObjectType) values,
for example, "`ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, `CATEGORY`, `DISCOUNT`, `TAX`, `IMAGE`".

__Important:__ ListCatalog does not return deleted catalog items. To retrieve
deleted catalog items, use [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects)
and set the `include_deleted_objects` attribute value to `true`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListCatalogRequest{
        Cursor: square.String(
            "cursor",
        ),
        Types: square.String(
            "types",
        ),
        CatalogVersion: square.Int64(
            1000000,
        ),
    }
client.Catalog.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

The pagination cursor returned in the previous response. Leave unset for an initial request.
The page size is currently set to be 100.
See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**types:** `*string` 

An optional case-insensitive, comma-separated list of object types to retrieve.

The valid values are defined in the [CatalogObjectType](entity:CatalogObjectType) enum, for example,
`ITEM`, `ITEM_VARIATION`, `CATEGORY`, `DISCOUNT`, `TAX`,
`MODIFIER`, `MODIFIER_LIST`, `IMAGE`, etc.

If this is unspecified, the operation returns objects of all the top level types at the version
of the Square API used to make the request. Object types that are nested onto other object types
are not included in the defaults.

At the current API version the default object types are:
ITEM, CATEGORY, TAX, DISCOUNT, MODIFIER_LIST, 
PRICING_RULE, PRODUCT_SET, TIME_PERIOD, MEASUREMENT_UNIT,
SUBSCRIPTION_PLAN, ITEM_OPTION, CUSTOM_ATTRIBUTE_DEFINITION, QUICK_AMOUNT_SETTINGS.
    
</dd>
</dl>

<dl>
<dd>

**catalogVersion:** `*int64` 

The specific version of the catalog objects to be included in the response.
This allows you to retrieve historical versions of objects. The specified version value is matched against
the [CatalogObject](entity:CatalogObject)s' `version` attribute.  If not included, results will be from the
current version of the catalog.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.Search(request) -> *square.SearchCatalogObjectsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for [CatalogObject](entity:CatalogObject) of any type by matching supported search attribute values,
excluding custom attribute values on items or item variations, against one or more of the specified query filters.

This (`SearchCatalogObjects`) endpoint differs from the [SearchCatalogItems](api-endpoint:Catalog-SearchCatalogItems)
endpoint in the following aspects:

- `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
- `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
- `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
- The both endpoints have different call conventions, including the query filter formats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchCatalogObjectsRequest{
        ObjectTypes: []square.CatalogObjectType{
            square.CatalogObjectTypeItem,
        },
        Query: &square.CatalogQuery{
            PrefixQuery: &square.CatalogQueryPrefix{
                AttributeName: "name",
                AttributePrefix: "tea",
            },
        },
        Limit: square.Int(
            100,
        ),
    }
client.Catalog.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

The pagination cursor returned in the previous response. Leave unset for an initial request.
See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**objectTypes:** `[]*square.CatalogObjectType` 

The desired set of object types to appear in the search results.

If this is unspecified, the operation returns objects of all the top level types at the version
of the Square API used to make the request. Object types that are nested onto other object types
are not included in the defaults.

At the current API version the default object types are:
ITEM, CATEGORY, TAX, DISCOUNT, MODIFIER_LIST, 
PRICING_RULE, PRODUCT_SET, TIME_PERIOD, MEASUREMENT_UNIT,
SUBSCRIPTION_PLAN, ITEM_OPTION, CUSTOM_ATTRIBUTE_DEFINITION, QUICK_AMOUNT_SETTINGS.

Note that if you wish for the query to return objects belonging to nested types (i.e., COMPONENT, IMAGE,
ITEM_OPTION_VAL, ITEM_VARIATION, or MODIFIER), you must explicitly include all the types of interest
in this field.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedObjects:** `*bool` — If `true`, deleted objects will be included in the results. Defaults to `false`. Deleted objects will have their `is_deleted` field set to `true`. If `include_deleted_objects` is `true`, then the `include_category_path_to_root` request parameter must be `false`. Both properties cannot be `true` at the same time.
    
</dd>
</dl>

<dl>
<dd>

**includeRelatedObjects:** `*bool` 

If `true`, the response will include additional objects that are related to the
requested objects. Related objects are objects that are referenced by object ID by the objects
in the response. This is helpful if the objects are being fetched for immediate display to a user.
This process only goes one level deep. Objects referenced by the related objects will not be included.
For example:

If the `objects` field of the response contains a CatalogItem, its associated
CatalogCategory objects, CatalogTax objects, CatalogImage objects and
CatalogModifierLists will be returned in the `related_objects` field of the
response. If the `objects` field of the response contains a CatalogItemVariation,
its parent CatalogItem will be returned in the `related_objects` field of
the response.

Default value: `false`
    
</dd>
</dl>

<dl>
<dd>

**beginTime:** `*string` 

Return objects modified after this [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates), in RFC 3339
format, e.g., `2016-09-04T23:59:33.123Z`. The timestamp is exclusive - objects with a
timestamp equal to `begin_time` will not be included in the response.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*square.CatalogQuery` — A query to be used to filter or sort the results. If no query is specified, the entire catalog will be returned.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

A limit on the number of results to be returned in a single page. The limit is advisory -
the implementation may return more or fewer results. If the supplied limit is negative, zero, or
is higher than the maximum limit of 1,000, it will be ignored.
    
</dd>
</dl>

<dl>
<dd>

**includeCategoryPathToRoot:** `*bool` — Specifies whether or not to include the `path_to_root` list for each returned category instance. The `path_to_root` list consists of `CategoryPathToRootNode` objects and specifies the path that starts with the immediate parent category of the returned category and ends with its root category. If the returned category is a top-level category, the `path_to_root` list is empty and is not returned in the response payload. If `include_category_path_to_root` is `true`, then the `include_deleted_objects` request parameter must be `false`. Both properties cannot be `true` at the same time.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.SearchItems(request) -> *square.SearchCatalogItemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for catalog items or item variations by matching supported search attribute values, including
custom attribute values, against one or more of the specified query filters.

This (`SearchCatalogItems`) endpoint differs from the [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects)
endpoint in the following aspects:

- `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
- `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
- `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
- The both endpoints use different call conventions, including the query filter formats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchCatalogItemsRequest{
        TextFilter: square.String(
            "red",
        ),
        CategoryIDs: []string{
            "WINE_CATEGORY_ID",
        },
        StockLevels: []square.SearchCatalogItemsRequestStockLevel{
            square.SearchCatalogItemsRequestStockLevelOut,
            square.SearchCatalogItemsRequestStockLevelLow,
        },
        EnabledLocationIDs: []string{
            "ATL_LOCATION_ID",
        },
        Limit: square.Int(
            100,
        ),
        SortOrder: square.SortOrderAsc.Ptr(),
        ProductTypes: []square.CatalogItemProductType{
            square.CatalogItemProductTypeRegular,
        },
        CustomAttributeFilters: []*square.CustomAttributeFilter{
            &square.CustomAttributeFilter{
                CustomAttributeDefinitionID: square.String(
                    "VEGAN_DEFINITION_ID",
                ),
                BoolFilter: square.Bool(
                    true,
                ),
            },
            &square.CustomAttributeFilter{
                CustomAttributeDefinitionID: square.String(
                    "BRAND_DEFINITION_ID",
                ),
                StringFilter: square.String(
                    "Dark Horse",
                ),
            },
            &square.CustomAttributeFilter{
                Key: square.String(
                    "VINTAGE",
                ),
                NumberFilter: &square.Range{
                    Min: square.String(
                        "min",
                    ),
                    Max: square.String(
                        "max",
                    ),
                },
            },
            &square.CustomAttributeFilter{
                CustomAttributeDefinitionID: square.String(
                    "VARIETAL_DEFINITION_ID",
                ),
            },
        },
    }
client.Catalog.SearchItems(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**textFilter:** `*string` 

The text filter expression to return items or item variations containing specified text in
the `name`, `description`, or `abbreviation` attribute value of an item, or in
the `name`, `sku`, or `upc` attribute value of an item variation.
    
</dd>
</dl>

<dl>
<dd>

**categoryIDs:** `[]string` — The category id query expression to return items containing the specified category IDs.
    
</dd>
</dl>

<dl>
<dd>

**stockLevels:** `[]*square.SearchCatalogItemsRequestStockLevel` 

The stock-level query expression to return item variations with the specified stock levels.
See [SearchCatalogItemsRequestStockLevel](#type-searchcatalogitemsrequeststocklevel) for possible values
    
</dd>
</dl>

<dl>
<dd>

**enabledLocationIDs:** `[]string` — The enabled-location query expression to return items and item variations having specified enabled locations.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination token, returned in the previous response, used to fetch the next batch of pending results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to return per page. The default value is 100.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

The order to sort the results by item names. The default sort order is ascending (`ASC`).
See [SortOrder](#type-sortorder) for possible values
    
</dd>
</dl>

<dl>
<dd>

**productTypes:** `[]*square.CatalogItemProductType` — The product types query expression to return items or item variations having the specified product types.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeFilters:** `[]*square.CustomAttributeFilter` 

The customer-attribute filter to return items or item variations matching the specified
custom attribute expressions. A maximum number of 10 custom attribute expressions are supported in
a single call to the [SearchCatalogItems](api-endpoint:Catalog-SearchCatalogItems) endpoint.
    
</dd>
</dl>

<dl>
<dd>

**archivedState:** `*square.ArchivedState` — The query filter to return not archived (`ARCHIVED_STATE_NOT_ARCHIVED`), archived (`ARCHIVED_STATE_ARCHIVED`), or either type (`ARCHIVED_STATE_ALL`) of items.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.UpdateItemModifierLists(request) -> *square.UpdateItemModifierListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the [CatalogModifierList](entity:CatalogModifierList) objects
that apply to the targeted [CatalogItem](entity:CatalogItem) without having
to perform an upsert on the entire item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateItemModifierListsRequest{
        ItemIDs: []string{
            "H42BRLUJ5KTZTTMPVSLFAACQ",
            "2JXOBJIHCWBQ4NZ3RIXQGJA6",
        },
        ModifierListsToEnable: []string{
            "H42BRLUJ5KTZTTMPVSLFAACQ",
            "2JXOBJIHCWBQ4NZ3RIXQGJA6",
        },
        ModifierListsToDisable: []string{
            "7WRC16CJZDVLSNDQ35PP6YAD",
        },
    }
client.Catalog.UpdateItemModifierLists(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**itemIDs:** `[]string` — The IDs of the catalog items associated with the CatalogModifierList objects being updated.
    
</dd>
</dl>

<dl>
<dd>

**modifierListsToEnable:** `[]string` 

The IDs of the CatalogModifierList objects to enable for the CatalogItem.
At least one of `modifier_lists_to_enable` or `modifier_lists_to_disable` must be specified.
    
</dd>
</dl>

<dl>
<dd>

**modifierListsToDisable:** `[]string` 

The IDs of the CatalogModifierList objects to disable for the CatalogItem.
At least one of `modifier_lists_to_enable` or `modifier_lists_to_disable` must be specified.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.UpdateItemTaxes(request) -> *square.UpdateItemTaxesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the [CatalogTax](entity:CatalogTax) objects that apply to the
targeted [CatalogItem](entity:CatalogItem) without having to perform an
upsert on the entire item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateItemTaxesRequest{
        ItemIDs: []string{
            "H42BRLUJ5KTZTTMPVSLFAACQ",
            "2JXOBJIHCWBQ4NZ3RIXQGJA6",
        },
        TaxesToEnable: []string{
            "4WRCNHCJZDVLSNDQ35PP6YAD",
        },
        TaxesToDisable: []string{
            "AQCEGCEBBQONINDOHRGZISEX",
        },
    }
client.Catalog.UpdateItemTaxes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**itemIDs:** `[]string` 

IDs for the CatalogItems associated with the CatalogTax objects being updated.
No more than 1,000 IDs may be provided.
    
</dd>
</dl>

<dl>
<dd>

**taxesToEnable:** `[]string` 

IDs of the CatalogTax objects to enable.
At least one of `taxes_to_enable` or `taxes_to_disable` must be specified.
    
</dd>
</dl>

<dl>
<dd>

**taxesToDisable:** `[]string` 

IDs of the CatalogTax objects to disable.
At least one of `taxes_to_enable` or `taxes_to_disable` must be specified.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Channels
<details><summary><code>client.Channels.List() -> *square.ListChannelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>


</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListChannelsRequest{
        ReferenceType: square.ReferenceTypeUnknownType.Ptr(),
        ReferenceID: square.String(
            "reference_id",
        ),
        Status: square.ChannelStatusActive.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Channels.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**referenceType:** `*square.ReferenceType` — Type of reference associated to channel
    
</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` — id of reference associated to channel
    
</dd>
</dl>

<dl>
<dd>

**status:** `*square.ChannelStatus` — Status of channel
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to fetch the next result
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

Maximum number of results to return.
When not provided the returned results will be cap at 100 channels.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Channels.BulkRetrieve(request) -> *square.BulkRetrieveChannelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>


</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkRetrieveChannelsRequest{
        ChannelIDs: []string{
            "CH_9C03D0B59",
            "CH_6X139B5MN",
            "NOT_EXISTING",
        },
    }
client.Channels.BulkRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**channelIDs:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Channels.Get(ChannelID) -> *square.RetrieveChannelResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>


</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetChannelsRequest{
        ChannelID: "channel_id",
    }
client.Channels.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**channelID:** `string` — A channel id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers
<details><summary><code>client.Customers.List() -> *square.ListCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists customer profiles associated with a Square account.

Under normal operating conditions, newly created or updated customer profiles become available
for the listing operation in well under 30 seconds. Occasionally, propagation of the new or updated
profiles can take closer to one minute or longer, especially during network incidents and outages.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListCustomersRequest{
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
        SortField: square.CustomerSortFieldDefault.Ptr(),
        SortOrder: square.SortOrderDesc.Ptr(),
        Count: square.Bool(
            true,
        ),
    }
client.Customers.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results.
If the specified limit is less than 1 or greater than 100, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 100.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*square.CustomerSortField` 

Indicates how customers should be sorted.

The default value is `DEFAULT`.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

Indicates whether customers should be sorted in ascending (`ASC`) or
descending (`DESC`) order.

The default value is `ASC`.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*bool` 

Indicates whether to return the total count of customers in the `count` field of the response.

The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Create(request) -> *square.CreateCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new customer for a business.

You must provide at least one of the following values in your request to this
endpoint:

- `given_name`
- `family_name`
- `company_name`
- `email_address`
- `phone_number`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateCustomerRequest{
        GivenName: square.String(
            "Amelia",
        ),
        FamilyName: square.String(
            "Earhart",
        ),
        EmailAddress: square.String(
            "Amelia.Earhart@example.com",
        ),
        Address: &square.Address{
            AddressLine1: square.String(
                "500 Electric Ave",
            ),
            AddressLine2: square.String(
                "Suite 600",
            ),
            Locality: square.String(
                "New York",
            ),
            AdministrativeDistrictLevel1: square.String(
                "NY",
            ),
            PostalCode: square.String(
                "10003",
            ),
            Country: square.CountryUs.Ptr(),
        },
        PhoneNumber: square.String(
            "+1-212-555-4240",
        ),
        ReferenceID: square.String(
            "YOUR_REFERENCE_ID",
        ),
        Note: square.String(
            "a customer",
        ),
    }
client.Customers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` 

The idempotency key for the request.	For more information, see
[Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**givenName:** `*string` 

The given name (that is, the first name) associated with the customer profile.

The maximum length for this value is 300 characters.
    
</dd>
</dl>

<dl>
<dd>

**familyName:** `*string` 

The family name (that is, the last name) associated with the customer profile.

The maximum length for this value is 300 characters.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `*string` 

A business name associated with the customer profile.

The maximum length for this value is 500 characters.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `*string` 

A nickname for the customer profile.

The maximum length for this value is 100 characters.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` 

The email address associated with the customer profile.

The maximum length for this value is 254 characters.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*square.Address` 

The physical address associated with the customer profile. For maximum length constraints, see 
[Customer addresses](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#address).
The `first_name` and `last_name` fields are ignored if they are present in the request.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `*string` 

The phone number associated with the customer profile. The phone number must be valid and can contain
9–16 digits, with an optional `+` prefix and country code. For more information, see
[Customer phone numbers](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#phone-number).
    
</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` 

An optional second ID used to associate the customer profile with an
entity in another system.

The maximum length for this value is 100 characters.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` — A custom note associated with the customer profile.
    
</dd>
</dl>

<dl>
<dd>

**birthday:** `*string` 

The birthday associated with the customer profile, in `YYYY-MM-DD` or `MM-DD` format. For example,
specify `1998-09-21` for September 21, 1998, or `09-21` for September 21. Birthdays are returned in `YYYY-MM-DD`
format, where `YYYY` is the specified birth year or `0000` if a birth year is not specified.
    
</dd>
</dl>

<dl>
<dd>

**taxIDs:** `*square.CustomerTaxIDs` 

The tax ID associated with the customer profile. This field is available only for customers of sellers
in EU countries or the United Kingdom. For more information,
see [Customer tax IDs](https://developer.squareup.com/docs/customers-api/what-it-does#customer-tax-ids).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.BatchCreate(request) -> *square.BulkCreateCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates multiple [customer profiles](entity:Customer) for a business.

This endpoint takes a map of individual create requests and returns a map of responses.

You must provide at least one of the following values in each create request:

- `given_name`
- `family_name`
- `company_name`
- `email_address`
- `phone_number`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkCreateCustomersRequest{
        Customers: map[string]*square.BulkCreateCustomerData{
            "8bb76c4f-e35d-4c5b-90de-1194cd9179f0": &square.BulkCreateCustomerData{
                GivenName: square.String(
                    "Amelia",
                ),
                FamilyName: square.String(
                    "Earhart",
                ),
                EmailAddress: square.String(
                    "Amelia.Earhart@example.com",
                ),
                Address: &square.Address{
                    AddressLine1: square.String(
                        "500 Electric Ave",
                    ),
                    AddressLine2: square.String(
                        "Suite 600",
                    ),
                    Locality: square.String(
                        "New York",
                    ),
                    AdministrativeDistrictLevel1: square.String(
                        "NY",
                    ),
                    PostalCode: square.String(
                        "10003",
                    ),
                    Country: square.CountryUs.Ptr(),
                },
                PhoneNumber: square.String(
                    "+1-212-555-4240",
                ),
                ReferenceID: square.String(
                    "YOUR_REFERENCE_ID",
                ),
                Note: square.String(
                    "a customer",
                ),
            },
            "d1689f23-b25d-4932-b2f0-aed00f5e2029": &square.BulkCreateCustomerData{
                GivenName: square.String(
                    "Marie",
                ),
                FamilyName: square.String(
                    "Curie",
                ),
                EmailAddress: square.String(
                    "Marie.Curie@example.com",
                ),
                Address: &square.Address{
                    AddressLine1: square.String(
                        "500 Electric Ave",
                    ),
                    AddressLine2: square.String(
                        "Suite 601",
                    ),
                    Locality: square.String(
                        "New York",
                    ),
                    AdministrativeDistrictLevel1: square.String(
                        "NY",
                    ),
                    PostalCode: square.String(
                        "10003",
                    ),
                    Country: square.CountryUs.Ptr(),
                },
                PhoneNumber: square.String(
                    "+1-212-444-4240",
                ),
                ReferenceID: square.String(
                    "YOUR_REFERENCE_ID",
                ),
                Note: square.String(
                    "another customer",
                ),
            },
        },
    }
client.Customers.BatchCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customers:** `map[string]*square.BulkCreateCustomerData` 

A map of 1 to 100 individual create requests, represented by `idempotency key: { customer data }`
key-value pairs.

Each key is an [idempotency key](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency)
that uniquely identifies the create request. Each value contains the customer data used to create the
customer profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.BulkDeleteCustomers(request) -> *square.BulkDeleteCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes multiple customer profiles.

The endpoint takes a list of customer IDs and returns a map of responses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkDeleteCustomersRequest{
        CustomerIDs: []string{
            "8DDA5NZVBZFGAX0V3HPF81HHE0",
            "N18CPRVXR5214XPBBA6BZQWF3C",
            "2GYD7WNXF7BJZW1PMGNXZ3Y8M8",
        },
    }
client.Customers.BulkDeleteCustomers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerIDs:** `[]string` — The IDs of the [customer profiles](entity:Customer) to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.BulkRetrieveCustomers(request) -> *square.BulkRetrieveCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves multiple customer profiles.

This endpoint takes a list of customer IDs and returns a map of responses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkRetrieveCustomersRequest{
        CustomerIDs: []string{
            "8DDA5NZVBZFGAX0V3HPF81HHE0",
            "N18CPRVXR5214XPBBA6BZQWF3C",
            "2GYD7WNXF7BJZW1PMGNXZ3Y8M8",
        },
    }
client.Customers.BulkRetrieveCustomers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerIDs:** `[]string` — The IDs of the [customer profiles](entity:Customer) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.BulkUpdateCustomers(request) -> *square.BulkUpdateCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates multiple customer profiles.

This endpoint takes a map of individual update requests and returns a map of responses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkUpdateCustomersRequest{
        Customers: map[string]*square.BulkUpdateCustomerData{
            "8DDA5NZVBZFGAX0V3HPF81HHE0": &square.BulkUpdateCustomerData{
                EmailAddress: square.String(
                    "New.Amelia.Earhart@example.com",
                ),
                Note: square.String(
                    "updated customer note",
                ),
                Version: square.Int64(
                    2,
                ),
            },
            "N18CPRVXR5214XPBBA6BZQWF3C": &square.BulkUpdateCustomerData{
                GivenName: square.String(
                    "Marie",
                ),
                FamilyName: square.String(
                    "Curie",
                ),
                Version: square.Int64(
                    0,
                ),
            },
        },
    }
client.Customers.BulkUpdateCustomers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customers:** `map[string]*square.BulkUpdateCustomerData` 

A map of 1 to 100 individual update requests, represented by `customer ID: { customer data }`
key-value pairs.

Each key is the ID of the [customer profile](entity:Customer) to update. To update a customer profile
that was created by merging existing profiles, provide the ID of the newly created profile.

Each value contains the updated customer data. Only new or changed fields are required. To add or
update a field, specify the new value. To remove a field, specify `null`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Search(request) -> *square.SearchCustomersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches the customer profiles associated with a Square account using one or more supported query filters.

Calling `SearchCustomers` without any explicit query filter returns all
customer profiles ordered alphabetically based on `given_name` and
`family_name`.

Under normal operating conditions, newly created or updated customer profiles become available
for the search operation in well under 30 seconds. Occasionally, propagation of the new or updated
profiles can take closer to one minute or longer, especially during network incidents and outages.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchCustomersRequest{
        Limit: square.Int64(
            2,
        ),
        Query: &square.CustomerQuery{
            Filter: &square.CustomerFilter{
                CreationSource: &square.CustomerCreationSourceFilter{
                    Values: []square.CustomerCreationSource{
                        square.CustomerCreationSourceThirdParty,
                    },
                    Rule: square.CustomerInclusionExclusionInclude.Ptr(),
                },
                CreatedAt: &square.TimeRange{
                    StartAt: square.String(
                        "2018-01-01T00:00:00-00:00",
                    ),
                    EndAt: square.String(
                        "2018-02-01T00:00:00-00:00",
                    ),
                },
                EmailAddress: &square.CustomerTextFilter{
                    Fuzzy: square.String(
                        "example.com",
                    ),
                },
                GroupIDs: &square.FilterValue{
                    All: []string{
                        "545AXB44B4XXWMVQ4W8SBT3HHF",
                    },
                },
            },
            Sort: &square.CustomerSort{
                Field: square.CustomerSortFieldCreatedAt.Ptr(),
                Order: square.SortOrderAsc.Ptr(),
            },
        },
    }
client.Customers.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

Include the pagination cursor in subsequent calls to this endpoint to retrieve
the next set of results associated with the original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int64` 

The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results.
If the specified limit is invalid, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 100.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**query:** `*square.CustomerQuery` 

The filtering and sorting criteria for the search request. If a query is not specified,
Square returns all customer profiles ordered alphabetically by `given_name` and `family_name`.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*bool` 

Indicates whether to return the total count of matching customers in the `count` field of the response.

The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Get(CustomerID) -> *square.GetCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns details for a single customer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetCustomersRequest{
        CustomerID: "customer_id",
    }
client.Customers.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Update(CustomerID, request) -> *square.UpdateCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a customer profile. This endpoint supports sparse updates, so only new or changed fields are required in the request.
To add or update a field, specify the new value. To remove a field, specify `null`.

To update a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateCustomerRequest{
        CustomerID: "customer_id",
        EmailAddress: square.String(
            "New.Amelia.Earhart@example.com",
        ),
        Note: square.String(
            "updated customer note",
        ),
        Version: square.Int64(
            2,
        ),
    }
client.Customers.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to update.
    
</dd>
</dl>

<dl>
<dd>

**givenName:** `*string` 

The given name (that is, the first name) associated with the customer profile.

The maximum length for this value is 300 characters.
    
</dd>
</dl>

<dl>
<dd>

**familyName:** `*string` 

The family name (that is, the last name) associated with the customer profile.

The maximum length for this value is 300 characters.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `*string` 

A business name associated with the customer profile.

The maximum length for this value is 500 characters.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `*string` 

A nickname for the customer profile.

The maximum length for this value is 100 characters.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` 

The email address associated with the customer profile.

The maximum length for this value is 254 characters.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*square.Address` 

The physical address associated with the customer profile. Only new or changed fields are required in the request.

For maximum length constraints, see [Customer addresses](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#address).
The `first_name` and `last_name` fields are ignored if they are present in the request.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `*string` 

The phone number associated with the customer profile. The phone number must be valid and can contain
9–16 digits, with an optional `+` prefix and country code. For more information, see
[Customer phone numbers](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#phone-number).
    
</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` 

An optional second ID used to associate the customer profile with an
entity in another system.

The maximum length for this value is 100 characters.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` — A custom note associated with the customer profile.
    
</dd>
</dl>

<dl>
<dd>

**birthday:** `*string` 

The birthday associated with the customer profile, in `YYYY-MM-DD` or `MM-DD` format. For example,
specify `1998-09-21` for September 21, 1998, or `09-21` for September 21. Birthdays are returned in `YYYY-MM-DD`
format, where `YYYY` is the specified birth year or `0000` if a birth year is not specified.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` 

The current version of the customer profile.

As a best practice, you should include this field to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control. For more information, see [Update a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#update-a-customer-profile).
    
</dd>
</dl>

<dl>
<dd>

**taxIDs:** `*square.CustomerTaxIDs` 

The tax ID associated with the customer profile. This field is available only for customers of sellers
in EU countries or the United Kingdom. For more information,
see [Customer tax IDs](https://developer.squareup.com/docs/customers-api/what-it-does#customer-tax-ids).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Delete(CustomerID) -> *square.DeleteCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a customer profile from a business.

To delete a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteCustomersRequest{
        CustomerID: "customer_id",
        Version: square.Int64(
            1000000,
        ),
    }
client.Customers.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to delete.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` 

The current version of the customer profile.

As a best practice, you should include this parameter to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control.  For more information, see [Delete a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#delete-customer-profile).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Devices
<details><summary><code>client.Devices.List() -> *square.ListDevicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List devices associated with the merchant. Currently, only Terminal API
devices are supported.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListDevicesRequest{
        Cursor: square.String(
            "cursor",
        ),
        SortOrder: square.SortOrderDesc.Ptr(),
        Limit: square.Int(
            1,
        ),
        LocationID: square.String(
            "location_id",
        ),
    }
client.Devices.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

The order in which results are listed.
- `ASC` - Oldest to newest.
- `DESC` - Newest to oldest (default).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of results to return in a single page.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` — If present, only returns devices at the target location.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Devices.Get(DeviceID) -> *square.GetDeviceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves Device with the associated `device_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetDevicesRequest{
        DeviceID: "device_id",
    }
client.Devices.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**deviceID:** `string` — The unique ID for the desired `Device`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Disputes
<details><summary><code>client.Disputes.List() -> *square.ListDisputesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of disputes associated with a particular account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListDisputesRequest{
        Cursor: square.String(
            "cursor",
        ),
        States: square.DisputeStateInquiryEvidenceRequired.Ptr(),
        LocationID: square.String(
            "location_id",
        ),
    }
client.Disputes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**states:** `*square.DisputeState` — The dispute states used to filter the result. If not specified, the endpoint returns all disputes.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

The ID of the location for which to return a list of disputes.
If not specified, the endpoint returns disputes associated with all locations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Get(DisputeID) -> *square.GetDisputeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns details about a specific dispute.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetDisputesRequest{
        DisputeID: "dispute_id",
    }
client.Disputes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute you want more details about.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Accept(DisputeID) -> *square.AcceptDisputeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Accepts the loss on a dispute. Square returns the disputed amount to the cardholder and
updates the dispute state to ACCEPTED.

Square debits the disputed amount from the seller’s Square account. If the Square account
does not have sufficient funds, Square debits the associated bank account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.AcceptDisputesRequest{
        DisputeID: "dispute_id",
    }
client.Disputes.Accept(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute you want to accept.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.CreateEvidenceFile(DisputeID, request) -> *square.CreateDisputeEvidenceFileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads a file to use as evidence in a dispute challenge. The endpoint accepts HTTP
multipart/form-data file uploads in HEIC, HEIF, JPEG, PDF, PNG, and TIFF formats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateEvidenceFileDisputesRequest{
        DisputeID: "dispute_id",
    }
client.Disputes.CreateEvidenceFile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute for which you want to upload evidence.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.CreateEvidenceText(DisputeID, request) -> *square.CreateDisputeEvidenceTextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads text to use as evidence for a dispute challenge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateDisputeEvidenceTextRequest{
        DisputeID: "dispute_id",
        IdempotencyKey: "ed3ee3933d946f1514d505d173c82648",
        EvidenceType: square.DisputeEvidenceTypeTrackingNumber.Ptr(),
        EvidenceText: "1Z8888888888888888",
    }
client.Disputes.CreateEvidenceText(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute for which you want to upload evidence.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key identifying the request. For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**evidenceType:** `*square.DisputeEvidenceType` 

The type of evidence you are uploading.
See [DisputeEvidenceType](#type-disputeevidencetype) for possible values
    
</dd>
</dl>

<dl>
<dd>

**evidenceText:** `string` — The evidence string.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.SubmitEvidence(DisputeID) -> *square.SubmitEvidenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submits evidence to the cardholder's bank.

The evidence submitted by this endpoint includes evidence uploaded
using the [CreateDisputeEvidenceFile](api-endpoint:Disputes-CreateDisputeEvidenceFile) and
[CreateDisputeEvidenceText](api-endpoint:Disputes-CreateDisputeEvidenceText) endpoints and
evidence automatically provided by Square, when available. Evidence cannot be removed from
a dispute after submission.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SubmitEvidenceDisputesRequest{
        DisputeID: "dispute_id",
    }
client.Disputes.SubmitEvidence(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute for which you want to submit evidence.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Employees
<details><summary><code>client.Employees.List() -> *square.ListEmployeesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>


</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListEmployeesRequest{
        LocationID: square.String(
            "location_id",
        ),
        Status: square.EmployeeStatusActive.Ptr(),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Employees.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `*string` — 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*square.EmployeeStatus` — Specifies the EmployeeStatus to filter the employee by.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of employees to be returned on each page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The token required to retrieve the specified page of results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Employees.Get(ID) -> *square.GetEmployeeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>


</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetEmployeesRequest{
        ID: "id",
    }
client.Employees.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — UUID for the employee that was requested.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Events
<details><summary><code>client.Events.SearchEvents(request) -> *square.SearchEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for Square API events that occur within a 28-day timeframe.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchEventsRequest{}
client.Events.SearchEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of events for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of events to return in a single page. The response might contain fewer events. The default value is 100, which is also the maximum allowed value.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).

Default: 100
    
</dd>
</dl>

<dl>
<dd>

**query:** `*square.SearchEventsQuery` — The filtering and sorting criteria for the search request. To retrieve additional pages using a cursor, you must use the original query.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.DisableEvents() -> *square.DisableEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disables events to prevent them from being searchable.
All events are disabled by default. You must enable events to make them searchable.
Disabling events for a specific time period prevents them from being searchable, even if you re-enable them later.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Events.DisableEvents(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.EnableEvents() -> *square.EnableEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enables events to make them searchable. Only events that occur while in the enabled state are searchable.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Events.EnableEvents(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.ListEventTypes() -> *square.ListEventTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all event types that you can subscribe to as webhooks or query using the Events API.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListEventTypesRequest{
        APIVersion: square.String(
            "api_version",
        ),
    }
client.Events.ListEventTypes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiVersion:** `*string` — The API version for which to list event types. Setting this field overrides the default version used by the application.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## GiftCards
<details><summary><code>client.GiftCards.List() -> *square.ListGiftCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all gift cards. You can specify optional filters to retrieve 
a subset of the gift cards. Results are sorted by `created_at` in ascending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListGiftCardsRequest{
        Type: square.String(
            "type",
        ),
        State: square.String(
            "state",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        CustomerID: square.String(
            "customer_id",
        ),
    }
client.GiftCards.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `*string` 

If a [type](entity:GiftCardType) is provided, the endpoint returns gift cards of the specified type.
Otherwise, the endpoint returns gift cards of all types.
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` 

If a [state](entity:GiftCardStatus) is provided, the endpoint returns the gift cards in the specified state.
Otherwise, the endpoint returns the gift cards of all states.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

If a limit is provided, the endpoint returns only the specified number of results per page.
The maximum value is 200. The default value is 30.
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
If a cursor is not provided, the endpoint returns the first page of the results. 
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` — If a customer ID is provided, the endpoint returns only the gift cards linked to the specified customer.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.Create(request) -> *square.CreateGiftCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a digital gift card or registers a physical (plastic) gift card. The resulting gift card
has a `PENDING` state. To activate a gift card so that it can be redeemed for purchases, call
[CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity) and create an `ACTIVATE`
activity with the initial balance. Alternatively, you can use [RefundPayment](api-endpoint:Refunds-RefundPayment)
to refund a payment to the new gift card.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateGiftCardRequest{
        IdempotencyKey: "NC9Tm69EjbjtConu",
        LocationID: "81FN9BNFZTKS4",
        GiftCard: &square.GiftCard{
            Type: square.GiftCardTypeDigital,
        },
    }
client.GiftCards.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique identifier for this request, used to ensure idempotency. For more information, 
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` 

The ID of the [location](entity:Location) where the gift card should be registered for 
reporting purposes. Gift cards can be redeemed at any of the seller's locations.
    
</dd>
</dl>

<dl>
<dd>

**giftCard:** `*square.GiftCard` 

The gift card to create. The `type` field is required for this request. The `gan_source` 
and `gan` fields are included as follows: 

To direct Square to generate a 16-digit GAN, omit `gan_source` and `gan`.

To provide a custom GAN, include `gan_source` and `gan`.
- For `gan_source`, specify `OTHER`. 
- For `gan`, provide a custom GAN containing 8 to 20 alphanumeric characters. The GAN must be 
unique for the seller and cannot start with the same bank identification number (BIN) as major 
credit cards. Do not use GANs that are easy to guess (such as 12345678) because they greatly 
increase the risk of fraud. It is the responsibility of the developer to ensure the security 
of their custom GANs. For more information, see 
[Custom GANs](https://developer.squareup.com/docs/gift-cards/using-gift-cards-api#custom-gans). 

To register an unused, physical gift card that the seller previously ordered from Square, 
include `gan` and provide the GAN that is printed on the gift card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.GetFromGan(request) -> *square.GetGiftCardFromGanResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a gift card using the gift card account number (GAN).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetGiftCardFromGanRequest{
        Gan: "7783320001001635",
    }
client.GiftCards.GetFromGan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**gan:** `string` 

The gift card account number (GAN) of the gift card to retrieve.
The maximum length of a GAN is 255 digits to account for third-party GANs that have been imported.
Square-issued gift cards have 16-digit GANs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.GetFromNonce(request) -> *square.GetGiftCardFromNonceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a gift card using a secure payment token that represents the gift card.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetGiftCardFromNonceRequest{
        Nonce: "cnon:7783322135245171",
    }
client.GiftCards.GetFromNonce(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**nonce:** `string` 

The payment token of the gift card to retrieve. Payment tokens are generated by the 
Web Payments SDK or In-App Payments SDK.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.LinkCustomer(GiftCardID, request) -> *square.LinkCustomerToGiftCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Links a customer to a gift card, which is also referred to as adding a card on file.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.LinkCustomerToGiftCardRequest{
        GiftCardID: "gift_card_id",
        CustomerID: "GKY0FZ3V717AH8Q2D821PNT2ZW",
    }
client.GiftCards.LinkCustomer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**giftCardID:** `string` — The ID of the gift card to be linked.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to link to the gift card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.UnlinkCustomer(GiftCardID, request) -> *square.UnlinkCustomerFromGiftCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unlinks a customer from a gift card, which is also referred to as removing a card on file.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UnlinkCustomerFromGiftCardRequest{
        GiftCardID: "gift_card_id",
        CustomerID: "GKY0FZ3V717AH8Q2D821PNT2ZW",
    }
client.GiftCards.UnlinkCustomer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**giftCardID:** `string` — The ID of the gift card to be unlinked.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to unlink from the gift card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.Get(ID) -> *square.GetGiftCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a gift card using the gift card ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetGiftCardsRequest{
        ID: "id",
    }
client.GiftCards.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the gift card to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Inventory
<details><summary><code>client.Inventory.DeprecatedGetAdjustment(AdjustmentID) -> *square.GetInventoryAdjustmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated version of [RetrieveInventoryAdjustment](api-endpoint:Inventory-RetrieveInventoryAdjustment) after the endpoint URL
is updated to conform to the standard convention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeprecatedGetAdjustmentInventoryRequest{
        AdjustmentID: "adjustment_id",
    }
client.Inventory.DeprecatedGetAdjustment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**adjustmentID:** `string` — ID of the [InventoryAdjustment](entity:InventoryAdjustment) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.GetAdjustment(AdjustmentID) -> *square.GetInventoryAdjustmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the [InventoryAdjustment](entity:InventoryAdjustment) object
with the provided `adjustment_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetAdjustmentInventoryRequest{
        AdjustmentID: "adjustment_id",
    }
client.Inventory.GetAdjustment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**adjustmentID:** `string` — ID of the [InventoryAdjustment](entity:InventoryAdjustment) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.DeprecatedBatchChange(request) -> *square.BatchChangeInventoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated version of [BatchChangeInventory](api-endpoint:Inventory-BatchChangeInventory) after the endpoint URL
is updated to conform to the standard convention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchChangeInventoryRequest{
        IdempotencyKey: "8fc6a5b0-9fe8-4b46-b46b-2ef95793abbe",
        Changes: []*square.InventoryChange{
            &square.InventoryChange{
                Type: square.InventoryChangeTypePhysicalCount.Ptr(),
                PhysicalCount: &square.InventoryPhysicalCount{
                    ReferenceID: square.String(
                        "1536bfbf-efed-48bf-b17d-a197141b2a92",
                    ),
                    CatalogObjectID: square.String(
                        "W62UWFY35CWMYGVWK6TWJDNI",
                    ),
                    State: square.InventoryStateInStock.Ptr(),
                    LocationID: square.String(
                        "C6W5YS5QM06F5",
                    ),
                    Quantity: square.String(
                        "53",
                    ),
                    TeamMemberID: square.String(
                        "LRK57NSQ5X7PUD05",
                    ),
                    OccurredAt: square.String(
                        "2016-11-16T22:25:24.878Z",
                    ),
                },
            },
        },
        IgnoreUnchangedCounts: square.Bool(
            true,
        ),
    }
client.Inventory.DeprecatedBatchChange(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchChangeInventoryRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.DeprecatedBatchGetChanges(request) -> *square.BatchGetInventoryChangesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated version of [BatchRetrieveInventoryChanges](api-endpoint:Inventory-BatchRetrieveInventoryChanges) after the endpoint URL
is updated to conform to the standard convention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchRetrieveInventoryChangesRequest{
        CatalogObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
        },
        LocationIDs: []string{
            "C6W5YS5QM06F5",
        },
        Types: []square.InventoryChangeType{
            square.InventoryChangeTypePhysicalCount,
        },
        States: []square.InventoryState{
            square.InventoryStateInStock,
        },
        UpdatedAfter: square.String(
            "2016-11-01T00:00:00.000Z",
        ),
        UpdatedBefore: square.String(
            "2016-12-01T00:00:00.000Z",
        ),
    }
client.Inventory.DeprecatedBatchGetChanges(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchRetrieveInventoryChangesRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.DeprecatedBatchGetCounts(request) -> *square.BatchGetInventoryCountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated version of [BatchRetrieveInventoryCounts](api-endpoint:Inventory-BatchRetrieveInventoryCounts) after the endpoint URL
is updated to conform to the standard convention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchGetInventoryCountsRequest{
        CatalogObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
        },
        LocationIDs: []string{
            "59TNP9SA8VGDA",
        },
        UpdatedAfter: square.String(
            "2016-11-16T00:00:00.000Z",
        ),
    }
client.Inventory.DeprecatedBatchGetCounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchGetInventoryCountsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.BatchCreateChanges(request) -> *square.BatchChangeInventoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Applies adjustments and counts to the provided item quantities.

On success: returns the current calculated counts for all objects
referenced in the request.
On failure: returns a list of related errors.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchChangeInventoryRequest{
        IdempotencyKey: "8fc6a5b0-9fe8-4b46-b46b-2ef95793abbe",
        Changes: []*square.InventoryChange{
            &square.InventoryChange{
                Type: square.InventoryChangeTypePhysicalCount.Ptr(),
                PhysicalCount: &square.InventoryPhysicalCount{
                    ReferenceID: square.String(
                        "1536bfbf-efed-48bf-b17d-a197141b2a92",
                    ),
                    CatalogObjectID: square.String(
                        "W62UWFY35CWMYGVWK6TWJDNI",
                    ),
                    State: square.InventoryStateInStock.Ptr(),
                    LocationID: square.String(
                        "C6W5YS5QM06F5",
                    ),
                    Quantity: square.String(
                        "53",
                    ),
                    TeamMemberID: square.String(
                        "LRK57NSQ5X7PUD05",
                    ),
                    OccurredAt: square.String(
                        "2016-11-16T22:25:24.878Z",
                    ),
                },
            },
        },
        IgnoreUnchangedCounts: square.Bool(
            true,
        ),
    }
client.Inventory.BatchCreateChanges(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchChangeInventoryRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.BatchGetChanges(request) -> *square.BatchGetInventoryChangesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns historical physical counts and adjustments based on the
provided filter criteria.

Results are paginated and sorted in ascending order according their
`occurred_at` timestamp (oldest first).

BatchRetrieveInventoryChanges is a catch-all query endpoint for queries
that cannot be handled by other, simpler endpoints.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchRetrieveInventoryChangesRequest{
        CatalogObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
        },
        LocationIDs: []string{
            "C6W5YS5QM06F5",
        },
        Types: []square.InventoryChangeType{
            square.InventoryChangeTypePhysicalCount,
        },
        States: []square.InventoryState{
            square.InventoryStateInStock,
        },
        UpdatedAfter: square.String(
            "2016-11-01T00:00:00.000Z",
        ),
        UpdatedBefore: square.String(
            "2016-12-01T00:00:00.000Z",
        ),
    }
client.Inventory.BatchGetChanges(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchRetrieveInventoryChangesRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.BatchGetCounts(request) -> *square.BatchGetInventoryCountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns current counts for the provided
[CatalogObject](entity:CatalogObject)s at the requested
[Location](entity:Location)s.

Results are paginated and sorted in descending order according to their
`calculated_at` timestamp (newest first).

When `updated_after` is specified, only counts that have changed since that
time (based on the server timestamp for the most recent change) are
returned. This allows clients to perform a "sync" operation, for example
in response to receiving a Webhook notification.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchGetInventoryCountsRequest{
        CatalogObjectIDs: []string{
            "W62UWFY35CWMYGVWK6TWJDNI",
        },
        LocationIDs: []string{
            "59TNP9SA8VGDA",
        },
        UpdatedAfter: square.String(
            "2016-11-16T00:00:00.000Z",
        ),
    }
client.Inventory.BatchGetCounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.BatchGetInventoryCountsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.DeprecatedGetPhysicalCount(PhysicalCountID) -> *square.GetInventoryPhysicalCountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated version of [RetrieveInventoryPhysicalCount](api-endpoint:Inventory-RetrieveInventoryPhysicalCount) after the endpoint URL
is updated to conform to the standard convention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeprecatedGetPhysicalCountInventoryRequest{
        PhysicalCountID: "physical_count_id",
    }
client.Inventory.DeprecatedGetPhysicalCount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**physicalCountID:** `string` 

ID of the
[InventoryPhysicalCount](entity:InventoryPhysicalCount) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.GetPhysicalCount(PhysicalCountID) -> *square.GetInventoryPhysicalCountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the [InventoryPhysicalCount](entity:InventoryPhysicalCount)
object with the provided `physical_count_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetPhysicalCountInventoryRequest{
        PhysicalCountID: "physical_count_id",
    }
client.Inventory.GetPhysicalCount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**physicalCountID:** `string` 

ID of the
[InventoryPhysicalCount](entity:InventoryPhysicalCount) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.GetTransfer(TransferID) -> *square.GetInventoryTransferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the [InventoryTransfer](entity:InventoryTransfer) object
with the provided `transfer_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetTransferInventoryRequest{
        TransferID: "transfer_id",
    }
client.Inventory.GetTransfer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferID:** `string` — ID of the [InventoryTransfer](entity:InventoryTransfer) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.Get(CatalogObjectID) -> *square.GetInventoryCountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the current calculated stock count for a given
[CatalogObject](entity:CatalogObject) at a given set of
[Location](entity:Location)s. Responses are paginated and unsorted.
For more sophisticated queries, use a batch endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetInventoryRequest{
        CatalogObjectID: "catalog_object_id",
        LocationIDs: square.String(
            "location_ids",
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Inventory.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**catalogObjectID:** `string` — ID of the [CatalogObject](entity:CatalogObject) to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**locationIDs:** `*string` 

The [Location](entity:Location) IDs to look up as a comma-separated
list. An empty list queries all locations.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for the original query.

See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Inventory.Changes(CatalogObjectID) -> *square.GetInventoryChangesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a set of physical counts and inventory adjustments for the
provided [CatalogObject](entity:CatalogObject) at the requested
[Location](entity:Location)s.

You can achieve the same result by calling [BatchRetrieveInventoryChanges](api-endpoint:Inventory-BatchRetrieveInventoryChanges)
and having the `catalog_object_ids` list contain a single element of the `CatalogObject` ID.

Results are paginated and sorted in descending order according to their
`occurred_at` timestamp (newest first).

There are no limits on how far back the caller can page. This endpoint can be
used to display recent changes for a specific item. For more
sophisticated queries, use a batch endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ChangesInventoryRequest{
        CatalogObjectID: "catalog_object_id",
        LocationIDs: square.String(
            "location_ids",
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Inventory.Changes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**catalogObjectID:** `string` — ID of the [CatalogObject](entity:CatalogObject) to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**locationIDs:** `*string` 

The [Location](entity:Location) IDs to look up as a comma-separated
list. An empty list queries all locations.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for the original query.

See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Invoices
<details><summary><code>client.Invoices.List() -> *square.ListInvoicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of invoices for a given location. The response 
is paginated. If truncated, the response includes a `cursor` that you    
use in a subsequent request to retrieve the next set of invoices.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListInvoicesRequest{
        LocationID: "location_id",
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Invoices.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location for which to list invoices.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint. 
Provide this cursor to retrieve the next set of results for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of invoices to return (200 is the maximum `limit`). 
If not provided, the server uses a default limit of 100 invoices.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Create(request) -> *square.CreateInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a draft [invoice](entity:Invoice) 
for an order created using the Orders API.

A draft invoice remains in your account and no action is taken. 
You must publish the invoice before Square can process it (send it to the customer's email address or charge the customer’s card on file).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateInvoiceRequest{
        Invoice: &square.Invoice{
            LocationID: square.String(
                "ES0RJRZYEC39A",
            ),
            OrderID: square.String(
                "CAISENgvlJ6jLWAzERDzjyHVybY",
            ),
            PrimaryRecipient: &square.InvoiceRecipient{
                CustomerID: square.String(
                    "JDKYHBWT1D4F8MFH63DBMEN8Y4",
                ),
            },
            PaymentRequests: []*square.InvoicePaymentRequest{
                &square.InvoicePaymentRequest{
                    RequestType: square.InvoiceRequestTypeBalance.Ptr(),
                    DueDate: square.String(
                        "2030-01-24",
                    ),
                    TippingEnabled: square.Bool(
                        true,
                    ),
                    AutomaticPaymentSource: square.InvoiceAutomaticPaymentSourceNone.Ptr(),
                    Reminders: []*square.InvoicePaymentReminder{
                        &square.InvoicePaymentReminder{
                            RelativeScheduledDays: square.Int(
                                -1,
                            ),
                            Message: square.String(
                                "Your invoice is due tomorrow",
                            ),
                        },
                    },
                },
            },
            DeliveryMethod: square.InvoiceDeliveryMethodEmail.Ptr(),
            InvoiceNumber: square.String(
                "inv-100",
            ),
            Title: square.String(
                "Event Planning Services",
            ),
            Description: square.String(
                "We appreciate your business!",
            ),
            ScheduledAt: square.String(
                "2030-01-13T10:00:00Z",
            ),
            AcceptedPaymentMethods: &square.InvoiceAcceptedPaymentMethods{
                Card: square.Bool(
                    true,
                ),
                SquareGiftCard: square.Bool(
                    false,
                ),
                BankAccount: square.Bool(
                    false,
                ),
                BuyNowPayLater: square.Bool(
                    false,
                ),
                CashAppPay: square.Bool(
                    false,
                ),
            },
            CustomFields: []*square.InvoiceCustomField{
                &square.InvoiceCustomField{
                    Label: square.String(
                        "Event Reference Number",
                    ),
                    Value: square.String(
                        "Ref. #1234",
                    ),
                    Placement: square.InvoiceCustomFieldPlacementAboveLineItems.Ptr(),
                },
                &square.InvoiceCustomField{
                    Label: square.String(
                        "Terms of Service",
                    ),
                    Value: square.String(
                        "The terms of service are...",
                    ),
                    Placement: square.InvoiceCustomFieldPlacementBelowLineItems.Ptr(),
                },
            },
            SaleOrServiceDate: square.String(
                "2030-01-24",
            ),
            StorePaymentMethodEnabled: square.Bool(
                false,
            ),
        },
        IdempotencyKey: square.String(
            "ce3748f9-5fc1-4762-aa12-aae5e843f1f4",
        ),
    }
client.Invoices.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoice:** `*square.Invoice` — The invoice to create.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique string that identifies the `CreateInvoice` request. If you do not 
provide `idempotency_key` (or provide an empty string as the value), the endpoint 
treats each request as independent.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Search(request) -> *square.SearchInvoicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for invoices from a location specified in 
the filter. You can optionally specify customers in the filter for whom to 
retrieve invoices. In the current implementation, you can only specify one location and 
optionally one customer.

The response is paginated. If truncated, the response includes a `cursor` 
that you use in a subsequent request to retrieve the next set of invoices.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchInvoicesRequest{
        Query: &square.InvoiceQuery{
            Filter: &square.InvoiceFilter{
                LocationIDs: []string{
                    "ES0RJRZYEC39A",
                },
                CustomerIDs: []string{
                    "JDKYHBWT1D4F8MFH63DBMEN8Y4",
                },
            },
            Sort: &square.InvoiceSort{
                Field: square.InvoiceSortField(
                    "INVOICE_SORT_DATE",
                ),
                Order: square.SortOrderDesc.Ptr(),
            },
        },
        Limit: square.Int(
            100,
        ),
    }
client.Invoices.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.InvoiceQuery` — Describes the query criteria for searching invoices.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of invoices to return (200 is the maximum `limit`). 
If not provided, the server uses a default limit of 100 invoices.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint. 
Provide this cursor to retrieve the next set of results for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Get(InvoiceID) -> *square.GetInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an invoice by invoice ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetInvoicesRequest{
        InvoiceID: "invoice_id",
    }
client.Invoices.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the invoice to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Update(InvoiceID, request) -> *square.UpdateInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an invoice. This endpoint supports sparse updates, so you only need
to specify the fields you want to change along with the required `version` field.
Some restrictions apply to updating invoices. For example, you cannot change the
`order_id` or `location_id` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateInvoiceRequest{
        InvoiceID: "invoice_id",
        Invoice: &square.Invoice{
            Version: square.Int(
                1,
            ),
            PaymentRequests: []*square.InvoicePaymentRequest{
                &square.InvoicePaymentRequest{
                    UID: square.String(
                        "2da7964f-f3d2-4f43-81e8-5aa220bf3355",
                    ),
                    TippingEnabled: square.Bool(
                        false,
                    ),
                },
            },
        },
        IdempotencyKey: square.String(
            "4ee82288-0910-499e-ab4c-5d0071dad1be",
        ),
    }
client.Invoices.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the invoice to update.
    
</dd>
</dl>

<dl>
<dd>

**invoice:** `*square.Invoice` 

The invoice fields to add, change, or clear. Fields can be cleared using
null values or the `remove` field (for individual payment requests or reminders).
The current invoice `version` is also required. For more information, including requirements,
limitations, and more examples, see [Update an Invoice](https://developer.squareup.com/docs/invoices-api/update-invoices).
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique string that identifies the `UpdateInvoice` request. If you do not
provide `idempotency_key` (or provide an empty string as the value), the endpoint
treats each request as independent.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**fieldsToClear:** `[]string` 

The list of fields to clear. Although this field is currently supported, we
recommend using null values or the `remove` field when possible. For examples, see
[Update an Invoice](https://developer.squareup.com/docs/invoices-api/update-invoices).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Delete(InvoiceID) -> *square.DeleteInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the specified invoice. When an invoice is deleted, the 
associated order status changes to CANCELED. You can only delete a draft 
invoice (you cannot delete a published invoice, including one that is scheduled for processing).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteInvoicesRequest{
        InvoiceID: "invoice_id",
        Version: square.Int(
            1,
        ),
    }
client.Invoices.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the invoice to delete.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The version of the [invoice](entity:Invoice) to delete.
If you do not know the version, you can call [GetInvoice](api-endpoint:Invoices-GetInvoice) or 
[ListInvoices](api-endpoint:Invoices-ListInvoices).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.CreateInvoiceAttachment(InvoiceID, request) -> *square.CreateInvoiceAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads a file and attaches it to an invoice. This endpoint accepts HTTP multipart/form-data file uploads
with a JSON `request` part and a `file` part. The `file` part must be a `readable stream` that contains a file
in a supported format: GIF, JPEG, PNG, TIFF, BMP, or PDF.

Invoices can have up to 10 attachments with a total file size of 25 MB. Attachments can be added only to invoices
in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.

__NOTE:__ When testing in the Sandbox environment, the total file size is limited to 1 KB.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateInvoiceAttachmentRequest{
        InvoiceID: "invoice_id",
    }
client.Invoices.CreateInvoiceAttachment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the [invoice](entity:Invoice) to attach the file to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.DeleteInvoiceAttachment(InvoiceID, AttachmentID) -> *square.DeleteInvoiceAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes an attachment from an invoice and permanently deletes the file. Attachments can be removed only
from invoices in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteInvoiceAttachmentRequest{
        InvoiceID: "invoice_id",
        AttachmentID: "attachment_id",
    }
client.Invoices.DeleteInvoiceAttachment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the [invoice](entity:Invoice) to delete the attachment from.
    
</dd>
</dl>

<dl>
<dd>

**attachmentID:** `string` — The ID of the [attachment](entity:InvoiceAttachment) to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Cancel(InvoiceID, request) -> *square.CancelInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels an invoice. The seller cannot collect payments for 
the canceled invoice.

You cannot cancel an invoice in the `DRAFT` state or in a terminal state: `PAID`, `REFUNDED`, `CANCELED`, or `FAILED`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelInvoiceRequest{
        InvoiceID: "invoice_id",
        Version: 0,
    }
client.Invoices.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the [invoice](entity:Invoice) to cancel.
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` 

The version of the [invoice](entity:Invoice) to cancel.
If you do not know the version, you can call 
[GetInvoice](api-endpoint:Invoices-GetInvoice) or [ListInvoices](api-endpoint:Invoices-ListInvoices).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Publish(InvoiceID, request) -> *square.PublishInvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publishes the specified draft invoice. 

After an invoice is published, Square 
follows up based on the invoice configuration. For example, Square 
sends the invoice to the customer's email address, charges the customer's card on file, or does 
nothing. Square also makes the invoice available on a Square-hosted invoice page. 

The invoice `status` also changes from `DRAFT` to a status 
based on the invoice configuration. For example, the status changes to `UNPAID` if 
Square emails the invoice or `PARTIALLY_PAID` if Square charges a card on file for a portion of the 
invoice amount.

In addition to the required `ORDERS_WRITE` and `INVOICES_WRITE` permissions, `CUSTOMERS_READ`
and `PAYMENTS_WRITE` are required when publishing invoices configured for card-on-file payments.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.PublishInvoiceRequest{
        InvoiceID: "invoice_id",
        Version: 1,
        IdempotencyKey: square.String(
            "32da42d0-1997-41b0-826b-f09464fc2c2e",
        ),
    }
client.Invoices.Publish(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invoiceID:** `string` — The ID of the invoice to publish.
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` 

The version of the [invoice](entity:Invoice) to publish.
This must match the current version of the invoice; otherwise, the request is rejected.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique string that identifies the `PublishInvoice` request. If you do not 
provide `idempotency_key` (or provide an empty string as the value), the endpoint 
treats each request as independent.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor
<details><summary><code>client.Labor.CreateScheduledShift(request) -> *square.CreateScheduledShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a scheduled shift by providing draft shift details such as job ID,
team member assignment, and start and end times.

The following `draft_shift_details` fields are required:
- `location_id`
- `job_id`
- `start_at`
- `end_at`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateScheduledShiftRequest{
        IdempotencyKey: square.String(
            "HIDSNG5KS478L",
        ),
        ScheduledShift: &square.ScheduledShift{
            DraftShiftDetails: &square.ScheduledShiftDetails{
                TeamMemberID: square.String(
                    "ormj0jJJZ5OZIzxrZYJI",
                ),
                LocationID: square.String(
                    "PAA1RJZZKXBFG",
                ),
                JobID: square.String(
                    "FzbJAtt9qEWncK1BWgVCxQ6M",
                ),
                StartAt: square.String(
                    "2019-01-25T03:11:00-05:00",
                ),
                EndAt: square.String(
                    "2019-01-25T13:11:00-05:00",
                ),
                Notes: square.String(
                    "Dont forget to prep the vegetables",
                ),
                IsDeleted: square.Bool(
                    false,
                ),
            },
        },
    }
client.Labor.CreateScheduledShift(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for the `CreateScheduledShift` request, used to ensure the
[idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency)
of the operation.
    
</dd>
</dl>

<dl>
<dd>

**scheduledShift:** `*square.ScheduledShift` 

The scheduled shift with `draft_shift_details`.
If needed, call [ListLocations](api-endpoint:Locations-ListLocations) to get location IDs,
[ListJobs](api-endpoint:Team-ListJobs) to get job IDs, and [SearchTeamMembers](api-endpoint:Team-SearchTeamMembers)
to get team member IDs and current job assignments.

The `start_at` and `end_at` timestamps must be provided in the time zone + offset of the
shift location specified in `location_id`. Example for Pacific Standard Time: 2024-10-31T12:30:00-08:00
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.BulkPublishScheduledShifts(request) -> *square.BulkPublishScheduledShiftsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publishes 1 - 100 scheduled shifts. This endpoint takes a map of individual publish
requests and returns a map of responses. When a scheduled shift is published, Square keeps
the `draft_shift_details` field as is and copies it to the `published_shift_details` field.

The minimum `start_at` and maximum `end_at` timestamps of all shifts in a
`BulkPublishScheduledShifts` request must fall within a two-week period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkPublishScheduledShiftsRequest{
        ScheduledShifts: map[string]*square.BulkPublishScheduledShiftsData{
            "key": &square.BulkPublishScheduledShiftsData{},
        },
        ScheduledShiftNotificationAudience: square.ScheduledShiftNotificationAudienceAffected.Ptr(),
    }
client.Labor.BulkPublishScheduledShifts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**scheduledShifts:** `map[string]*square.BulkPublishScheduledShiftsData` 

A map of 1 to 100 key-value pairs that represent individual publish requests.

- Each key is the ID of a scheduled shift you want to publish.
- Each value is a `BulkPublishScheduledShiftsData` object that contains the
`version` field or is an empty object.
    
</dd>
</dl>

<dl>
<dd>

**scheduledShiftNotificationAudience:** `*square.ScheduledShiftNotificationAudience` 

Indicates whether Square should send email notifications to team members and
which team members should receive the notifications. This setting applies to all shifts
specified in the bulk operation. The default value is `AFFECTED`.
See [ScheduledShiftNotificationAudience](#type-scheduledshiftnotificationaudience) for possible values
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.SearchScheduledShifts(request) -> *square.SearchScheduledShiftsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of scheduled shifts, with optional filter and sort settings.
By default, results are sorted by `start_at` in ascending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchScheduledShiftsRequest{
        Query: &square.ScheduledShiftQuery{
            Filter: &square.ScheduledShiftFilter{
                AssignmentStatus: square.ScheduledShiftFilterAssignmentStatusAssigned.Ptr(),
            },
            Sort: &square.ScheduledShiftSort{
                Field: square.ScheduledShiftSortFieldCreatedAt.Ptr(),
                Order: square.SortOrderAsc.Ptr(),
            },
        },
        Limit: square.Int(
            2,
        ),
        Cursor: square.String(
            "xoxp-1234-5678-90123",
        ),
    }
client.Labor.SearchScheduledShifts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.ScheduledShiftQuery` — Query conditions used to filter and sort the results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to return in a single response page. The default value is 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The pagination cursor returned by the previous call to this endpoint. Provide
this cursor to retrieve the next page of results for your original request. For more
information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.RetrieveScheduledShift(ID) -> *square.RetrieveScheduledShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a scheduled shift by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RetrieveScheduledShiftRequest{
        ID: "id",
    }
client.Labor.RetrieveScheduledShift(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the scheduled shift to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.UpdateScheduledShift(ID, request) -> *square.UpdateScheduledShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the draft shift details for a scheduled shift. This endpoint supports
sparse updates, so only new, changed, or removed fields are required in the request.
You must publish the shift to make updates public.

You can make the following updates to `draft_shift_details`:
- Change the `location_id`, `job_id`, `start_at`, and `end_at` fields.
- Add, change, or clear the `team_member_id` and `notes` fields. To clear these fields,
set the value to null.
- Change the `is_deleted` field. To delete a scheduled shift, set `is_deleted` to true
and then publish the shift.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateScheduledShiftRequest{
        ID: "id",
        ScheduledShift: &square.ScheduledShift{
            DraftShiftDetails: &square.ScheduledShiftDetails{
                TeamMemberID: square.String(
                    "ormj0jJJZ5OZIzxrZYJI",
                ),
                LocationID: square.String(
                    "PAA1RJZZKXBFG",
                ),
                JobID: square.String(
                    "FzbJAtt9qEWncK1BWgVCxQ6M",
                ),
                StartAt: square.String(
                    "2019-03-25T03:11:00-05:00",
                ),
                EndAt: square.String(
                    "2019-03-25T13:18:00-05:00",
                ),
                Notes: square.String(
                    "Dont forget to prep the vegetables",
                ),
                IsDeleted: square.Bool(
                    false,
                ),
            },
            Version: square.Int(
                1,
            ),
        },
    }
client.Labor.UpdateScheduledShift(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the scheduled shift to update.
    
</dd>
</dl>

<dl>
<dd>

**scheduledShift:** `*square.ScheduledShift` 

The scheduled shift with any updates in the `draft_shift_details` field.
If needed, call [ListLocations](api-endpoint:Locations-ListLocations) to get location IDs,
[ListJobs](api-endpoint:Team-ListJobs) to get job IDs, and [SearchTeamMembers](api-endpoint:Team-SearchTeamMembers)
to get team member IDs and current job assignments. Updates made to `published_shift_details`
are ignored.

If provided, the `start_at` and `end_at` timestamps must be in the time zone + offset of the
shift location specified in `location_id`. Example for Pacific Standard Time: 2024-10-31T12:30:00-08:00

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control for the request, provide the current version of the shift in the `version` field.
If the provided version doesn't match the server version, the request fails. If `version` is
omitted, Square executes a blind write, potentially overwriting data from another publish request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.PublishScheduledShift(ID, request) -> *square.PublishScheduledShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publishes a scheduled shift. When a scheduled shift is published, Square keeps the
`draft_shift_details` field as is and copies it to the `published_shift_details` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.PublishScheduledShiftRequest{
        ID: "id",
        IdempotencyKey: "HIDSNG5KS478L",
        Version: square.Int(
            2,
        ),
        ScheduledShiftNotificationAudience: square.ScheduledShiftNotificationAudienceAll.Ptr(),
    }
client.Labor.PublishScheduledShift(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the scheduled shift to publish.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique identifier for the `PublishScheduledShift` request, used to ensure the
[idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency)
of the operation.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the scheduled shift, used to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control. If the provided version doesn't match the server version, the request fails.
If omitted, Square executes a blind write, potentially overwriting data from another publish request.
    
</dd>
</dl>

<dl>
<dd>

**scheduledShiftNotificationAudience:** `*square.ScheduledShiftNotificationAudience` 

Indicates whether Square should send an email notification to team members and
which team members should receive the notification. The default value is `AFFECTED`.
See [ScheduledShiftNotificationAudience](#type-scheduledshiftnotificationaudience) for possible values
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.CreateTimecard(request) -> *square.CreateTimecardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new `Timecard`.

A `Timecard` represents a complete workday for a single team member.
You must provide the following values in your request to this
endpoint:

- `location_id`
- `team_member_id`
- `start_at`

An attempt to create a new `Timecard` can result in a `BAD_REQUEST` error when:
- The `status` of the new `Timecard` is `OPEN` and the team member has another
timecard with an `OPEN` status.
- The `start_at` date is in the future.
- The `start_at` or `end_at` date overlaps another timecard for the same team member.
- The `Break` instances are set in the request and a break `start_at`
is before the `Timecard.start_at`, a break `end_at` is after
the `Timecard.end_at`, or both.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateTimecardRequest{
        IdempotencyKey: square.String(
            "HIDSNG5KS478L",
        ),
        Timecard: &square.Timecard{
            LocationID: "PAA1RJZZKXBFG",
            StartAt: "2019-01-25T03:11:00-05:00",
            EndAt: square.String(
                "2019-01-25T13:11:00-05:00",
            ),
            Wage: &square.TimecardWage{
                Title: square.String(
                    "Barista",
                ),
                HourlyRate: &square.Money{
                    Amount: square.Int64(
                        1100,
                    ),
                    Currency: square.CurrencyUsd.Ptr(),
                },
                TipEligible: square.Bool(
                    true,
                ),
            },
            Breaks: []*square.Break{
                &square.Break{
                    StartAt: "2019-01-25T06:11:00-05:00",
                    EndAt: square.String(
                        "2019-01-25T06:16:00-05:00",
                    ),
                    BreakTypeID: "REGS1EQR1TPZ5",
                    Name: "Tea Break",
                    ExpectedDuration: "PT5M",
                    IsPaid: true,
                },
            },
            TeamMemberID: "ormj0jJJZ5OZIzxrZYJI",
            DeclaredCashTipMoney: &square.Money{
                Amount: square.Int64(
                    500,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
        },
    }
client.Labor.CreateTimecard(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique string value to ensure the idempotency of the operation.
    
</dd>
</dl>

<dl>
<dd>

**timecard:** `*square.Timecard` — The `Timecard` to be created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.SearchTimecards(request) -> *square.SearchTimecardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `Timecard` records for a business.
The list to be returned can be filtered by:
- Location IDs
- Team member IDs
- Timecard status (`OPEN` or `CLOSED`)
- Timecard start
- Timecard end
- Workday details

The list can be sorted by:
- `START_AT`
- `END_AT`
- `CREATED_AT`
- `UPDATED_AT`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchTimecardsRequest{
        Query: &square.TimecardQuery{
            Filter: &square.TimecardFilter{
                Workday: &square.TimecardWorkday{
                    DateRange: &square.DateRange{
                        StartDate: square.String(
                            "2019-01-20",
                        ),
                        EndDate: square.String(
                            "2019-02-03",
                        ),
                    },
                    MatchTimecardsBy: square.TimecardWorkdayMatcherStartAt.Ptr(),
                    DefaultTimezone: square.String(
                        "America/Los_Angeles",
                    ),
                },
            },
        },
        Limit: square.Int(
            100,
        ),
    }
client.Labor.SearchTimecards(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.TimecardQuery` — Query filters.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of resources in a page (200 by default).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — An opaque cursor for fetching the next page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.RetrieveTimecard(ID) -> *square.RetrieveTimecardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single `Timecard` specified by `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RetrieveTimecardRequest{
        ID: "id",
    }
client.Labor.RetrieveTimecard(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `Timecard` being retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.UpdateTimecard(ID, request) -> *square.UpdateTimecardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing `Timecard`.

When adding a `Break` to a `Timecard`, any earlier `Break` instances in the `Timecard` have
the `end_at` property set to a valid RFC-3339 datetime string.

When closing a `Timecard`, all `Break` instances in the `Timecard` must be complete with `end_at`
set on each `Break`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateTimecardRequest{
        ID: "id",
        Timecard: &square.Timecard{
            LocationID: "PAA1RJZZKXBFG",
            StartAt: "2019-01-25T03:11:00-05:00",
            EndAt: square.String(
                "2019-01-25T13:11:00-05:00",
            ),
            Wage: &square.TimecardWage{
                Title: square.String(
                    "Bartender",
                ),
                HourlyRate: &square.Money{
                    Amount: square.Int64(
                        1500,
                    ),
                    Currency: square.CurrencyUsd.Ptr(),
                },
                TipEligible: square.Bool(
                    true,
                ),
            },
            Breaks: []*square.Break{
                &square.Break{
                    ID: square.String(
                        "X7GAQYVVRRG6P",
                    ),
                    StartAt: "2019-01-25T06:11:00-05:00",
                    EndAt: square.String(
                        "2019-01-25T06:16:00-05:00",
                    ),
                    BreakTypeID: "REGS1EQR1TPZ5",
                    Name: "Tea Break",
                    ExpectedDuration: "PT5M",
                    IsPaid: true,
                },
            },
            Status: square.TimecardStatusClosed.Ptr(),
            Version: square.Int(
                1,
            ),
            TeamMemberID: "ormj0jJJZ5OZIzxrZYJI",
            DeclaredCashTipMoney: &square.Money{
                Amount: square.Int64(
                    500,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
        },
    }
client.Labor.UpdateTimecard(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the object being updated.
    
</dd>
</dl>

<dl>
<dd>

**timecard:** `*square.Timecard` — The updated `Timecard` object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.DeleteTimecard(ID) -> *square.DeleteTimecardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a `Timecard`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteTimecardRequest{
        ID: "id",
    }
client.Labor.DeleteTimecard(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `Timecard` being deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Locations
<details><summary><code>client.Locations.List() -> *square.ListLocationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides details about all of the seller's [locations](https://developer.squareup.com/docs/locations-api),
including those with an inactive status. Locations are listed alphabetically by `name`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Locations.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Create(request) -> *square.CreateLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a [location](https://developer.squareup.com/docs/locations-api).
Creating new locations allows for separate configuration of receipt layouts, item prices,
and sales reports. Developers can use locations to separate sales activity through applications
that integrate with Square from sales activity elsewhere in a seller's account.
Locations created programmatically with the Locations API last forever and
are visible to the seller for their own management. Therefore, ensure that
each location has a sensible and unique name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateLocationRequest{
        Location: &square.Location{
            Name: square.String(
                "Midtown",
            ),
            Address: &square.Address{
                AddressLine1: square.String(
                    "1234 Peachtree St. NE",
                ),
                Locality: square.String(
                    "Atlanta",
                ),
                AdministrativeDistrictLevel1: square.String(
                    "GA",
                ),
                PostalCode: square.String(
                    "30309",
                ),
            },
            Description: square.String(
                "Midtown Atlanta store",
            ),
        },
    }
client.Locations.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**location:** `*square.Location` 

The initial values of the location being created. The `name` field is required and must be unique within a seller account.
All other fields are optional, but any information you care about for the location should be included.
The remaining fields are automatically added based on the data from the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Get(LocationID) -> *square.GetLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves details of a single location. Specify "main"
as the location ID to retrieve details of the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetLocationsRequest{
        LocationID: "location_id",
    }
client.Locations.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` 

The ID of the location to retrieve. Specify the string
"main" to return the main location.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Update(LocationID, request) -> *square.UpdateLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a [location](https://developer.squareup.com/docs/locations-api).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateLocationRequest{
        LocationID: "location_id",
        Location: &square.Location{
            BusinessHours: &square.BusinessHours{
                Periods: []*square.BusinessHoursPeriod{
                    &square.BusinessHoursPeriod{
                        DayOfWeek: square.DayOfWeekFri.Ptr(),
                        StartLocalTime: square.String(
                            "07:00",
                        ),
                        EndLocalTime: square.String(
                            "18:00",
                        ),
                    },
                    &square.BusinessHoursPeriod{
                        DayOfWeek: square.DayOfWeekSat.Ptr(),
                        StartLocalTime: square.String(
                            "07:00",
                        ),
                        EndLocalTime: square.String(
                            "18:00",
                        ),
                    },
                    &square.BusinessHoursPeriod{
                        DayOfWeek: square.DayOfWeekSun.Ptr(),
                        StartLocalTime: square.String(
                            "09:00",
                        ),
                        EndLocalTime: square.String(
                            "15:00",
                        ),
                    },
                },
            },
            Description: square.String(
                "Midtown Atlanta store - Open weekends",
            ),
        },
    }
client.Locations.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location to update.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*square.Location` — The `Location` object with only the fields to update.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Checkouts(LocationID, request) -> *square.CreateCheckoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Links a `checkoutId` to a `checkout_page_url` that customers are
directed to in order to provide their payment information using a
payment processing workflow hosted on connect.squareup.com. 


NOTE: The Checkout API has been updated with new features. 
For more information, see [Checkout API highlights](https://developer.squareup.com/docs/checkout-api#checkout-api-highlights).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateCheckoutRequest{
        LocationID: "location_id",
        IdempotencyKey: "86ae1696-b1e3-4328-af6d-f1e04d947ad6",
        Order: &square.CreateOrderRequest{
            Order: &square.Order{
                LocationID: "location_id",
                ReferenceID: square.String(
                    "reference_id",
                ),
                CustomerID: square.String(
                    "customer_id",
                ),
                LineItems: []*square.OrderLineItem{
                    &square.OrderLineItem{
                        Name: square.String(
                            "Printed T Shirt",
                        ),
                        Quantity: "2",
                        AppliedTaxes: []*square.OrderLineItemAppliedTax{
                            &square.OrderLineItemAppliedTax{
                                TaxUID: "38ze1696-z1e3-5628-af6d-f1e04d947fg3",
                            },
                        },
                        AppliedDiscounts: []*square.OrderLineItemAppliedDiscount{
                            &square.OrderLineItemAppliedDiscount{
                                DiscountUID: "56ae1696-z1e3-9328-af6d-f1e04d947gd4",
                            },
                        },
                        BasePriceMoney: &square.Money{
                            Amount: square.Int64(
                                1500,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                    },
                    &square.OrderLineItem{
                        Name: square.String(
                            "Slim Jeans",
                        ),
                        Quantity: "1",
                        BasePriceMoney: &square.Money{
                            Amount: square.Int64(
                                2500,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                    },
                    &square.OrderLineItem{
                        Name: square.String(
                            "Woven Sweater",
                        ),
                        Quantity: "3",
                        BasePriceMoney: &square.Money{
                            Amount: square.Int64(
                                3500,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                    },
                },
                Taxes: []*square.OrderLineItemTax{
                    &square.OrderLineItemTax{
                        UID: square.String(
                            "38ze1696-z1e3-5628-af6d-f1e04d947fg3",
                        ),
                        Type: square.OrderLineItemTaxTypeInclusive.Ptr(),
                        Percentage: square.String(
                            "7.75",
                        ),
                        Scope: square.OrderLineItemTaxScopeLineItem.Ptr(),
                    },
                },
                Discounts: []*square.OrderLineItemDiscount{
                    &square.OrderLineItemDiscount{
                        UID: square.String(
                            "56ae1696-z1e3-9328-af6d-f1e04d947gd4",
                        ),
                        Type: square.OrderLineItemDiscountTypeFixedAmount.Ptr(),
                        AmountMoney: &square.Money{
                            Amount: square.Int64(
                                100,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                        Scope: square.OrderLineItemDiscountScopeLineItem.Ptr(),
                    },
                },
            },
            IdempotencyKey: square.String(
                "12ae1696-z1e3-4328-af6d-f1e04d947gd4",
            ),
        },
        AskForShippingAddress: square.Bool(
            true,
        ),
        MerchantSupportEmail: square.String(
            "merchant+support@website.com",
        ),
        PrePopulateBuyerEmail: square.String(
            "example@email.com",
        ),
        PrePopulateShippingAddress: &square.Address{
            AddressLine1: square.String(
                "1455 Market St.",
            ),
            AddressLine2: square.String(
                "Suite 600",
            ),
            Locality: square.String(
                "San Francisco",
            ),
            AdministrativeDistrictLevel1: square.String(
                "CA",
            ),
            PostalCode: square.String(
                "94103",
            ),
            Country: square.CountryUs.Ptr(),
            FirstName: square.String(
                "Jane",
            ),
            LastName: square.String(
                "Doe",
            ),
        },
        RedirectURL: square.String(
            "https://merchant.website.com/order-confirm",
        ),
        AdditionalRecipients: []*square.ChargeRequestAdditionalRecipient{
            &square.ChargeRequestAdditionalRecipient{
                LocationID: "057P5VYJ4A5X1",
                Description: "Application fees",
                AmountMoney: &square.Money{
                    Amount: square.Int64(
                        60,
                    ),
                    Currency: square.CurrencyUsd.Ptr(),
                },
            },
        },
    }
client.Locations.Checkouts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the business location to associate the checkout with.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this checkout among others you have created. It can be
any valid string but must be unique for every order sent to Square Checkout for a given location ID.

The idempotency key is used to avoid processing the same order more than once. If you are 
unsure whether a particular checkout was created successfully, you can attempt it again with
the same idempotency key and all the same other parameters without worrying about creating duplicates.

You should use a random number/string generator native to the language
you are working in to generate strings for your idempotency keys.

For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**order:** `*square.CreateOrderRequest` — The order including line items to be checked out.
    
</dd>
</dl>

<dl>
<dd>

**askForShippingAddress:** `*bool` 

If `true`, Square Checkout collects shipping information on your behalf and stores 
that information with the transaction information in the Square Seller Dashboard.

Default: `false`.
    
</dd>
</dl>

<dl>
<dd>

**merchantSupportEmail:** `*string` 

The email address to display on the Square Checkout confirmation page
and confirmation email that the buyer can use to contact the seller.

If this value is not set, the confirmation page and email display the
primary email address associated with the seller's Square account.

Default: none; only exists if explicitly set.
    
</dd>
</dl>

<dl>
<dd>

**prePopulateBuyerEmail:** `*string` 

If provided, the buyer's email is prepopulated on the checkout page
as an editable text field.

Default: none; only exists if explicitly set.
    
</dd>
</dl>

<dl>
<dd>

**prePopulateShippingAddress:** `*square.Address` 

If provided, the buyer's shipping information is prepopulated on the
checkout page as editable text fields.

Default: none; only exists if explicitly set.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `*string` 

The URL to redirect to after the checkout is completed with `checkoutId`,
`transactionId`, and `referenceId` appended as URL parameters. For example,
if the provided redirect URL is `http://www.example.com/order-complete`, a
successful transaction redirects the customer to:

`http://www.example.com/order-complete?checkoutId=xxxxxx&amp;referenceId=xxxxxx&amp;transactionId=xxxxxx`

If you do not provide a redirect URL, Square Checkout displays an order
confirmation page on your behalf; however, it is strongly recommended that
you provide a redirect URL so you can verify the transaction results and
finalize the order through your existing/normal confirmation workflow.

Default: none; only exists if explicitly set.
    
</dd>
</dl>

<dl>
<dd>

**additionalRecipients:** `[]*square.ChargeRequestAdditionalRecipient` 

The basic primitive of a multi-party transaction. The value is optional.
The transaction facilitated by you can be split from here.

If you provide this value, the `amount_money` value in your `additional_recipients` field
cannot be more than 90% of the `total_money` calculated by Square for your order.
The `location_id` must be a valid seller location where the checkout is occurring.

This field requires `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission.

This field is currently not supported in the Square Sandbox.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 

An optional note to associate with the `checkout` object.

This value cannot exceed 60 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loyalty
<details><summary><code>client.Loyalty.SearchEvents(request) -> *square.SearchLoyaltyEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for loyalty events.

A Square loyalty program maintains a ledger of events that occur during the lifetime of a
buyer's loyalty account. Each change in the point balance
(for example, points earned, points redeemed, and points expired) is
recorded in the ledger. Using this endpoint, you can search the ledger for events.

Search results are sorted by `created_at` in descending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchLoyaltyEventsRequest{
        Query: &square.LoyaltyEventQuery{
            Filter: &square.LoyaltyEventFilter{
                OrderFilter: &square.LoyaltyEventOrderFilter{
                    OrderID: "PyATxhYLfsMqpVkcKJITPydgEYfZY",
                },
            },
        },
        Limit: square.Int(
            30,
        ),
    }
client.Loyalty.SearchEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.LoyaltyEventQuery` 

A set of one or more predefined query filters to apply when 
searching for loyalty events. The endpoint performs a logical AND to 
evaluate multiple filters and performs a logical OR on arrays  
that specifies multiple field values.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to include in the response. 
The last page might contain fewer events. 
The default is 30 events.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for your original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Merchants
<details><summary><code>client.Merchants.List() -> *square.ListMerchantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides details about the merchant associated with a given access token.

The access token used to connect your application to a Square seller is associated
with a single merchant. That means that `ListMerchants` returns a list
with a single `Merchant` object. You can specify your personal access token
to get your own merchant information or specify an OAuth token to get the
information for the merchant that granted your application access.

If you know the merchant ID, you can also use the [RetrieveMerchant](api-endpoint:Merchants-RetrieveMerchant)
endpoint to retrieve the merchant information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListMerchantsRequest{
        Cursor: square.Int(
            1,
        ),
    }
client.Merchants.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*int` — The cursor generated by the previous response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.Get(MerchantID) -> *square.GetMerchantResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the `Merchant` object for the given `merchant_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetMerchantsRequest{
        MerchantID: "merchant_id",
    }
client.Merchants.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 

The ID of the merchant to retrieve. If the string "me" is supplied as the ID,
then retrieve the merchant that is currently accessible to this call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Checkout
<details><summary><code>client.Checkout.RetrieveLocationSettings(LocationID) -> *square.RetrieveLocationSettingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the location-level settings for a Square-hosted checkout page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RetrieveLocationSettingsRequest{
        LocationID: "location_id",
    }
client.Checkout.RetrieveLocationSettings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location for which to retrieve settings.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.UpdateLocationSettings(LocationID, request) -> *square.UpdateLocationSettingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the location-level settings for a Square-hosted checkout page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateLocationSettingsRequest{
        LocationID: "location_id",
        LocationSettings: &square.CheckoutLocationSettings{},
    }
client.Checkout.UpdateLocationSettings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location for which to retrieve settings.
    
</dd>
</dl>

<dl>
<dd>

**locationSettings:** `*square.CheckoutLocationSettings` — Describe your updates using the `location_settings` object. Make sure it contains only the fields that have changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.RetrieveMerchantSettings() -> *square.RetrieveMerchantSettingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the merchant-level settings for a Square-hosted checkout page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Checkout.RetrieveMerchantSettings(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.UpdateMerchantSettings(request) -> *square.UpdateMerchantSettingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the merchant-level settings for a Square-hosted checkout page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateMerchantSettingsRequest{
        MerchantSettings: &square.CheckoutMerchantSettings{},
    }
client.Checkout.UpdateMerchantSettings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantSettings:** `*square.CheckoutMerchantSettings` — Describe your updates using the `merchant_settings` object. Make sure it contains only the fields that have changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Orders
<details><summary><code>client.Orders.Create(request) -> *square.CreateOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new [order](entity:Order) that can include information about products for
purchase and settings to apply to the purchase.

To pay for a created order, see
[Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).

You can modify open orders using the [UpdateOrder](api-endpoint:Orders-UpdateOrder) endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateOrderRequest{
        Order: &square.Order{
            LocationID: "057P5VYJ4A5X1",
            ReferenceID: square.String(
                "my-order-001",
            ),
            LineItems: []*square.OrderLineItem{
                &square.OrderLineItem{
                    Name: square.String(
                        "New York Strip Steak",
                    ),
                    Quantity: "1",
                    BasePriceMoney: &square.Money{
                        Amount: square.Int64(
                            1599,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                },
                &square.OrderLineItem{
                    Quantity: "2",
                    CatalogObjectID: square.String(
                        "BEMYCSMIJL46OCDV4KYIKXIB",
                    ),
                    Modifiers: []*square.OrderLineItemModifier{
                        &square.OrderLineItemModifier{
                            CatalogObjectID: square.String(
                                "CHQX7Y4KY6N5KINJKZCFURPZ",
                            ),
                        },
                    },
                    AppliedDiscounts: []*square.OrderLineItemAppliedDiscount{
                        &square.OrderLineItemAppliedDiscount{
                            DiscountUID: "one-dollar-off",
                        },
                    },
                },
            },
            Taxes: []*square.OrderLineItemTax{
                &square.OrderLineItemTax{
                    UID: square.String(
                        "state-sales-tax",
                    ),
                    Name: square.String(
                        "State Sales Tax",
                    ),
                    Percentage: square.String(
                        "9",
                    ),
                    Scope: square.OrderLineItemTaxScopeOrder.Ptr(),
                },
            },
            Discounts: []*square.OrderLineItemDiscount{
                &square.OrderLineItemDiscount{
                    UID: square.String(
                        "labor-day-sale",
                    ),
                    Name: square.String(
                        "Labor Day Sale",
                    ),
                    Percentage: square.String(
                        "5",
                    ),
                    Scope: square.OrderLineItemDiscountScopeOrder.Ptr(),
                },
                &square.OrderLineItemDiscount{
                    UID: square.String(
                        "membership-discount",
                    ),
                    CatalogObjectID: square.String(
                        "DB7L55ZH2BGWI4H23ULIWOQ7",
                    ),
                    Scope: square.OrderLineItemDiscountScopeOrder.Ptr(),
                },
                &square.OrderLineItemDiscount{
                    UID: square.String(
                        "one-dollar-off",
                    ),
                    Name: square.String(
                        "Sale - $1.00 off",
                    ),
                    AmountMoney: &square.Money{
                        Amount: square.Int64(
                            100,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                    Scope: square.OrderLineItemDiscountScopeLineItem.Ptr(),
                },
            },
        },
        IdempotencyKey: square.String(
            "8193148c-9586-11e6-99f9-28cfe92138cf",
        ),
    }
client.Orders.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.CreateOrderRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.BatchGet(request) -> *square.BatchGetOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a set of [orders](entity:Order) by their IDs.

If a given order ID does not exist, the ID is ignored instead of generating an error.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchGetOrdersRequest{
        LocationID: square.String(
            "057P5VYJ4A5X1",
        ),
        OrderIDs: []string{
            "CAISEM82RcpmcFBM0TfOyiHV3es",
            "CAISENgvlJ6jLWAzERDzjyHVybY",
        },
    }
client.Orders.BatchGet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `*string` 

The ID of the location for these orders. This field is optional: omit it to retrieve
orders within the scope of the current authorization's merchant ID.
    
</dd>
</dl>

<dl>
<dd>

**orderIDs:** `[]string` — The IDs of the orders to retrieve. A maximum of 100 orders can be retrieved per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Calculate(request) -> *square.CalculateOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enables applications to preview order pricing without creating an order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CalculateOrderRequest{
        Order: &square.Order{
            LocationID: "D7AVYMEAPJ3A3",
            LineItems: []*square.OrderLineItem{
                &square.OrderLineItem{
                    Name: square.String(
                        "Item 1",
                    ),
                    Quantity: "1",
                    BasePriceMoney: &square.Money{
                        Amount: square.Int64(
                            500,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                },
                &square.OrderLineItem{
                    Name: square.String(
                        "Item 2",
                    ),
                    Quantity: "2",
                    BasePriceMoney: &square.Money{
                        Amount: square.Int64(
                            300,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                },
            },
            Discounts: []*square.OrderLineItemDiscount{
                &square.OrderLineItemDiscount{
                    Name: square.String(
                        "50% Off",
                    ),
                    Percentage: square.String(
                        "50",
                    ),
                    Scope: square.OrderLineItemDiscountScopeOrder.Ptr(),
                },
            },
        },
    }
client.Orders.Calculate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**order:** `*square.Order` — The order to be calculated. Expects the entire order, not a sparse update.
    
</dd>
</dl>

<dl>
<dd>

**proposedRewards:** `[]*square.OrderReward` 

Identifies one or more loyalty reward tiers to apply during the order calculation.
The discounts defined by the reward tiers are added to the order only to preview the
effect of applying the specified rewards. The rewards do not correspond to actual
redemptions; that is, no `reward`s are created. Therefore, the reward `id`s are
random strings used only to reference the reward tier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Clone(request) -> *square.CloneOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new order, in the `DRAFT` state, by duplicating an existing order. The newly created order has
only the core fields (such as line items, taxes, and discounts) copied from the original order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CloneOrderRequest{
        OrderID: "ZAISEM52YcpmcWAzERDOyiWS123",
        Version: square.Int(
            3,
        ),
        IdempotencyKey: square.String(
            "UNIQUE_STRING",
        ),
    }
client.Orders.Clone(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the order to clone.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

An optional order version for concurrency protection.

If a version is provided, it must match the latest stored version of the order to clone.
If a version is not provided, the API clones the latest version.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A value you specify that uniquely identifies this clone request.

If you are unsure whether a particular order was cloned successfully,
you can reattempt the call with the same idempotency key without
worrying about creating duplicate cloned orders.
The originally cloned order is returned.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Search(request) -> *square.SearchOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search all orders for one or more locations. Orders include all sales,
returns, and exchanges regardless of how or when they entered the Square
ecosystem (such as Point of Sale, Invoices, and Connect APIs).

`SearchOrders` requests need to specify which locations to search and define a
[SearchOrdersQuery](entity:SearchOrdersQuery) object that controls
how to sort or filter the results. Your `SearchOrdersQuery` can:

  Set filter criteria.
  Set the sort order.
  Determine whether to return results as complete `Order` objects or as
[OrderEntry](entity:OrderEntry) objects.

Note that details for orders processed with Square Point of Sale while in
offline mode might not be transmitted to Square for up to 72 hours. Offline
orders have a `created_at` value that reflects the time the order was created,
not the time it was subsequently transmitted to Square.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchOrdersRequest{
        LocationIDs: []string{
            "057P5VYJ4A5X1",
            "18YC4JDH91E1H",
        },
        Query: &square.SearchOrdersQuery{
            Filter: &square.SearchOrdersFilter{
                StateFilter: &square.SearchOrdersStateFilter{
                    States: []square.OrderState{
                        square.OrderStateCompleted,
                    },
                },
                DateTimeFilter: &square.SearchOrdersDateTimeFilter{
                    ClosedAt: &square.TimeRange{
                        StartAt: square.String(
                            "2018-03-03T20:00:00+00:00",
                        ),
                        EndAt: square.String(
                            "2019-03-04T21:54:45+00:00",
                        ),
                    },
                },
            },
            Sort: &square.SearchOrdersSort{
                SortField: square.SearchOrdersSortFieldClosedAt,
                SortOrder: square.SortOrderDesc.Ptr(),
            },
        },
        Limit: square.Int(
            3,
        ),
        ReturnEntries: square.Bool(
            true,
        ),
    }
client.Orders.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationIDs:** `[]string` 

The location IDs for the orders to query. All locations must belong to
the same merchant.

Max: 10 location IDs.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for your original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**query:** `*square.SearchOrdersQuery` 

Query conditions used to filter or sort the results. Note that when
retrieving additional pages using a cursor, you must use the original query.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.

Default: `500`
Max: `1000`
    
</dd>
</dl>

<dl>
<dd>

**returnEntries:** `*bool` 

A Boolean that controls the format of the search results. If `true`,
`SearchOrders` returns [OrderEntry](entity:OrderEntry) objects. If `false`, `SearchOrders`
returns complete order objects.

Default: `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Get(OrderID) -> *square.GetOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an [Order](entity:Order) by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetOrdersRequest{
        OrderID: "order_id",
    }
client.Orders.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the order to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Update(OrderID, request) -> *square.UpdateOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an open [order](entity:Order) by adding, replacing, or deleting
fields. Orders with a `COMPLETED` or `CANCELED` state cannot be updated.

An `UpdateOrder` request requires the following:

- The `order_id` in the endpoint path, identifying the order to update.
- The latest `version` of the order to update.
- The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#sparse-order-objects)
containing only the fields to update and the version to which the update is
being applied.
- If deleting fields, the [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#identifying-fields-to-delete)
identifying the fields to clear.

To pay for an order, see
[Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateOrderRequest{
        OrderID: "order_id",
        Order: &square.Order{
            LocationID: "location_id",
            LineItems: []*square.OrderLineItem{
                &square.OrderLineItem{
                    UID: square.String(
                        "cookie_uid",
                    ),
                    Name: square.String(
                        "COOKIE",
                    ),
                    Quantity: "2",
                    BasePriceMoney: &square.Money{
                        Amount: square.Int64(
                            200,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                },
            },
            Version: square.Int(
                1,
            ),
        },
        FieldsToClear: []string{
            "discounts",
        },
        IdempotencyKey: square.String(
            "UNIQUE_STRING",
        ),
    }
client.Orders.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the order to update.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*square.Order` 

The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#sparse-order-objects)
containing only the fields to update and the version to which the update is
being applied.
    
</dd>
</dl>

<dl>
<dd>

**fieldsToClear:** `[]string` 

The [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#identifying-fields-to-delete)
fields to clear. For example, `line_items[uid].note`.
For more information, see [Deleting fields](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#deleting-fields).
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A value you specify that uniquely identifies this update request.

If you are unsure whether a particular update was applied to an order successfully,
you can reattempt it with the same idempotency key without
worrying about creating duplicate updates to the order.
The latest order version is returned.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Pay(OrderID, request) -> *square.PayOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pay for an [order](entity:Order) using one or more approved [payments](entity:Payment)
or settle an order with a total of `0`.

The total of the `payment_ids` listed in the request must be equal to the order
total. Orders with a total amount of `0` can be marked as paid by specifying an empty
array of `payment_ids` in the request.

To be used with `PayOrder`, a payment must:

- Reference the order by specifying the `order_id` when [creating the payment](api-endpoint:Payments-CreatePayment).
Any approved payments that reference the same `order_id` not specified in the
`payment_ids` is canceled.
- Be approved with [delayed capture](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture).
Using a delayed capture payment with `PayOrder` completes the approved payment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.PayOrderRequest{
        OrderID: "order_id",
        IdempotencyKey: "c043a359-7ad9-4136-82a9-c3f1d66dcbff",
        PaymentIDs: []string{
            "EnZdNAlWCmfh6Mt5FMNST1o7taB",
            "0LRiVlbXVwe8ozu4KbZxd12mvaB",
        },
    }
client.Orders.Pay(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the order being paid.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A value you specify that uniquely identifies this request among requests you have sent. If
you are unsure whether a particular payment request was completed successfully, you can reattempt
it with the same idempotency key without worrying about duplicate payments.

For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**orderVersion:** `*int` — The version of the order being paid. If not supplied, the latest version will be paid.
    
</dd>
</dl>

<dl>
<dd>

**paymentIDs:** `[]string` 

The IDs of the [payments](entity:Payment) to collect.
The payment total must match the order total.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payments
<details><summary><code>client.Payments.List() -> *square.ListPaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of payments taken by the account making the request.

Results are eventually consistent, and new payments or changes to payments might take several
seconds to appear.

The maximum results per page is 100.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListPaymentsRequest{
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        SortOrder: square.String(
            "sort_order",
        ),
        Cursor: square.String(
            "cursor",
        ),
        LocationID: square.String(
            "location_id",
        ),
        Total: square.Int64(
            1000000,
        ),
        Last4: square.String(
            "last_4",
        ),
        CardBrand: square.String(
            "card_brand",
        ),
        Limit: square.Int(
            1,
        ),
        IsOfflinePayment: square.Bool(
            true,
        ),
        OfflineBeginTime: square.String(
            "offline_begin_time",
        ),
        OfflineEndTime: square.String(
            "offline_end_time",
        ),
        UpdatedAtBeginTime: square.String(
            "updated_at_begin_time",
        ),
        UpdatedAtEndTime: square.String(
            "updated_at_end_time",
        ),
        SortField: square.ListPaymentsRequestSortFieldCreatedAt.Ptr(),
    }
client.Payments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**beginTime:** `*string` 

Indicates the start of the time range to retrieve payments for, in RFC 3339 format.
The range is determined using the `created_at` field for each Payment.
Inclusive. Default: The current time minus one year.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 

Indicates the end of the time range to retrieve payments for, in RFC 3339 format.  The
range is determined using the `created_at` field for each Payment.

Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*string` 

The order in which results are listed by `ListPaymentsRequest.sort_field`:
- `ASC` - Oldest to newest.
- `DESC` - Newest to oldest (default).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

Limit results to the location supplied. By default, results are returned
for the default (main) location associated with the seller.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*int64` — The exact amount in the `total_money` for a payment.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — The last four digits of a payment card.
    
</dd>
</dl>

<dl>
<dd>

**cardBrand:** `*string` — The brand of the payment card (for example, VISA).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.
It is possible to receive fewer results than the specified limit on a given page.

The default value of 100 is also the maximum allowed value. If the provided value is 
greater than 100, it is ignored and the default value is used instead.

Default: `100`
    
</dd>
</dl>

<dl>
<dd>

**isOfflinePayment:** `*bool` — Whether the payment was taken offline or not.
    
</dd>
</dl>

<dl>
<dd>

**offlineBeginTime:** `*string` 

Indicates the start of the time range for which to retrieve offline payments, in RFC 3339
format for timestamps. The range is determined using the
`offline_payment_details.client_created_at` field for each Payment. If set, payments without a
value set in `offline_payment_details.client_created_at` will not be returned.

Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**offlineEndTime:** `*string` 

Indicates the end of the time range for which to retrieve offline payments, in RFC 3339
format for timestamps. The range is determined using the
`offline_payment_details.client_created_at` field for each Payment. If set, payments without a
value set in `offline_payment_details.client_created_at` will not be returned.

Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtBeginTime:** `*string` 

Indicates the start of the time range to retrieve payments for, in RFC 3339 format.  The
range is determined using the `updated_at` field for each Payment.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtEndTime:** `*string` 

Indicates the end of the time range to retrieve payments for, in RFC 3339 format.  The
range is determined using the `updated_at` field for each Payment.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*square.ListPaymentsRequestSortField` — The field used to sort results by. The default is `CREATED_AT`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Create(request) -> *square.CreatePaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a payment using the provided source. You can use this endpoint 
to charge a card (credit/debit card or    
Square gift card) or record a payment that the seller received outside of Square 
(cash payment from a buyer or a payment that an external entity 
processed on behalf of the seller).

The endpoint creates a 
`Payment` object and returns it in the response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreatePaymentRequest{
        SourceID: "ccof:GaJGNaZa8x4OgDJn4GB",
        IdempotencyKey: "7b0f3ec5-086a-4871-8f13-3c81b3875218",
        AmountMoney: &square.Money{
            Amount: square.Int64(
                1000,
            ),
            Currency: square.CurrencyUsd.Ptr(),
        },
        AppFeeMoney: &square.Money{
            Amount: square.Int64(
                10,
            ),
            Currency: square.CurrencyUsd.Ptr(),
        },
        Autocomplete: square.Bool(
            true,
        ),
        CustomerID: square.String(
            "W92WH6P11H4Z77CTET0RNTGFW8",
        ),
        LocationID: square.String(
            "L88917AVBK2S5",
        ),
        ReferenceID: square.String(
            "123456",
        ),
        Note: square.String(
            "Brief description",
        ),
    }
client.Payments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sourceID:** `string` 

The ID for the source of funds for this payment.
This could be a payment token generated by the Web Payments SDK for any of its
[supported methods](https://developer.squareup.com/docs/web-payments/overview#explore-payment-methods),
including cards, bank transfers, Afterpay or Cash App Pay. If recording a payment
that the seller received outside of Square, specify either "CASH" or "EXTERNAL".
For more information, see
[Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreatePayment` request. Keys can be any valid string
but must be unique for every `CreatePayment` request.

Note: The number of allowed characters might be less than the stated maximum, if multi-byte
characters are used.

For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**amountMoney:** `*square.Money` 

The amount of money to accept for this payment, not including `tip_money`.

The amount must be specified in the smallest denomination of the applicable currency
(for example, US dollar amounts are specified in cents). For more information, see
[Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).

The currency code must match the currency associated with the business
that is accepting the payment.
    
</dd>
</dl>

<dl>
<dd>

**tipMoney:** `*square.Money` 

The amount designated as a tip, in addition to `amount_money`.

The amount must be specified in the smallest denomination of the applicable currency
(for example, US dollar amounts are specified in cents). For more information, see
[Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).

The currency code must match the currency associated with the business
that is accepting the payment.
    
</dd>
</dl>

<dl>
<dd>

**appFeeMoney:** `*square.Money` 

The amount of money that the developer is taking as a fee
for facilitating the payment on behalf of the seller.

The amount cannot be more than 90% of the total amount of the payment.

The amount must be specified in the smallest denomination of the applicable currency
(for example, US dollar amounts are specified in cents). For more information, see
[Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).

The fee currency code must match the currency associated with the seller
that is accepting the payment. The application must be from a developer
account in the same country and using the same currency code as the seller.

For more information about the application fee scenario, see
[Take Payments and Collect Fees](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees).

To set this field, `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required.
For more information, see [Permissions](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees#permissions).
    
</dd>
</dl>

<dl>
<dd>

**delayDuration:** `*string` 

The duration of time after the payment's creation when Square automatically
either completes or cancels the payment depending on the `delay_action` field value.
For more information, see
[Time threshold](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture#time-threshold).

This parameter should be specified as a time duration, in RFC 3339 format.

Note: This feature is only supported for card payments. This parameter can only be set for a delayed
capture payment (`autocomplete=false`).

Default:

- Card-present payments: "PT36H" (36 hours) from the creation time.
- Card-not-present payments: "P7D" (7 days) from the creation time.
    
</dd>
</dl>

<dl>
<dd>

**delayAction:** `*string` 

The action to be applied to the payment when the `delay_duration` has elapsed. The action must be
CANCEL or COMPLETE. For more information, see
[Time Threshold](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture#time-threshold).

Default: CANCEL
    
</dd>
</dl>

<dl>
<dd>

**autocomplete:** `*bool` 

If set to `true`, this payment will be completed when possible. If
set to `false`, this payment is held in an approved state until either
explicitly completed (captured) or canceled (voided). For more information, see
[Delayed capture](https://developer.squareup.com/docs/payments-api/take-payments/card-payments#delayed-capture-of-a-card-payment).

Default: true
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*string` — Associates a previously created order with this payment.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` 

The [Customer](entity:Customer) ID of the customer associated with the payment.

This is required if the `source_id` refers to a card on file created using the Cards API.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

The location ID to associate with the payment. If not specified, the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location) is
used.
    
</dd>
</dl>

<dl>
<dd>

**teamMemberID:** `*string` 

An optional [TeamMember](entity:TeamMember) ID to associate with
this payment.
    
</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` 

A user-defined ID to associate with the payment.

You can use this field to associate the payment to an entity in an external system
(for example, you might specify an order ID that is generated by a third-party shopping cart).
    
</dd>
</dl>

<dl>
<dd>

**verificationToken:** `*string` 

An identifying token generated by [payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer).
Verification tokens encapsulate customer device information and 3-D Secure
challenge results to indicate that Square has verified the buyer identity.

For more information, see [SCA Overview](https://developer.squareup.com/docs/sca-overview).
    
</dd>
</dl>

<dl>
<dd>

**acceptPartialAuthorization:** `*bool` 

If set to `true` and charging a Square Gift Card, a payment might be returned with
`amount_money` equal to less than what was requested. For example, a request for $20 when charging
a Square Gift Card with a balance of $5 results in an APPROVED payment of $5. You might choose
to prompt the buyer for an additional payment to cover the remainder or cancel the Gift Card
payment. This field cannot be `true` when `autocomplete = true`.

For more information, see
[Partial amount with Square Gift Cards](https://developer.squareup.com/docs/payments-api/take-payments#partial-payment-gift-card).

Default: false
    
</dd>
</dl>

<dl>
<dd>

**buyerEmailAddress:** `*string` — The buyer's email address.
    
</dd>
</dl>

<dl>
<dd>

**buyerPhoneNumber:** `*string` 

The buyer's phone number.
Must follow the following format:
1. A leading + symbol (followed by a country code)
2. The phone number can contain spaces and the special characters `(` , `)` , `-` , and `.`.
Alphabetical characters aren't allowed.
3. The phone number must contain between 9 and 16 digits.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*square.Address` — The buyer's billing address.
    
</dd>
</dl>

<dl>
<dd>

**shippingAddress:** `*square.Address` — The buyer's shipping address.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` — An optional note to be entered by the developer when creating a payment.
    
</dd>
</dl>

<dl>
<dd>

**statementDescriptionIdentifier:** `*string` 

Optional additional payment information to include on the customer's card statement
as part of the statement description. This can be, for example, an invoice number, ticket number,
or short description that uniquely identifies the purchase.

Note that the `statement_description_identifier` might get truncated on the statement description
to fit the required information including the Square identifier (SQ *) and name of the
seller taking the payment.
    
</dd>
</dl>

<dl>
<dd>

**cashDetails:** `*square.CashPaymentDetails` — Additional details required when recording a cash payment (`source_id` is CASH).
    
</dd>
</dl>

<dl>
<dd>

**externalDetails:** `*square.ExternalPaymentDetails` — Additional details required when recording an external payment (`source_id` is EXTERNAL).
    
</dd>
</dl>

<dl>
<dd>

**customerDetails:** `*square.CustomerDetails` — Details about the customer making the payment.
    
</dd>
</dl>

<dl>
<dd>

**offlinePaymentDetails:** `*square.OfflinePaymentDetails` 

An optional field for specifying the offline payment details. This is intended for
internal 1st-party callers only.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.CancelByIdempotencyKey(request) -> *square.CancelPaymentByIdempotencyKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels (voids) a payment identified by the idempotency key that is specified in the
request.

Use this method when the status of a `CreatePayment` request is unknown (for example, after you send a
`CreatePayment` request, a network error occurs and you do not get a response). In this case, you can
direct Square to cancel the payment using this endpoint. In the request, you provide the same
idempotency key that you provided in your `CreatePayment` request that you want to cancel. After
canceling the payment, you can submit your `CreatePayment` request again.

Note that if no payment with the specified idempotency key is found, no action is taken and the endpoint
returns successfully.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelPaymentByIdempotencyKeyRequest{
        IdempotencyKey: "a7e36d40-d24b-11e8-b568-0800200c9a66",
    }
client.Payments.CancelByIdempotencyKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — The `idempotency_key` identifying the payment to be canceled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Get(PaymentID) -> *square.GetPaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves details for a specific payment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetPaymentsRequest{
        PaymentID: "payment_id",
    }
client.Payments.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — A unique ID for the desired payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Update(PaymentID, request) -> *square.UpdatePaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a payment with the APPROVED status.
You can update the `amount_money` and `tip_money` using this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdatePaymentRequest{
        PaymentID: "payment_id",
        Payment: &square.Payment{
            AmountMoney: &square.Money{
                Amount: square.Int64(
                    1000,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
            TipMoney: &square.Money{
                Amount: square.Int64(
                    100,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
            VersionToken: square.String(
                "ODhwVQ35xwlzRuoZEwKXucfu7583sPTzK48c5zoGd0g6o",
            ),
        },
        IdempotencyKey: "956f8b13-e4ec-45d6-85e8-d1d95ef0c5de",
    }
client.Payments.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — The ID of the payment to update.
    
</dd>
</dl>

<dl>
<dd>

**payment:** `*square.Payment` — The updated `Payment` object.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `UpdatePayment` request. Keys can be any valid string
but must be unique for every `UpdatePayment` request.

For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Cancel(PaymentID) -> *square.CancelPaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels (voids) a payment. You can use this endpoint to cancel a payment with 
the APPROVED `status`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelPaymentsRequest{
        PaymentID: "payment_id",
    }
client.Payments.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — The ID of the payment to cancel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Complete(PaymentID, request) -> *square.CompletePaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Completes (captures) a payment.
By default, payments are set to complete immediately after they are created.

You can use this endpoint to complete a payment with the APPROVED `status`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CompletePaymentRequest{
        PaymentID: "payment_id",
    }
client.Payments.Complete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — The unique ID identifying the payment to be completed.
    
</dd>
</dl>

<dl>
<dd>

**versionToken:** `*string` 

Used for optimistic concurrency. This opaque token identifies the current `Payment`
version that the caller expects. If the server has a different version of the Payment,
the update fails and a response with a VERSION_MISMATCH error is returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payouts
<details><summary><code>client.Payouts.List() -> *square.ListPayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of all payouts for the default location.
You can filter payouts by location ID, status, time range, and order them in ascending or descending order.
To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListPayoutsRequest{
        LocationID: square.String(
            "location_id",
        ),
        Status: square.PayoutStatusSent.Ptr(),
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        SortOrder: square.SortOrderDesc.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Payouts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `*string` 

The ID of the location for which to list the payouts.
By default, payouts are returned for the default (main) location associated with the seller.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*square.PayoutStatus` — If provided, only payouts with the given status are returned.
    
</dd>
</dl>

<dl>
<dd>

**beginTime:** `*string` 

The timestamp for the beginning of the payout creation time, in RFC 3339 format.
Inclusive. Default: The current time minus one year.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 

The timestamp for the end of the payout creation time, in RFC 3339 format.
Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` — The order in which payouts are listed.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
If request parameters change between requests, subsequent results may contain duplicates or missing records.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.
It is possible to receive fewer results than the specified limit on a given page.
The default value of 100 is also the maximum allowed value. If the provided value is
greater than 100, it is ignored and the default value is used instead.
Default: `100`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Get(PayoutID) -> *square.GetPayoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves details of a specific payout identified by a payout ID.
To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetPayoutsRequest{
        PayoutID: "payout_id",
    }
client.Payouts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**payoutID:** `string` — The ID of the payout to retrieve the information for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.ListEntries(PayoutID) -> *square.ListPayoutEntriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of all payout entries for a specific payout.
To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListEntriesPayoutsRequest{
        PayoutID: "payout_id",
        SortOrder: square.SortOrderDesc.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Payouts.ListEntries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**payoutID:** `string` — The ID of the payout to retrieve the information for.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` — The order in which payout entries are listed.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
If request parameters change between requests, subsequent results may contain duplicates or missing records.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.
It is possible to receive fewer results than the specified limit on a given page.
The default value of 100 is also the maximum allowed value. If the provided value is
greater than 100, it is ignored and the default value is used instead.
Default: `100`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Refunds
<details><summary><code>client.Refunds.List() -> *square.ListPaymentRefundsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of refunds for the account making the request.

Results are eventually consistent, and new refunds or changes to refunds might take several
seconds to appear.

The maximum results per page is 100.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListRefundsRequest{
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        SortOrder: square.String(
            "sort_order",
        ),
        Cursor: square.String(
            "cursor",
        ),
        LocationID: square.String(
            "location_id",
        ),
        Status: square.String(
            "status",
        ),
        SourceType: square.String(
            "source_type",
        ),
        Limit: square.Int(
            1,
        ),
        UpdatedAtBeginTime: square.String(
            "updated_at_begin_time",
        ),
        UpdatedAtEndTime: square.String(
            "updated_at_end_time",
        ),
        SortField: square.ListPaymentRefundsRequestSortFieldCreatedAt.Ptr(),
    }
client.Refunds.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**beginTime:** `*string` 

Indicates the start of the time range to retrieve each `PaymentRefund` for, in RFC 3339 
format.  The range is determined using the `created_at` field for each `PaymentRefund`. 

Default: The current time minus one year.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 

Indicates the end of the time range to retrieve each `PaymentRefund` for, in RFC 3339 
format.  The range is determined using the `created_at` field for each `PaymentRefund`.

Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*string` 

The order in which results are listed by `PaymentRefund.created_at`:
- `ASC` - Oldest to newest.
- `DESC` - Newest to oldest (default).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

Limit results to the location supplied. By default, results are returned
for all locations associated with the seller.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 

If provided, only refunds with the given status are returned.
For a list of refund status values, see [PaymentRefund](entity:PaymentRefund).

Default: If omitted, refunds are returned regardless of their status.
    
</dd>
</dl>

<dl>
<dd>

**sourceType:** `*string` 

If provided, only returns refunds whose payments have the indicated source type.
Current values include `CARD`, `BANK_ACCOUNT`, `WALLET`, `CASH`, and `EXTERNAL`.
For information about these payment source types, see
[Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).

Default: If omitted, refunds are returned regardless of the source type.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.

It is possible to receive fewer results than the specified limit on a given page.

If the supplied value is greater than 100, no more than 100 results are returned.

Default: 100
    
</dd>
</dl>

<dl>
<dd>

**updatedAtBeginTime:** `*string` 

Indicates the start of the time range to retrieve each `PaymentRefund` for, in RFC 3339
format.  The range is determined using the `updated_at` field for each `PaymentRefund`.

Default: If omitted, the time range starts at `begin_time`.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtEndTime:** `*string` 

Indicates the end of the time range to retrieve each `PaymentRefund` for, in RFC 3339
format.  The range is determined using the `updated_at` field for each `PaymentRefund`.

Default: The current time.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*square.ListPaymentRefundsRequestSortField` — The field used to sort results by. The default is `CREATED_AT`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Refunds.RefundPayment(request) -> *square.RefundPaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refunds a payment. You can refund the entire payment amount or a
portion of it. You can use this endpoint to refund a card payment or record a 
refund of a cash or external payment. For more information, see
[Refund Payment](https://developer.squareup.com/docs/payments-api/refund-payments).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RefundPaymentRequest{
        IdempotencyKey: "9b7f2dcf-49da-4411-b23e-a2d6af21333a",
        AmountMoney: &square.Money{
            Amount: square.Int64(
                1000,
            ),
            Currency: square.CurrencyUsd.Ptr(),
        },
        AppFeeMoney: &square.Money{
            Amount: square.Int64(
                10,
            ),
            Currency: square.CurrencyUsd.Ptr(),
        },
        PaymentID: square.String(
            "R2B3Z8WMVt3EAmzYWLZvz7Y69EbZY",
        ),
        Reason: square.String(
            "Example",
        ),
    }
client.Refunds.RefundPayment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

 A unique string that identifies this `RefundPayment` request. The key can be any valid string
but must be unique for every `RefundPayment` request.

Keys are limited to a max of 45 characters - however, the number of allowed characters might be
less than 45, if multi-byte characters are used.

For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**amountMoney:** `*square.Money` 

The amount of money to refund.

This amount cannot be more than the `total_money` value of the payment minus the total
amount of all previously completed refunds for this payment.

This amount must be specified in the smallest denomination of the applicable currency
(for example, US dollar amounts are specified in cents). For more information, see
[Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).

The currency code must match the currency associated with the business
that is charging the card.
    
</dd>
</dl>

<dl>
<dd>

**appFeeMoney:** `*square.Money` 

The amount of money the developer contributes to help cover the refunded amount.
This amount is specified in the smallest denomination of the applicable currency (for example,
US dollar amounts are specified in cents).

The value cannot be more than the `amount_money`.

You can specify this parameter in a refund request only if the same parameter was also included
when taking the payment. This is part of the application fee scenario the API supports. For more
information, see [Take Payments and Collect Fees](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees).

To set this field, `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required.
For more information, see [Permissions](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees#permissions).
    
</dd>
</dl>

<dl>
<dd>

**paymentID:** `*string` 

The unique ID of the payment being refunded.
Required when unlinked=false, otherwise must not be set.
    
</dd>
</dl>

<dl>
<dd>

**destinationID:** `*string` 

The ID indicating where funds will be refunded to. Required for unlinked refunds. For more
information, see [Process an Unlinked Refund](https://developer.squareup.com/docs/refunds-api/unlinked-refunds).

For refunds linked to Square payments, `destination_id` is usually omitted; in this case, funds
will be returned to the original payment source. The field may be specified in order to request
a cross-method refund to a gift card. For more information,
see [Cross-method refunds to gift cards](https://developer.squareup.com/docs/payments-api/refund-payments#cross-method-refunds-to-gift-cards).
    
</dd>
</dl>

<dl>
<dd>

**unlinked:** `*bool` 

Indicates that the refund is not linked to a Square payment.
If set to true, `destination_id` and `location_id` must be supplied while `payment_id` must not
be provided.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

The location ID associated with the unlinked refund.
Required for requests specifying `unlinked=true`.
Otherwise, if included when `unlinked=false`, will throw an error.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` 

The [Customer](entity:Customer) ID of the customer associated with the refund.
This is required if the `destination_id` refers to a card on file created using the Cards
API. Only allowed when `unlinked=true`.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` — A description of the reason for the refund.
    
</dd>
</dl>

<dl>
<dd>

**paymentVersionToken:** `*string` 

 Used for optimistic concurrency. This opaque token identifies the current `Payment`
version that the caller expects. If the server has a different version of the Payment,
the update fails and a response with a VERSION_MISMATCH error is returned.
If the versions match, or the field is not provided, the refund proceeds as normal.
    
</dd>
</dl>

<dl>
<dd>

**teamMemberID:** `*string` — An optional [TeamMember](entity:TeamMember) ID to associate with this refund.
    
</dd>
</dl>

<dl>
<dd>

**cashDetails:** `*square.DestinationDetailsCashRefundDetails` — Additional details required when recording an unlinked cash refund (`destination_id` is CASH).
    
</dd>
</dl>

<dl>
<dd>

**externalDetails:** `*square.DestinationDetailsExternalRefundDetails` 

Additional details required when recording an unlinked external refund
(`destination_id` is EXTERNAL).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Refunds.Get(RefundID) -> *square.GetPaymentRefundResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specific refund using the `refund_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetRefundsRequest{
        RefundID: "refund_id",
    }
client.Refunds.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundID:** `string` — The unique ID for the desired `PaymentRefund`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sites
<details><summary><code>client.Sites.List() -> *square.ListSitesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the Square Online sites that belong to a seller. Sites are listed in descending order by the `created_at` date.


__Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sites.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Snippets
<details><summary><code>client.Snippets.Get(SiteID) -> *square.GetSnippetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves your snippet from a Square Online site. A site can contain snippets from multiple snippet applications, but you can retrieve only the snippet that was added by your application.

You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.


__Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetSnippetsRequest{
        SiteID: "site_id",
    }
client.Snippets.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**siteID:** `string` — The ID of the site that contains the snippet.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Snippets.Upsert(SiteID, request) -> *square.UpsertSnippetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a snippet to a Square Online site or updates the existing snippet on the site. 
The snippet code is appended to the end of the `head` element on every page of the site, except checkout pages. A snippet application can add one snippet to a given site. 

You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.


__Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpsertSnippetRequest{
        SiteID: "site_id",
        Snippet: &square.Snippet{
            Content: "<script>var js = 1;</script>",
        },
    }
client.Snippets.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**siteID:** `string` — The ID of the site where you want to add or update the snippet.
    
</dd>
</dl>

<dl>
<dd>

**snippet:** `*square.Snippet` — The snippet for the site.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Snippets.Delete(SiteID) -> *square.DeleteSnippetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes your snippet from a Square Online site.

You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.


__Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteSnippetsRequest{
        SiteID: "site_id",
    }
client.Snippets.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**siteID:** `string` — The ID of the site that contains the snippet.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscriptions
<details><summary><code>client.Subscriptions.Create(request) -> *square.CreateSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enrolls a customer in a subscription.

If you provide a card on file in the request, Square charges the card for
the subscription. Otherwise, Square sends an invoice to the customer's email
address. The subscription starts immediately, unless the request includes
the optional `start_date`. Each individual subscription is associated with a particular location.

For more information, see [Create a subscription](https://developer.squareup.com/docs/subscriptions-api/manage-subscriptions#create-a-subscription).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateSubscriptionRequest{
        IdempotencyKey: square.String(
            "8193148c-9586-11e6-99f9-28cfe92138cf",
        ),
        LocationID: "S8GWD5R9QB376",
        PlanVariationID: square.String(
            "6JHXF3B2CW3YKHDV4XEM674H",
        ),
        CustomerID: "CHFGVKYY8RSV93M5KCYTG4PN0G",
        StartDate: square.String(
            "2023-06-20",
        ),
        CardID: square.String(
            "ccof:qy5x8hHGYsgLrp4Q4GB",
        ),
        Timezone: square.String(
            "America/Los_Angeles",
        ),
        Source: &square.SubscriptionSource{
            Name: square.String(
                "My Application",
            ),
        },
        Phases: []*square.Phase{
            &square.Phase{
                Ordinal: square.Int64(
                    0,
                ),
                OrderTemplateID: square.String(
                    "U2NaowWxzXwpsZU697x7ZHOAnCNZY",
                ),
            },
        },
    }
client.Subscriptions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique string that identifies this `CreateSubscription` request.
If you do not provide a unique string (or provide an empty string as the value),
the endpoint treats each request as independent.

For more information, see [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The ID of the location the subscription is associated with.
    
</dd>
</dl>

<dl>
<dd>

**planVariationID:** `*string` — The ID of the [subscription plan variation](https://developer.squareup.com/docs/subscriptions-api/plans-and-variations#plan-variations) created using the Catalog API.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The ID of the [customer](entity:Customer) subscribing to the subscription plan variation.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` 

The `YYYY-MM-DD`-formatted date to start the subscription. 
If it is unspecified, the subscription starts immediately.
    
</dd>
</dl>

<dl>
<dd>

**canceledDate:** `*string` 

The `YYYY-MM-DD`-formatted date when the newly created subscription is scheduled for cancellation. 

This date overrides the cancellation date set in the plan variation configuration.
If the cancellation date is earlier than the end date of a subscription cycle, the subscription stops
at the canceled date and the subscriber is sent a prorated invoice at the beginning of the canceled cycle. 

When the subscription plan of the newly created subscription has a fixed number of cycles and the `canceled_date`
occurs before the subscription plan completes, the specified `canceled_date` sets the date when the subscription
stops through the end of the last cycle.
    
</dd>
</dl>

<dl>
<dd>

**taxPercentage:** `*string` 

The tax to add when billing the subscription.
The percentage is expressed in decimal form, using a `'.'` as the decimal
separator and without a `'%'` sign. For example, a value of 7.5
corresponds to 7.5%.
    
</dd>
</dl>

<dl>
<dd>

**priceOverrideMoney:** `*square.Money` 

A custom price which overrides the cost of a subscription plan variation with `STATIC` pricing.
This field does not affect itemized subscriptions with `RELATIVE` pricing. Instead, 
you should edit the Subscription's [order template](https://developer.squareup.com/docs/subscriptions-api/manage-subscriptions#phases-and-order-templates).
    
</dd>
</dl>

<dl>
<dd>

**cardID:** `*string` 

The ID of the [subscriber's](entity:Customer) [card](entity:Card) to charge.
If it is not specified, the subscriber receives an invoice via email with a link to pay for their subscription.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` 

The timezone that is used in date calculations for the subscription. If unset, defaults to
the location timezone. If a timezone is not configured for the location, defaults to "America/New_York".
Format: the IANA Timezone Database identifier for the location timezone. For
a list of time zones, see [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*square.SubscriptionSource` — The origination details of the subscription.
    
</dd>
</dl>

<dl>
<dd>

**monthlyBillingAnchorDate:** `*int` — The day-of-the-month to change the billing date to.
    
</dd>
</dl>

<dl>
<dd>

**phases:** `[]*square.Phase` — array of phases for this subscription
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.BulkSwapPlan(request) -> *square.BulkSwapPlanResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedules a plan variation change for all active subscriptions under a given plan
variation. For more information, see [Swap Subscription Plan Variations](https://developer.squareup.com/docs/subscriptions-api/swap-plan-variations).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BulkSwapPlanRequest{
        NewPlanVariationID: "FQ7CDXXWSLUJRPM3GFJSJGZ7",
        OldPlanVariationID: "6JHXF3B2CW3YKHDV4XEM674H",
        LocationID: "S8GWD5R9QB376",
    }
client.Subscriptions.BulkSwapPlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**newPlanVariationID:** `string` 

The ID of the new subscription plan variation.

This field is required.
    
</dd>
</dl>

<dl>
<dd>

**oldPlanVariationID:** `string` 

The ID of the plan variation whose subscriptions should be swapped. Active subscriptions
using this plan variation will be subscribed to the new plan variation on their next billing
day.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The ID of the location to associate with the swapped subscriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Search(request) -> *square.SearchSubscriptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for subscriptions.

Results are ordered chronologically by subscription creation date. If
the request specifies more than one location ID,
the endpoint orders the result
by location ID, and then by creation date within each location. If no locations are given
in the query, all locations are searched.

You can also optionally specify `customer_ids` to search by customer.
If left unset, all customers
associated with the specified locations are returned.
If the request specifies customer IDs, the endpoint orders results
first by location, within location by customer ID, and within
customer by subscription creation date.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchSubscriptionsRequest{
        Query: &square.SearchSubscriptionsQuery{
            Filter: &square.SearchSubscriptionsFilter{
                CustomerIDs: []string{
                    "CHFGVKYY8RSV93M5KCYTG4PN0G",
                },
                LocationIDs: []string{
                    "S8GWD5R9QB376",
                },
                SourceNames: []string{
                    "My App",
                },
            },
        },
    }
client.Subscriptions.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

When the total number of resulting subscriptions exceeds the limit of a paged response, 
specify the cursor returned from a preceding response here to fetch the next set of results.
If the cursor is unset, the response contains the last page of the results.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The upper limit on the number of subscriptions to return
in a paged response.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*square.SearchSubscriptionsQuery` 

A subscription query consisting of specified filtering conditions.

If this `query` field is unspecified, the `SearchSubscriptions` call will return all subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**include:** `[]string` 

An option to include related information in the response. 

The supported values are: 

- `actions`: to include scheduled actions on the targeted subscriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Get(SubscriptionID) -> *square.GetSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specific subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetSubscriptionsRequest{
        SubscriptionID: "subscription_id",
        Include: square.String(
            "include",
        ),
    }
client.Subscriptions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` 

A query parameter to specify related information to be included in the response. 

The supported query parameter values are: 

- `actions`: to include scheduled actions on the targeted subscription.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Update(SubscriptionID, request) -> *square.UpdateSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a subscription by modifying or clearing `subscription` field values.
To clear a field, set its value to `null`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateSubscriptionRequest{
        SubscriptionID: "subscription_id",
        Subscription: &square.Subscription{
            CardID: square.String(
                "{NEW CARD ID}",
            ),
        },
    }
client.Subscriptions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to update.
    
</dd>
</dl>

<dl>
<dd>

**subscription:** `*square.Subscription` 

The subscription object containing the current version, and fields to update.
Unset fields will be left at their current server values, and JSON `null` values will
be treated as a request to clear the relevant data.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.DeleteAction(SubscriptionID, ActionID) -> *square.DeleteSubscriptionActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a scheduled action for a subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteActionSubscriptionsRequest{
        SubscriptionID: "subscription_id",
        ActionID: "action_id",
    }
client.Subscriptions.DeleteAction(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription the targeted action is to act upon.
    
</dd>
</dl>

<dl>
<dd>

**actionID:** `string` — The ID of the targeted action to be deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.ChangeBillingAnchorDate(SubscriptionID, request) -> *square.ChangeBillingAnchorDateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes the [billing anchor date](https://developer.squareup.com/docs/subscriptions-api/subscription-billing#billing-dates)
for a subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ChangeBillingAnchorDateRequest{
        SubscriptionID: "subscription_id",
        MonthlyBillingAnchorDate: square.Int(
            1,
        ),
    }
client.Subscriptions.ChangeBillingAnchorDate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to update the billing anchor date.
    
</dd>
</dl>

<dl>
<dd>

**monthlyBillingAnchorDate:** `*int` — The anchor day for the billing cycle.
    
</dd>
</dl>

<dl>
<dd>

**effectiveDate:** `*string` 

The `YYYY-MM-DD`-formatted date when the scheduled `BILLING_ANCHOR_CHANGE` action takes
place on the subscription.

When this date is unspecified or falls within the current billing cycle, the billing anchor date
is changed immediately.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Cancel(SubscriptionID) -> *square.CancelSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedules a `CANCEL` action to cancel an active subscription. This 
sets the `canceled_date` field to the end of the active billing period. After this date, 
the subscription status changes from ACTIVE to CANCELED.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelSubscriptionsRequest{
        SubscriptionID: "subscription_id",
    }
client.Subscriptions.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to cancel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.ListEvents(SubscriptionID) -> *square.ListSubscriptionEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all [events](https://developer.squareup.com/docs/subscriptions-api/actions-events) for a specific subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListEventsSubscriptionsRequest{
        SubscriptionID: "subscription_id",
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Subscriptions.ListEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to retrieve the events for.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

When the total number of resulting subscription events exceeds the limit of a paged response, 
specify the cursor returned from a preceding response here to fetch the next set of results.
If the cursor is unset, the response contains the last page of the results.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The upper limit on the number of subscription events to return
in a paged response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Pause(SubscriptionID, request) -> *square.PauseSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedules a `PAUSE` action to pause an active subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.PauseSubscriptionRequest{
        SubscriptionID: "subscription_id",
    }
client.Subscriptions.Pause(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to pause.
    
</dd>
</dl>

<dl>
<dd>

**pauseEffectiveDate:** `*string` 

The `YYYY-MM-DD`-formatted date when the scheduled `PAUSE` action takes place on the subscription.

When this date is unspecified or falls within the current billing cycle, the subscription is paused
on the starting date of the next billing cycle.
    
</dd>
</dl>

<dl>
<dd>

**pauseCycleDuration:** `*int64` 

The number of billing cycles the subscription will be paused before it is reactivated. 

When this is set, a `RESUME` action is also scheduled to take place on the subscription at 
the end of the specified pause cycle duration. In this case, neither `resume_effective_date` 
nor `resume_change_timing` may be specified.
    
</dd>
</dl>

<dl>
<dd>

**resumeEffectiveDate:** `*string` 

The date when the subscription is reactivated by a scheduled `RESUME` action. 
This date must be at least one billing cycle ahead of `pause_effective_date`.
    
</dd>
</dl>

<dl>
<dd>

**resumeChangeTiming:** `*square.ChangeTiming` 

The timing whether the subscription is reactivated immediately or at the end of the billing cycle, relative to 
`resume_effective_date`.
See [ChangeTiming](#type-changetiming) for possible values
    
</dd>
</dl>

<dl>
<dd>

**pauseReason:** `*string` — The user-provided reason to pause the subscription.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.Resume(SubscriptionID, request) -> *square.ResumeSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedules a `RESUME` action to resume a paused or a deactivated subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ResumeSubscriptionRequest{
        SubscriptionID: "subscription_id",
    }
client.Subscriptions.Resume(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to resume.
    
</dd>
</dl>

<dl>
<dd>

**resumeEffectiveDate:** `*string` — The `YYYY-MM-DD`-formatted date when the subscription reactivated.
    
</dd>
</dl>

<dl>
<dd>

**resumeChangeTiming:** `*square.ChangeTiming` 

The timing to resume a subscription, relative to the specified
`resume_effective_date` attribute value.
See [ChangeTiming](#type-changetiming) for possible values
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscriptions.SwapPlan(SubscriptionID, request) -> *square.SwapPlanResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedules a `SWAP_PLAN` action to swap a subscription plan variation in an existing subscription. 
For more information, see [Swap Subscription Plan Variations](https://developer.squareup.com/docs/subscriptions-api/swap-plan-variations).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SwapPlanRequest{
        SubscriptionID: "subscription_id",
        NewPlanVariationID: square.String(
            "FQ7CDXXWSLUJRPM3GFJSJGZ7",
        ),
        Phases: []*square.PhaseInput{
            &square.PhaseInput{
                Ordinal: 0,
                OrderTemplateID: square.String(
                    "uhhnjH9osVv3shUADwaC0b3hNxQZY",
                ),
            },
        },
    }
client.Subscriptions.SwapPlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — The ID of the subscription to swap the subscription plan for.
    
</dd>
</dl>

<dl>
<dd>

**newPlanVariationID:** `*string` 

The ID of the new subscription plan variation.

This field is required.
    
</dd>
</dl>

<dl>
<dd>

**phases:** `[]*square.PhaseInput` — A list of PhaseInputs, to pass phase-specific information used in the swap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TeamMembers
<details><summary><code>client.TeamMembers.Create(request) -> *square.CreateTeamMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a single `TeamMember` object. The `TeamMember` object is returned on successful creates.
You must provide the following values in your request to this endpoint:
- `given_name`
- `family_name`

Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#createteammember).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateTeamMemberRequest{
        IdempotencyKey: square.String(
            "idempotency-key-0",
        ),
        TeamMember: &square.TeamMember{
            ReferenceID: square.String(
                "reference_id_1",
            ),
            Status: square.TeamMemberStatusActive.Ptr(),
            GivenName: square.String(
                "Joe",
            ),
            FamilyName: square.String(
                "Doe",
            ),
            EmailAddress: square.String(
                "joe_doe@gmail.com",
            ),
            PhoneNumber: square.String(
                "+14159283333",
            ),
            AssignedLocations: &square.TeamMemberAssignedLocations{
                AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeExplicitLocations.Ptr(),
                LocationIDs: []string{
                    "YSGH2WBKG94QZ",
                    "GA2Y9HSJ8KRYT",
                },
            },
            WageSetting: &square.WageSetting{
                JobAssignments: []*square.JobAssignment{
                    &square.JobAssignment{
                        PayType: square.JobAssignmentPayTypeSalary,
                        AnnualRate: &square.Money{
                            Amount: square.Int64(
                                3000000,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                        WeeklyHours: square.Int(
                            40,
                        ),
                        JobID: square.String(
                            "FjS8x95cqHiMenw4f1NAUH4P",
                        ),
                    },
                    &square.JobAssignment{
                        PayType: square.JobAssignmentPayTypeHourly,
                        HourlyRate: &square.Money{
                            Amount: square.Int64(
                                2000,
                            ),
                            Currency: square.CurrencyUsd.Ptr(),
                        },
                        JobID: square.String(
                            "VDNpRv8da51NU8qZFC5zDWpF",
                        ),
                    },
                },
                IsOvertimeExempt: square.Bool(
                    true,
                ),
            },
        },
    }
client.TeamMembers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*square.CreateTeamMemberRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.BatchCreate(request) -> *square.BatchCreateTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates multiple `TeamMember` objects. The created `TeamMember` objects are returned on successful creates.
This process is non-transactional and processes as much of the request as possible. If one of the creates in
the request cannot be successfully processed, the request is not marked as failed, but the body of the response
contains explicit error information for the failed create.

Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-create-team-members).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchCreateTeamMembersRequest{
        TeamMembers: map[string]*square.CreateTeamMemberRequest{
            "idempotency-key-1": &square.CreateTeamMemberRequest{
                TeamMember: &square.TeamMember{
                    ReferenceID: square.String(
                        "reference_id_1",
                    ),
                    GivenName: square.String(
                        "Joe",
                    ),
                    FamilyName: square.String(
                        "Doe",
                    ),
                    EmailAddress: square.String(
                        "joe_doe@gmail.com",
                    ),
                    PhoneNumber: square.String(
                        "+14159283333",
                    ),
                    AssignedLocations: &square.TeamMemberAssignedLocations{
                        AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeExplicitLocations.Ptr(),
                        LocationIDs: []string{
                            "YSGH2WBKG94QZ",
                            "GA2Y9HSJ8KRYT",
                        },
                    },
                },
            },
            "idempotency-key-2": &square.CreateTeamMemberRequest{
                TeamMember: &square.TeamMember{
                    ReferenceID: square.String(
                        "reference_id_2",
                    ),
                    GivenName: square.String(
                        "Jane",
                    ),
                    FamilyName: square.String(
                        "Smith",
                    ),
                    EmailAddress: square.String(
                        "jane_smith@gmail.com",
                    ),
                    PhoneNumber: square.String(
                        "+14159223334",
                    ),
                    AssignedLocations: &square.TeamMemberAssignedLocations{
                        AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeAllCurrentAndFutureLocations.Ptr(),
                    },
                },
            },
        },
    }
client.TeamMembers.BatchCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMembers:** `map[string]*square.CreateTeamMemberRequest` 

The data used to create the `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`.
The maximum number of create objects is 25.

If you include a team member's `wage_setting`, you must provide `job_id` for each job assignment. To get job IDs,
call [ListJobs](api-endpoint:Team-ListJobs).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.BatchUpdate(request) -> *square.BatchUpdateTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates multiple `TeamMember` objects. The updated `TeamMember` objects are returned on successful updates.
This process is non-transactional and processes as much of the request as possible. If one of the updates in
the request cannot be successfully processed, the request is not marked as failed, but the body of the response
contains explicit error information for the failed update.
Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-update-team-members).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchUpdateTeamMembersRequest{
        TeamMembers: map[string]*square.UpdateTeamMemberRequest{
            "AFMwA08kR-MIF-3Vs0OE": &square.UpdateTeamMemberRequest{
                TeamMember: &square.TeamMember{
                    ReferenceID: square.String(
                        "reference_id_2",
                    ),
                    IsOwner: square.Bool(
                        false,
                    ),
                    Status: square.TeamMemberStatusActive.Ptr(),
                    GivenName: square.String(
                        "Jane",
                    ),
                    FamilyName: square.String(
                        "Smith",
                    ),
                    EmailAddress: square.String(
                        "jane_smith@gmail.com",
                    ),
                    PhoneNumber: square.String(
                        "+14159223334",
                    ),
                    AssignedLocations: &square.TeamMemberAssignedLocations{
                        AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeAllCurrentAndFutureLocations.Ptr(),
                    },
                },
            },
            "fpgteZNMaf0qOK-a4t6P": &square.UpdateTeamMemberRequest{
                TeamMember: &square.TeamMember{
                    ReferenceID: square.String(
                        "reference_id_1",
                    ),
                    IsOwner: square.Bool(
                        false,
                    ),
                    Status: square.TeamMemberStatusActive.Ptr(),
                    GivenName: square.String(
                        "Joe",
                    ),
                    FamilyName: square.String(
                        "Doe",
                    ),
                    EmailAddress: square.String(
                        "joe_doe@gmail.com",
                    ),
                    PhoneNumber: square.String(
                        "+14159283333",
                    ),
                    AssignedLocations: &square.TeamMemberAssignedLocations{
                        AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeExplicitLocations.Ptr(),
                        LocationIDs: []string{
                            "YSGH2WBKG94QZ",
                            "GA2Y9HSJ8KRYT",
                        },
                    },
                },
            },
        },
    }
client.TeamMembers.BatchUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMembers:** `map[string]*square.UpdateTeamMemberRequest` 

The data used to update the `TeamMember` objects. Each key is the `team_member_id` that maps to the `UpdateTeamMemberRequest`.
The maximum number of update objects is 25.

For each team member, include the fields to add, change, or clear. Fields can be cleared using a null value.
To update `wage_setting.job_assignments`, you must provide the complete list of job assignments. If needed,
call [ListJobs](api-endpoint:Team-ListJobs) to get the required `job_id` values.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Search(request) -> *square.SearchTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `TeamMember` objects for a business. 
The list can be filtered by location IDs, `ACTIVE` or `INACTIVE` status, or whether
the team member is the Square account owner.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchTeamMembersRequest{
        Query: &square.SearchTeamMembersQuery{
            Filter: &square.SearchTeamMembersFilter{
                LocationIDs: []string{
                    "0G5P3VGACMMQZ",
                },
                Status: square.TeamMemberStatusActive.Ptr(),
            },
        },
        Limit: square.Int(
            10,
        ),
    }
client.TeamMembers.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.SearchTeamMembersQuery` — The query parameters.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of `TeamMember` objects in a page (100 by default).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The opaque cursor for fetching the next page. For more information, see
[pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Get(TeamMemberID) -> *square.GetTeamMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a `TeamMember` object for the given `TeamMember.id`.
Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrieve-a-team-member).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetTeamMembersRequest{
        TeamMemberID: "team_member_id",
    }
client.TeamMembers.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `string` — The ID of the team member to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Update(TeamMemberID, request) -> *square.UpdateTeamMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a single `TeamMember` object. The `TeamMember` object is returned on successful updates.
Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#update-a-team-member).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateTeamMembersRequest{
        TeamMemberID: "team_member_id",
        Body: &square.UpdateTeamMemberRequest{
            TeamMember: &square.TeamMember{
                ReferenceID: square.String(
                    "reference_id_1",
                ),
                Status: square.TeamMemberStatusActive.Ptr(),
                GivenName: square.String(
                    "Joe",
                ),
                FamilyName: square.String(
                    "Doe",
                ),
                EmailAddress: square.String(
                    "joe_doe@gmail.com",
                ),
                PhoneNumber: square.String(
                    "+14159283333",
                ),
                AssignedLocations: &square.TeamMemberAssignedLocations{
                    AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeExplicitLocations.Ptr(),
                    LocationIDs: []string{
                        "YSGH2WBKG94QZ",
                        "GA2Y9HSJ8KRYT",
                    },
                },
                WageSetting: &square.WageSetting{
                    JobAssignments: []*square.JobAssignment{
                        &square.JobAssignment{
                            PayType: square.JobAssignmentPayTypeSalary,
                            AnnualRate: &square.Money{
                                Amount: square.Int64(
                                    3000000,
                                ),
                                Currency: square.CurrencyUsd.Ptr(),
                            },
                            WeeklyHours: square.Int(
                                40,
                            ),
                            JobID: square.String(
                                "FjS8x95cqHiMenw4f1NAUH4P",
                            ),
                        },
                        &square.JobAssignment{
                            PayType: square.JobAssignmentPayTypeHourly,
                            HourlyRate: &square.Money{
                                Amount: square.Int64(
                                    1200,
                                ),
                                Currency: square.CurrencyUsd.Ptr(),
                            },
                            JobID: square.String(
                                "VDNpRv8da51NU8qZFC5zDWpF",
                            ),
                        },
                    },
                    IsOvertimeExempt: square.Bool(
                        true,
                    ),
                },
            },
        },
    }
client.TeamMembers.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `string` — The ID of the team member to update.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*square.UpdateTeamMemberRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Team
<details><summary><code>client.Team.ListJobs() -> *square.ListJobsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists jobs in a seller account. Results are sorted by title in ascending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ListJobsRequest{
        Cursor: square.String(
            "cursor",
        ),
    }
client.Team.ListJobs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

The pagination cursor returned by the previous call to this endpoint. Provide this
cursor to retrieve the next page of results for your original request. For more information,
see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.CreateJob(request) -> *square.CreateJobResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a job in a seller account. A job defines a title and tip eligibility. Note that
compensation is defined in a [job assignment](entity:JobAssignment) in a team member's wage setting.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateJobRequest{
        Job: &square.Job{
            Title: square.String(
                "Cashier",
            ),
            IsTipEligible: square.Bool(
                true,
            ),
        },
        IdempotencyKey: "idempotency-key-0",
    }
client.Team.CreateJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**job:** `*square.Job` — The job to create. The `title` field is required and `is_tip_eligible` defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique identifier for the `CreateJob` request. Keys can be any valid string,
but must be unique for each request. For more information, see
[Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.RetrieveJob(JobID) -> *square.RetrieveJobResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specified job.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.RetrieveJobRequest{
        JobID: "job_id",
    }
client.Team.RetrieveJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobID:** `string` — The ID of the job to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.UpdateJob(JobID, request) -> *square.UpdateJobResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the title or tip eligibility of a job. Changes to the title propagate to all
`JobAssignment`, `Shift`, and `TeamMemberWage` objects that reference the job ID. Changes to
tip eligibility propagate to all `TeamMemberWage` objects that reference the job ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateJobRequest{
        JobID: "job_id",
        Job: &square.Job{
            Title: square.String(
                "Cashier 1",
            ),
            IsTipEligible: square.Bool(
                true,
            ),
        },
    }
client.Team.UpdateJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobID:** `string` — The ID of the job to update.
    
</dd>
</dl>

<dl>
<dd>

**job:** `*square.Job` 

The job with the updated fields, either `title`, `is_tip_eligible`, or both. Only changed fields need
to be included in the request. Optionally include `version` to enable optimistic concurrency control.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Terminal
<details><summary><code>client.Terminal.DismissTerminalAction(ActionID) -> *square.DismissTerminalActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Dismisses a Terminal action request if the status and type of the request permits it.

See [Link and Dismiss Actions](https://developer.squareup.com/docs/terminal-api/advanced-features/custom-workflows/link-and-dismiss-actions) for more details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DismissTerminalActionRequest{
        ActionID: "action_id",
    }
client.Terminal.DismissTerminalAction(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — Unique ID for the `TerminalAction` associated with the action to be dismissed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.DismissTerminalCheckout(CheckoutID) -> *square.DismissTerminalCheckoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Dismisses a Terminal checkout request if the status and type of the request permits it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DismissTerminalCheckoutRequest{
        CheckoutID: "checkout_id",
    }
client.Terminal.DismissTerminalCheckout(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**checkoutID:** `string` — Unique ID for the `TerminalCheckout` associated with the checkout to be dismissed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.DismissTerminalRefund(TerminalRefundID) -> *square.DismissTerminalRefundResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Dismisses a Terminal refund request if the status and type of the request permits it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DismissTerminalRefundRequest{
        TerminalRefundID: "terminal_refund_id",
    }
client.Terminal.DismissTerminalRefund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**terminalRefundID:** `string` — Unique ID for the `TerminalRefund` associated with the refund to be dismissed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TransferOrders
<details><summary><code>client.TransferOrders.Create(request) -> *square.CreateTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new transfer order in [DRAFT](entity:TransferOrderStatus) status. A transfer order represents the intent 
to move [CatalogItemVariation](entity:CatalogItemVariation)s from one [Location](entity:Location) to another. 
The source and destination locations must be different and must belong to your Square account.

In [DRAFT](entity:TransferOrderStatus) status, you can:
- Add or remove items
- Modify quantities
- Update shipping information
- Delete the entire order via [DeleteTransferOrder](api-endpoint:TransferOrders-DeleteTransferOrder)

The request requires source_location_id and destination_location_id.
Inventory levels are not affected until the order is started via 
[StartTransferOrder](api-endpoint:TransferOrders-StartTransferOrder).

Common integration points:
- Sync with warehouse management systems
- Automate regular stock transfers
- Initialize transfers from inventory optimization systems

Creates a [transfer_order.created](webhook:transfer_order.created) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateTransferOrderRequest{
        IdempotencyKey: "65cc0586-3e82-384s-b524-3885cffd52",
        TransferOrder: &square.CreateTransferOrderData{
            SourceLocationID: "EXAMPLE_SOURCE_LOCATION_ID_123",
            DestinationLocationID: "EXAMPLE_DEST_LOCATION_ID_456",
            ExpectedAt: square.String(
                "2025-11-09T05:00:00Z",
            ),
            Notes: square.String(
                "Example transfer order for inventory redistribution between locations",
            ),
            TrackingNumber: square.String(
                "TRACK123456789",
            ),
            CreatedByTeamMemberID: square.String(
                "EXAMPLE_TEAM_MEMBER_ID_789",
            ),
            LineItems: []*square.CreateTransferOrderLineData{
                &square.CreateTransferOrderLineData{
                    ItemVariationID: "EXAMPLE_ITEM_VARIATION_ID_001",
                    QuantityOrdered: "5",
                },
                &square.CreateTransferOrderLineData{
                    ItemVariationID: "EXAMPLE_ITEM_VARIATION_ID_002",
                    QuantityOrdered: "3",
                },
            },
        },
    }
client.TransferOrders.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this CreateTransferOrder request. Keys can be
any valid string but must be unique for every CreateTransferOrder request.
    
</dd>
</dl>

<dl>
<dd>

**transferOrder:** `*square.CreateTransferOrderData` — The transfer order to create
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Search(request) -> *square.SearchTransferOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for transfer orders using filters. Returns a paginated list of matching
[TransferOrder](entity:TransferOrder)s sorted by creation date.

Common search scenarios:
- Find orders for a source [Location](entity:Location)
- Find orders for a destination [Location](entity:Location)
- Find orders in a particular [TransferOrderStatus](entity:TransferOrderStatus)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchTransferOrdersRequest{
        Query: &square.TransferOrderQuery{
            Filter: &square.TransferOrderFilter{
                SourceLocationIDs: []string{
                    "EXAMPLE_SOURCE_LOCATION_ID_123",
                },
                DestinationLocationIDs: []string{
                    "EXAMPLE_DEST_LOCATION_ID_456",
                },
                Statuses: []square.TransferOrderStatus{
                    square.TransferOrderStatusStarted,
                    square.TransferOrderStatusPartiallyReceived,
                },
            },
            Sort: &square.TransferOrderSort{
                Field: square.TransferOrderSortFieldUpdatedAt.Ptr(),
                Order: square.SortOrderDesc.Ptr(),
            },
        },
        Cursor: square.String(
            "eyJsYXN0X3VwZGF0ZWRfYXQiOjE3NTMxMTg2NjQ4NzN9",
        ),
        Limit: square.Int(
            10,
        ),
    }
client.TransferOrders.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.TransferOrderQuery` — The search query
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Pagination cursor from a previous search response
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of results to return (1-100)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Get(TransferOrderID) -> *square.RetrieveTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specific [TransferOrder](entity:TransferOrder) by ID. Returns the complete
order details including:

- Basic information (status, dates, notes)
- Line items with ordered and received quantities
- Source and destination [Location](entity:Location)s
- Tracking information (if available)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetTransferOrdersRequest{
        TransferOrderID: "transfer_order_id",
    }
client.TransferOrders.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Update(TransferOrderID, request) -> *square.UpdateTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing transfer order. This endpoint supports sparse updates,
allowing you to modify specific fields without affecting others.

Creates a [transfer_order.updated](webhook:transfer_order.updated) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateTransferOrderRequest{
        TransferOrderID: "transfer_order_id",
        IdempotencyKey: "f47ac10b-58cc-4372-a567-0e02b2c3d479",
        TransferOrder: &square.UpdateTransferOrderData{
            SourceLocationID: square.String(
                "EXAMPLE_SOURCE_LOCATION_ID_789",
            ),
            DestinationLocationID: square.String(
                "EXAMPLE_DEST_LOCATION_ID_101",
            ),
            ExpectedAt: square.String(
                "2025-11-10T08:00:00Z",
            ),
            Notes: square.String(
                "Updated: Priority transfer due to low stock at destination",
            ),
            TrackingNumber: square.String(
                "TRACK987654321",
            ),
            LineItems: []*square.UpdateTransferOrderLineData{
                &square.UpdateTransferOrderLineData{
                    UID: square.String(
                        "1",
                    ),
                    QuantityOrdered: square.String(
                        "7",
                    ),
                },
                &square.UpdateTransferOrderLineData{
                    ItemVariationID: square.String(
                        "EXAMPLE_NEW_ITEM_VARIATION_ID_003",
                    ),
                    QuantityOrdered: square.String(
                        "2",
                    ),
                },
                &square.UpdateTransferOrderLineData{
                    UID: square.String(
                        "2",
                    ),
                    Remove: square.Bool(
                        true,
                    ),
                },
            },
        },
        Version: square.Int64(
            1753109537351,
        ),
    }
client.TransferOrders.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to update
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique string that identifies this UpdateTransferOrder request. Keys must contain only alphanumeric characters, dashes and underscores
    
</dd>
</dl>

<dl>
<dd>

**transferOrder:** `*square.UpdateTransferOrderData` — The transfer order updates to apply
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` — Version for optimistic concurrency
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Delete(TransferOrderID) -> *square.DeleteTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a transfer order in [DRAFT](entity:TransferOrderStatus) status.
Only draft orders can be deleted. Once an order is started via 
[StartTransferOrder](api-endpoint:TransferOrders-StartTransferOrder), it can no longer be deleted.

Creates a [transfer_order.deleted](webhook:transfer_order.deleted) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.DeleteTransferOrdersRequest{
        TransferOrderID: "transfer_order_id",
        Version: square.Int64(
            1000000,
        ),
    }
client.TransferOrders.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to delete
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` — Version for optimistic concurrency
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Cancel(TransferOrderID, request) -> *square.CancelTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a transfer order in [STARTED](entity:TransferOrderStatus) or 
[PARTIALLY_RECEIVED](entity:TransferOrderStatus) status. Any unreceived quantities will no
longer be receivable and will be immediately returned to the source [Location](entity:Location)'s inventory.

Common reasons for cancellation:
- Items no longer needed at destination
- Source location needs the inventory
- Order created in error

Creates a [transfer_order.updated](webhook:transfer_order.updated) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CancelTransferOrderRequest{
        TransferOrderID: "transfer_order_id",
        IdempotencyKey: "65cc0586-3e82-4d08-b524-3885cffd52",
        Version: square.Int64(
            1753117449752,
        ),
    }
client.TransferOrders.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to cancel. Must be in STARTED or PARTIALLY_RECEIVED status.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this UpdateTransferOrder request. Keys can be
any valid string but must be unique for every UpdateTransferOrder request.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` — Version for optimistic concurrency
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Receive(TransferOrderID, request) -> *square.ReceiveTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Records receipt of [CatalogItemVariation](entity:CatalogItemVariation)s for a transfer order.
This endpoint supports partial receiving - you can receive items in multiple batches.

For each line item, you can specify:
- Quantity received in good condition (added to destination inventory with [InventoryState](entity:InventoryState) of IN_STOCK)
- Quantity damaged during transit/handling (added to destination inventory with [InventoryState](entity:InventoryState) of WASTE)
- Quantity canceled (returned to source location's inventory)

The order must be in [STARTED](entity:TransferOrderStatus) or [PARTIALLY_RECEIVED](entity:TransferOrderStatus) status.
Received quantities are added to the destination [Location](entity:Location)'s inventory according to their condition.
Canceled quantities are immediately returned to the source [Location](entity:Location)'s inventory.

When all items are either received, damaged, or canceled, the order moves to
[COMPLETED](entity:TransferOrderStatus) status.

Creates a [transfer_order.updated](webhook:transfer_order.updated) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.ReceiveTransferOrderRequest{
        TransferOrderID: "transfer_order_id",
        IdempotencyKey: "EXAMPLE_IDEMPOTENCY_KEY_101",
        Receipt: &square.TransferOrderGoodsReceipt{
            LineItems: []*square.TransferOrderGoodsReceiptLineItem{
                &square.TransferOrderGoodsReceiptLineItem{
                    TransferOrderLineUID: "transfer_order_line_uid",
                    QuantityReceived: square.String(
                        "3",
                    ),
                    QuantityDamaged: square.String(
                        "1",
                    ),
                    QuantityCanceled: square.String(
                        "1",
                    ),
                },
                &square.TransferOrderGoodsReceiptLineItem{
                    TransferOrderLineUID: "transfer_order_line_uid",
                    QuantityReceived: square.String(
                        "2",
                    ),
                    QuantityCanceled: square.String(
                        "1",
                    ),
                },
            },
        },
        Version: square.Int64(
            1753118664873,
        ),
    }
client.TransferOrders.Receive(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to receive items for
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key to make this request idempotent
    
</dd>
</dl>

<dl>
<dd>

**receipt:** `*square.TransferOrderGoodsReceipt` — The receipt details
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` — Version for optimistic concurrency
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransferOrders.Start(TransferOrderID, request) -> *square.StartTransferOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes a [DRAFT](entity:TransferOrderStatus) transfer order to [STARTED](entity:TransferOrderStatus) status.
This decrements inventory at the source [Location](entity:Location) and marks it as in-transit.

The order must be in [DRAFT](entity:TransferOrderStatus) status and have all required fields populated.
Once started, the order can no longer be deleted, but it can be canceled via 
[CancelTransferOrder](api-endpoint:TransferOrders-CancelTransferOrder).

Creates a [transfer_order.updated](webhook:transfer_order.updated) webhook event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.StartTransferOrderRequest{
        TransferOrderID: "transfer_order_id",
        IdempotencyKey: "EXAMPLE_IDEMPOTENCY_KEY_789",
        Version: square.Int64(
            1753109537351,
        ),
    }
client.TransferOrders.Start(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transferOrderID:** `string` — The ID of the transfer order to start. Must be in DRAFT status.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this UpdateTransferOrder request. Keys can be
any valid string but must be unique for every UpdateTransferOrder request.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int64` — Version for optimistic concurrency
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Vendors
<details><summary><code>client.Vendors.BatchCreate(request) -> *square.BatchCreateVendorsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates one or more [Vendor](entity:Vendor) objects to represent suppliers to a seller.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchCreateVendorsRequest{
        Vendors: map[string]*square.Vendor{
            "8fc6a5b0-9fe8-4b46-b46b-2ef95793abbe": &square.Vendor{
                Name: square.String(
                    "Joe's Fresh Seafood",
                ),
                Address: &square.Address{
                    AddressLine1: square.String(
                        "505 Electric Ave",
                    ),
                    AddressLine2: square.String(
                        "Suite 600",
                    ),
                    Locality: square.String(
                        "New York",
                    ),
                    AdministrativeDistrictLevel1: square.String(
                        "NY",
                    ),
                    PostalCode: square.String(
                        "10003",
                    ),
                    Country: square.CountryUs.Ptr(),
                },
                Contacts: []*square.VendorContact{
                    &square.VendorContact{
                        Name: square.String(
                            "Joe Burrow",
                        ),
                        EmailAddress: square.String(
                            "joe@joesfreshseafood.com",
                        ),
                        PhoneNumber: square.String(
                            "1-212-555-4250",
                        ),
                        Ordinal: 1,
                    },
                },
                AccountNumber: square.String(
                    "4025391",
                ),
                Note: square.String(
                    "a vendor",
                ),
            },
        },
    }
client.Vendors.BatchCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vendors:** `map[string]*square.Vendor` — Specifies a set of new [Vendor](entity:Vendor) objects as represented by a collection of idempotency-key/`Vendor`-object pairs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.BatchGet(request) -> *square.BatchGetVendorsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one or more vendors of specified [Vendor](entity:Vendor) IDs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchGetVendorsRequest{
        VendorIDs: []string{
            "INV_V_JDKYHBWT1D4F8MFH63DBMEN8Y4",
        },
    }
client.Vendors.BatchGet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vendorIDs:** `[]string` — IDs of the [Vendor](entity:Vendor) objects to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.BatchUpdate(request) -> *square.BatchUpdateVendorsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates one or more of existing [Vendor](entity:Vendor) objects as suppliers to a seller.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.BatchUpdateVendorsRequest{
        Vendors: map[string]*square.UpdateVendorRequest{
            "FMCYHBWT1TPL8MFH52PBMEN92A": &square.UpdateVendorRequest{
                Vendor: &square.Vendor{},
            },
            "INV_V_JDKYHBWT1D4F8MFH63DBMEN8Y4": &square.UpdateVendorRequest{
                Vendor: &square.Vendor{},
            },
        },
    }
client.Vendors.BatchUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vendors:** `map[string]*square.UpdateVendorRequest` 

A set of [UpdateVendorRequest](entity:UpdateVendorRequest) objects encapsulating to-be-updated [Vendor](entity:Vendor)
objects. The set is represented by  a collection of `Vendor`-ID/`UpdateVendorRequest`-object pairs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.Create(request) -> *square.CreateVendorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a single [Vendor](entity:Vendor) object to represent a supplier to a seller.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.CreateVendorRequest{
        IdempotencyKey: "8fc6a5b0-9fe8-4b46-b46b-2ef95793abbe",
        Vendor: &square.Vendor{
            Name: square.String(
                "Joe's Fresh Seafood",
            ),
            Address: &square.Address{
                AddressLine1: square.String(
                    "505 Electric Ave",
                ),
                AddressLine2: square.String(
                    "Suite 600",
                ),
                Locality: square.String(
                    "New York",
                ),
                AdministrativeDistrictLevel1: square.String(
                    "NY",
                ),
                PostalCode: square.String(
                    "10003",
                ),
                Country: square.CountryUs.Ptr(),
            },
            Contacts: []*square.VendorContact{
                &square.VendorContact{
                    Name: square.String(
                        "Joe Burrow",
                    ),
                    EmailAddress: square.String(
                        "joe@joesfreshseafood.com",
                    ),
                    PhoneNumber: square.String(
                        "1-212-555-4250",
                    ),
                    Ordinal: 1,
                },
            },
            AccountNumber: square.String(
                "4025391",
            ),
            Note: square.String(
                "a vendor",
            ),
        },
    }
client.Vendors.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A client-supplied, universally unique identifier (UUID) to make this [CreateVendor](api-endpoint:Vendors-CreateVendor) call idempotent.

See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) in the
[API Development 101](https://developer.squareup.com/docs/buildbasics) section for more
information.
    
</dd>
</dl>

<dl>
<dd>

**vendor_:** `*square.Vendor` — The requested [Vendor](entity:Vendor) to be created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.Search(request) -> *square.SearchVendorsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for vendors using a filter against supported [Vendor](entity:Vendor) properties and a supported sorter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.SearchVendorsRequest{}
client.Vendors.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**filter:** `*square.SearchVendorsRequestFilter` — Specifies a filter used to search for vendors.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*square.SearchVendorsRequestSort` — Specifies a sorter used to sort the returned vendors.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for the original query.

See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.Get(VendorID) -> *square.GetVendorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the vendor of a specified [Vendor](entity:Vendor) ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.GetVendorsRequest{
        VendorID: "vendor_id",
    }
client.Vendors.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vendorID:** `string` — ID of the [Vendor](entity:Vendor) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vendors.Update(VendorID, request) -> *square.UpdateVendorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing [Vendor](entity:Vendor) object as a supplier to a seller.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &square.UpdateVendorsRequest{
        VendorID: "vendor_id",
        Body: &square.UpdateVendorRequest{
            IdempotencyKey: square.String(
                "8fc6a5b0-9fe8-4b46-b46b-2ef95793abbe",
            ),
            Vendor: &square.Vendor{
                ID: square.String(
                    "INV_V_JDKYHBWT1D4F8MFH63DBMEN8Y4",
                ),
                Name: square.String(
                    "Jack's Chicken Shack",
                ),
                Version: square.Int(
                    1,
                ),
                Status: square.VendorStatusActive.Ptr(),
            },
        },
    }
client.Vendors.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**vendorID:** `string` — ID of the [Vendor](entity:Vendor) to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*square.UpdateVendorRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bookings CustomAttributeDefinitions
<details><summary><code>client.Bookings.CustomAttributeDefinitions.List() -> *square.ListBookingCustomAttributeDefinitionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all bookings custom attribute definitions.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.ListCustomAttributeDefinitionsRequest{
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Bookings.CustomAttributeDefinitions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributeDefinitions.Create(request) -> *square.CreateBookingCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a bookings custom attribute definition.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.CreateBookingCustomAttributeDefinitionRequest{
        CustomAttributeDefinition: &square.CustomAttributeDefinition{},
    }
client.Bookings.CustomAttributeDefinitions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition to create, with the following fields:

- `key`

- `name`. If provided, `name` must be unique (case-sensitive) across all visible booking-related custom attribute
definitions for the seller.

- `description`

- `visibility`. Note that all custom attributes are visible in exported booking data, including those set to
`VISIBILITY_HIDDEN`.

- `schema`. With the exception of the `Selection` data type, the `schema` is specified as a
simple URL to the JSON schema definition hosted on the Square CDN. For more information, see
[Specifying the schema](https://developer.squareup.com/docs/booking-custom-attributes-api/custom-attribute-definitions#specify-schema).
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributeDefinitions.Get(Key) -> *square.RetrieveBookingCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a bookings custom attribute definition.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.GetCustomAttributeDefinitionsRequest{
        Key: "key",
        Version: square.Int(
            1,
        ),
    }
client.Bookings.CustomAttributeDefinitions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute definition to retrieve. If the requesting application
is not the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute definition, which is used for strongly consistent
reads to guarantee that you receive the most up-to-date data. When included in the request,
Square returns the specified version or a higher version if one exists. If the specified version
is higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributeDefinitions.Update(Key, request) -> *square.UpdateBookingCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a bookings custom attribute definition.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.UpdateBookingCustomAttributeDefinitionRequest{
        Key: "key",
        CustomAttributeDefinition: &square.CustomAttributeDefinition{},
    }
client.Bookings.CustomAttributeDefinitions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to update.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition that contains the fields to update. Only the following fields can be updated:
- `name`
- `description`
- `visibility`
- `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
selections are supported.

For more information, see
[Updatable definition fields](https://developer.squareup.com/docs/booking-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control, include the optional `version` field and specify the current version of the custom attribute definition.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributeDefinitions.Delete(Key) -> *square.DeleteBookingCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a bookings custom attribute definition.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.DeleteCustomAttributeDefinitionsRequest{
        Key: "key",
    }
client.Bookings.CustomAttributeDefinitions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bookings CustomAttributes
<details><summary><code>client.Bookings.CustomAttributes.BatchDelete(request) -> *square.BulkDeleteBookingCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Bulk deletes bookings custom attributes.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.BulkDeleteBookingCustomAttributesRequest{
        Values: map[string]*square.BookingCustomAttributeDeleteRequest{
            "key": &square.BookingCustomAttributeDeleteRequest{
                BookingID: "booking_id",
                Key: "key",
            },
        },
    }
client.Bookings.CustomAttributes.BatchDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BookingCustomAttributeDeleteRequest` 

A map containing 1 to 25 individual Delete requests. For each request, provide an
arbitrary ID that is unique for this `BulkDeleteBookingCustomAttributes` request and the
information needed to delete a custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributes.BatchUpsert(request) -> *square.BulkUpsertBookingCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Bulk upserts bookings custom attributes.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.BulkUpsertBookingCustomAttributesRequest{
        Values: map[string]*square.BookingCustomAttributeUpsertRequest{
            "key": &square.BookingCustomAttributeUpsertRequest{
                BookingID: "booking_id",
                CustomAttribute: &square.CustomAttribute{},
            },
        },
    }
client.Bookings.CustomAttributes.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BookingCustomAttributeUpsertRequest` 

A map containing 1 to 25 individual upsert requests. For each request, provide an
arbitrary ID that is unique for this `BulkUpsertBookingCustomAttributes` request and the
information needed to create or update a custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributes.List(BookingID) -> *square.ListBookingCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists a booking's custom attributes.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.ListCustomAttributesRequest{
        BookingID: "booking_id",
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        WithDefinitions: square.Bool(
            true,
        ),
    }
client.Bookings.CustomAttributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the target [booking](entity:Booking).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request. For more
information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**withDefinitions:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
custom attribute. Set this parameter to `true` to get the name and description of each custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributes.Get(BookingID, Key) -> *square.RetrieveBookingCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a bookings custom attribute.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.GetCustomAttributesRequest{
        BookingID: "booking_id",
        Key: "key",
        WithDefinition: square.Bool(
            true,
        ),
        Version: square.Int(
            1,
        ),
    }
client.Bookings.CustomAttributes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the target [booking](entity:Booking).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to retrieve. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**withDefinition:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of
the custom attribute. Set this parameter to `true` to get the name and description of the custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute, which is used for strongly consistent reads to
guarantee that you receive the most up-to-date data. When included in the request, Square
returns the specified version or a higher version if one exists. If the specified version is
higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributes.Upsert(BookingID, Key, request) -> *square.UpsertBookingCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upserts a bookings custom attribute.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.UpsertBookingCustomAttributeRequest{
        BookingID: "booking_id",
        Key: "key",
        CustomAttribute: &square.CustomAttribute{},
    }
client.Bookings.CustomAttributes.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the target [booking](entity:Booking).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to create or update. This key must match the `key` of a
custom attribute definition in the Square seller account. If the requesting application is not
the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**customAttribute:** `*square.CustomAttribute` 

The custom attribute to create or update, with the following fields:

- `value`. This value must conform to the `schema` specified by the definition.
For more information, see [Value data types](https://developer.squareup.com/docs/booking-custom-attributes-api/custom-attributes#value-data-types).

- `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control for an update operation, include this optional field and specify the current version
of the custom attribute.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.CustomAttributes.Delete(BookingID, Key) -> *square.DeleteBookingCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a bookings custom attribute.

To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope.
To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.

For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus*
or *Appointments Premium*.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.DeleteCustomAttributesRequest{
        BookingID: "booking_id",
        Key: "key",
    }
client.Bookings.CustomAttributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookingID:** `string` — The ID of the target [booking](entity:Booking).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to delete. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bookings LocationProfiles
<details><summary><code>client.Bookings.LocationProfiles.List() -> *square.ListLocationBookingProfilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists location booking profiles of a seller.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.ListLocationProfilesRequest{
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Bookings.LocationProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to return in a paged response.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bookings TeamMemberProfiles
<details><summary><code>client.Bookings.TeamMemberProfiles.List() -> *square.ListTeamMemberBookingProfilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists booking profiles for team members.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.ListTeamMemberProfilesRequest{
        BookableOnly: square.Bool(
            true,
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        LocationID: square.String(
            "location_id",
        ),
    }
client.Bookings.TeamMemberProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bookableOnly:** `*bool` — Indicates whether to include only bookable team members in the returned result (`true`) or not (`false`).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to return in a paged response.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` — Indicates whether to include only team members enabled at the given location in the returned result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bookings.TeamMemberProfiles.Get(TeamMemberID) -> *square.GetTeamMemberBookingProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a team member's booking profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bookings.GetTeamMemberProfilesRequest{
        TeamMemberID: "team_member_id",
    }
client.Bookings.TeamMemberProfiles.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `string` — The ID of the team member to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CashDrawers Shifts
<details><summary><code>client.CashDrawers.Shifts.List() -> *square.ListCashDrawerShiftsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides the details for all of the cash drawer shifts for a location
in a date range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cashdrawers.ListShiftsRequest{
        LocationID: "location_id",
        SortOrder: square.SortOrderDesc.Ptr(),
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.CashDrawers.Shifts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location to query for a list of cash drawer shifts.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

The order in which cash drawer shifts are listed in the response,
based on their opened_at field. Default value: ASC
    
</dd>
</dl>

<dl>
<dd>

**beginTime:** `*string` — The inclusive start time of the query on opened_at, in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` — The exclusive end date of the query on opened_at, in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

Number of cash drawer shift events in a page of results (200 by
default, 1000 max).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque cursor for fetching the next page of results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CashDrawers.Shifts.Get(ShiftID) -> *square.GetCashDrawerShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides the summary details for a single cash drawer shift. See
[ListCashDrawerShiftEvents](api-endpoint:CashDrawers-ListCashDrawerShiftEvents) for a list of cash drawer shift events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cashdrawers.GetShiftsRequest{
        ShiftID: "shift_id",
        LocationID: "location_id",
    }
client.CashDrawers.Shifts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**shiftID:** `string` — The shift ID.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The ID of the location to retrieve cash drawer shifts from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CashDrawers.Shifts.ListEvents(ShiftID) -> *square.ListCashDrawerShiftEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides a paginated list of events for a single cash drawer shift.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cashdrawers.ListEventsShiftsRequest{
        ShiftID: "shift_id",
        LocationID: "location_id",
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.CashDrawers.Shifts.ListEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**shiftID:** `string` — The shift ID.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The ID of the location to list cash drawer shifts for.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

Number of resources to be returned in a page of results (200 by
default, 1000 max).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque cursor for fetching the next page of results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Catalog Images
<details><summary><code>client.Catalog.Images.Create(request) -> *square.CreateCatalogImageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads an image file to be represented by a [CatalogImage](entity:CatalogImage) object that can be linked to an existing
[CatalogObject](entity:CatalogObject) instance. The resulting `CatalogImage` is unattached to any `CatalogObject` if the `object_id`
is not specified.

This `CreateCatalogImage` endpoint accepts HTTP multipart/form-data requests with a JSON part and an image file part in
JPEG, PJPEG, PNG, or GIF format. The maximum file size is 15MB.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &catalog.CreateImagesRequest{}
client.Catalog.Images.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.Images.Update(ImageID, request) -> *square.UpdateCatalogImageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads a new image file to replace the existing one in the specified [CatalogImage](entity:CatalogImage) object.

This `UpdateCatalogImage` endpoint accepts HTTP multipart/form-data requests with a JSON part and an image file part in
JPEG, PJPEG, PNG, or GIF format. The maximum file size is 15MB.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &catalog.UpdateImagesRequest{
        ImageID: "image_id",
    }
client.Catalog.Images.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**imageID:** `string` — The ID of the `CatalogImage` object to update the encapsulated image file.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Catalog Object
<details><summary><code>client.Catalog.Object.Upsert(request) -> *square.UpsertCatalogObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new or updates the specified [CatalogObject](entity:CatalogObject).

To ensure consistency, only one update request is processed at a time per seller account.
While one (batch or non-batch) update request is being processed, other (batched and non-batched)
update requests are rejected with the `429` error code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &catalog.UpsertCatalogObjectRequest{
        IdempotencyKey: "af3d1afc-7212-4300-b463-0bfc5314a5ae",
        Object: &square.CatalogObject{
            Item: &square.CatalogObjectItem{
                ID: "id",
            },
        },
    }
client.Catalog.Object.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A value you specify that uniquely identifies this
request among all your requests. A common way to create
a valid idempotency key is to use a Universally unique
identifier (UUID).

If you're unsure whether a particular request was successful,
you can reattempt it with the same idempotency key without
worrying about creating duplicate objects.

See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**object:** `*square.CatalogObject` 

A CatalogObject to be created or updated.

- For updates, the object must be active (the `is_deleted` field is not `true`).
- For creates, the object ID must start with `#`. The provided ID is replaced with a server-generated ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.Object.Get(ObjectID) -> *square.GetCatalogObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single [CatalogItem](entity:CatalogItem) as a
[CatalogObject](entity:CatalogObject) based on the provided ID. The returned
object includes all of the relevant [CatalogItem](entity:CatalogItem)
information including: [CatalogItemVariation](entity:CatalogItemVariation)
children, references to its
[CatalogModifierList](entity:CatalogModifierList) objects, and the ids of
any [CatalogTax](entity:CatalogTax) objects that apply to it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &catalog.GetObjectRequest{
        ObjectID: "object_id",
        IncludeRelatedObjects: square.Bool(
            true,
        ),
        CatalogVersion: square.Int64(
            1000000,
        ),
        IncludeCategoryPathToRoot: square.Bool(
            true,
        ),
    }
client.Catalog.Object.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**objectID:** `string` — The object ID of any type of catalog objects to be retrieved.
    
</dd>
</dl>

<dl>
<dd>

**includeRelatedObjects:** `*bool` 

If `true`, the response will include additional objects that are related to the
requested objects. Related objects are defined as any objects referenced by ID by the results in the `objects` field
of the response. These objects are put in the `related_objects` field. Setting this to `true` is
helpful when the objects are needed for immediate display to a user.
This process only goes one level deep. Objects referenced by the related objects will not be included. For example,

if the `objects` field of the response contains a CatalogItem, its associated
CatalogCategory objects, CatalogTax objects, CatalogImage objects and
CatalogModifierLists will be returned in the `related_objects` field of the
response. If the `objects` field of the response contains a CatalogItemVariation,
its parent CatalogItem will be returned in the `related_objects` field of
the response.

Default value: `false`
    
</dd>
</dl>

<dl>
<dd>

**catalogVersion:** `*int64` 

Requests objects as of a specific version of the catalog. This allows you to retrieve historical
versions of objects. The value to retrieve a specific version of an object can be found
in the version field of [CatalogObject](entity:CatalogObject)s. If not included, results will
be from the current version of the catalog.
    
</dd>
</dl>

<dl>
<dd>

**includeCategoryPathToRoot:** `*bool` 

Specifies whether or not to include the `path_to_root` list for each returned category instance. The `path_to_root` list consists
of `CategoryPathToRootNode` objects and specifies the path that starts with the immediate parent category of the returned category
and ends with its root category. If the returned category is a top-level category, the `path_to_root` list is empty and is not returned
in the response payload.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.Object.Delete(ObjectID) -> *square.DeleteCatalogObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a single [CatalogObject](entity:CatalogObject) based on the
provided ID and returns the set of successfully deleted IDs in the response.
Deletion is a cascading event such that all children of the targeted object
are also deleted. For example, deleting a [CatalogItem](entity:CatalogItem)
will also delete all of its
[CatalogItemVariation](entity:CatalogItemVariation) children.

To ensure consistency, only one delete request is processed at a time per seller account.
While one (batch or non-batch) delete request is being processed, other (batched and non-batched)
delete requests are rejected with the `429` error code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &catalog.DeleteObjectRequest{
        ObjectID: "object_id",
    }
client.Catalog.Object.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**objectID:** `string` 

The ID of the catalog object to be deleted. When an object is deleted, other
objects in the graph that depend on that object will be deleted as well (for example, deleting a
catalog item will delete its catalog item variations).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Checkout PaymentLinks
<details><summary><code>client.Checkout.PaymentLinks.List() -> *square.ListPaymentLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all payment links.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &checkout.ListPaymentLinksRequest{
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Checkout.PaymentLinks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
If a cursor is not provided, the endpoint returns the first page of the results.
For more  information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

A limit on the number of results to return per page. The limit is advisory and
the implementation might return more or less results. If the supplied limit is negative, zero, or
greater than the maximum limit of 1000, it is ignored.

Default value: `100`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PaymentLinks.Create(request) -> *square.CreatePaymentLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a Square-hosted checkout page. Applications can share the resulting payment link with their buyer to pay for goods and services.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &checkout.CreatePaymentLinkRequest{
        IdempotencyKey: square.String(
            "cd9e25dc-d9f2-4430-aedb-61605070e95f",
        ),
        QuickPay: &square.QuickPay{
            Name: "Auto Detailing",
            PriceMoney: &square.Money{
                Amount: square.Int64(
                    10000,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
            LocationID: "A9Y43N9ABXZBP",
        },
    }
client.Checkout.PaymentLinks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique string that identifies this `CreatePaymentLinkRequest` request.
If you do not provide a unique string (or provide an empty string as the value),
the endpoint treats each request as independent.

For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 

A description of the payment link. You provide this optional description that is useful in your
application context. It is not used anywhere.
    
</dd>
</dl>

<dl>
<dd>

**quickPay:** `*square.QuickPay` 

Describes an ad hoc item and price for which to generate a quick pay checkout link.
For more information,
see [Quick Pay Checkout](https://developer.squareup.com/docs/checkout-api/quick-pay-checkout).
    
</dd>
</dl>

<dl>
<dd>

**order:** `*square.Order` 

Describes the `Order` for which to create a checkout link.
For more information,
see [Square Order Checkout](https://developer.squareup.com/docs/checkout-api/square-order-checkout).
    
</dd>
</dl>

<dl>
<dd>

**checkoutOptions:** `*square.CheckoutOptions` 

Describes optional fields to add to the resulting checkout page.
For more information,
see [Optional Checkout Configurations](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations).
    
</dd>
</dl>

<dl>
<dd>

**prePopulatedData:** `*square.PrePopulatedData` 

Describes fields to prepopulate in the resulting checkout page.
For more information, see [Prepopulate the shipping address](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations#prepopulate-the-shipping-address).
    
</dd>
</dl>

<dl>
<dd>

**paymentNote:** `*string` — A note for the payment. After processing the payment, Square adds this note to the resulting `Payment`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PaymentLinks.Get(ID) -> *square.GetPaymentLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a payment link.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &checkout.GetPaymentLinksRequest{
        ID: "id",
    }
client.Checkout.PaymentLinks.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of link to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PaymentLinks.Update(ID, request) -> *square.UpdatePaymentLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a payment link. You can update the `payment_link` fields such as
`description`, `checkout_options`, and  `pre_populated_data`.
You cannot update other fields such as the `order_id`, `version`, `URL`, or `timestamp` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &checkout.UpdatePaymentLinkRequest{
        ID: "id",
        PaymentLink: &square.PaymentLink{
            Version: 1,
            CheckoutOptions: &square.CheckoutOptions{
                AskForShippingAddress: square.Bool(
                    true,
                ),
            },
        },
    }
client.Checkout.PaymentLinks.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the payment link to update.
    
</dd>
</dl>

<dl>
<dd>

**paymentLink:** `*square.PaymentLink` 

The `payment_link` object describing the updates to apply.
For more information, see [Update a payment link](https://developer.squareup.com/docs/checkout-api/manage-checkout#update-a-payment-link).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PaymentLinks.Delete(ID) -> *square.DeletePaymentLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a payment link.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &checkout.DeletePaymentLinksRequest{
        ID: "id",
    }
client.Checkout.PaymentLinks.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the payment link to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers CustomAttributeDefinitions
<details><summary><code>client.Customers.CustomAttributeDefinitions.List() -> *square.ListCustomerCustomAttributeDefinitionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the customer-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.

When all response pages are retrieved, the results include all custom attribute definitions
that are visible to the requesting application, including those that are created by other
applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that
seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.ListCustomAttributeDefinitionsRequest{
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Customers.CustomAttributeDefinitions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributeDefinitions.Create(request) -> *square.CreateCustomerCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
Use this endpoint to define a custom attribute that can be associated with customer profiles.

A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties
for a custom attribute. After the definition is created, you can call
[UpsertCustomerCustomAttribute](api-endpoint:CustomerCustomAttributes-UpsertCustomerCustomAttribute) or
[BulkUpsertCustomerCustomAttributes](api-endpoint:CustomerCustomAttributes-BulkUpsertCustomerCustomAttributes)
to set the custom attribute for customer profiles in the seller's Customer Directory.

Sellers can view all custom attributes in exported customer data, including those set to
`VISIBILITY_HIDDEN`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.CreateCustomerCustomAttributeDefinitionRequest{
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Key: square.String(
                "favoritemovie",
            ),
            Schema: map[string]any{
                "\$ref": "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.String",
            },
            Name: square.String(
                "Favorite Movie",
            ),
            Description: square.String(
                "The favorite movie of the customer.",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityHidden.Ptr(),
        },
    }
client.Customers.CustomAttributeDefinitions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition to create. Note the following:
- With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
definition hosted on the Square CDN. For more information, including supported values and constraints, see
[Specifying the schema](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#specify-schema).
- If provided, `name` must be unique (case-sensitive) across all visible customer-related custom attribute definitions for the seller.
- All custom attributes are visible in exported customer data, including those set to `VISIBILITY_HIDDEN`.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributeDefinitions.Get(Key) -> *square.GetCustomerCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.

To retrieve a custom attribute definition created by another application, the `visibility`
setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.GetCustomAttributeDefinitionsRequest{
        Key: "key",
        Version: square.Int(
            1,
        ),
    }
client.Customers.CustomAttributeDefinitions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute definition to retrieve. If the requesting application
is not the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute definition, which is used for strongly consistent
reads to guarantee that you receive the most up-to-date data. When included in the request,
Square returns the specified version or a higher version if one exists. If the specified version
is higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributeDefinitions.Update(Key, request) -> *square.UpdateCustomerCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.

Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the
`schema` for a `Selection` data type.

Only the definition owner can update a custom attribute definition. Note that sellers can view
all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.UpdateCustomerCustomAttributeDefinitionRequest{
        Key: "key",
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Description: square.String(
                "Update the description as desired.",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadOnly.Ptr(),
        },
    }
client.Customers.CustomAttributeDefinitions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to update.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition that contains the fields to update. This endpoint
supports sparse updates, so only new or changed fields need to be included in the request.
Only the following fields can be updated:

- `name`
- `description`
- `visibility`
- `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
selections are supported.

For more information, see
[Updatable definition fields](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) 
control, include the optional `version` field and specify the current version of the custom attribute definition.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributeDefinitions.Delete(Key) -> *square.DeleteCustomerCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.

Deleting a custom attribute definition also deletes the corresponding custom attribute from
all customer profiles in the seller's Customer Directory.

Only the definition owner can delete a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.DeleteCustomAttributeDefinitionsRequest{
        Key: "key",
    }
client.Customers.CustomAttributeDefinitions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributeDefinitions.BatchUpsert(request) -> *square.BatchUpsertCustomerCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates [custom attributes](entity:CustomAttribute) for customer profiles as a bulk operation.

Use this endpoint to set the value of one or more custom attributes for one or more customer profiles.
A custom attribute is based on a custom attribute definition in a Square seller account, which is
created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.

This `BulkUpsertCustomerCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
requests and returns a map of individual upsert responses. Each upsert request has a unique ID
and provides a customer ID and custom attribute. Each upsert response is returned with the ID
of the corresponding request.

To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.BatchUpsertCustomerCustomAttributesRequest{
        Values: map[string]*square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
            "id1": &square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
                CustomerID: "N3NCVYY3WS27HF0HKANA3R9FP8",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "favoritemovie",
                    ),
                    Value: "Dune",
                },
            },
            "id2": &square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
                CustomerID: "SY8EMWRNDN3TQDP2H4KS1QWMMM",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "ownsmovie",
                    ),
                    Value: false,
                },
            },
            "id3": &square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
                CustomerID: "SY8EMWRNDN3TQDP2H4KS1QWMMM",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "favoritemovie",
                    ),
                    Value: "Star Wars",
                },
            },
            "id4": &square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
                CustomerID: "N3NCVYY3WS27HF0HKANA3R9FP8",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "square:a0f1505a-2aa1-490d-91a8-8d31ff181808",
                    ),
                    Value: "10.5",
                },
            },
            "id5": &square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest{
                CustomerID: "70548QG1HN43B05G0KCZ4MMC1G",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "sq0ids-0evKIskIGaY45fCyNL66aw:backupemail",
                    ),
                    Value: "fake-email@squareup.com",
                },
            },
        },
    }
client.Customers.CustomAttributeDefinitions.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BatchUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest` 

A map containing 1 to 25 individual upsert requests. For each request, provide an
arbitrary ID that is unique for this `BulkUpsertCustomerCustomAttributes` request and the
information needed to create or update a custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers Groups
<details><summary><code>client.Customers.Groups.List() -> *square.ListCustomerGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of customer groups of a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.ListGroupsRequest{
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Customers.Groups.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results.
If the limit is less than 1 or greater than 50, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 50.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Create(request) -> *square.CreateCustomerGroupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new customer group for a business.

The request must include the `name` value of the group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.CreateCustomerGroupRequest{
        Group: &square.CustomerGroup{
            Name: "Loyal Customers",
        },
    }
client.Customers.Groups.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — The idempotency key for the request. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**group:** `*square.CustomerGroup` — The customer group to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Get(GroupID) -> *square.GetCustomerGroupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specific customer group as identified by the `group_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.GetGroupsRequest{
        GroupID: "group_id",
    }
client.Customers.Groups.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupID:** `string` — The ID of the customer group to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Update(GroupID, request) -> *square.UpdateCustomerGroupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a customer group as identified by the `group_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.UpdateCustomerGroupRequest{
        GroupID: "group_id",
        Group: &square.CustomerGroup{
            Name: "Loyal Customers",
        },
    }
client.Customers.Groups.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupID:** `string` — The ID of the customer group to update.
    
</dd>
</dl>

<dl>
<dd>

**group:** `*square.CustomerGroup` — The `CustomerGroup` object including all the updates you want to make.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Delete(GroupID) -> *square.DeleteCustomerGroupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a customer group as identified by the `group_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.DeleteGroupsRequest{
        GroupID: "group_id",
    }
client.Customers.Groups.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupID:** `string` — The ID of the customer group to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Add(CustomerID, GroupID) -> *square.AddGroupToCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a group membership to a customer.

The customer is identified by the `customer_id` value
and the customer group is identified by the `group_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.AddGroupsRequest{
        CustomerID: "customer_id",
        GroupID: "group_id",
    }
client.Customers.Groups.Add(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to add to a group.
    
</dd>
</dl>

<dl>
<dd>

**groupID:** `string` — The ID of the customer group to add the customer to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Groups.Remove(CustomerID, GroupID) -> *square.RemoveGroupFromCustomerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a group membership from a customer.

The customer is identified by the `customer_id` value
and the customer group is identified by the `group_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.RemoveGroupsRequest{
        CustomerID: "customer_id",
        GroupID: "group_id",
    }
client.Customers.Groups.Remove(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer to remove from the group.
    
</dd>
</dl>

<dl>
<dd>

**groupID:** `string` — The ID of the customer group to remove the customer from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers Segments
<details><summary><code>client.Customers.Segments.List() -> *square.ListCustomerSegmentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of customer segments of a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.ListSegmentsRequest{
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Customers.Segments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by previous calls to `ListCustomerSegments`.
This cursor is used to retrieve the next set of query results.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results.
If the specified limit is less than 1 or greater than 50, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 50.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Segments.Get(SegmentID) -> *square.GetCustomerSegmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a specific customer segment as identified by the `segment_id` value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.GetSegmentsRequest{
        SegmentID: "segment_id",
    }
client.Customers.Segments.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**segmentID:** `string` — The Square-issued ID of the customer segment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers Cards
<details><summary><code>client.Customers.Cards.Create(CustomerID, request) -> *square.CreateCustomerCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a card on file to an existing customer.

As with charges, calls to `CreateCustomerCard` are idempotent. Multiple
calls with the same card nonce return the same card record that was created
with the provided nonce during the _first_ call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.CreateCustomerCardRequest{
        CustomerID: "customer_id",
        CardNonce: "YOUR_CARD_NONCE",
        BillingAddress: &square.Address{
            AddressLine1: square.String(
                "500 Electric Ave",
            ),
            AddressLine2: square.String(
                "Suite 600",
            ),
            Locality: square.String(
                "New York",
            ),
            AdministrativeDistrictLevel1: square.String(
                "NY",
            ),
            PostalCode: square.String(
                "10003",
            ),
            Country: square.CountryUs.Ptr(),
        },
        CardholderName: square.String(
            "Amelia Earhart",
        ),
    }
client.Customers.Cards.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The Square ID of the customer profile the card is linked to.
    
</dd>
</dl>

<dl>
<dd>

**cardNonce:** `string` 

A card nonce representing the credit card to link to the customer.

Card nonces are generated by the Square payment form when customers enter
their card information. For more information, see
[Walkthrough: Integrate Square Payments in a Website](https://developer.squareup.com/docs/web-payments/take-card-payment).

__NOTE:__ Card nonces generated by digital wallets (such as Apple Pay)
cannot be used to create a customer card.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*square.Address` 

Address information for the card on file.

__NOTE:__ If a billing address is provided in the request, the
`CreateCustomerCardRequest.billing_address.postal_code` must match
the postal code encoded in the card nonce.
    
</dd>
</dl>

<dl>
<dd>

**cardholderName:** `*string` — The full name printed on the credit card.
    
</dd>
</dl>

<dl>
<dd>

**verificationToken:** `*string` 

An identifying token generated by [Payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer).
Verification tokens encapsulate customer device information and 3-D Secure
challenge results to indicate that Square has verified the buyer identity.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.Cards.Delete(CustomerID, CardID) -> *square.DeleteCustomerCardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a card on file from a customer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.DeleteCardsRequest{
        CustomerID: "customer_id",
        CardID: "card_id",
    }
client.Customers.Cards.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the customer that the card on file belongs to.
    
</dd>
</dl>

<dl>
<dd>

**cardID:** `string` — The ID of the card on file to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customers CustomAttributes
<details><summary><code>client.Customers.CustomAttributes.List(CustomerID) -> *square.ListCustomerCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the [custom attributes](entity:CustomAttribute) associated with a customer profile.

You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
in the same call.

When all response pages are retrieved, the results include all custom attributes that are
visible to the requesting application, including those that are owned by other applications
and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.ListCustomAttributesRequest{
        CustomerID: "customer_id",
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        WithDefinitions: square.Bool(
            true,
        ),
    }
client.Customers.CustomAttributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the target [customer profile](entity:Customer).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request. For more
information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**withDefinitions:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
custom attribute. Set this parameter to `true` to get the name and description of each custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributes.Get(CustomerID, Key) -> *square.GetCustomerCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a [custom attribute](entity:CustomAttribute) associated with a customer profile.

You can use the `with_definition` query parameter to also retrieve the custom attribute definition
in the same call.

To retrieve a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.GetCustomAttributesRequest{
        CustomerID: "customer_id",
        Key: "key",
        WithDefinition: square.Bool(
            true,
        ),
        Version: square.Int(
            1,
        ),
    }
client.Customers.CustomAttributes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the target [customer profile](entity:Customer).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to retrieve. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**withDefinition:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of
the custom attribute. Set this parameter to `true` to get the name and description of the custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute, which is used for strongly consistent reads to
guarantee that you receive the most up-to-date data. When included in the request, Square
returns the specified version or a higher version if one exists. If the specified version is
higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributes.Upsert(CustomerID, Key, request) -> *square.UpsertCustomerCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates a [custom attribute](entity:CustomAttribute) for a customer profile.

Use this endpoint to set the value of a custom attribute for a specified customer profile.
A custom attribute is based on a custom attribute definition in a Square seller account, which
is created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.

To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.UpsertCustomerCustomAttributeRequest{
        CustomerID: "customer_id",
        Key: "key",
        CustomAttribute: &square.CustomAttribute{
            Value: "Dune",
        },
    }
client.Customers.CustomAttributes.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the target [customer profile](entity:Customer).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to create or update. This key must match the `key` of a
custom attribute definition in the Square seller account. If the requesting application is not
the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**customAttribute:** `*square.CustomAttribute` 

The custom attribute to create or update, with the following fields:

- `value`. This value must conform to the `schema` specified by the definition. 
For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).

- `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control for an update operation, include this optional field and specify the current version
of the custom attribute.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customers.CustomAttributes.Delete(CustomerID, Key) -> *square.DeleteCustomerCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a [custom attribute](entity:CustomAttribute) associated with a customer profile.

To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &customers.DeleteCustomAttributesRequest{
        CustomerID: "customer_id",
        Key: "key",
    }
client.Customers.CustomAttributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerID:** `string` — The ID of the target [customer profile](entity:Customer).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to delete. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Devices Codes
<details><summary><code>client.Devices.Codes.List() -> *square.ListDeviceCodesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all DeviceCodes associated with the merchant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &devices.ListCodesRequest{
        Cursor: square.String(
            "cursor",
        ),
        LocationID: square.String(
            "location_id",
        ),
        ProductType: &square.ProductType(
            "TERMINAL_API",
        ),
        Status: square.DeviceCodeStatusUnknown.Ptr(),
    }
client.Devices.Codes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for your original query.

See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

If specified, only returns DeviceCodes of the specified location.
Returns DeviceCodes of all locations if empty.
    
</dd>
</dl>

<dl>
<dd>

**productType:** `*square.ProductType` 

If specified, only returns DeviceCodes targeting the specified product type.
Returns DeviceCodes of all product types if empty.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*square.DeviceCodeStatus` 

If specified, returns DeviceCodes with the specified statuses.
Returns DeviceCodes of status `PAIRED` and `UNPAIRED` if empty.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Devices.Codes.Create(request) -> *square.CreateDeviceCodeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a DeviceCode that can be used to login to a Square Terminal device to enter the connected
terminal mode.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &devices.CreateDeviceCodeRequest{
        IdempotencyKey: "01bb00a6-0c86-4770-94ed-f5fca973cd56",
        DeviceCode: &square.DeviceCode{
            Name: square.String(
                "Counter 1",
            ),
            ProductType: square.ProductType(
                "TERMINAL_API",
            ),
            LocationID: square.String(
                "B5E4484SHHNYH",
            ),
        },
    }
client.Devices.Codes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this CreateDeviceCode request. Keys can
be any valid string but must be unique for every CreateDeviceCode request.

See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**deviceCode:** `*square.DeviceCode` — The device code to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Devices.Codes.Get(ID) -> *square.GetDeviceCodeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves DeviceCode with the associated ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &devices.GetCodesRequest{
        ID: "id",
    }
client.Devices.Codes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier for the device code.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Disputes Evidence
<details><summary><code>client.Disputes.Evidence.List(DisputeID) -> *square.ListDisputeEvidenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of evidence associated with a dispute.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &disputes.ListEvidenceRequest{
        DisputeID: "dispute_id",
        Cursor: square.String(
            "cursor",
        ),
    }
client.Disputes.Evidence.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Evidence.Get(DisputeID, EvidenceID) -> *square.GetDisputeEvidenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the metadata for the evidence specified in the request URL path.

You must maintain a copy of any evidence uploaded if you want to reference it later. Evidence cannot be downloaded after you upload it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &disputes.GetEvidenceRequest{
        DisputeID: "dispute_id",
        EvidenceID: "evidence_id",
    }
client.Disputes.Evidence.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute from which you want to retrieve evidence metadata.
    
</dd>
</dl>

<dl>
<dd>

**evidenceID:** `string` — The ID of the evidence to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Evidence.Delete(DisputeID, EvidenceID) -> *square.DeleteDisputeEvidenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes specified evidence from a dispute.
Square does not send the bank any evidence that is removed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &disputes.DeleteEvidenceRequest{
        DisputeID: "dispute_id",
        EvidenceID: "evidence_id",
    }
client.Disputes.Evidence.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeID:** `string` — The ID of the dispute from which you want to remove evidence.
    
</dd>
</dl>

<dl>
<dd>

**evidenceID:** `string` — The ID of the evidence you want to remove.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## GiftCards Activities
<details><summary><code>client.GiftCards.Activities.List() -> *square.ListGiftCardActivitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists gift card activities. By default, you get gift card activities for all
gift cards in the seller's account. You can optionally specify query parameters to
filter the list. For example, you can get a list of gift card activities for a gift card,
for all gift cards in a specific region, or for activities within a time window.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &giftcards.ListActivitiesRequest{
        GiftCardID: square.String(
            "gift_card_id",
        ),
        Type: square.String(
            "type",
        ),
        LocationID: square.String(
            "location_id",
        ),
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        SortOrder: square.String(
            "sort_order",
        ),
    }
client.GiftCards.Activities.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**giftCardID:** `*string` 

If a gift card ID is provided, the endpoint returns activities related 
to the specified gift card. Otherwise, the endpoint returns all gift card activities for 
the seller.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 

If a [type](entity:GiftCardActivityType) is provided, the endpoint returns gift card activities of the specified type. 
Otherwise, the endpoint returns all types of gift card activities.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 

If a location ID is provided, the endpoint returns gift card activities for the specified location. 
Otherwise, the endpoint returns gift card activities for all locations.
    
</dd>
</dl>

<dl>
<dd>

**beginTime:** `*string` 

The timestamp for the beginning of the reporting period, in RFC 3339 format.
This start time is inclusive. The default value is the current time minus one year.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 

The timestamp for the end of the reporting period, in RFC 3339 format.
This end time is inclusive. The default value is the current time.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

If a limit is provided, the endpoint returns the specified number 
of results (or fewer) per page. The maximum value is 100. The default value is 50.
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
If a cursor is not provided, the endpoint returns the first page of the results.
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*string` 

The order in which the endpoint returns the activities, based on `created_at`.
- `ASC` - Oldest to newest.
- `DESC` - Newest to oldest (default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GiftCards.Activities.Create(request) -> *square.CreateGiftCardActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a gift card activity to manage the balance or state of a [gift card](entity:GiftCard).
For example, create an `ACTIVATE` activity to activate a gift card with an initial balance before first use.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &giftcards.CreateGiftCardActivityRequest{
        IdempotencyKey: "U16kfr-kA70er-q4Rsym-7U7NnY",
        GiftCardActivity: &square.GiftCardActivity{
            Type: square.GiftCardActivityTypeActivate,
            LocationID: "81FN9BNFZTKS4",
            GiftCardID: square.String(
                "gftc:6d55a72470d940c6ba09c0ab8ad08d20",
            ),
            ActivateActivityDetails: &square.GiftCardActivityActivate{
                OrderID: square.String(
                    "jJNGHm4gLI6XkFbwtiSLqK72KkAZY",
                ),
                LineItemUID: square.String(
                    "eIWl7X0nMuO9Ewbh0ChIx",
                ),
            },
        },
    }
client.GiftCards.Activities.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — A unique string that identifies the `CreateGiftCardActivity` request.
    
</dd>
</dl>

<dl>
<dd>

**giftCardActivity:** `*square.GiftCardActivity` 

The activity to create for the gift card. This activity must specify `gift_card_id` or `gift_card_gan` for the target
gift card, the `location_id` where the activity occurred, and the activity `type` along with the corresponding activity details.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor BreakTypes
<details><summary><code>client.Labor.BreakTypes.List() -> *square.ListBreakTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `BreakType` instances for a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.ListBreakTypesRequest{
        LocationID: square.String(
            "location_id",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Labor.BreakTypes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `*string` 

Filter the returned `BreakType` results to only those that are associated with the
specified location.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of `BreakType` results to return per page. The number can range between 1
and 200. The default is 200.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A pointer to the next page of `BreakType` results to fetch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.BreakTypes.Create(request) -> *square.CreateBreakTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new `BreakType`.

A `BreakType` is a template for creating `Break` objects.
You must provide the following values in your request to this
endpoint:

- `location_id`
- `break_name`
- `expected_duration`
- `is_paid`

You can only have three `BreakType` instances per location. If you attempt to add a fourth
`BreakType` for a location, an `INVALID_REQUEST_ERROR` "Exceeded limit of 3 breaks per location."
is returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.CreateBreakTypeRequest{
        IdempotencyKey: square.String(
            "PAD3NG5KSN2GL",
        ),
        BreakType: &square.BreakType{
            LocationID: "CGJN03P1D08GF",
            BreakName: "Lunch Break",
            ExpectedDuration: "PT30M",
            IsPaid: true,
        },
    }
client.Labor.BreakTypes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique string value to ensure the idempotency of the operation.
    
</dd>
</dl>

<dl>
<dd>

**breakType:** `*square.BreakType` — The `BreakType` to be created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.BreakTypes.Get(ID) -> *square.GetBreakTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single `BreakType` specified by `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.GetBreakTypesRequest{
        ID: "id",
    }
client.Labor.BreakTypes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `BreakType` being retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.BreakTypes.Update(ID, request) -> *square.UpdateBreakTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing `BreakType`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.UpdateBreakTypeRequest{
        ID: "id",
        BreakType: &square.BreakType{
            LocationID: "26M7H24AZ9N6R",
            BreakName: "Lunch",
            ExpectedDuration: "PT50M",
            IsPaid: true,
            Version: square.Int(
                1,
            ),
        },
    }
client.Labor.BreakTypes.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` —  The UUID for the `BreakType` being updated.
    
</dd>
</dl>

<dl>
<dd>

**breakType:** `*square.BreakType` — The updated `BreakType`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.BreakTypes.Delete(ID) -> *square.DeleteBreakTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an existing `BreakType`.

A `BreakType` can be deleted even if it is referenced from a `Shift`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.DeleteBreakTypesRequest{
        ID: "id",
    }
client.Labor.BreakTypes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `BreakType` being deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor EmployeeWages
<details><summary><code>client.Labor.EmployeeWages.List() -> *square.ListEmployeeWagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `EmployeeWage` instances for a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.ListEmployeeWagesRequest{
        EmployeeID: square.String(
            "employee_id",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Labor.EmployeeWages.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**employeeID:** `*string` — Filter the returned wages to only those that are associated with the specified employee.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of `EmployeeWage` results to return per page. The number can range between
1 and 200. The default is 200.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A pointer to the next page of `EmployeeWage` results to fetch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.EmployeeWages.Get(ID) -> *square.GetEmployeeWageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single `EmployeeWage` specified by `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.GetEmployeeWagesRequest{
        ID: "id",
    }
client.Labor.EmployeeWages.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `EmployeeWage` being retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor Shifts
<details><summary><code>client.Labor.Shifts.Create(request) -> *square.CreateShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new `Shift`.

A `Shift` represents a complete workday for a single team member.
You must provide the following values in your request to this
endpoint:

- `location_id`
- `team_member_id`
- `start_at`

An attempt to create a new `Shift` can result in a `BAD_REQUEST` error when:
- The `status` of the new `Shift` is `OPEN` and the team member has another
shift with an `OPEN` status.
- The `start_at` date is in the future.
- The `start_at` or `end_at` date overlaps another shift for the same team member.
- The `Break` instances are set in the request and a break `start_at`
is before the `Shift.start_at`, a break `end_at` is after
the `Shift.end_at`, or both.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.CreateShiftRequest{
        IdempotencyKey: square.String(
            "HIDSNG5KS478L",
        ),
        Shift: &square.Shift{
            LocationID: "PAA1RJZZKXBFG",
            StartAt: "2019-01-25T03:11:00-05:00",
            EndAt: square.String(
                "2019-01-25T13:11:00-05:00",
            ),
            Wage: &square.ShiftWage{
                Title: square.String(
                    "Barista",
                ),
                HourlyRate: &square.Money{
                    Amount: square.Int64(
                        1100,
                    ),
                    Currency: square.CurrencyUsd.Ptr(),
                },
                TipEligible: square.Bool(
                    true,
                ),
            },
            Breaks: []*square.Break{
                &square.Break{
                    StartAt: "2019-01-25T06:11:00-05:00",
                    EndAt: square.String(
                        "2019-01-25T06:16:00-05:00",
                    ),
                    BreakTypeID: "REGS1EQR1TPZ5",
                    Name: "Tea Break",
                    ExpectedDuration: "PT5M",
                    IsPaid: true,
                },
            },
            TeamMemberID: square.String(
                "ormj0jJJZ5OZIzxrZYJI",
            ),
            DeclaredCashTipMoney: &square.Money{
                Amount: square.Int64(
                    500,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
        },
    }
client.Labor.Shifts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique string value to ensure the idempotency of the operation.
    
</dd>
</dl>

<dl>
<dd>

**shift:** `*square.Shift` — The `Shift` to be created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.Shifts.Search(request) -> *square.SearchShiftsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `Shift` records for a business.
The list to be returned can be filtered by:
- Location IDs
- Team member IDs
- Shift status (`OPEN` or `CLOSED`)
- Shift start
- Shift end
- Workday details

The list can be sorted by:
- `START_AT`
- `END_AT`
- `CREATED_AT`
- `UPDATED_AT`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.SearchShiftsRequest{
        Query: &square.ShiftQuery{
            Filter: &square.ShiftFilter{
                Workday: &square.ShiftWorkday{
                    DateRange: &square.DateRange{
                        StartDate: square.String(
                            "2019-01-20",
                        ),
                        EndDate: square.String(
                            "2019-02-03",
                        ),
                    },
                    MatchShiftsBy: square.ShiftWorkdayMatcherStartAt.Ptr(),
                    DefaultTimezone: square.String(
                        "America/Los_Angeles",
                    ),
                },
            },
        },
        Limit: square.Int(
            100,
        ),
    }
client.Labor.Shifts.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.ShiftQuery` — Query filters.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of resources in a page (200 by default).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — An opaque cursor for fetching the next page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.Shifts.Get(ID) -> *square.GetShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single `Shift` specified by `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.GetShiftsRequest{
        ID: "id",
    }
client.Labor.Shifts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `Shift` being retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.Shifts.Update(ID, request) -> *square.UpdateShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing `Shift`.

When adding a `Break` to a `Shift`, any earlier `Break` instances in the `Shift` have
the `end_at` property set to a valid RFC-3339 datetime string.

When closing a `Shift`, all `Break` instances in the `Shift` must be complete with `end_at`
set on each `Break`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.UpdateShiftRequest{
        ID: "id",
        Shift: &square.Shift{
            LocationID: "PAA1RJZZKXBFG",
            StartAt: "2019-01-25T03:11:00-05:00",
            EndAt: square.String(
                "2019-01-25T13:11:00-05:00",
            ),
            Wage: &square.ShiftWage{
                Title: square.String(
                    "Bartender",
                ),
                HourlyRate: &square.Money{
                    Amount: square.Int64(
                        1500,
                    ),
                    Currency: square.CurrencyUsd.Ptr(),
                },
                TipEligible: square.Bool(
                    true,
                ),
            },
            Breaks: []*square.Break{
                &square.Break{
                    ID: square.String(
                        "X7GAQYVVRRG6P",
                    ),
                    StartAt: "2019-01-25T06:11:00-05:00",
                    EndAt: square.String(
                        "2019-01-25T06:16:00-05:00",
                    ),
                    BreakTypeID: "REGS1EQR1TPZ5",
                    Name: "Tea Break",
                    ExpectedDuration: "PT5M",
                    IsPaid: true,
                },
            },
            Version: square.Int(
                1,
            ),
            TeamMemberID: square.String(
                "ormj0jJJZ5OZIzxrZYJI",
            ),
            DeclaredCashTipMoney: &square.Money{
                Amount: square.Int64(
                    500,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
        },
    }
client.Labor.Shifts.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the object being updated.
    
</dd>
</dl>

<dl>
<dd>

**shift:** `*square.Shift` — The updated `Shift` object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.Shifts.Delete(ID) -> *square.DeleteShiftResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a `Shift`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.DeleteShiftsRequest{
        ID: "id",
    }
client.Labor.Shifts.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `Shift` being deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor TeamMemberWages
<details><summary><code>client.Labor.TeamMemberWages.List() -> *square.ListTeamMemberWagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of `TeamMemberWage` instances for a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.ListTeamMemberWagesRequest{
        TeamMemberID: square.String(
            "team_member_id",
        ),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Labor.TeamMemberWages.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `*string` 

Filter the returned wages to only those that are associated with the
specified team member.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of `TeamMemberWage` results to return per page. The number can range between
1 and 200. The default is 200.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A pointer to the next page of `EmployeeWage` results to fetch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.TeamMemberWages.Get(ID) -> *square.GetTeamMemberWageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single `TeamMemberWage` specified by `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.GetTeamMemberWagesRequest{
        ID: "id",
    }
client.Labor.TeamMemberWages.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `TeamMemberWage` being retrieved.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Labor WorkweekConfigs
<details><summary><code>client.Labor.WorkweekConfigs.List() -> *square.ListWorkweekConfigsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WorkweekConfig` instances for a business.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.ListWorkweekConfigsRequest{
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Labor.WorkweekConfigs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — The maximum number of `WorkweekConfigs` results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — A pointer to the next page of `WorkweekConfig` results to fetch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Labor.WorkweekConfigs.Get(ID, request) -> *square.UpdateWorkweekConfigResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `WorkweekConfig`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &labor.UpdateWorkweekConfigRequest{
        ID: "id",
        WorkweekConfig: &square.WorkweekConfig{
            StartOfWeek: square.WeekdayMon,
            StartOfDayLocalTime: "10:00",
            Version: square.Int(
                10,
            ),
        },
    }
client.Labor.WorkweekConfigs.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID for the `WorkweekConfig` object being updated.
    
</dd>
</dl>

<dl>
<dd>

**workweekConfig:** `*square.WorkweekConfig` — The updated `WorkweekConfig` object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Locations CustomAttributeDefinitions
<details><summary><code>client.Locations.CustomAttributeDefinitions.List() -> *square.ListLocationCustomAttributeDefinitionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the location-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.
When all response pages are retrieved, the results include all custom attribute definitions
that are visible to the requesting application, including those that are created by other
applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.ListCustomAttributeDefinitionsRequest{
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Locations.CustomAttributeDefinitions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Filters the `CustomAttributeDefinition` results by their `visibility` values.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributeDefinitions.Create(request) -> *square.CreateLocationCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a location-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
Use this endpoint to define a custom attribute that can be associated with locations.
A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties
for a custom attribute. After the definition is created, you can call
[UpsertLocationCustomAttribute](api-endpoint:LocationCustomAttributes-UpsertLocationCustomAttribute) or
[BulkUpsertLocationCustomAttributes](api-endpoint:LocationCustomAttributes-BulkUpsertLocationCustomAttributes)
to set the custom attribute for locations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.CreateLocationCustomAttributeDefinitionRequest{
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Key: square.String(
                "bestseller",
            ),
            Schema: map[string]any{
                "\$ref": "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.String",
            },
            Name: square.String(
                "Bestseller",
            ),
            Description: square.String(
                "Bestselling item at location",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadWriteValues.Ptr(),
        },
    }
client.Locations.CustomAttributeDefinitions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition to create. Note the following:
- With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
definition hosted on the Square CDN. For more information, including supported values and constraints, see
[Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
- `name` is required unless `visibility` is set to `VISIBILITY_HIDDEN`.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributeDefinitions.Get(Key) -> *square.RetrieveLocationCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a location-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
To retrieve a custom attribute definition created by another application, the `visibility`
setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.GetCustomAttributeDefinitionsRequest{
        Key: "key",
        Version: square.Int(
            1,
        ),
    }
client.Locations.CustomAttributeDefinitions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute definition to retrieve. If the requesting application
is not the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute definition, which is used for strongly consistent
reads to guarantee that you receive the most up-to-date data. When included in the request,
Square returns the specified version or a higher version if one exists. If the specified version
is higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributeDefinitions.Update(Key, request) -> *square.UpdateLocationCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a location-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the
`schema` for a `Selection` data type.
Only the definition owner can update a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.UpdateLocationCustomAttributeDefinitionRequest{
        Key: "key",
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Description: square.String(
                "Update the description as desired.",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadOnly.Ptr(),
        },
    }
client.Locations.CustomAttributeDefinitions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to update.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition that contains the fields to update. This endpoint
supports sparse updates, so only new or changed fields need to be included in the request.
Only the following fields can be updated:
- `name`
- `description`
- `visibility`
- `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
selections are supported.

For more information, see
[Update a location custom attribute definition](https://developer.squareup.com/docs/location-custom-attributes-api/custom-attribute-definitions#update-custom-attribute-definition).
To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control, specify the current version of the custom attribute definition. 
If this is not important for your application, `version` can be set to -1.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributeDefinitions.Delete(Key) -> *square.DeleteLocationCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a location-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
Deleting a custom attribute definition also deletes the corresponding custom attribute from
all locations.
Only the definition owner can delete a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.DeleteCustomAttributeDefinitionsRequest{
        Key: "key",
    }
client.Locations.CustomAttributeDefinitions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Locations CustomAttributes
<details><summary><code>client.Locations.CustomAttributes.BatchDelete(request) -> *square.BulkDeleteLocationCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes [custom attributes](entity:CustomAttribute) for locations as a bulk operation.
To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.BulkDeleteLocationCustomAttributesRequest{
        Values: map[string]*square.BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest{
            "id1": &square.BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest{
                Key: square.String(
                    "bestseller",
                ),
            },
            "id2": &square.BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest{
                Key: square.String(
                    "bestseller",
                ),
            },
            "id3": &square.BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest{
                Key: square.String(
                    "phone-number",
                ),
            },
        },
    }
client.Locations.CustomAttributes.BatchDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest` 

The data used to update the `CustomAttribute` objects.
The keys must be unique and are used to map to the corresponding response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributes.BatchUpsert(request) -> *square.BulkUpsertLocationCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates [custom attributes](entity:CustomAttribute) for locations as a bulk operation.
Use this endpoint to set the value of one or more custom attributes for one or more locations.
A custom attribute is based on a custom attribute definition in a Square seller account, which is
created using the [CreateLocationCustomAttributeDefinition](api-endpoint:LocationCustomAttributes-CreateLocationCustomAttributeDefinition) endpoint.
This `BulkUpsertLocationCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
requests and returns a map of individual upsert responses. Each upsert request has a unique ID
and provides a location ID and custom attribute. Each upsert response is returned with the ID
of the corresponding request.
To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.BulkUpsertLocationCustomAttributesRequest{
        Values: map[string]*square.BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest{
            "id1": &square.BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest{
                LocationID: "L0TBCBTB7P8RQ",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "bestseller",
                    ),
                    Value: "hot cocoa",
                },
            },
            "id2": &square.BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest{
                LocationID: "L9XMD04V3STJX",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "bestseller",
                    ),
                    Value: "berry smoothie",
                },
            },
            "id3": &square.BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest{
                LocationID: "L0TBCBTB7P8RQ",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "phone-number",
                    ),
                    Value: "+12223334444",
                },
            },
        },
    }
client.Locations.CustomAttributes.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest` 

A map containing 1 to 25 individual upsert requests. For each request, provide an
arbitrary ID that is unique for this `BulkUpsertLocationCustomAttributes` request and the
information needed to create or update a custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributes.List(LocationID) -> *square.ListLocationCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the [custom attributes](entity:CustomAttribute) associated with a location.
You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
in the same call.
When all response pages are retrieved, the results include all custom attributes that are
visible to the requesting application, including those that are owned by other applications
and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.ListCustomAttributesRequest{
        LocationID: "location_id",
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        WithDefinitions: square.Bool(
            true,
        ),
    }
client.Locations.CustomAttributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the target [location](entity:Location).
    
</dd>
</dl>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Filters the `CustomAttributeDefinition` results by their `visibility` values.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request. For more
information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**withDefinitions:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
custom attribute. Set this parameter to `true` to get the name and description of each custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributes.Get(LocationID, Key) -> *square.RetrieveLocationCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a [custom attribute](entity:CustomAttribute) associated with a location.
You can use the `with_definition` query parameter to also retrieve the custom attribute definition
in the same call.
To retrieve a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.GetCustomAttributesRequest{
        LocationID: "location_id",
        Key: "key",
        WithDefinition: square.Bool(
            true,
        ),
        Version: square.Int(
            1,
        ),
    }
client.Locations.CustomAttributes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the target [location](entity:Location).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to retrieve. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**withDefinition:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of
the custom attribute. Set this parameter to `true` to get the name and description of the custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute, which is used for strongly consistent reads to
guarantee that you receive the most up-to-date data. When included in the request, Square
returns the specified version or a higher version if one exists. If the specified version is
higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributes.Upsert(LocationID, Key, request) -> *square.UpsertLocationCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates a [custom attribute](entity:CustomAttribute) for a location.
Use this endpoint to set the value of a custom attribute for a specified location.
A custom attribute is based on a custom attribute definition in a Square seller account, which
is created using the [CreateLocationCustomAttributeDefinition](api-endpoint:LocationCustomAttributes-CreateLocationCustomAttributeDefinition) endpoint.
To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.UpsertLocationCustomAttributeRequest{
        LocationID: "location_id",
        Key: "key",
        CustomAttribute: &square.CustomAttribute{
            Value: "hot cocoa",
        },
    }
client.Locations.CustomAttributes.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the target [location](entity:Location).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to create or update. This key must match the `key` of a
custom attribute definition in the Square seller account. If the requesting application is not
the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**customAttribute:** `*square.CustomAttribute` 

The custom attribute to create or update, with the following fields:
- `value`. This value must conform to the `schema` specified by the definition.
For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
- `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control for an update operation, include the current version of the custom attribute.
If this is not important for your application, version can be set to -1.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CustomAttributes.Delete(LocationID, Key) -> *square.DeleteLocationCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a [custom attribute](entity:CustomAttribute) associated with a location.
To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.DeleteCustomAttributesRequest{
        LocationID: "location_id",
        Key: "key",
    }
client.Locations.CustomAttributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the target [location](entity:Location).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to delete. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Locations Transactions
<details><summary><code>client.Locations.Transactions.List(LocationID) -> *square.ListTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists transactions for a particular location.

Transactions include payment information from sales and exchanges and refund
information from returns and exchanges.

Max results per [page](https://developer.squareup.com/docs/working-with-apis/pagination): 50
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.ListTransactionsRequest{
        LocationID: "location_id",
        BeginTime: square.String(
            "begin_time",
        ),
        EndTime: square.String(
            "end_time",
        ),
        SortOrder: square.SortOrderDesc.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Locations.Transactions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the location to list transactions for.
    
</dd>
</dl>

<dl>
<dd>

**beginTime:** `*string` 

The beginning of the requested reporting period, in RFC 3339 format.

See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.

Default value: The current time minus one year.
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 

The end of the requested reporting period, in RFC 3339 format.

See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.

Default value: The current time.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

The order in which results are listed in the response (`ASC` for
oldest first, `DESC` for newest first).

Default value: `DESC`
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for your original query.

See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Transactions.Get(LocationID, TransactionID) -> *square.GetTransactionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves details for a single transaction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.GetTransactionsRequest{
        LocationID: "location_id",
        TransactionID: "transaction_id",
    }
client.Locations.Transactions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — The ID of the transaction's associated location.
    
</dd>
</dl>

<dl>
<dd>

**transactionID:** `string` — The ID of the transaction to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Transactions.Capture(LocationID, TransactionID) -> *square.CaptureTransactionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Captures a transaction that was created with the [Charge](api-endpoint:Transactions-Charge)
endpoint with a `delay_capture` value of `true`.


See [Delayed capture transactions](https://developer.squareup.com/docs/payments/transactions/overview#delayed-capture)
for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.CaptureTransactionsRequest{
        LocationID: "location_id",
        TransactionID: "transaction_id",
    }
client.Locations.Transactions.Capture(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — 
    
</dd>
</dl>

<dl>
<dd>

**transactionID:** `string` — 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.Transactions.Void(LocationID, TransactionID) -> *square.VoidTransactionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a transaction that was created with the [Charge](api-endpoint:Transactions-Charge)
endpoint with a `delay_capture` value of `true`.


See [Delayed capture transactions](https://developer.squareup.com/docs/payments/transactions/overview#delayed-capture)
for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &locations.VoidTransactionsRequest{
        LocationID: "location_id",
        TransactionID: "transaction_id",
    }
client.Locations.Transactions.Void(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**locationID:** `string` — 
    
</dd>
</dl>

<dl>
<dd>

**transactionID:** `string` — 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loyalty Accounts
<details><summary><code>client.Loyalty.Accounts.Create(request) -> *square.CreateLoyaltyAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a loyalty account. To create a loyalty account, you must provide the `program_id` and a `mapping` with the `phone_number` of the buyer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.CreateLoyaltyAccountRequest{
        LoyaltyAccount: &square.LoyaltyAccount{
            ProgramID: "d619f755-2d17-41f3-990d-c04ecedd64dd",
            Mapping: &square.LoyaltyAccountMapping{
                PhoneNumber: square.String(
                    "+14155551234",
                ),
            },
        },
        IdempotencyKey: "ec78c477-b1c3-4899-a209-a4e71337c996",
    }
client.Loyalty.Accounts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loyaltyAccount:** `*square.LoyaltyAccount` — The loyalty account to create.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreateLoyaltyAccount` request. 
Keys can be any valid string, but must be unique for every request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Accounts.Search(request) -> *square.SearchLoyaltyAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for loyalty accounts in a loyalty program.

You can search for a loyalty account using the phone number or customer ID associated with the account. To return all loyalty accounts, specify an empty `query` object or omit it entirely.

Search results are sorted by `created_at` in ascending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.SearchLoyaltyAccountsRequest{
        Query: &square.SearchLoyaltyAccountsRequestLoyaltyAccountQuery{
            Mappings: []*square.LoyaltyAccountMapping{
                &square.LoyaltyAccountMapping{
                    PhoneNumber: square.String(
                        "+14155551234",
                    ),
                },
            },
        },
        Limit: square.Int(
            10,
        ),
    }
client.Loyalty.Accounts.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.SearchLoyaltyAccountsRequestLoyaltyAccountQuery` — The search criteria for the request.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to include in the response. The default value is 30.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to 
this endpoint. Provide this to retrieve the next set of 
results for the original query.

For more information, 
see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Accounts.Get(AccountID) -> *square.GetLoyaltyAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a loyalty account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.GetAccountsRequest{
        AccountID: "account_id",
    }
client.Loyalty.Accounts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The ID of the [loyalty account](entity:LoyaltyAccount) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Accounts.AccumulatePoints(AccountID, request) -> *square.AccumulateLoyaltyPointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds points earned from a purchase to a [loyalty account](entity:LoyaltyAccount).

- If you are using the Orders API to manage orders, provide the `order_id`. Square reads the order
to compute the points earned from both the base loyalty program and an associated
[loyalty promotion](entity:LoyaltyPromotion). For purchases that qualify for multiple accrual
rules, Square computes points based on the accrual rule that grants the most points.
For purchases that qualify for multiple promotions, Square computes points based on the most
recently created promotion. A purchase must first qualify for program points to be eligible for promotion points.

- If you are not using the Orders API to manage orders, provide `points` with the number of points to add.
You must first perform a client-side computation of the points earned from the loyalty program and
loyalty promotion. For spend-based and visit-based programs, you can call [CalculateLoyaltyPoints](api-endpoint:Loyalty-CalculateLoyaltyPoints)
to compute the points earned from the base loyalty program. For information about computing points earned from a loyalty promotion, see
[Calculating promotion points](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#calculate-promotion-points).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.AccumulateLoyaltyPointsRequest{
        AccountID: "account_id",
        AccumulatePoints: &square.LoyaltyEventAccumulatePoints{
            OrderID: square.String(
                "RFZfrdtm3mhO1oGzf5Cx7fEMsmGZY",
            ),
        },
        IdempotencyKey: "58b90739-c3e8-4b11-85f7-e636d48d72cb",
        LocationID: "P034NEENMD09F",
    }
client.Loyalty.Accounts.AccumulatePoints(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The ID of the target [loyalty account](entity:LoyaltyAccount).
    
</dd>
</dl>

<dl>
<dd>

**accumulatePoints:** `*square.LoyaltyEventAccumulatePoints` 

The points to add to the account. 
If you are using the Orders API to manage orders, specify the order ID.
Otherwise, specify the points to add.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies the `AccumulateLoyaltyPoints` request. 
Keys can be any valid string but must be unique for every request.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The [location](entity:Location) where the purchase was made.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Accounts.Adjust(AccountID, request) -> *square.AdjustLoyaltyPointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds points to or subtracts points from a buyer's account.

Use this endpoint only when you need to manually adjust points. Otherwise, in your application flow, you call
[AccumulateLoyaltyPoints](api-endpoint:Loyalty-AccumulateLoyaltyPoints)
to add points when a buyer pays for the purchase.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.AdjustLoyaltyPointsRequest{
        AccountID: "account_id",
        IdempotencyKey: "bc29a517-3dc9-450e-aa76-fae39ee849d1",
        AdjustPoints: &square.LoyaltyEventAdjustPoints{
            Points: 10,
            Reason: square.String(
                "Complimentary points",
            ),
        },
    }
client.Loyalty.Accounts.Adjust(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The ID of the target [loyalty account](entity:LoyaltyAccount).
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `AdjustLoyaltyPoints` request. 
Keys can be any valid string, but must be unique for every request.
    
</dd>
</dl>

<dl>
<dd>

**adjustPoints:** `*square.LoyaltyEventAdjustPoints` 

The points to add or subtract and the reason for the adjustment. To add points, specify a positive integer.
To subtract points, specify a negative integer.
    
</dd>
</dl>

<dl>
<dd>

**allowNegativeBalance:** `*bool` 

Indicates whether to allow a negative adjustment to result in a negative balance. If `true`, a negative
balance is allowed when subtracting points. If `false`, Square returns a `BAD_REQUEST` error when subtracting
the specified number of points would result in a negative balance. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loyalty Programs
<details><summary><code>client.Loyalty.Programs.List() -> *square.ListLoyaltyProgramsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of loyalty programs in the seller's account.
Loyalty programs define how buyers can earn points and redeem points for rewards. Square sellers can have only one loyalty program, which is created and managed from the Seller Dashboard. For more information, see [Loyalty Program Overview](https://developer.squareup.com/docs/loyalty/overview).


Replaced with [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) when used with the keyword `main`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Loyalty.Programs.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Programs.Get(ProgramID) -> *square.GetLoyaltyProgramResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the loyalty program in a seller's account, specified by the program ID or the keyword `main`.

Loyalty programs define how buyers can earn points and redeem points for rewards. Square sellers can have only one loyalty program, which is created and managed from the Seller Dashboard. For more information, see [Loyalty Program Overview](https://developer.squareup.com/docs/loyalty/overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.GetProgramsRequest{
        ProgramID: "program_id",
    }
client.Loyalty.Programs.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` — The ID of the loyalty program or the keyword `main`. Either value can be used to retrieve the single loyalty program that belongs to the seller.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Programs.Calculate(ProgramID, request) -> *square.CalculateLoyaltyPointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Calculates the number of points a buyer can earn from a purchase. Applications might call this endpoint
to display the points to the buyer.

- If you are using the Orders API to manage orders, provide the `order_id` and (optional) `loyalty_account_id`.
Square reads the order to compute the points earned from the base loyalty program and an associated
[loyalty promotion](entity:LoyaltyPromotion).

- If you are not using the Orders API to manage orders, provide `transaction_amount_money` with the
purchase amount. Square uses this amount to calculate the points earned from the base loyalty program,
but not points earned from a loyalty promotion. For spend-based and visit-based programs, the `tax_mode`
setting of the accrual rule indicates how taxes should be treated for loyalty points accrual.
If the purchase qualifies for program points, call
[ListLoyaltyPromotions](api-endpoint:Loyalty-ListLoyaltyPromotions) and perform a client-side computation
to calculate whether the purchase also qualifies for promotion points. For more information, see
[Calculating promotion points](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#calculate-promotion-points).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.CalculateLoyaltyPointsRequest{
        ProgramID: "program_id",
        OrderID: square.String(
            "RFZfrdtm3mhO1oGzf5Cx7fEMsmGZY",
        ),
        LoyaltyAccountID: square.String(
            "79b807d2-d786-46a9-933b-918028d7a8c5",
        ),
    }
client.Loyalty.Programs.Calculate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` — The ID of the [loyalty program](entity:LoyaltyProgram), which defines the rules for accruing points.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*string` 

The [order](entity:Order) ID for which to calculate the points.
Specify this field if your application uses the Orders API to process orders.
Otherwise, specify the `transaction_amount_money`.
    
</dd>
</dl>

<dl>
<dd>

**transactionAmountMoney:** `*square.Money` 

The purchase amount for which to calculate the points. 
Specify this field if your application does not use the Orders API to process orders.
Otherwise, specify the `order_id`.
    
</dd>
</dl>

<dl>
<dd>

**loyaltyAccountID:** `*string` 

The ID of the target [loyalty account](entity:LoyaltyAccount). Optionally specify this field
if your application uses the Orders API to process orders.

If specified, the `promotion_points` field in the response shows the number of points the buyer would
earn from the purchase. In this case, Square uses the account ID to determine whether the promotion's
`trigger_limit` (the maximum number of times that a buyer can trigger the promotion) has been reached.
If not specified, the `promotion_points` field shows the number of points the purchase qualifies
for regardless of the trigger limit.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loyalty Rewards
<details><summary><code>client.Loyalty.Rewards.Create(request) -> *square.CreateLoyaltyRewardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a loyalty reward. In the process, the endpoint does following:

- Uses the `reward_tier_id` in the request to determine the number of points
to lock for this reward.
- If the request includes `order_id`, it adds the reward and related discount to the order.

After a reward is created, the points are locked and
not available for the buyer to redeem another reward.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.CreateLoyaltyRewardRequest{
        Reward: &square.LoyaltyReward{
            LoyaltyAccountID: "5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd",
            RewardTierID: "e1b39225-9da5-43d1-a5db-782cdd8ad94f",
            OrderID: square.String(
                "RFZfrdtm3mhO1oGzf5Cx7fEMsmGZY",
            ),
        },
        IdempotencyKey: "18c2e5ea-a620-4b1f-ad60-7b167285e451",
    }
client.Loyalty.Rewards.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**reward:** `*square.LoyaltyReward` — The reward to create.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreateLoyaltyReward` request. 
Keys can be any valid string, but must be unique for every request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Rewards.Search(request) -> *square.SearchLoyaltyRewardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for loyalty rewards. This endpoint accepts a request with no query filters and returns results for all loyalty accounts.
If you include a `query` object, `loyalty_account_id` is required and `status` is  optional.

If you know a reward ID, use the
[RetrieveLoyaltyReward](api-endpoint:Loyalty-RetrieveLoyaltyReward) endpoint.

Search results are sorted by `updated_at` in descending order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.SearchLoyaltyRewardsRequest{
        Query: &square.SearchLoyaltyRewardsRequestLoyaltyRewardQuery{
            LoyaltyAccountID: "5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd",
        },
        Limit: square.Int(
            10,
        ),
    }
client.Loyalty.Rewards.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.SearchLoyaltyRewardsRequestLoyaltyRewardQuery` 

The search criteria for the request. 
If empty, the endpoint retrieves all loyalty rewards in the loyalty program.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The maximum number of results to return in the response. The default value is 30.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to 
this endpoint. Provide this to retrieve the next set of 
results for the original query.
For more information, 
see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Rewards.Get(RewardID) -> *square.GetLoyaltyRewardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a loyalty reward.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.GetRewardsRequest{
        RewardID: "reward_id",
    }
client.Loyalty.Rewards.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**rewardID:** `string` — The ID of the [loyalty reward](entity:LoyaltyReward) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Rewards.Delete(RewardID) -> *square.DeleteLoyaltyRewardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a loyalty reward by doing the following:

- Returns the loyalty points back to the loyalty account.
- If an order ID was specified when the reward was created
(see [CreateLoyaltyReward](api-endpoint:Loyalty-CreateLoyaltyReward)),
it updates the order by removing the reward and related
discounts.

You cannot delete a reward that has reached the terminal state (REDEEMED).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.DeleteRewardsRequest{
        RewardID: "reward_id",
    }
client.Loyalty.Rewards.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**rewardID:** `string` — The ID of the [loyalty reward](entity:LoyaltyReward) to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Rewards.Redeem(RewardID, request) -> *square.RedeemLoyaltyRewardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Redeems a loyalty reward.

The endpoint sets the reward to the `REDEEMED` terminal state.

If you are using your own order processing system (not using the
Orders API), you call this endpoint after the buyer paid for the
purchase.

After the reward reaches the terminal state, it cannot be deleted.
In other words, points used for the reward cannot be returned
to the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &loyalty.RedeemLoyaltyRewardRequest{
        RewardID: "reward_id",
        IdempotencyKey: "98adc7f7-6963-473b-b29c-f3c9cdd7d994",
        LocationID: "P034NEENMD09F",
    }
client.Loyalty.Rewards.Redeem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**rewardID:** `string` — The ID of the [loyalty reward](entity:LoyaltyReward) to redeem.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `RedeemLoyaltyReward` request. 
Keys can be any valid string, but must be unique for every request.
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` — The ID of the [location](entity:Location) where the reward is redeemed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loyalty Programs Promotions
<details><summary><code>client.Loyalty.Programs.Promotions.List(ProgramID) -> *square.ListLoyaltyPromotionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the loyalty promotions associated with a [loyalty program](entity:LoyaltyProgram).
Results are sorted by the `created_at` date in descending order (newest to oldest).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &programs.ListPromotionsRequest{
        ProgramID: "program_id",
        Status: square.LoyaltyPromotionStatusActive.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Loyalty.Programs.Promotions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` 

The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID,
call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the `main` keyword.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*square.LoyaltyPromotionStatus` 

The status to filter the results by. If a status is provided, only loyalty promotions
with the specified status are returned. Otherwise, all loyalty promotions associated with
the loyalty program are returned.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response.
The minimum value is 1 and the maximum value is 30. The default value is 30.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Programs.Promotions.Create(ProgramID, request) -> *square.CreateLoyaltyPromotionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a loyalty promotion for a [loyalty program](entity:LoyaltyProgram). A loyalty promotion
enables buyers to earn points in addition to those earned from the base loyalty program.

This endpoint sets the loyalty promotion to the `ACTIVE` or `SCHEDULED` status, depending on the
`available_time` setting. A loyalty program can have a maximum of 10 loyalty promotions with an
`ACTIVE` or `SCHEDULED` status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &programs.CreateLoyaltyPromotionRequest{
        ProgramID: "program_id",
        LoyaltyPromotion: &square.LoyaltyPromotion{
            Name: "Tuesday Happy Hour Promo",
            Incentive: &square.LoyaltyPromotionIncentive{
                Type: square.LoyaltyPromotionIncentiveTypePointsMultiplier,
                PointsMultiplierData: &square.LoyaltyPromotionIncentivePointsMultiplierData{
                    Multiplier: square.String(
                        "3.0",
                    ),
                },
            },
            AvailableTime: &square.LoyaltyPromotionAvailableTimeData{
                TimePeriods: []string{
                    `BEGIN:VEVENT
                    DTSTART:20220816T160000
                    DURATION:PT2H
                    RRULE:FREQ=WEEKLY;BYDAY=TU
                    END:VEVENT`,
                },
            },
            TriggerLimit: &square.LoyaltyPromotionTriggerLimit{
                Times: 1,
                Interval: square.LoyaltyPromotionTriggerLimitIntervalDay.Ptr(),
            },
            MinimumSpendAmountMoney: &square.Money{
                Amount: square.Int64(
                    2000,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
            QualifyingCategoryIDs: []string{
                "XTQPYLR3IIU9C44VRCB3XD12",
            },
        },
        IdempotencyKey: "ec78c477-b1c3-4899-a209-a4e71337c996",
    }
client.Loyalty.Programs.Promotions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` 

The ID of the [loyalty program](entity:LoyaltyProgram) to associate with the promotion.
To get the program ID, call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram)
using the `main` keyword.
    
</dd>
</dl>

<dl>
<dd>

**loyaltyPromotion:** `*square.LoyaltyPromotion` — The loyalty promotion to create.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique identifier for this request, which is used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Programs.Promotions.Get(ProgramID, PromotionID) -> *square.GetLoyaltyPromotionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a loyalty promotion.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &programs.GetPromotionsRequest{
        ProgramID: "program_id",
        PromotionID: "promotion_id",
    }
client.Loyalty.Programs.Promotions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` 

The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID,
call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the `main` keyword.
    
</dd>
</dl>

<dl>
<dd>

**promotionID:** `string` — The ID of the [loyalty promotion](entity:LoyaltyPromotion) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loyalty.Programs.Promotions.Cancel(ProgramID, PromotionID) -> *square.CancelLoyaltyPromotionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a loyalty promotion. Use this endpoint to cancel an `ACTIVE` promotion earlier than the
end date, cancel an `ACTIVE` promotion when an end date is not specified, or cancel a `SCHEDULED` promotion.
Because updating a promotion is not supported, you can also use this endpoint to cancel a promotion before
you create a new one.

This endpoint sets the loyalty promotion to the `CANCELED` state
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &programs.CancelPromotionsRequest{
        ProgramID: "program_id",
        PromotionID: "promotion_id",
    }
client.Loyalty.Programs.Promotions.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**programID:** `string` — The ID of the base [loyalty program](entity:LoyaltyProgram).
    
</dd>
</dl>

<dl>
<dd>

**promotionID:** `string` 

The ID of the [loyalty promotion](entity:LoyaltyPromotion) to cancel. You can cancel a
promotion that has an `ACTIVE` or `SCHEDULED` status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Merchants CustomAttributeDefinitions
<details><summary><code>client.Merchants.CustomAttributeDefinitions.List() -> *square.ListMerchantCustomAttributeDefinitionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the merchant-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.
When all response pages are retrieved, the results include all custom attribute definitions
that are visible to the requesting application, including those that are created by other
applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.ListCustomAttributeDefinitionsRequest{
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
    }
client.Merchants.CustomAttributeDefinitions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Filters the `CustomAttributeDefinition` results by their `visibility` values.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request.
For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributeDefinitions.Create(request) -> *square.CreateMerchantCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a merchant-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
Use this endpoint to define a custom attribute that can be associated with a merchant connecting to your application.
A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties
for a custom attribute. After the definition is created, you can call
[UpsertMerchantCustomAttribute](api-endpoint:MerchantCustomAttributes-UpsertMerchantCustomAttribute) or
[BulkUpsertMerchantCustomAttributes](api-endpoint:MerchantCustomAttributes-BulkUpsertMerchantCustomAttributes)
to set the custom attribute for a merchant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.CreateMerchantCustomAttributeDefinitionRequest{
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Key: square.String(
                "alternative_seller_name",
            ),
            Schema: map[string]any{
                "\$ref": "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.String",
            },
            Name: square.String(
                "Alternative Merchant Name",
            ),
            Description: square.String(
                "This is the other name this merchant goes by.",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadOnly.Ptr(),
        },
    }
client.Merchants.CustomAttributeDefinitions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition to create. Note the following:
- With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
definition hosted on the Square CDN. For more information, including supported values and constraints, see
[Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
- `name` is required unless `visibility` is set to `VISIBILITY_HIDDEN`.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributeDefinitions.Get(Key) -> *square.RetrieveMerchantCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a merchant-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
To retrieve a custom attribute definition created by another application, the `visibility`
setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.GetCustomAttributeDefinitionsRequest{
        Key: "key",
        Version: square.Int(
            1,
        ),
    }
client.Merchants.CustomAttributeDefinitions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute definition to retrieve. If the requesting application
is not the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute definition, which is used for strongly consistent
reads to guarantee that you receive the most up-to-date data. When included in the request,
Square returns the specified version or a higher version if one exists. If the specified version
is higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributeDefinitions.Update(Key, request) -> *square.UpdateMerchantCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a merchant-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.
Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the
`schema` for a `Selection` data type.
Only the definition owner can update a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.UpdateMerchantCustomAttributeDefinitionRequest{
        Key: "key",
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Description: square.String(
                "Update the description as desired.",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadOnly.Ptr(),
        },
    }
client.Merchants.CustomAttributeDefinitions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to update.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition that contains the fields to update. This endpoint
supports sparse updates, so only new or changed fields need to be included in the request.
Only the following fields can be updated:
- `name`
- `description`
- `visibility`
- `schema` for a `Selection` data type. Only changes to the named options or the maximum number of allowed
selections are supported.
For more information, see
[Update a merchant custom attribute definition](https://developer.squareup.com/docs/merchant-custom-attributes-api/custom-attribute-definitions#update-custom-attribute-definition).
The version field must match the current version of the custom attribute definition to enable
[optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
If this is not important for your application, version can be set to -1. For any other values, the request fails with a BAD_REQUEST error.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributeDefinitions.Delete(Key) -> *square.DeleteMerchantCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a merchant-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.
Deleting a custom attribute definition also deletes the corresponding custom attribute from
the merchant.
Only the definition owner can delete a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.DeleteCustomAttributeDefinitionsRequest{
        Key: "key",
    }
client.Merchants.CustomAttributeDefinitions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Merchants CustomAttributes
<details><summary><code>client.Merchants.CustomAttributes.BatchDelete(request) -> *square.BulkDeleteMerchantCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes [custom attributes](entity:CustomAttribute) for a merchant as a bulk operation.
To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.BulkDeleteMerchantCustomAttributesRequest{
        Values: map[string]*square.BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest{
            "id1": &square.BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest{
                Key: square.String(
                    "alternative_seller_name",
                ),
            },
            "id2": &square.BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest{
                Key: square.String(
                    "has_seen_tutorial",
                ),
            },
        },
    }
client.Merchants.CustomAttributes.BatchDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkDeleteMerchantCustomAttributesRequestMerchantCustomAttributeDeleteRequest` 

The data used to update the `CustomAttribute` objects.
The keys must be unique and are used to map to the corresponding response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributes.BatchUpsert(request) -> *square.BulkUpsertMerchantCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates [custom attributes](entity:CustomAttribute) for a merchant as a bulk operation.
Use this endpoint to set the value of one or more custom attributes for a merchant.
A custom attribute is based on a custom attribute definition in a Square seller account, which is
created using the [CreateMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-CreateMerchantCustomAttributeDefinition) endpoint.
This `BulkUpsertMerchantCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
requests and returns a map of individual upsert responses. Each upsert request has a unique ID
and provides a merchant ID and custom attribute. Each upsert response is returned with the ID
of the corresponding request.
To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.BulkUpsertMerchantCustomAttributesRequest{
        Values: map[string]*square.BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest{
            "id1": &square.BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest{
                MerchantID: "DM7VKY8Q63GNP",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "alternative_seller_name",
                    ),
                    Value: "Ultimate Sneaker Store",
                },
            },
            "id2": &square.BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest{
                MerchantID: "DM7VKY8Q63GNP",
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "has_seen_tutorial",
                    ),
                    Value: true,
                },
            },
        },
    }
client.Merchants.CustomAttributes.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkUpsertMerchantCustomAttributesRequestMerchantCustomAttributeUpsertRequest` 

A map containing 1 to 25 individual upsert requests. For each request, provide an
arbitrary ID that is unique for this `BulkUpsertMerchantCustomAttributes` request and the
information needed to create or update a custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributes.List(MerchantID) -> *square.ListMerchantCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the [custom attributes](entity:CustomAttribute) associated with a merchant.
You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
in the same call.
When all response pages are retrieved, the results include all custom attributes that are
visible to the requesting application, including those that are owned by other applications
and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.ListCustomAttributesRequest{
        MerchantID: "merchant_id",
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Limit: square.Int(
            1,
        ),
        Cursor: square.String(
            "cursor",
        ),
        WithDefinitions: square.Bool(
            true,
        ),
    }
client.Merchants.CustomAttributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` — The ID of the target [merchant](entity:Merchant).
    
</dd>
</dl>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Filters the `CustomAttributeDefinition` results by their `visibility` values.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory.
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.
The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint.
Provide this cursor to retrieve the next page of results for your original request. For more
information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**withDefinitions:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
custom attribute. Set this parameter to `true` to get the name and description of each custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributes.Get(MerchantID, Key) -> *square.RetrieveMerchantCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a [custom attribute](entity:CustomAttribute) associated with a merchant.
You can use the `with_definition` query parameter to also retrieve the custom attribute definition
in the same call.
To retrieve a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.GetCustomAttributesRequest{
        MerchantID: "merchant_id",
        Key: "key",
        WithDefinition: square.Bool(
            true,
        ),
        Version: square.Int(
            1,
        ),
    }
client.Merchants.CustomAttributes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` — The ID of the target [merchant](entity:Merchant).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to retrieve. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**withDefinition:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of
the custom attribute. Set this parameter to `true` to get the name and description of the custom
attribute, information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

The current version of the custom attribute, which is used for strongly consistent reads to
guarantee that you receive the most up-to-date data. When included in the request, Square
returns the specified version or a higher version if one exists. If the specified version is
higher than the current version, Square returns a `BAD_REQUEST` error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributes.Upsert(MerchantID, Key, request) -> *square.UpsertMerchantCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates a [custom attribute](entity:CustomAttribute) for a merchant.
Use this endpoint to set the value of a custom attribute for a specified merchant.
A custom attribute is based on a custom attribute definition in a Square seller account, which
is created using the [CreateMerchantCustomAttributeDefinition](api-endpoint:MerchantCustomAttributes-CreateMerchantCustomAttributeDefinition) endpoint.
To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.UpsertMerchantCustomAttributeRequest{
        MerchantID: "merchant_id",
        Key: "key",
        CustomAttribute: &square.CustomAttribute{
            Value: "Ultimate Sneaker Store",
        },
    }
client.Merchants.CustomAttributes.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` — The ID of the target [merchant](entity:Merchant).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to create or update. This key must match the `key` of a
custom attribute definition in the Square seller account. If the requesting application is not
the definition owner, you must use the qualified key.
    
</dd>
</dl>

<dl>
<dd>

**customAttribute:** `*square.CustomAttribute` 

The custom attribute to create or update, with the following fields:
- `value`. This value must conform to the `schema` specified by the definition.
For more information, see [Supported data types](https://developer.squareup.com/docs/devtools/customattributes/overview#supported-data-types).
- The version field must match the current version of the custom attribute definition to enable
[optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
If this is not important for your application, version can be set to -1. For any other values, the request fails with a BAD_REQUEST error.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. For more information,
see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Merchants.CustomAttributes.Delete(MerchantID, Key) -> *square.DeleteMerchantCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a [custom attribute](entity:CustomAttribute) associated with a merchant.
To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchants.DeleteCustomAttributesRequest{
        MerchantID: "merchant_id",
        Key: "key",
    }
client.Merchants.CustomAttributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` — The ID of the target [merchant](entity:Merchant).
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 

The key of the custom attribute to delete. This key must match the `key` of a custom
attribute definition in the Square seller account. If the requesting application is not the
definition owner, you must use the qualified key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Orders CustomAttributeDefinitions
<details><summary><code>client.Orders.CustomAttributeDefinitions.List() -> *square.ListOrderCustomAttributeDefinitionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the order-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.

When all response pages are retrieved, the results include all custom attribute definitions
that are visible to the requesting application, including those that are created by other
applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that
seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.ListCustomAttributeDefinitionsRequest{
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
    }
client.Orders.CustomAttributeDefinitions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Requests that all of the custom attributes be returned, or only those that are read-only or read-write.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint. 
Provide this cursor to retrieve the next page of results for your original request. 
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory. 
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. 
The default value is 20.
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributeDefinitions.Create(request) -> *square.CreateOrderCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an order-related custom attribute definition.  Use this endpoint to
define a custom attribute that can be associated with orders.

After creating a custom attribute definition, you can set the custom attribute for orders
in the Square seller account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.CreateOrderCustomAttributeDefinitionRequest{
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Key: square.String(
                "cover-count",
            ),
            Schema: map[string]any{
                "\$ref": "https://developer-production-s.squarecdn.com/schemas/v1/common.json#squareup.common.Number",
            },
            Name: square.String(
                "Cover count",
            ),
            Description: square.String(
                "The number of people seated at a table",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadWriteValues.Ptr(),
        },
        IdempotencyKey: square.String(
            "IDEMPOTENCY_KEY",
        ),
    }
client.Orders.CustomAttributeDefinitions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition to create. Note the following:
- With the exception of the `Selection` data type, the `schema` is specified as a simple URL to the JSON schema
definition hosted on the Square CDN. For more information, including supported values and constraints, see
[Specifying the schema](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attribute-definitions#specify-schema).
- If provided, `name` must be unique (case-sensitive) across all visible customer-related custom attribute definitions for the seller.
- All custom attributes are visible in exported customer data, including those set to `VISIBILITY_HIDDEN`.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. 
For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributeDefinitions.Get(Key) -> *square.RetrieveOrderCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.

To retrieve a custom attribute definition created by another application, the `visibility`
setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.GetCustomAttributeDefinitionsRequest{
        Key: "key",
        Version: square.Int(
            1,
        ),
    }
client.Orders.CustomAttributeDefinitions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control, include this optional field and specify the current version of the custom attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributeDefinitions.Update(Key, request) -> *square.UpdateOrderCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an order-related custom attribute definition for a Square seller account.

Only the definition owner can update a custom attribute definition. Note that sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.UpdateOrderCustomAttributeDefinitionRequest{
        Key: "key",
        CustomAttributeDefinition: &square.CustomAttributeDefinition{
            Key: square.String(
                "cover-count",
            ),
            Visibility: square.CustomAttributeDefinitionVisibilityVisibilityReadOnly.Ptr(),
            Version: square.Int(
                1,
            ),
        },
        IdempotencyKey: square.String(
            "IDEMPOTENCY_KEY",
        ),
    }
client.Orders.CustomAttributeDefinitions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to update.
    
</dd>
</dl>

<dl>
<dd>

**customAttributeDefinition:** `*square.CustomAttributeDefinition` 

The custom attribute definition that contains the fields to update. This endpoint supports sparse updates, 
so only new or changed fields need to be included in the request.  For more information, see 
[Updatable definition fields](https://developer.squareup.com/docs/orders-custom-attributes-api/custom-attribute-definitions#updatable-definition-fields).

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control, include the optional `version` field and specify the current version of the custom attribute definition.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. 
For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributeDefinitions.Delete(Key) -> *square.DeleteOrderCustomAttributeDefinitionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an order-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.

Only the definition owner can delete a custom attribute definition.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.DeleteCustomAttributeDefinitionsRequest{
        Key: "key",
    }
client.Orders.CustomAttributeDefinitions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — The key of the custom attribute definition to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Orders CustomAttributes
<details><summary><code>client.Orders.CustomAttributes.BatchDelete(request) -> *square.BulkDeleteOrderCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes order [custom attributes](entity:CustomAttribute) as a bulk operation.

Use this endpoint to delete one or more custom attributes from one or more orders.
A custom attribute is based on a custom attribute definition in a Square seller account.  (To create a
custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)

This `BulkDeleteOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual delete
requests and returns a map of individual delete responses. Each delete request has a unique ID
and provides an order ID and custom attribute. Each delete response is returned with the ID
of the corresponding request.

To delete a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.BulkDeleteOrderCustomAttributesRequest{
        Values: map[string]*square.BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute{
            "cover-count": &square.BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute{
                Key: square.String(
                    "cover-count",
                ),
                OrderID: "7BbXGEIWNldxAzrtGf9GPVZTwZ4F",
            },
            "table-number": &square.BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute{
                Key: square.String(
                    "table-number",
                ),
                OrderID: "7BbXGEIWNldxAzrtGf9GPVZTwZ4F",
            },
        },
    }
client.Orders.CustomAttributes.BatchDelete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute` — A map of requests that correspond to individual delete operations for custom attributes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributes.BatchUpsert(request) -> *square.BulkUpsertOrderCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates order [custom attributes](entity:CustomAttribute) as a bulk operation.

Use this endpoint to delete one or more custom attributes from one or more orders.
A custom attribute is based on a custom attribute definition in a Square seller account.  (To create a
custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)

This `BulkUpsertOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert
requests and returns a map of individual upsert responses. Each upsert request has a unique ID
and provides an order ID and custom attribute. Each upsert response is returned with the ID
of the corresponding request.

To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.BulkUpsertOrderCustomAttributesRequest{
        Values: map[string]*square.BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute{
            "cover-count": &square.BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute{
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "cover-count",
                    ),
                    Value: "6",
                    Version: square.Int(
                        2,
                    ),
                },
                OrderID: "7BbXGEIWNldxAzrtGf9GPVZTwZ4F",
            },
            "table-number": &square.BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute{
                CustomAttribute: &square.CustomAttribute{
                    Key: square.String(
                        "table-number",
                    ),
                    Value: "11",
                    Version: square.Int(
                        4,
                    ),
                },
                OrderID: "7BbXGEIWNldxAzrtGf9GPVZTwZ4F",
            },
        },
    }
client.Orders.CustomAttributes.BatchUpsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]*square.BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute` — A map of requests that correspond to individual upsert operations for custom attributes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributes.List(OrderID) -> *square.ListOrderCustomAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the [custom attributes](entity:CustomAttribute) associated with an order.

You can use the `with_definitions` query parameter to also retrieve custom attribute definitions
in the same call.

When all response pages are retrieved, the results include all custom attributes that are
visible to the requesting application, including those that are owned by other applications
and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.ListCustomAttributesRequest{
        OrderID: "order_id",
        VisibilityFilter: square.VisibilityFilterAll.Ptr(),
        Cursor: square.String(
            "cursor",
        ),
        Limit: square.Int(
            1,
        ),
        WithDefinitions: square.Bool(
            true,
        ),
    }
client.Orders.CustomAttributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the target [order](entity:Order).
    
</dd>
</dl>

<dl>
<dd>

**visibilityFilter:** `*square.VisibilityFilter` — Requests that all of the custom attributes be returned, or only those that are read-only or read-write.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

The cursor returned in the paged response from the previous call to this endpoint. 
Provide this cursor to retrieve the next page of results for your original request. 
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to return in a single paged response. This limit is advisory. 
The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. 
The default value is 20.
For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
    
</dd>
</dl>

<dl>
<dd>

**withDefinitions:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each
custom attribute. Set this parameter to `true` to get the name and description of each custom attribute, 
information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributes.Get(OrderID, CustomAttributeKey) -> *square.RetrieveOrderCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a [custom attribute](entity:CustomAttribute) associated with an order.

You can use the `with_definition` query parameter to also retrieve the custom attribute definition
in the same call.

To retrieve a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.GetCustomAttributesRequest{
        OrderID: "order_id",
        CustomAttributeKey: "custom_attribute_key",
        Version: square.Int(
            1,
        ),
        WithDefinition: square.Bool(
            true,
        ),
    }
client.Orders.CustomAttributes.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the target [order](entity:Order).
    
</dd>
</dl>

<dl>
<dd>

**customAttributeKey:** `string` 

The key of the custom attribute to retrieve.  This key must match the key of an
existing custom attribute definition.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*int` 

To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control, include this optional field and specify the current version of the custom attribute.
    
</dd>
</dl>

<dl>
<dd>

**withDefinition:** `*bool` 

Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the `definition` field of each 
custom attribute. Set this parameter to `true` to get the name and description of each custom attribute, 
information about the data type, or other definition details. The default value is `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributes.Upsert(OrderID, CustomAttributeKey, request) -> *square.UpsertOrderCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates a [custom attribute](entity:CustomAttribute) for an order.

Use this endpoint to set the value of a custom attribute for a specific order.
A custom attribute is based on a custom attribute definition in a Square seller account. (To create a
custom attribute definition, use the [CreateOrderCustomAttributeDefinition](api-endpoint:OrderCustomAttributes-CreateOrderCustomAttributeDefinition) endpoint.)

To create or update a custom attribute owned by another application, the `visibility` setting
must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.UpsertOrderCustomAttributeRequest{
        OrderID: "order_id",
        CustomAttributeKey: "custom_attribute_key",
        CustomAttribute: &square.CustomAttribute{
            Key: square.String(
                "table-number",
            ),
            Value: "42",
            Version: square.Int(
                1,
            ),
        },
    }
client.Orders.CustomAttributes.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the target [order](entity:Order).
    
</dd>
</dl>

<dl>
<dd>

**customAttributeKey:** `string` 

The key of the custom attribute to create or update.  This key must match the key 
of an existing custom attribute definition.
    
</dd>
</dl>

<dl>
<dd>

**customAttribute:** `*square.CustomAttribute` 

The custom attribute to create or update, with the following fields:

- `value`. This value must conform to the `schema` specified by the definition. 
For more information, see [Value data types](https://developer.squareup.com/docs/customer-custom-attributes-api/custom-attributes#value-data-types).

- `version`. To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency)
control, include this optional field and specify the current version of the custom attribute.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 

A unique identifier for this request, used to ensure idempotency. 
For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CustomAttributes.Delete(OrderID, CustomAttributeKey) -> *square.DeleteOrderCustomAttributeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a [custom attribute](entity:CustomAttribute) associated with a customer profile.

To delete a custom attribute owned by another application, the `visibility` setting must be
`VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes
(also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &orders.DeleteCustomAttributesRequest{
        OrderID: "order_id",
        CustomAttributeKey: "custom_attribute_key",
    }
client.Orders.CustomAttributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The ID of the target [order](entity:Order).
    
</dd>
</dl>

<dl>
<dd>

**customAttributeKey:** `string` 

The key of the custom attribute to delete.  This key must match the key of an
existing custom attribute definition.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TeamMembers WageSetting
<details><summary><code>client.TeamMembers.WageSetting.Get(TeamMemberID) -> *square.GetWageSettingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a `WageSetting` object for a team member specified
by `TeamMember.id`. For more information, see
[Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrievewagesetting).

Square recommends using [RetrieveTeamMember](api-endpoint:Team-RetrieveTeamMember) or [SearchTeamMembers](api-endpoint:Team-SearchTeamMembers)
to get this information directly from the `TeamMember.wage_setting` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &teammembers.GetWageSettingRequest{
        TeamMemberID: "team_member_id",
    }
client.TeamMembers.WageSetting.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `string` — The ID of the team member for which to retrieve the wage setting.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.WageSetting.Update(TeamMemberID, request) -> *square.UpdateWageSettingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates a `WageSetting` object. The object is created if a
`WageSetting` with the specified `team_member_id` doesn't exist. Otherwise,
it fully replaces the `WageSetting` object for the team member.
The `WageSetting` is returned on a successful update. For more information, see
[Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#create-or-update-a-wage-setting).

Square recommends using [CreateTeamMember](api-endpoint:Team-CreateTeamMember) or [UpdateTeamMember](api-endpoint:Team-UpdateTeamMember)
to manage the `TeamMember.wage_setting` field directly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &teammembers.UpdateWageSettingRequest{
        TeamMemberID: "team_member_id",
        WageSetting: &square.WageSetting{
            JobAssignments: []*square.JobAssignment{
                &square.JobAssignment{
                    JobTitle: square.String(
                        "Manager",
                    ),
                    PayType: square.JobAssignmentPayTypeSalary,
                    AnnualRate: &square.Money{
                        Amount: square.Int64(
                            3000000,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                    WeeklyHours: square.Int(
                        40,
                    ),
                },
                &square.JobAssignment{
                    JobTitle: square.String(
                        "Cashier",
                    ),
                    PayType: square.JobAssignmentPayTypeHourly,
                    HourlyRate: &square.Money{
                        Amount: square.Int64(
                            2000,
                        ),
                        Currency: square.CurrencyUsd.Ptr(),
                    },
                },
            },
            IsOvertimeExempt: square.Bool(
                true,
            ),
        },
    }
client.TeamMembers.WageSetting.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamMemberID:** `string` — The ID of the team member for which to update the `WageSetting` object.
    
</dd>
</dl>

<dl>
<dd>

**wageSetting:** `*square.WageSetting` 

The complete `WageSetting` object. For all job assignments, specify one of the following:
- `job_id` (recommended) - If needed, call [ListJobs](api-endpoint:Team-ListJobs) to get a list of all jobs.
Requires Square API version 2024-12-18 or later.
- `job_title` - Use the exact, case-sensitive spelling of an existing title unless you want to create a new job.
This value is ignored if `job_id` is also provided.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Terminal Actions
<details><summary><code>client.Terminal.Actions.Create(request) -> *square.CreateTerminalActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a Terminal action request and sends it to the specified device.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CreateTerminalActionRequest{
        IdempotencyKey: "thahn-70e75c10-47f7-4ab6-88cc-aaa4076d065e",
        Action: &square.TerminalAction{
            DeviceID: square.String(
                "{{DEVICE_ID}}",
            ),
            DeadlineDuration: square.String(
                "PT5M",
            ),
            Type: square.TerminalActionActionTypeSaveCard.Ptr(),
            SaveCardOptions: &square.SaveCardOptions{
                CustomerID: "{{CUSTOMER_ID}}",
                ReferenceID: square.String(
                    "user-id-1",
                ),
            },
        },
    }
client.Terminal.Actions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreateAction` request. Keys can be any valid string
but must be unique for every `CreateAction` request.

See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more
information.
    
</dd>
</dl>

<dl>
<dd>

**action:** `*square.TerminalAction` — The Action to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Actions.Search(request) -> *square.SearchTerminalActionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a filtered list of Terminal action requests created by the account making the request. Terminal action requests are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.SearchTerminalActionsRequest{
        Query: &square.TerminalActionQuery{
            Filter: &square.TerminalActionQueryFilter{
                CreatedAt: &square.TimeRange{
                    StartAt: square.String(
                        "2022-04-01T00:00:00.000Z",
                    ),
                },
            },
            Sort: &square.TerminalActionQuerySort{
                SortOrder: square.SortOrderDesc.Ptr(),
            },
        },
        Limit: square.Int(
            2,
        ),
    }
client.Terminal.Actions.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.TerminalActionQuery` 

Queries terminal actions based on given conditions and sort order.
Leaving this unset will return all actions with the default sort order.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for the original query.
See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more
information.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the number of results returned for a single request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Actions.Get(ActionID) -> *square.GetTerminalActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a Terminal action request by `action_id`. Terminal action requests are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.GetActionsRequest{
        ActionID: "action_id",
    }
client.Terminal.Actions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — Unique ID for the desired `TerminalAction`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Actions.Cancel(ActionID) -> *square.CancelTerminalActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a Terminal action request if the status of the request permits it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CancelActionsRequest{
        ActionID: "action_id",
    }
client.Terminal.Actions.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — Unique ID for the desired `TerminalAction`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Terminal Checkouts
<details><summary><code>client.Terminal.Checkouts.Create(request) -> *square.CreateTerminalCheckoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a Terminal checkout request and sends it to the specified device to take a payment
for the requested amount.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CreateTerminalCheckoutRequest{
        IdempotencyKey: "28a0c3bc-7839-11ea-bc55-0242ac130003",
        Checkout: &square.TerminalCheckout{
            AmountMoney: &square.Money{
                Amount: square.Int64(
                    2610,
                ),
                Currency: square.CurrencyUsd.Ptr(),
            },
            ReferenceID: square.String(
                "id11572",
            ),
            Note: square.String(
                "A brief note",
            ),
            DeviceOptions: &square.DeviceCheckoutOptions{
                DeviceID: "dbb5d83a-7838-11ea-bc55-0242ac130003",
            },
        },
    }
client.Terminal.Checkouts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreateCheckout` request. Keys can be any valid string but
must be unique for every `CreateCheckout` request.

See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**checkout:** `*square.TerminalCheckout` — The checkout to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Checkouts.Search(request) -> *square.SearchTerminalCheckoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a filtered list of Terminal checkout requests created by the application making the request. Only Terminal checkout requests created for the merchant scoped to the OAuth token are returned. Terminal checkout requests are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.SearchTerminalCheckoutsRequest{
        Query: &square.TerminalCheckoutQuery{
            Filter: &square.TerminalCheckoutQueryFilter{
                Status: square.String(
                    "COMPLETED",
                ),
            },
        },
        Limit: square.Int(
            2,
        ),
    }
client.Terminal.Checkouts.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.TerminalCheckoutQuery` 

Queries Terminal checkouts based on given conditions and the sort order.
Leaving these unset returns all checkouts with the default sort order.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limits the number of results returned for a single request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Checkouts.Get(CheckoutID) -> *square.GetTerminalCheckoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a Terminal checkout request by `checkout_id`. Terminal checkout requests are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.GetCheckoutsRequest{
        CheckoutID: "checkout_id",
    }
client.Terminal.Checkouts.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**checkoutID:** `string` — The unique ID for the desired `TerminalCheckout`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Checkouts.Cancel(CheckoutID) -> *square.CancelTerminalCheckoutResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a Terminal checkout request if the status of the request permits it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CancelCheckoutsRequest{
        CheckoutID: "checkout_id",
    }
client.Terminal.Checkouts.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**checkoutID:** `string` — The unique ID for the desired `TerminalCheckout`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Terminal Refunds
<details><summary><code>client.Terminal.Refunds.Create(request) -> *square.CreateTerminalRefundResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a request to refund an Interac payment completed on a Square Terminal. Refunds for Interac payments on a Square Terminal are supported only for Interac debit cards in Canada. Other refunds for Terminal payments should use the Refunds API. For more information, see [Refunds API](api:Refunds).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CreateTerminalRefundRequest{
        IdempotencyKey: "402a640b-b26f-401f-b406-46f839590c04",
        Refund: &square.TerminalRefund{
            PaymentID: "5O5OvgkcNUhl7JBuINflcjKqUzXZY",
            AmountMoney: &square.Money{
                Amount: square.Int64(
                    111,
                ),
                Currency: square.CurrencyCad.Ptr(),
            },
            Reason: "Returning items",
            DeviceID: "f72dfb8e-4d65-4e56-aade-ec3fb8d33291",
        },
    }
client.Terminal.Refunds.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` 

A unique string that identifies this `CreateRefund` request. Keys can be any valid string but
must be unique for every `CreateRefund` request.

See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
    
</dd>
</dl>

<dl>
<dd>

**refund:** `*square.TerminalRefund` — The refund to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Refunds.Search(request) -> *square.SearchTerminalRefundsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a filtered list of Interac Terminal refund requests created by the seller making the request. Terminal refund requests are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.SearchTerminalRefundsRequest{
        Query: &square.TerminalRefundQuery{
            Filter: &square.TerminalRefundQueryFilter{
                Status: square.String(
                    "COMPLETED",
                ),
            },
        },
        Limit: square.Int(
            1,
        ),
    }
client.Terminal.Refunds.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*square.TerminalRefundQuery` 

Queries the Terminal refunds based on given conditions and the sort order. Calling
`SearchTerminalRefunds` without an explicit query parameter returns all available
refunds with the default sort order.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this cursor to retrieve the next set of results for the original query.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limits the number of results returned for a single request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Refunds.Get(TerminalRefundID) -> *square.GetTerminalRefundResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an Interac Terminal refund object by ID. Terminal refund objects are available for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.GetRefundsRequest{
        TerminalRefundID: "terminal_refund_id",
    }
client.Terminal.Refunds.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**terminalRefundID:** `string` — The unique ID for the desired `TerminalRefund`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Terminal.Refunds.Cancel(TerminalRefundID) -> *square.CancelTerminalRefundResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels an Interac Terminal refund request by refund request ID if the status of the request permits it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &terminal.CancelRefundsRequest{
        TerminalRefundID: "terminal_refund_id",
    }
client.Terminal.Refunds.Cancel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**terminalRefundID:** `string` — The unique ID for the desired `TerminalRefund`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks EventTypes
<details><summary><code>client.Webhooks.EventTypes.List() -> *square.ListWebhookEventTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all webhook event types that can be subscribed to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.ListEventTypesRequest{
        APIVersion: square.String(
            "api_version",
        ),
    }
client.Webhooks.EventTypes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiVersion:** `*string` — The API version for which to list event types. Setting this field overrides the default version used by the application.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks Subscriptions
<details><summary><code>client.Webhooks.Subscriptions.List() -> *square.ListWebhookSubscriptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all webhook subscriptions owned by your application.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.ListSubscriptionsRequest{
        Cursor: square.String(
            "cursor",
        ),
        IncludeDisabled: square.Bool(
            true,
        ),
        SortOrder: square.SortOrderDesc.Ptr(),
        Limit: square.Int(
            1,
        ),
    }
client.Webhooks.Subscriptions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 

A pagination cursor returned by a previous call to this endpoint.
Provide this to retrieve the next set of results for your original query.

For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
    
</dd>
</dl>

<dl>
<dd>

**includeDisabled:** `*bool` 

Includes disabled [Subscription](entity:WebhookSubscription)s.
By default, all enabled [Subscription](entity:WebhookSubscription)s are returned.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*square.SortOrder` 

Sorts the returned list by when the [Subscription](entity:WebhookSubscription) was created with the specified order.
This field defaults to ASC.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 

The maximum number of results to be returned in a single page.
It is possible to receive fewer results than the specified limit on a given page.
The default value of 100 is also the maximum allowed value.

Default: 100
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.Create(request) -> *square.CreateWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a webhook subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.CreateWebhookSubscriptionRequest{
        IdempotencyKey: square.String(
            "63f84c6c-2200-4c99-846c-2670a1311fbf",
        ),
        Subscription: &square.WebhookSubscription{
            Name: square.String(
                "Example Webhook Subscription",
            ),
            EventTypes: []string{
                "payment.created",
                "payment.updated",
            },
            NotificationURL: square.String(
                "https://example-webhook-url.com",
            ),
            APIVersion: square.String(
                "2021-12-15",
            ),
        },
    }
client.Webhooks.Subscriptions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique string that identifies the [CreateWebhookSubscription](api-endpoint:WebhookSubscriptions-CreateWebhookSubscription) request.
    
</dd>
</dl>

<dl>
<dd>

**subscription:** `*square.WebhookSubscription` — The [Subscription](entity:WebhookSubscription) to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.Get(SubscriptionID) -> *square.GetWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a webhook subscription identified by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.GetSubscriptionsRequest{
        SubscriptionID: "subscription_id",
    }
client.Webhooks.Subscriptions.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.Update(SubscriptionID, request) -> *square.UpdateWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a webhook subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.UpdateWebhookSubscriptionRequest{
        SubscriptionID: "subscription_id",
        Subscription: &square.WebhookSubscription{
            Name: square.String(
                "Updated Example Webhook Subscription",
            ),
            Enabled: square.Bool(
                false,
            ),
        },
    }
client.Webhooks.Subscriptions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
    
</dd>
</dl>

<dl>
<dd>

**subscription:** `*square.WebhookSubscription` — The [Subscription](entity:WebhookSubscription) to update.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.Delete(SubscriptionID) -> *square.DeleteWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a webhook subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.DeleteSubscriptionsRequest{
        SubscriptionID: "subscription_id",
    }
client.Webhooks.Subscriptions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.UpdateSignatureKey(SubscriptionID, request) -> *square.UpdateWebhookSubscriptionSignatureKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a webhook subscription by replacing the existing signature key with a new one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.UpdateWebhookSubscriptionSignatureKeyRequest{
        SubscriptionID: "subscription_id",
        IdempotencyKey: square.String(
            "ed80ae6b-0654-473b-bbab-a39aee89a60d",
        ),
    }
client.Webhooks.Subscriptions.UpdateSignatureKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — A unique string that identifies the [UpdateWebhookSubscriptionSignatureKey](api-endpoint:WebhookSubscriptions-UpdateWebhookSubscriptionSignatureKey) request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Subscriptions.Test(SubscriptionID, request) -> *square.TestWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Tests a webhook subscription by sending a test event to the notification URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &webhooks.TestWebhookSubscriptionRequest{
        SubscriptionID: "subscription_id",
        EventType: square.String(
            "payment.created",
        ),
    }
client.Webhooks.Subscriptions.Test(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to test.
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` 

The event type that will be used to test the [Subscription](entity:WebhookSubscription). The event type must be
contained in the list of event types in the [Subscription](entity:WebhookSubscription).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
