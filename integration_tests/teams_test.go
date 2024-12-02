//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/square/square-go-sdk"
)

// Teams API integration tests.
func TestTeamsAPI(t *testing.T) {
	t.Run("bulk update team members with mix of successes and failures", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		locationId, err := getDefaultLocationID(squareClient)
		require.NoError(t, err)

		// SETUP: Create 3 team members (should always be successful).
		createMembersResp, err := squareClient.TeamMembers.BatchCreate(
			context.Background(),
			&square.BatchCreateTeamMembersRequest{
				TeamMembers: map[string]*square.CreateTeamMemberRequest{
					newTestUUID(): {
						TeamMember: newTestTeamMember([]string{locationId}),
					},
					newTestUUID(): {
						TeamMember: newTestTeamMember([]string{locationId}),
					},
					newTestUUID(): {
						TeamMember: newTestTeamMember([]string{locationId}),
					},
				},
			},
		)
		require.NoError(t, err)

		createdMemberIds := make([]string, 0, len(createMembersResp.TeamMembers))
		for k := range createMembersResp.TeamMembers {
			memberId := createMembersResp.TeamMembers[k].TeamMember.ID
			createdMemberIds = append(createdMemberIds, *memberId)

			require.Nil(t, createMembersResp.TeamMembers[k].Errors)
		}

		// Update 3 team members in a bulk request, with 2 successful updates and 1
		// invalid update (location ID is invalid). This should result in a 200
		// response, with 2 nested success responses and 1 nested error response.
		updateTeamMembersResp, err := squareClient.TeamMembers.BatchUpdate(
			context.Background(),
			&square.BatchUpdateTeamMembersRequest{
				TeamMembers: map[string]*square.UpdateTeamMemberRequest{
					createdMemberIds[0]: {
						TeamMember: newTestTeamMember([]string{locationId}),
					},
					createdMemberIds[1]: {
						TeamMember: newTestTeamMember([]string{locationId}),
					},
					createdMemberIds[2]: {
						TeamMember: newTestTeamMember([]string{"INVALID_LocationID"}),
					},
				},
			},
		)
		require.NoError(t, err)

		teamMembers := updateTeamMembersResp.TeamMembers
		require.Len(t, teamMembers, 3)

		member1Errors := teamMembers[createdMemberIds[0]].Errors
		member2Errors := teamMembers[createdMemberIds[1]].Errors
		member3Errors := teamMembers[createdMemberIds[2]].Errors
		assert.Empty(t, member1Errors)
		assert.Empty(t, member2Errors)
		assert.NotEmpty(t, member3Errors)
		assert.Equal(t, square.ErrorCodeInvalidValue, member3Errors[0].Code)
		assert.Equal(t, "Expected the assigned_locations.location_ids to be valid", *member3Errors[0].Detail)

	})
}

// newTestTeamMember creates a [square.TeamMember] with the given location IDs.
func newTestTeamMember(
	locationIDs []string,
) *square.TeamMember {
	return &square.TeamMember{
		AssignedLocations: &square.TeamMemberAssignedLocations{
			AssignmentType: square.TeamMemberAssignedLocationsAssignmentTypeExplicitLocations.Ptr(),
			LocationIDs:    locationIDs,
		},
		FamilyName: square.String("Doe"),
		GivenName:  square.String("Jane"),
	}
}
