# TrajectoryItem

Schema for reading a trajectory item, including the linked message or tool call.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Order`                                                     | *int64*                                                     | :heavy_check_mark:                                          | N/A                                                         |
| `CreatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `UpdatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `TestCaseID`                                                | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Message`                                                   | [*components.Message](../../models/components/message.md)   | :heavy_minus_sign:                                          | N/A                                                         |
| `ToolCall`                                                  | [*components.ToolCall](../../models/components/toolcall.md) | :heavy_minus_sign:                                          | N/A                                                         |
| `ItemType`                                                  | [components.ItemType](../../models/components/itemtype.md)  | :heavy_check_mark:                                          | Derives the type of the trajectory item.                    |