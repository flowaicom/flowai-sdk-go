# APIKeyInfo


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `KeyName`                                  | **string*                                  | :heavy_minus_sign:                         | Optional user-friendly name for the key    |
| `ID`                                       | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `UserID`                                   | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `KeyPrefix`                                | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |
| `UpdatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |
| `LastUsedAt`                               | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `IsActive`                                 | *bool*                                     | :heavy_check_mark:                         | N/A                                        |