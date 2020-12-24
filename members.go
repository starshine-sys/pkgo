package pkgo

import "strings"

// Members holds multiple members, which can be sorted by creation date
type Members []Member

func (m Members) Len() int           { return len(m) }
func (m Members) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Members) Less(i, j int) bool { return m[i].Created.Before(m[j].Created) }

// ToMaps returns two maps, one of name-ID, one of ID-name.
// Note that the name-ID map may be incomplete if multiple members have the same name
func (m Members) ToMaps() (map[string]string, map[string]string) {
	nameMap := make(map[string]string)
	idMap := make(map[string]string)
	for _, member := range m {
		nameMap[strings.ToLower(member.Name)] = member.ID
		idMap[member.ID] = member.Name
	}
	return nameMap, idMap
}

// GetMembers gets all members of a system
func (s *Session) GetMembers(id string) (Members, error) {
	if id == "" && (!s.authorized || s.token == "") {
		return nil, ErrNoToken
	}

	var m Members
	if id == "" {
		id = s.system
	}
	err := s.getEndpoint("/s/"+id+"/members", &m)
	return m, err
}
