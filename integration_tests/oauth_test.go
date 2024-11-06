package integration_tests

//func TestRetrieveTokenStatus(t *testing.T) {
//	squareClient := initSquareClient()
//
//	// Step 1: Retrieve the status of the OAuth token
//	retrieveTokenStatusResp, err := squareClient.OAuth.RetrieveTokenStatus(context.TODO())
//	if err != nil {
//		t.Fatalf("failed to retrieve token status: %v", err)
//	}
//
//	//squareClient.OAuth.Authorize()
//
//	squareClient.OAuth.RenewToken(context.TODO(),
//		*retrieveTokenStatusResp.ClientID,
//		option.WithHTTPHeader(
//			http.Header{
//				"Authorization": []string{fmt.Sprintf("Bearer %s", os.Getenv("SQUARE_SANDBOX_TOKEN"))},
//			},
//		)
//	)
//
//	//fmt.Printf("Retrieved token status: %v\n", retrieveTokenStatusResp)
//}
