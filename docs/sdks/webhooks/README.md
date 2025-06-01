# Webhooks
(*Webhooks*)

## Overview

### Available Operations

* [Handle](#handle) - Handle Clerk Webhook Events

## Handle

Handles incoming webhook events from Clerk.
Verification is done by the `verify_clerk_webhook` dependency.
Retrieves the verified event from request.state for further processing.

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

    res, err := s.Webhooks.Handle(ctx, nil, nil, nil)
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
| `svixID`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `svixTimestamp`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `svixSignature`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.HandleClerkWebhookEndpointClerkWebhooksPostResponse](../../models/operations/handleclerkwebhookendpointclerkwebhookspostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |