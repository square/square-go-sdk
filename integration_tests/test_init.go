package integration_tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	square "github.com/fern-demo/square-go-sdk"
	client "github.com/fern-demo/square-go-sdk/client"
	"github.com/fern-demo/square-go-sdk/option"
)

const LOCATION_NAME = "test-location-name"

func initSquareClient() *client.Client {
	return client.NewClient(
		option.WithToken(os.Getenv("SQUARE_SANDBOX_TOKEN")),
		option.WithBaseURL("https://connect.squareupsandbox.com"))
}

// Run me to insert necessary objects into the test account
func InitTestAccount(t *testing.T) {
	squareClient := initSquareClient()

	// check that location exists
	listLocationsResp, err := squareClient.Locations.List(context.TODO())
	if err != nil {
		t.Fatalf("failed to list locations: %v", err)
	}
	// iterate through, looking for location with LOCATION_NAME
	foundLocation := false
	for _, location := range listLocationsResp.Locations {
		if *location.Name == LOCATION_NAME {
			foundLocation = true
			break
		}
	}

	if !foundLocation {
		fmt.Println("Location does not exist. Creating it...")
		// create location that we modify
		createLocationResp, err := squareClient.Locations.Create(context.TODO(), &square.CreateLocationRequest{
			Location: &square.Location{
				Name: square.String(LOCATION_NAME),
			},
		})
		if err != nil {
			t.Fatalf("failed to create location: %v. square error: %v", err, createLocationResp.Errors)
		}
	} else {
		fmt.Println("Location already exists. Skipping create")
	}
}
