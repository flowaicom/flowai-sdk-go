# Users
(*Users*)

## Overview

### Available Operations

* [GetMe](#getme) - Read Users Me
* [UpdateMe](#updateme) - Update User Me
* [UpdateRole](#updaterole) - Update User Role
* [GetBasicInfo](#getbasicinfo) - Get Current User Basic Info

## GetMe

Read Users Me

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

    res, err := s.Users.GetMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserRead != nil {
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

**[*operations.ReadUsersMeUsersMeGetResponse](../../models/operations/readusersmeusersmegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMe

Update User Me

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

    res, err := s.Users.UpdateMe(ctx, "<value>", "<value>", components.UserUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UserRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `args`                                                         | *any*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `kwargs`                                                       | *any*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `userUpdate`                                                   | [components.UserUpdate](../../models/components/userupdate.md) | :heavy_check_mark:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.UpdateUserMeUsersMePatchResponse](../../models/operations/updateusermeusersmepatchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UpdateRole

Update the user's role in Clerk's public metadata.
Only allows setting role to either 'user' or 'validator'.
Returns 204 No Content on successful update.

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

    res, err := s.Users.UpdateRole(ctx, "<value>", "<value>", components.UserRoleUpdate{
        Role: components.RoleUser,
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `args`                                                                 | *any*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    |
| `kwargs`                                                               | *any*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    |
| `userRoleUpdate`                                                       | [components.UserRoleUpdate](../../models/components/userroleupdate.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.UpdateUserRoleUsersRolePatchResponse](../../models/operations/updateuserroleusersrolepatchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetBasicInfo

Get Current User Basic Info

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

    res, err := s.Users.GetBasicInfo(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCurrentUserBasicInfoUsersMeBasicInfoGetResponse](../../models/operations/getcurrentuserbasicinfousersmebasicinfogetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |