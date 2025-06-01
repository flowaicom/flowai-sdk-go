# Datasets
(*Datasets*)

## Overview

### Available Operations

* [GetItems](#getitems) - Get Dataset Items
* [Delete](#delete) - Delete Dataset

## GetItems

Retrieves a paginated list of dataset items belonging to a specific dataset.

### Example Usage

```go
package main

import(
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := flowaisdkgo.New()

    res, err := s.Datasets.GetItems(ctx, "ebd3200c-226f-403c-9ac9-78e04cf26055", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasetItemListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `datasetID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Maximum number of dataset items to return.               |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Number of dataset items to skip for pagination.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDatasetItemsDatasetsDatasetIDItemsGetResponse](../../models/operations/getdatasetitemsdatasetsdatasetiditemsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Permanently deletes a specific dataset and its associated items.

### Example Usage

```go
package main

import(
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := flowaisdkgo.New()

    res, err := s.Datasets.Delete(ctx, "e503657d-340f-4f49-9559-db4bf8be136c")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `datasetID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteDatasetDatasetsDatasetIDDeleteResponse](../../models/operations/deletedatasetdatasetsdatasetiddeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |