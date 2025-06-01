# ValidationTasks
(*ValidationTasks*)

## Overview

### Available Operations

* [GetStatus](#getstatus) - Check Validation Task Status

## GetStatus

Retrieves the current status and basic information of a specific validation task.

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

    res, err := s.ValidationTasks.GetStatus(ctx, "512391e4-a78a-48d4-9b2b-7e2623c3162c")
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidationTaskRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `validationTaskID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CheckValidationTaskStatusValidationTasksValidationTaskIDGetResponse](../../models/operations/checkvalidationtaskstatusvalidationtasksvalidationtaskidgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |