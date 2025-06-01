# TestCases
(*TestCases*)

## Overview

### Available Operations

* [List](#list) - List Test Cases
* [Create](#create) - Create Test Case
* [Get](#get) - Get Test Case
* [Update](#update) - Update Test Case
* [Delete](#delete) - Delete Test Case
* [GetValidation](#getvalidation) - Get Single Validation From Batch

## List

List Test Cases

### Example Usage

```go
package main

import(
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/models/components"
	"github.com/flowaicom/flowai-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := flowaisdkgo.New()

    res, err := s.TestCases.List(ctx, operations.ListTestCasesTestCasesGetRequest{
        JWTPayload: components.JWTPayload{
            Sub: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseListTestCasesTestCasesGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListTestCasesTestCasesGetRequest](../../models/operations/listtestcasestestcasesgetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListTestCasesTestCasesGetResponse](../../models/operations/listtestcasestestcasesgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create Test Case

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

    res, err := s.TestCases.Create(ctx, components.TestCaseCreate{
        ExpectedOutput: "<value>",
        Status: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestCaseRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `testCaseCreate`                                                       | [components.TestCaseCreate](../../models/components/testcasecreate.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `apiKeyUserID`                                                         | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateTestCaseTestCasesPostResponse](../../models/operations/createtestcasetestcasespostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get Test Case

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

    res, err := s.TestCases.Get(ctx, "ba925364-fbf9-4671-8f4e-768dba5a68a5", components.JWTPayload{
        Sub: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TestCaseRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `testCaseID`                                                   | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `jwtPayload`                                                   | [components.JWTPayload](../../models/components/jwtpayload.md) | :heavy_check_mark:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.GetTestCaseTestCasesTestCaseIDGetResponse](../../models/operations/gettestcasetestcasestestcaseidgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update Test Case

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

    res, err := s.TestCases.Update(ctx, "729a6e02-489b-48c0-a515-9fcf8f007759", components.BodyUpdateTestCaseTestCasesTestCaseIDPut{
        TestCaseIn: components.TestCaseUpdate{},
        CurrentUserPayload: components.JWTPayload{
            Sub: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TestCaseRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `testCaseID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `bodyUpdateTestCaseTestCasesTestCaseIDPut`                                                                                 | [components.BodyUpdateTestCaseTestCasesTestCaseIDPut](../../models/components/bodyupdatetestcasetestcasestestcaseidput.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateTestCaseTestCasesTestCaseIDPutResponse](../../models/operations/updatetestcasetestcasestestcaseidputresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete Test Case

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

    res, err := s.TestCases.Delete(ctx, "7d408309-6274-42fd-8416-ea3e553e02ae", components.JWTPayload{
        Sub: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `testCaseID`                                                   | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `jwtPayload`                                                   | [components.JWTPayload](../../models/components/jwtpayload.md) | :heavy_check_mark:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.DeleteTestCaseTestCasesTestCaseIDDeleteResponse](../../models/operations/deletetestcasetestcasestestcaseiddeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetValidation

Get Single Validation From Batch

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

    res, err := s.TestCases.GetValidation(ctx, "187fe75d-c338-4707-b27a-cf7628ebe85e")
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
| `testCaseID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSingleValidationFromBatchTestCasesTestCaseIDValidationGetResponse](../../models/operations/getsinglevalidationfrombatchtestcasestestcaseidvalidationgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |