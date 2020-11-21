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
	if !s.Authorized || s.Token == "" {
		return nil, &ErrNoToken{}
	}

	err = s.GetEndpoint("/s", &sys)
	if err != nil {
		return
	}
	if s.system == "" {
		s.system = sys.ID
	}
	return
}
