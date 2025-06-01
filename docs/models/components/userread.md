# UserRead

Schema for reading/returning user data, aligned with DB model.


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `ID`                                      | *string*                                  | :heavy_check_mark:                        | N/A                                       |
| `Email`                                   | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `FirstName`                               | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `LastName`                                | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `Username`                                | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `ImageURL`                                | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `OrgID`                                   | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `IsActive`                                | **bool*                                   | :heavy_minus_sign:                        | N/A                                       |
| `ClerkID`                                 | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `UpdatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |