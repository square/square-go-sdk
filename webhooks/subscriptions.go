// This file was auto-generated by Fern from our API Definition.

package webhooks

import (
	v40 "github.com/square/square-go-sdk/v40"
)

type CreateWebhookSubscriptionRequest struct {
	// A unique string that identifies the [CreateWebhookSubscription](api-endpoint:WebhookSubscriptions-CreateWebhookSubscription) request.
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The [Subscription](entity:WebhookSubscription) to create.
	Subscription *v40.WebhookSubscription `json:"subscription,omitempty" url:"-"`
}

type SubscriptionsDeleteRequest struct {
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to delete.
	SubscriptionID string `json:"-" url:"-"`
}

type SubscriptionsGetRequest struct {
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to retrieve.
	SubscriptionID string `json:"-" url:"-"`
}

type SubscriptionsListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Includes disabled [Subscription](entity:WebhookSubscription)s.
	// By default, all enabled [Subscription](entity:WebhookSubscription)s are returned.
	IncludeDisabled *bool `json:"-" url:"include_disabled,omitempty"`
	// Sorts the returned list by when the [Subscription](entity:WebhookSubscription) was created with the specified order.
	// This field defaults to ASC.
	SortOrder *v40.SortOrder `json:"-" url:"sort_order,omitempty"`
	// The maximum number of results to be returned in a single page.
	// It is possible to receive fewer results than the specified limit on a given page.
	// The default value of 100 is also the maximum allowed value.
	//
	// Default: 100
	Limit *int `json:"-" url:"limit,omitempty"`
}

type TestWebhookSubscriptionRequest struct {
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to test.
	SubscriptionID string `json:"-" url:"-"`
	// The event type that will be used to test the [Subscription](entity:WebhookSubscription). The event type must be
	// contained in the list of event types in the [Subscription](entity:WebhookSubscription).
	EventType *string `json:"event_type,omitempty" url:"-"`
}

type UpdateWebhookSubscriptionRequest struct {
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
	SubscriptionID string `json:"-" url:"-"`
	// The [Subscription](entity:WebhookSubscription) to update.
	Subscription *v40.WebhookSubscription `json:"subscription,omitempty" url:"-"`
}

type UpdateWebhookSubscriptionSignatureKeyRequest struct {
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
	SubscriptionID string `json:"-" url:"-"`
	// A unique string that identifies the [UpdateWebhookSubscriptionSignatureKey](api-endpoint:WebhookSubscriptions-UpdateWebhookSubscriptionSignatureKey) request.
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
}
