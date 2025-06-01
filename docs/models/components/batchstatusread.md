# BatchStatusRead

Schema for returning batch status, primarily derived from its PipelineRun.


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `BatchID`                                 | *string*                                  | :heavy_check_mark:                        | N/A                                       |
| `PipelineRunID`                           | *string*                                  | :heavy_check_mark:                        | N/A                                       |
| `Status`                                  | *string*                                  | :heavy_check_mark:                        | N/A                                       |
| `ExternalRunID`                           | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `PipelineRunUpdatedAt`                    | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |