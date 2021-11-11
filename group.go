package pkgo

import (
	"net/url"

	"github.com/google/uuid"
)

// Group is a member group.
type Group struct {
	ID          string    `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name,omitempty"`
	Description string    `json:"description,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	Color       Color     `json:"color"`

	Privacy *GroupPrivacy `json:"privacy,omitempty"`

	// Only returned in GroupsWithMembers
	Members []uuid.UUID `json:"members,omitempty"`
}

// GroupPrivacy is a group privacy object.
type GroupPrivacy struct {
	DescriptionPrivacy Privacy `json:"description_privacy,omitempty"`
	IconPrivacy        Privacy `json:"icon_privacy,omitempty"`
	ListPrivacy        Privacy `json:"list_privacy,omitempty"`
	Visibility         Privacy `json:"visibility,omitempty"`
}

// Groups gets a system's groups.
func (s *Session) Groups(systemID string) (g []Group, err error) {
	err = s.RequestJSON("GET", "/systems/"+systemID+"/groups", &g)
	return
}

// GroupsWithMembers gets a system's groups, with the Members field filled.
func (s *Session) GroupsWithMembers(systemID string) (g []Group, err error) {
	err = s.RequestJSON("GET", "/systems/"+systemID+"/groups", &g, WithURLValues(url.Values{
		"with_members": {"true"},
	}))
	return
}

// CreateGroupData is the data for s.CreateGroup.
type CreateGroupData struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Color       Color  `json:"color"`

	Privacy *GroupPrivacy `json:"privacy,omitempty"`
}

// CreateGroup creates a group.
func (s *Session) CreateGroup(data CreateGroupData) (g Group, err error) {
	err = s.RequestJSON("POST", "/groups", &g, WithJSONBody(data))
	return
}

// Group returns a group by its ID or UUID.
func (s *Session) Group(id string) (g Group, err error) {
	err = s.RequestJSON("GET", "/groups/"+id, &g)
	return
}

// GroupByUUID returns a group by its UUID.
func (s *Session) GroupByUUID(id uuid.UUID) (g Group, err error) {
	return s.Group(id.String())
}

// MemberGroups returns the groups the given member is in.
func (s *Session) MemberGroups(id string) (g []Group, err error) {
	err = s.RequestJSON("GET", "/members/"+id+"/groups", &g)
	return
}

// EditGroupData is the data for s.EditGroup.
type EditGroupData struct {
	Name        NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"display_name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Icon        NullableString `json:"icon,omitempty"`
	Banner      NullableString `json:"banner,omitempty"`
	Color       *Color         `json:"color"`

	Privacy *GroupPrivacy `json:"privacy,omitempty"`
}

// EditGroup edits the given group.
func (s *Session) EditGroup(id string, dat EditGroupData) (g Group, err error) {
	err = s.RequestJSON("PATCH", "/groups/"+id, &g, WithJSONBody(dat))
	return g, err
}

// DeleteGroup deletes a group. Requires authentication.
func (s *Session) DeleteGroup(id string) (err error) {
	_, err = s.Request("DELETE", "/groups/"+id)
	return
}

// AddGroupMembers adds the given member IDs to the given group.
func (s *Session) AddGroupMembers(groupID string, memberIDs []string) (err error) {
	_, err = s.Request("POST", "/groups/"+groupID+"/members/add", WithJSONBody(memberIDs))
	return
}

// RemoveGroupMembers removes the given member IDs from the given group.
func (s *Session) RemoveGroupMembers(groupID string, memberIDs []string) (err error) {
	_, err = s.Request("POST", "/groups/"+groupID+"/members/remove", WithJSONBody(memberIDs))
	return
}

// OverwriteGroupMembers overwrites a group's members with the given member IDs.
func (s *Session) OverwriteGroupMembers(groupID string, memberIDs []string) (err error) {
	_, err = s.Request("POST", "/groups/"+groupID+"/members/overwrite", WithJSONBody(memberIDs))
	return
}
