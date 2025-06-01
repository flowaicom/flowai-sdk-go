# Batches
(*Batches*)

## Overview

### Available Operations

* [ListMine](#listmine) - List My Batches
* [Get](#get) - Get Batch By Id
* [Delete](#delete) - Delete Batch By Id
* [GetTestCases](#gettestcases) - Get Batch Test Cases
* [GetValidations](#getvalidations) - Get Validations For Batch
* [CreateValidationTask](#createvalidationtask) - Create Validation Task For Batch Endpoint
* [GetValidationTask](#getvalidationtask) - Get Validation Task For Batch Endpoint
* [GetStatus](#getstatus) - Get Batch Status Endpoint

## ListMine

List My Batches

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

    res, err := s.Batches.ListMine(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseListMyBatchesBatchesMineGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListMyBatchesBatchesMineGetResponse](../../models/operations/listmybatchesbatchesminegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get Batch By Id

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

    res, err := s.Batches.Get(ctx, "c1be57c5-73d0-4192-8460-c8feb4a1d0ea")
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetBatchByIDBatchesBatchIDGetResponse](../../models/operations/getbatchbyidbatchesbatchidgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete Batch By Id

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

    res, err := s.Batches.Delete(ctx, "a9a1c03c-b3f4-423f-8840-e8e17909e384")
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
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteBatchByIDBatchesBatchIDDeleteResponse](../../models/operations/deletebatchbyidbatchesbatchiddeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetTestCases

Get Batch Test Cases

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

    res, err := s.Batches.GetTestCases(ctx, "a065bb56-ca76-4b09-9cdd-06c06f722ffc")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseGetBatchTestCasesBatchesBatchIDTestcasesGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetBatchTestCasesBatchesBatchIDTestcasesGetResponse](../../models/operations/getbatchtestcasesbatchesbatchidtestcasesgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetValidations

Get Validations For Batch

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

    res, err := s.Batches.GetValidations(ctx, "59088369-c9bc-43ee-835a-04250b7613a8")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseGetValidationsForBatchBatchesBatchIDValidationsGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetValidationsForBatchBatchesBatchIDValidationsGetResponse](../../models/operations/getvalidationsforbatchbatchesbatchidvalidationsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CreateValidationTask

Create Validation Task For Batch Endpoint

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

    res, err := s.Batches.CreateValidationTask(ctx, "e512837d-3b13-42b2-a375-d1eced4c1aec")
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateValidationTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateValidationTaskForBatchEndpointBatchesBatchIDValidationTasksPostResponse](../../models/operations/createvalidationtaskforbatchendpointbatchesbatchidvalidationtaskspostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetValidationTask

Get Validation Task For Batch Endpoint

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

    res, err := s.Batches.GetValidationTask(ctx, "050301f2-e370-4717-9c8e-4aec119dea2f")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetValidationTaskForBatchResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetValidationTaskForBatchEndpointBatchesBatchIDValidationTaskGetResponse](../../models/operations/getvalidationtaskforbatchendpointbatchesbatchidvalidationtaskgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetStatus

Get Batch Status Endpoint

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

    res, err := s.Batches.GetStatus(ctx, "f9433f8a-a7a8-45f8-beca-1962cf314022")
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchStatusRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `batchID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetBatchStatusEndpointBatchesBatchIDStatusGetResponse](../../models/operations/getbatchstatusendpointbatchesbatchidstatusgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |