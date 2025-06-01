# CreateValidationTaskResponse

Response schema for creating a validation task for a batch.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ValidationTaskID`                                                           | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `BatchID`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `Status`                                                                     | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `CreatedAt`                                                                  | [time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_check_mark:                                                           | N/A                                                                          |
| `ValidatorUrls`                                                              | [][components.ValidatorURLInfo](../../models/components/validatorurlinfo.md) | :heavy_check_mark:                                                           | N/A                                                                          |