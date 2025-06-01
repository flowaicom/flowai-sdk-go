# Auth
(*Auth*)

## Overview

### Available Operations

* [SDKLogin](#sdklogin) - Sdk Login
* [CompleteValidatorSignup](#completevalidatorsignup) - Complete Validator Signup

## SDKLogin

Allows a user to obtain a session token for SDK access.
If no active session exists, creates a new one.
Returns a session token that can be used for SDK authentication.

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

    res, err := s.Auth.SDKLogin(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.JWTToken != nil {
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

**[*operations.SDKLoginAuthSDKLoginPostResponse](../../models/operations/sdkloginauthsdkloginpostresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CompleteValidatorSignup

Complete Validator Signup

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

    res, err := s.Auth.CompleteValidatorSignup(ctx, components.ValidatorCompleteSignupRequest{
        ValidationTaskID: "e8926b0b-0947-42ed-b09e-90f8ae03f6b2",
        ValidatorID: "9e7bf08a-82e7-4b5d-b801-30059c23c117",
        UniqueURLKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseCompleteValidatorSignupAuthValidatorsCompleteSignupPost != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.ValidatorCompleteSignupRequest](../../models/components/validatorcompletesignuprequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CompleteValidatorSignupAuthValidatorsCompleteSignupPostResponse](../../models/operations/completevalidatorsignupauthvalidatorscompletesignuppostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |