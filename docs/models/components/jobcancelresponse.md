# JobCancelResponse


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `ID`                                       | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `ProjectID`                                | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `Status`                                   | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `StartTime`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `EndTime`                                  | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Error`                                    | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Config`                                   | map[string]*any*                           | :heavy_minus_sign:                         | N/A                                        |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |
| `UpdatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |