# Integrations
(*Integrations*)

## Overview

### Available Operations

* [HandlePrefectWebhook](#handleprefectwebhook) - Prefect Webhook

## HandlePrefectWebhook

Receives notifications from Prefect about flow run state changes.
If COMPLETED, enqueues a task to RQ for result processing.

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

    res, err := s.Integrations.HandlePrefectWebhook(ctx, components.PrefectWebhookPayload{
        FlowRun: components.PrefectFlowRunInfo{
            PrefectFlowRunID: "<id>",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `prefectWebhookPayload`                                                              | [components.PrefectWebhookPayload](../../models/components/prefectwebhookpayload.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `queueName`                                                                          | **string*                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PrefectWebhookIntegrationsPrefectWebhookPostResponse](../../models/operations/prefectwebhookintegrationsprefectwebhookpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |