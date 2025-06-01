# ListValidationsValidationsGetRequest


## Fields

| Field                                | Type                                 | Required                             | Description                          |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `CurrentUserPayload`                 | *any*                                | :heavy_check_mark:                   | N/A                                  |
| `Skip`                               | **int64*                             | :heavy_minus_sign:                   | N/A                                  |
| `Limit`                              | **int64*                             | :heavy_minus_sign:                   | N/A                                  |
| `SortBy`                             | **string*                            | :heavy_minus_sign:                   | Field to sort by                     |
| `SortOrder`                          | **string*                            | :heavy_minus_sign:                   | Sort direction (asc or desc)         |
| `TestCaseID`                         | **string*                            | :heavy_minus_sign:                   | Filter by test case ID               |
| `ValidatorID`                        | **string*                            | :heavy_minus_sign:                   | Filter by validator user ID          |
| `IsAccepted`                         | **bool*                              | :heavy_minus_sign:                   | Filter by acceptance status          |
| `CreatedAfter`                       | **string*                            | :heavy_minus_sign:                   | Filter by creation date (ISO format) |
| `CreatedBefore`                      | **string*                            | :heavy_minus_sign:                   | Filter by creation date (ISO format) |