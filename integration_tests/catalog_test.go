//go:build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	square "github.com/square/square-go-sdk/v2"
	"github.com/square/square-go-sdk/v2/catalog"
	"github.com/square/square-go-sdk/v2/core"
)

// Catalog API integration tests.
func TestCatalogAPI(t *testing.T) {
	t.Run("bulk create and iterate through paginated catalog objects", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		// Setup: Create 100 catalog objects with 1 CatalogItemVariation each.
		catalogObjects := make([]*square.CatalogObject, 0, 100)
		for i := 0; i < 100; i++ {
			catalogObjects = append(catalogObjects, newTestCatalogObject())
		}

		// Create the catalog objects in a bulk request.
		createCatalogObjectsResp, err := squareClient.Catalog.BatchUpsert(
			context.Background(),
			&square.BatchUpsertCatalogObjectsRequest{
				IdempotencyKey: newTestUUID(),
				Batches: []*square.CatalogObjectBatch{
					{
						Objects: catalogObjects,
					},
				},
			},
		)
		require.NoError(t, err)
		assert.Len(t, createCatalogObjectsResp.Objects, 100)

		// List all catalog objects.
		catalogObjectsResp, err := squareClient.Catalog.List(
			context.Background(),
			&square.CatalogListRequest{},
		)
		require.NoError(t, err)
		assert.Len(t, catalogObjectsResp.Results, MaxCatalogPageSize)

		// Check if there's more than one page of catalog objects to iterate through,
		// and that the eventual error returned is `ErrNoPages`, which signals the
		// end of the iteration.
		numberOfPages := 0
		currentPage := catalogObjectsResp
		for {
			numberOfPages++
			currentPage, err = currentPage.GetNextPage(context.Background())
			if err != nil {
				assert.Equal(t, core.ErrNoPages, err)
				break
			}
		}
		assert.Greater(t, numberOfPages, 0)

		// Cleanup: Delete the catalog objects we created.
		var objectIDs []string
		for _, catalogObject := range createCatalogObjectsResp.Objects {
			objectIDs = append(objectIDs, catalogObject.Item.ID)
			require.GreaterOrEqual(t, len(catalogObject.Item.ItemData.Variations), 1)
			objectIDs = append(objectIDs, catalogObject.Item.ItemData.Variations[0].ItemVariation.ID)
		}

		deleteCatalogObjectsResp, err := squareClient.Catalog.BatchDelete(
			context.Background(),
			&square.BatchDeleteCatalogObjectsRequest{
				ObjectIDs: objectIDs,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, 200, len(deleteCatalogObjectsResp.DeletedObjectIDs))
	})

	t.Run("upload catalog image", func(t *testing.T) {
		squareClient := newTestSquareClient(t)

		// Setup: Load a test image file.
		file, err := os.Open("./testdata/image.jpeg")
		require.NoError(t, err)
		defer func() {
			require.NoError(t, file.Close())
		}()

		// Setup: Create a catalog object to associate the image with.
		catalogObject := newTestCatalogObject()
		createCatalogResp, err := squareClient.Catalog.BatchUpsert(
			context.Background(),
			&square.BatchUpsertCatalogObjectsRequest{
				IdempotencyKey: newTestUUID(),
				Batches: []*square.CatalogObjectBatch{
					{
						Objects: []*square.CatalogObject{catalogObject},
					},
				},
			},
		)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(createCatalogResp.Objects), 1)

		createdCatalogObject := createCatalogResp.Objects[0]
		require.NotNil(t, createdCatalogObject)

		// Create a new catalog image.
		imageName := fmt.Sprintf("Test Image %s", newTestUUID())
		createCatalogImageResp, err := squareClient.Catalog.Images.Create(
			context.Background(),
			&catalog.ImagesCreateRequest{
				ImageFile: file,
				Request: &square.CreateCatalogImageRequest{
					IdempotencyKey: newTestUUID(),
					ObjectID:       &createdCatalogObject.Item.ID,
					Image: &square.CatalogObject{
						Image: &square.CatalogObjectImage{
							ID: newTestSquareTempID(),
							ImageData: &square.CatalogImage{
								Name: &imageName,
							},
						},
					},
				},
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, createCatalogImageResp.Image)

		// Cleanup: Delete the created catalog object and image.
		_, err = squareClient.Catalog.BatchDelete(
			context.Background(),
			&square.BatchDeleteCatalogObjectsRequest{
				ObjectIDs: []string{
					createdCatalogObject.Item.ID,
					createCatalogImageResp.Image.Image.ID,
				},
			},
		)
		require.NoError(t, err)
	})
}

// Creates a CatalogObject with a single variation.
func newTestCatalogObject() *square.CatalogObject {
	return &square.CatalogObject{
		Item: &square.CatalogObjectItem{
			ID: newTestSquareTempID(),
			ItemData: &square.CatalogItem{
				Name: square.String(fmt.Sprintf("Item %s", newTestUUID())),
				Variations: []*square.CatalogObject{
					{
						ItemVariation: &square.CatalogObjectItemVariation{
							ID: newTestSquareTempID(),
							ItemVariationData: &square.CatalogItemVariation{
								Name:       square.String(fmt.Sprintf("Variation %s", newTestUUID())),
								PriceMoney: newTestMoney(1000),
							},
						},
					},
				},
			},
		},
	}
}
