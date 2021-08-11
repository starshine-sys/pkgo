package pkgo

import (
	"emperror.dev/errors"
)

// Member gets a member by their ID.
func (s *Session) Member(id string) (m Member, err error) {
	err = s.RequestJSON("GET", "/m/"+id, &m)
	return
}

// Members gets all members of a system.
// If the system's member list is set to private, requires authentication.
// If the request is not authenticated, only public members will be returned.
func (s *Session) Members(id string) ([]Member, error) {
	if id == "" {
		return nil, errors.Sentinel("pkgo: no ID provided")
	}

	m := []Member{}
	err := s.RequestJSON("GET", "/s/"+id+"/members", &m)
	return m, err
}

// EditMemberData is the data for s.EditMember.
type EditMemberData struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`

	Description string   `json:"description,omitempty"`
	Pronouns    string   `json:"pronouns,omitempty"`
	Color       Color    `json:"color,omitempty"`
	Birthday    Birthday `json:"birthday,omitempty"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	Banner      string   `json:"banner,omitempty"`

	ProxyTags []ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy bool       `json:"keep_proxy,omitempty"`

	Visibility         Privacy `json:"visibility,omitempty"`
	NamePrivacy        Privacy `json:"name_privacy,omitempty"`
	DescriptionPrivacy Privacy `json:"description_privacy,omitempty"`
	AvatarPrivacy      Privacy `json:"avatar_privacy,omitempty"`
	BirthdayPrivacy    Privacy `json:"birthday_privacy,omitempty"`
	PronounPrivacy     Privacy `json:"pronoun_privacy,omitempty"`
	MetadataPrivacy    Privacy `json:"metadata_privacy,omitempty"`
}

// EditMember edits a member by ID. Requires authentication.
func (s *Session) EditMember(id string, emd EditMemberData) (*Member, error) {
	m := &Member{}
	err := s.RequestJSON("PATCH", "/m/"+id, m, WithJSONBody(emd))
	if err != nil {
		return nil, err
	}
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

	Visibility         Privacy `json:"visibility,omitempty"`
	NamePrivacy        Privacy `json:"name_privacy,omitempty"`
	DescriptionPrivacy Privacy `json:"description_privacy,omitempty"`
	AvatarPrivacy      Privacy `json:"avatar_privacy,omitempty"`
	BirthdayPrivacy    Privacy `json:"birthday_privacy,omitempty"`
	PronounPrivacy     Privacy `json:"pronoun_privacy,omitempty"`
	MetadataPrivacy    Privacy `json:"metadata_privacy,omitempty"`
}

// CreateMember creates a member. Requires authentication.
func (s *Session) CreateMember(data CreateMemberData) (m Member, err error) {
	err = s.RequestJSON("POST", "/m", &m, WithJSONBody(data))
	return
}

// DeleteMember deletes a member. Requires authentication.
func (s *Session) DeleteMember(id string) (err error) {
	_, err = s.Request("DELETE", "/m/"+id)
	return
}
