package integration_tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/fern-demo/square-go-sdk"
	"github.com/fern-demo/square-go-sdk/catalog"
	"github.com/fern-demo/square-go-sdk/option"
	"github.com/google/uuid"
)

func TestCatalogItemLifecycle(t *testing.T) {
	squareClient := initSquareClient()

	// Step 1: Create a catalog item with two variations
	catalogItemName := "test-catalog-item-name"
	variation1Name := "test-catalog-variation-name-1"
	variation2Name := "test-catalog-variation-name-2"

	itemID := fmt.Sprintf("#%s", uuid.New().String())
	variation1ID := fmt.Sprintf("#%s", uuid.New().String())
	variation2ID := fmt.Sprintf("#%s", uuid.New().String())
	imageId := fmt.Sprintf("#%s", uuid.New().String())

	createCatalogItemResp, err := squareClient.Catalog.BatchUpsert(context.TODO(), &square.BatchUpsertCatalogObjectsRequest{
		IdempotencyKey: uuid.New().String(),
		Batches: []*square.CatalogObjectBatch{
			{
				Objects: []*square.CatalogObject{
					{
						Type: string(square.CatalogObjectTypeItem),
						Item: &square.CatalogObjectItem{
							ID: itemID,
							ItemData: &square.CatalogItem{
								Name: square.String(catalogItemName),
								Variations: []*square.CatalogObject{
									{
										Type: string(square.CatalogObjectTypeItemVariation),
										ItemVariation: &square.CatalogObjectItemVariation{
											ID: variation1ID,
											ItemVariationData: &square.CatalogItemVariation{
												Name: square.String(variation1Name),
												PriceMoney: &square.Money{
													Amount:   square.Int64(1000),
													Currency: square.CurrencyUsd.Ptr(),
												},
											},
										},
									},
									{
										Type: string(square.CatalogObjectTypeItemVariation),
										ItemVariation: &square.CatalogObjectItemVariation{
											ID: variation2ID,
											ItemVariationData: &square.CatalogItemVariation{
												Name: square.String(variation2Name),
												PriceMoney: &square.Money{
													Amount:   square.Int64(2000),
													Currency: square.CurrencyUsd.Ptr(),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create catalog item and variations: %v", err)
	}
	fmt.Println(createCatalogItemResp.String())

	// Step 2: Search for the created catalog item
	searchCatalogResp, err := squareClient.Catalog.Search(context.TODO(), &square.SearchCatalogObjectsRequest{
		ObjectTypes: []square.CatalogObjectType{square.CatalogObjectTypeItem},
		Query: &square.CatalogQuery{
			ExactQuery: &square.CatalogQueryExact{
				AttributeName:  "name",
				AttributeValue: catalogItemName,
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to search catalog item: %v", err)
	}
	if len(searchCatalogResp.Objects) == 0 {
		t.Fatalf("catalog item not found")
	}

	itemToUpdate := searchCatalogResp.Objects[0]
	// update item name
	itemToUpdate.Item.ItemData.Name = square.String(fmt.Sprintf("%s-updated", catalogItemName))

	// Step 3: Update the catalog item
	updateCatalogItemResp, err := squareClient.Catalog.BatchUpsert(context.TODO(), &square.BatchUpsertCatalogObjectsRequest{
		IdempotencyKey: uuid.New().String(),
		Batches: []*square.CatalogObjectBatch{
			{
				Objects: []*square.CatalogObject{
					itemToUpdate,
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to update catalog item: %v", err)
	}

	fmt.Println(updateCatalogItemResp.String())

	// Step 4: Create a catalog image
	file, err := os.Open("./test.jpeg")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	createCatalogImageResp, err := squareClient.Catalog.Images.Create(context.TODO(), file, &catalog.ImagesCreateRequest{
		Request: &square.CreateCatalogImageRequest{
			IdempotencyKey: uuid.New().String(),
			ObjectID:       &itemToUpdate.Item.ID,
			Image: &square.CatalogObject{
				Image: &square.CatalogObjectImage{
					ID: imageId,
					ImageData: &square.CatalogImage{
						Name: square.String("test-image"),
					},
				},
			},
		},
	}, option.WithHTTPHeader(http.Header{
		"Content-Type": []string{"image/jpeg"},
	}))

	if err != nil {
		t.Fatalf("failed to create catalog image: %v", err)
	}
	fmt.Println(createCatalogImageResp.String())

	// Step 5: Delete the catalog item and attached variations
	deleteCatalogItemResp, err := squareClient.Catalog.BatchDelete(context.TODO(), &square.BatchDeleteCatalogObjectsRequest{
		ObjectIDs: []string{createCatalogItemResp.Objects[0].Item.ID},
	})
	if err != nil {
		t.Fatalf("failed to delete catalog objects: %v", err)
	}

	// validate the deleted object ids
	if len(deleteCatalogItemResp.DeletedObjectIDs) != 3 {
		t.Fatalf("expected %d deleted objects, got %d", 3, len(deleteCatalogItemResp.DeletedObjectIDs))
	}

	// Step 6: Validate the catalog item is deleted
	searchCatalogResp, err = squareClient.Catalog.Search(context.TODO(), &square.SearchCatalogObjectsRequest{
		ObjectTypes: []square.CatalogObjectType{square.CatalogObjectTypeItem},
		Query: &square.CatalogQuery{
			ExactQuery: &square.CatalogQueryExact{
				AttributeName:  "name",
				AttributeValue: catalogItemName,
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to search catalog item: %v", err)
	}
	if len(searchCatalogResp.Objects) != 0 {
		t.Fatalf("catalog item was not deleted")
	}
}
