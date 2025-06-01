# ListTestCasesTestCasesGetRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Skip`                                                         | **int64*                                                       | :heavy_minus_sign:                                             | N/A                                                            |
| `Limit`                                                        | **int64*                                                       | :heavy_minus_sign:                                             | N/A                                                            |
| `SortBy`                                                       | **string*                                                      | :heavy_minus_sign:                                             | Field to sort by (e.g., 'name', 'created_at')                  |
| `SortOrder`                                                    | **string*                                                      | :heavy_minus_sign:                                             | Sort direction ('asc' or 'desc')                               |
| `Name`                                                         | **string*                                                      | :heavy_minus_sign:                                             | Filter by name (partial match)                                 |
| `Status`                                                       | **string*                                                      | :heavy_minus_sign:                                             | Filter by status                                               |
| `CreatedAfter`                                                 | **string*                                                      | :heavy_minus_sign:                                             | Filter by creation date (ISO format, YYYY-MM-DDTHH:MM:SSZ)     |
| `CreatedBefore`                                                | **string*                                                      | :heavy_minus_sign:                                             | Filter by creation date (ISO format, YYYY-MM-DDTHH:MM:SSZ)     |
| `BatchID`                                                      | **string*                                                      | :heavy_minus_sign:                                             | Filter by Batch ID                                             |
| `JWTPayload`                                                   | [components.JWTPayload](../../models/components/jwtpayload.md) | :heavy_check_mark:                                             | N/A                                                            |