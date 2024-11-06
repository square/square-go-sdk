package integration_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fern-demo/square-go-sdk"
	"github.com/google/uuid"
)

func TestLocationLifecycle(t *testing.T) {
	squareClient := initSquareClient()

	locationID := ""
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("failed to generate random uuid: %v", err)
	}

	// find the location in a list
	listLocationsResp, err := squareClient.Locations.List(context.TODO())
	if err != nil {
		t.Fatalf("failed to list locations: %v", err)
	}

	// iterate through until we find it
	foundLocation := false
	for _, location := range listLocationsResp.Locations {
		if *location.Name == LOCATION_NAME {
			foundLocation = true
			locationID = *location.ID
			break
		}
	}

	if !foundLocation {
		t.Fatalf("failed to find location in list. resp: %v", listLocationsResp.String())
	}

	// retrieve the location
	getLocationResp, err := squareClient.Locations.Get(context.TODO(), locationID)
	if err != nil {
		t.Fatalf("failed to get location: %v", err)
	}

	// validate name matches what we expect
	if *getLocationResp.Location.Name != LOCATION_NAME {
		t.Fatalf("expected location name %s, got %v", LOCATION_NAME, getLocationResp.Location.Name)
	}

	// update the location description
	updateLocationResp, err := squareClient.Locations.Update(context.TODO(), locationID, &square.UpdateLocationRequest{
		Location: &square.Location{
			ID:          &locationID,
			Name:        square.String(LOCATION_NAME),
			Description: square.String(fmt.Sprintf("updated-location-description-%s", randomUUID.String())),
		},
	})
	if err != nil {
		t.Fatalf("failed to update location: %v", err)
	}

	// validate the description was updated
	if *updateLocationResp.Location.Description == *getLocationResp.Location.Description {
		t.Fatalf("expected location description to be updated, got %v", updateLocationResp.Location.Description)
	}
}
