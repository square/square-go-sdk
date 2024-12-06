// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/internal"
)

type CreateCardRequest struct {
	// A unique string that identifies this CreateCard request. Keys can be
	// any valid string and must be unique for every request.
	//
	// Max: 45 characters
	//
	// See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The ID of the source which represents the card information to be stored. This can be a card nonce or a payment id.
	SourceID string `json:"source_id" url:"-"`
	// An identifying token generated by [Payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer).
	// Verification tokens encapsulate customer device information and 3-D Secure
	// challenge results to indicate that Square has verified the buyer identity.
	//
	// See the [SCA Overview](https://developer.squareup.com/docs/sca-overview).
	VerificationToken *string `json:"verification_token,omitempty" url:"-"`
	// Payment details associated with the card to be stored.
	Card *Card `json:"card,omitempty" url:"-"`
}

type CardsDisableRequest struct {
	// Unique ID for the desired Card.
	CardID string `json:"-" url:"-"`
}

type CardsGetRequest struct {
	// Unique ID for the desired Card.
	CardID string `json:"-" url:"-"`
}

type CardsListRequest struct {
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for your original query.
	//
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Limit results to cards associated with the customer supplied.
	// By default, all cards owned by the merchant are returned.
	CustomerID *string `json:"-" url:"customer_id,omitempty"`
	// Includes disabled cards.
	// By default, all enabled cards owned by the merchant are returned.
	IncludeDisabled *bool `json:"-" url:"include_disabled,omitempty"`
	// Limit results to cards associated with the reference_id supplied.
	ReferenceID *string `json:"-" url:"reference_id,omitempty"`
	// Sorts the returned list by when the card was created with the specified order.
	// This field defaults to ASC.
	SortOrder *SortOrder `json:"-" url:"sort_order,omitempty"`
}

// Defines the fields that are included in the response body of
// a request to the [CreateCard](api-endpoint:Cards-CreateCard) endpoint.
//
// Note: if there are errors processing the request, the card field will not be
// present.
type CreateCardResponse struct {
	// Errors resulting from the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The card created by the request.
	Card *Card `json:"card,omitempty" url:"card,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateCardResponse) GetErrors() []*Error {
	if c == nil {
		return nil
	}
	return c.Errors
}

func (c *CreateCardResponse) GetCard() *Card {
	if c == nil {
		return nil
	}
	return c.Card
}

func (c *CreateCardResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateCardResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCardResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateCardResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateCardResponse) String() string {
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
// a request to the [DisableCard](api-endpoint:Cards-DisableCard) endpoint.
//
// Note: if there are errors processing the request, the card field will not be
// present.
type DisableCardResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The retrieved card.
	Card *Card `json:"card,omitempty" url:"card,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DisableCardResponse) GetErrors() []*Error {
	if d == nil {
		return nil
	}
	return d.Errors
}

func (d *DisableCardResponse) GetCard() *Card {
	if d == nil {
		return nil
	}
	return d.Card
}

func (d *DisableCardResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DisableCardResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DisableCardResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DisableCardResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DisableCardResponse) String() string {
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

// Defines the fields that are included in the response body of
// a request to the [RetrieveCard](api-endpoint:Cards-RetrieveCard) endpoint.
//
// Note: if there are errors processing the request, the card field will not be
// present.
type GetCardResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The retrieved card.
	Card *Card `json:"card,omitempty" url:"card,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetCardResponse) GetErrors() []*Error {
	if g == nil {
		return nil
	}
	return g.Errors
}

func (g *GetCardResponse) GetCard() *Card {
	if g == nil {
		return nil
	}
	return g.Card
}

func (g *GetCardResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetCardResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCardResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCardResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCardResponse) String() string {
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

// Defines the fields that are included in the response body of
// a request to the [ListCards](api-endpoint:Cards-ListCards) endpoint.
//
// Note: if there are errors processing the request, the card field will not be
// present.
type ListCardsResponse struct {
	// Information on errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The requested list of `Card`s.
	Cards []*Card `json:"cards,omitempty" url:"cards,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty,
	// this is the final response.
	//
	// See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListCardsResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListCardsResponse) GetCards() []*Card {
	if l == nil {
		return nil
	}
	return l.Cards
}

func (l *ListCardsResponse) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListCardsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListCardsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListCardsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListCardsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListCardsResponse) String() string {
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
