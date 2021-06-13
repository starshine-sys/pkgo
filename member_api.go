package pkgo

import (
	"encoding/json"
)

// Member gets a member by their ID
func (s *Session) Member(id string) (m Member, err error) {
	err = s.getEndpoint("/m/"+id, &m)
	return
}

// Members gets all members of a system
func (s *Session) Members(id string) ([]Member, error) {
	if id == "" && (!s.authorized || s.token == "") {
		return nil, ErrNoToken
	}

	var m []Member
	err := s.getEndpoint("/s/"+id+"/members", &m)
	return m, err
}

// EditMemberData ...
type EditMemberData struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`

	Description string   `json:"description,omitempty"`
	Pronouns    string   `json:"pronouns,omitempty"`
	Color       Color    `json:"color,omitempty"`
	Birthday    Birthday `json:"birthday,omitempty"`
	AvatarURL   string   `json:"avatar_url,omitempty"`

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

// EditMember edits a member by ID
func (s *Session) EditMember(id string, emd EditMemberData) (*Member, error) {
	b, err := json.Marshal(emd)
	if err != nil {
		return nil, err
	}

	m := &Member{}
	err = s.patchEndpoint("/m/"+id, b, m)
	if err != nil {
		return nil, err
	}
	return m, err
}

// CreateMemberData ...
type CreateMemberData struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`

	Description string   `json:"description,omitempty"`
	Pronouns    string   `json:"pronouns,omitempty"`
	Color       Color    `json:"color,omitempty"`
	Birthday    Birthday `json:"birthday,omitempty"`
	AvatarURL   string   `json:"avatar_url,omitempty"`

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

// CreateMember creates a member
func (s *Session) CreateMember(data CreateMemberData) (m Member, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = s.postEndpoint("/m", b, &m)
	return
}
