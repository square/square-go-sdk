// This file was auto-generated by Fern from our API Definition.

package square

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/square/square-go-sdk/v2/internal"
)

// Represents a `ListSites` response. The response can include either `sites` or `errors`.
type ListSitesResponse struct {
	// Any errors that occurred during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The sites that belong to the seller.
	Sites []*Site `json:"sites,omitempty" url:"sites,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSitesResponse) GetErrors() []*Error {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListSitesResponse) GetSites() []*Site {
	if l == nil {
		return nil
	}
	return l.Sites
}

func (l *ListSitesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListSitesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSitesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSitesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSitesResponse) String() string {
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

// Represents a Square Online site, which is an online store for a Square seller.
type Site struct {
	// The Square-assigned ID of the site.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The title of the site.
	SiteTitle *string `json:"site_title,omitempty" url:"site_title,omitempty"`
	// The domain of the site (without the protocol). For example, `mysite1.square.site`.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// Indicates whether the site is published.
	IsPublished *bool `json:"is_published,omitempty" url:"is_published,omitempty"`
	// The timestamp of when the site was created, in RFC 3339 format.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The timestamp of when the site was last updated, in RFC 3339 format.
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *Site) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *Site) GetSiteTitle() *string {
	if s == nil {
		return nil
	}
	return s.SiteTitle
}

func (s *Site) GetDomain() *string {
	if s == nil {
		return nil
	}
	return s.Domain
}

func (s *Site) GetIsPublished() *bool {
	if s == nil {
		return nil
	}
	return s.IsPublished
}

func (s *Site) GetCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *Site) GetUpdatedAt() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *Site) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Site) UnmarshalJSON(data []byte) error {
	type unmarshaler Site
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Site(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *Site) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
