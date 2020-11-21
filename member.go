package pkgo

import "time"

// Member holds information for a specific system member
type Member struct {
	ID              string     `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	DisplayName     string     `json:"display_name,omitempty"`
	Description     string     `json:"description,omitempty"`
	Pronouns        string     `json:"pronouns,omitempty"`
	Color           Color      `json:"color,omitempty"`
	AvatarURL       string     `json:"avatar_url,omitempty"`
	Birthday        string     `json:"birthday,omitempty"`
	ProxyTags       []ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy       bool       `json:"keep_proxy,omitempty"`
	Created         time.Time  `json:"created"`
	Visibility      string     `json:"visibility,omitempty"`
	NamePrivacy     string     `json:"name_privacy,omitempty"`
	DescPrivacy     string     `json:"description_privacy,omitempty"`
	AvatarPrivacy   string     `json:"avatar_privacy,omitempty"`
	BirthdayPrivacy string     `json:"birthday_privacy,omitempty"`
	PronounPrivacy  string     `json:"pronoun_privacy,omitempty"`
	MetadataPrivacy string     `json:"metadata_privacy,omitempty"`
}

// Birthday format constants
const (
	BirthdaySource       = "2006-01-02"
	BirthdayMMDD         = "01-02"
	BirthdayDDMM         = "02-01"
	BirthdayMonthName    = "January 2"
	BirthdayMonthNameRev = "2 January"
)

// BirthdayString gives a member's birthday according to a format string
func (m *Member) BirthdayString(format string) string {
	if m.Birthday == "" {
		return ""
	}
	t, err := time.Parse(BirthdaySource, m.Birthday)
	if err != nil {
		return ""
	}
	return t.Format(format)
}

// SetDisplayName sets the displayname, honouring limits
func (m *Member) SetDisplayName(n string) *Member {
	if len(n) > memberNameLimit {
		n = n[:memberNameLimit]
	}
	m.DisplayName = n
	return m
}

// SetName sets the member's name, honouring limits
func (m *Member) SetName(n string) *Member {
	if len(n) > memberNameLimit {
		n = n[:memberNameLimit]
	}
	m.Name = n
	return m
}

// SetDescription sets the member's description, honouring limits
func (m *Member) SetDescription(desc string) *Member {
	if len(desc) > descLimit {
		desc = desc[:descLimit]
	}
	m.Description = desc
	return m
}

// SetPronouns sets the member's pronouns, honouring limits
func (m *Member) SetPronouns(p string) *Member {
	if len(p) > pronounLimit {
		p = p[:pronounLimit]
	}
	m.Pronouns = p
	return m
}

// Privacy is a privacy field
type Privacy int

// Privacy field constants
const (
	Visibility Privacy = iota
	NamePrivacy
	DescPrivacy
	AvatarPrivacy
	BirthdayPrivacy
	PronounPrivacy
	MetadataPrivacy
)

// Private will set the specified field(s) to private
func (m *Member) Private(fields ...Privacy) *Member {
	for _, f := range fields {
		switch f {
		case Visibility:
			m.Visibility = "private"
		case NamePrivacy:
			m.NamePrivacy = "private"
		case DescPrivacy:
			m.DescPrivacy = "private"
		case AvatarPrivacy:
			m.AvatarPrivacy = "private"
		case BirthdayPrivacy:
			m.BirthdayPrivacy = "private"
		case PronounPrivacy:
			m.PronounPrivacy = "private"
		case MetadataPrivacy:
			m.MetadataPrivacy = "private"
		default:
			continue
		}
	}
	return m
}

// Public will set the specified field(s) to public
func (m *Member) Public(fields ...Privacy) *Member {
	for _, f := range fields {
		switch f {
		case Visibility:
			m.Visibility = "public"
		case NamePrivacy:
			m.NamePrivacy = "public"
		case DescPrivacy:
			m.DescPrivacy = "public"
		case AvatarPrivacy:
			m.AvatarPrivacy = "public"
		case BirthdayPrivacy:
			m.BirthdayPrivacy = "public"
		case PronounPrivacy:
			m.PronounPrivacy = "public"
		case MetadataPrivacy:
			m.MetadataPrivacy = "public"
		default:
			continue
		}
	}
	return m
}
