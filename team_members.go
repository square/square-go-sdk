// This file was auto-generated by Fern from our API Definition.

package square

type BatchCreateTeamMembersRequest struct {
	// The data used to create the `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`. The maximum number of create objects is 25.
	TeamMembers map[string]*CreateTeamMemberRequest `json:"team_members,omitempty" url:"-"`
}

type BatchUpdateTeamMembersRequest struct {
	// The data used to update the `TeamMember` objects. Each key is the `team_member_id` that maps to the `UpdateTeamMemberRequest`. The maximum number of update objects is 25.
	TeamMembers map[string]*UpdateTeamMemberRequest `json:"team_members,omitempty" url:"-"`
}

type SearchTeamMembersRequest struct {
	// The query parameters.
	Query *SearchTeamMembersQuery `json:"query,omitempty" url:"-"`
	// The maximum number of `TeamMember` objects in a page (100 by default).
	Limit *int `json:"limit,omitempty" url:"-"`
	// The opaque cursor for fetching the next page. For more information, see
	// [pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
}
