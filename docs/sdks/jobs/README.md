# Jobs
(*Jobs*)

## Overview

### Available Operations

* [GetDetails](#getdetails) - Get Job Details
* [Cancel](#cancel) - Cancel Job
* [GenerateBatch](#generatebatch) - Generate Batch for Job
* [ListBatches](#listbatches) - List Batches for a Job
* [GenerateDataset](#generatedataset) - Generate Dataset for Job
* [TriggerPipeline](#triggerpipeline) - Trigger Pipeline From Stages

## GetDetails

Retrieves details for a specific job.

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

    res, err := s.Jobs.GetDetails(ctx, "beda9b0e-9d09-4a53-a3d8-8ba72834ab8e")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetJobDetailsJobsJobIDGetResponse](../../models/operations/getjobdetailsjobsjobidgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Cancel

Cancels a job if possible.

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

    res, err := s.Jobs.Cancel(ctx, "ecacb490-e165-4424-9a70-5bee2b7cb80c")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobCancelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CancelJobByIDJobsJobIDCancelPostResponse](../../models/operations/canceljobbyidjobsjobidcancelpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GenerateBatch

Starts batch generation for a job.

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

    res, err := s.Jobs.GenerateBatch(ctx, "7d4e5dfa-cf50-4b48-a68a-c7eca7b06a75")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobBatchGenerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GenerateBatchForJobJobsJobIDGenerateBatchPostResponse](../../models/operations/generatebatchforjobjobsjobidgeneratebatchpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListBatches

Retrieves batches associated with a job.

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

    res, err := s.Jobs.ListBatches(ctx, "5b14862a-6c07-4405-bce5-da60422c717f", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListJobBatchesForJobJobsJobIDBatchesGetResponse](../../models/operations/listjobbatchesforjobjobsjobidbatchesgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GenerateDataset

Starts dataset generation from a job.

### Example Usage

```go
package main

import(
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := flowaisdkgo.New()

    res, err := s.Jobs.GenerateDataset(ctx, "80bad737-ad60-4c43-aa9f-314f21fe6439", components.DatasetGenerationRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.JobDatasetGenerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `jobID`                                                                                    | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `datasetGenerationRequest`                                                                 | [components.DatasetGenerationRequest](../../models/components/datasetgenerationrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GenerateDatasetForJobJobsJobIDGenerateDatasetPostResponse](../../models/operations/generatedatasetforjobjobsjobidgeneratedatasetpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TriggerPipeline

Trigger Pipeline From Stages

### Example Usage

```go
package main

import(
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := flowaisdkgo.New()

    res, err := s.Jobs.TriggerPipeline(ctx, "d4ff7706-0465-4327-966d-f90f57254f7d", []components.StageDetailsResponse{
        components.StageDetailsResponse{
            ID: "86663fe7-41a6-49a7-a7ac-de4c44d473ae",
            ProjectID: "72247f39-8cd0-4c93-ac47-b0da59df86ea",
            S3Path: "<value>",
            SourceType: "<value>",
        },
        components.StageDetailsResponse{
            ID: "86663fe7-41a6-49a7-a7ac-de4c44d473ae",
            ProjectID: "72247f39-8cd0-4c93-ac47-b0da59df86ea",
            S3Path: "<value>",
            SourceType: "<value>",
        },
        components.StageDetailsResponse{
            ID: "86663fe7-41a6-49a7-a7ac-de4c44d473ae",
            ProjectID: "72247f39-8cd0-4c93-ac47-b0da59df86ea",
            S3Path: "<value>",
            SourceType: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.JobBatchGenerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `projectID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `requestBody`                                                                        | [][components.StageDetailsResponse](../../models/components/stagedetailsresponse.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.TriggerPipelineFromStagesJobsTriggerPipelineFromStagesPostResponse](../../models/operations/triggerpipelinefromstagesjobstriggerpipelinefromstagespostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |