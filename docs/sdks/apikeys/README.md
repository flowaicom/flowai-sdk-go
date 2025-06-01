# APIKeys
(*APIKeys*)

## Overview

### Available Operations

* [List](#list) - List API Keys for the authenticated user
* [Create](#create) - Generate a new API Key for the authenticated user
* [Update](#update) - Update an API Key
* [Revoke](#revoke) - Revoke an API Key

## List

Retrieves a list of non-sensitive information about the user's API keys.

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

    res, err := s.APIKeys.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserAPIKeysAPIKeysGetResponse](../../models/operations/listuserapikeysapikeysgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Generates a new API key, hashes it, stores the hash, and returns the plain key ONCE.

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

    res, err := s.APIKeys.Create(ctx, components.APIKeyCreate{
        KeyName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIKeyCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.APIKeyCreate](../../models/components/apikeycreate.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateAPIKeyForUserAPIKeysPostResponse](../../models/operations/createapikeyforuserapikeyspostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Updates the name or active status of a specific API key owned by the user.

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

    res, err := s.APIKeys.Update(ctx, "69ca862a-fb74-4f84-92d3-dd14874aae5c", components.APIKeyUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.APIKeyInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `keyID`                                                            | *string*                                                           | :heavy_check_mark:                                                 | The ID of the API key to update                                    |
| `apiKeyUpdate`                                                     | [components.APIKeyUpdate](../../models/components/apikeyupdate.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.UpdateAPIKeyDetailsAPIKeysKeyIDPutResponse](../../models/operations/updateapikeydetailsapikeyskeyidputresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Revoke

Marks a specific API key as inactive (revokes it).

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

    res, err := s.APIKeys.Revoke(ctx, "98be69b5-84e8-426f-bce1-a75f4a2eb71c")
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
| `keyID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The ID of the API key to revoke                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RevokeAPIKeyAPIKeysKeyIDDeleteResponse](../../models/operations/revokeapikeyapikeyskeyiddeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |