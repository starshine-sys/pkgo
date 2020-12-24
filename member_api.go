package pkgo

import "encoding/json"

// UpdateMember updates a member by ID
func (s *Session) UpdateMember(id string, member *Member) (*Member, error) {
	b, err := json.Marshal(member)
	if err != nil {
		return member, err
	}

	m := &Member{}
	err = s.patchEndpoint("/m/"+id, b, m)
	return m, err
}
