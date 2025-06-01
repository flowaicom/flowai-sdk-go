# BatchRead

Schema for reading/returning batch data


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ProjectID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `PipelineRunID`                                                      | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `Name`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Description`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `Error`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | N/A                                                                  |
| `UpdatedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | N/A                                                                  |
| `TestCases`                                                          | [][components.TestCaseRead](../../models/components/testcaseread.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `TotalTestCases`                                                     | *int64*                                                              | :heavy_check_mark:                                                   | N/A                                                                  |
| `CompletedTestCases`                                                 | *int64*                                                              | :heavy_check_mark:                                                   | N/A                                                                  |