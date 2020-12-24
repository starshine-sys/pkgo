package pkgo

import (
	"time"
)

// System holds all the data for a system
type System struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name,omitempty"`
	Description         string    `json:"description,omitempty"`
	Tag                 string    `json:"tag,omitempty"`
	AvatarURL           string    `json:"avatar_url,omitempty"`
	Timezone            string    `json:"tz,omitempty"`
	Created             time.Time `json:"created"`
	DescPrivacy         string    `json:"description_privacy,omitempty"`
	MemberListPrivacy   string    `json:"member_list_privacy,omitempty"`
	FrontPrivacy        string    `json:"front_privacy,omitempty"`
	FrontHistoryPrivacy string    `json:"front_history_privacy,omitempty"`
}

// GetSystem gets the current token's system
func (s *Session) GetSystem() (sys *System, err error) {
	if !s.authorized || s.token == "" {
		return nil, ErrNoToken
	}

	err = s.getEndpoint("/s", &sys)
	if err != nil {
		return
	}
	if s.system == "" {
		s.system = sys.ID
	}
	return
}

// GetSystemByID gets a system by its 5-character system ID
func (s *Session) GetSystemByID(id string) (sys *System, err error) {
	if !idRe.MatchString(id) {
		return nil, ErrInvalidID
	}
	err = s.getEndpoint("/s/"+id, &sys)
	return
}

// GetSystemByUserID gets a system by a Discord snowflake (user ID)
func (s *Session) GetSystemByUserID(id string) (sys *System, err error) {
	if !discordIDre.MatchString(id) {
		return nil, ErrInvalidSnowflake
	}
	err = s.getEndpoint("/a/"+id, &sys)
	return
}
