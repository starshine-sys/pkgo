package pkgo

import "time"

// Member holds information for a specific system member.
// Fields set to private are empty.
type Member struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`

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

// Validate will validate the member object.
// If any of the fields have invalid values, it returns an InvalidError.
func (m Member) Validate() error {
	if len(m.Name) > 50 || len(m.DisplayName) > 50 ||
		len(m.Description) > 1000 || len(m.Pronouns) > 100 {
		return &InvalidError{"Name", m.Name}
	}

	if len(m.DisplayName) > 50 {
		return &InvalidError{"DisplayName", m.DisplayName}
	}

	if len(m.Description) > 1000 {
		return &InvalidError{"Description", m.Description}
	}

	if len(m.Pronouns) > 100 {
		return &InvalidError{"Pronouns", m.Pronouns}
	}

	// Privacy fields can only be "public", "private", or null
	for _, f := range []Privacy{m.Visibility, m.NamePrivacy, m.DescriptionPrivacy, m.AvatarPrivacy, m.BirthdayPrivacy, m.PronounPrivacy, m.MetadataPrivacy} {
		if f != "private" && f != "public" && f != "" {
			return &InvalidError{"Privacy", string(f)}
		}
	}

	if !m.Color.IsValid() {
		return &InvalidError{"Color", string(m.Color)}
	}

	return nil
}

// String returns the member's displayed name--either DisplayName if one is set, otherwise Name.
func (m Member) String() string {
	if m.DisplayName != "" {
		return m.DisplayName
	}
	return m.Name
}
