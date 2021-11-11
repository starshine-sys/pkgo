package pkgo

import (
	"emperror.dev/errors"
	"github.com/google/uuid"
)

// Member gets a member by their ID.
func (s *Session) Member(id string) (m Member, err error) {
	err = s.RequestJSON("GET", "/members/"+id, &m)
	return
}

// MemberByUUID gets a member by their UUID.
func (s *Session) MemberByUUID(id uuid.UUID) (Member, error) {
	return s.Member(id.String())
}

// Members gets all members of a system.
// If the system's member list is set to private, requires authentication.
// If the request is not authenticated, only public members will be returned.
func (s *Session) Members(id string) ([]Member, error) {
	if id == "" {
		return nil, errors.Sentinel("pkgo: no ID provided")
	}

	m := []Member{}
	err := s.RequestJSON("GET", "/systems/"+id+"/members", &m)
	return m, err
}

// EditMemberData is the data for s.EditMember.
type EditMemberData struct {
	Name        NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"display_name,omitempty"`

	Description NullableString `json:"description,omitempty"`
	Pronouns    NullableString `json:"pronouns,omitempty"`
	Color       *Color         `json:"color,omitempty"`
	Birthday    *Birthday      `json:"birthday,omitempty"`
	AvatarURL   NullableString `json:"avatar_url,omitempty"`
	Banner      NullableString `json:"banner,omitempty"`

	ProxyTags []ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy bool       `json:"keep_proxy,omitempty"`

	Privacy *MemberPrivacy `json:"privacy,omitempty"`
}

// EditMember edits a member by ID. Requires authentication.
func (s *Session) EditMember(id string, emd EditMemberData) (Member, error) {
	var m Member
	err := s.RequestJSON("PATCH", "/members/"+id, &m, WithJSONBody(emd))
	return m, err
}

// CreateMemberData is the data for s.CreateMember.
type CreateMemberData struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`

	Description string   `json:"description,omitempty"`
	Pronouns    string   `json:"pronouns,omitempty"`
	Color       Color    `json:"color,omitempty"`
	Birthday    Birthday `json:"birthday,omitempty"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	Banner      string   `json:"banner,omitempty"`

	ProxyTags []ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy bool       `json:"keep_proxy"`

	Privacy *MemberPrivacy `json:"privacy,omitempty"`
}

// CreateMember creates a member. Requires authentication.
func (s *Session) CreateMember(data CreateMemberData) (m Member, err error) {
	err = s.RequestJSON("POST", "/members", &m, WithJSONBody(data))
	return
}

// DeleteMember deletes a member. Requires authentication.
func (s *Session) DeleteMember(id string) (err error) {
	_, err = s.Request("DELETE", "/members/"+id)
	return
}
