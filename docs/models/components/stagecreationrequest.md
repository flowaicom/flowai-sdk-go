# StageCreationRequest

Schema for creating stages via SDK (compatible with JSON)


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Files`                                                                                | [][components.FileUpload](../../models/components/fileupload.md)                       | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `SourceType`                                                                           | [components.StageFileTypeEnum](../../models/components/stagefiletypeenum.md)           | :heavy_check_mark:                                                                     | Enum for stage file types that inherits from str for better JSON/OpenAPI compatibility |
| `ProjectID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | UUID of the project                                                                    |