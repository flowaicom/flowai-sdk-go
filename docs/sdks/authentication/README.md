# Authentication
(*Authentication*)

## Overview

### Available Operations

* [VerifyValidatorAccess](#verifyvalidatoraccess) - Verify Validator Access

## VerifyValidatorAccess

Verify Validator Access

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

    res, err := s.Authentication.VerifyValidatorAccess(ctx, components.ValidatorVerifyAccessRequest{
        ValidationTaskID: "3092462e-7674-4322-89be-9e48f4298a69",
        ValidatorID: "e3fec0f9-7451-4f71-9146-632a5e7a343b",
        UniqueURLKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseVerifyValidatorAccessAuthValidatorsVerifyAccessPost != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.ValidatorVerifyAccessRequest](../../models/components/validatorverifyaccessrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.VerifyValidatorAccessAuthValidatorsVerifyAccessPostResponse](../../models/operations/verifyvalidatoraccessauthvalidatorsverifyaccesspostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |