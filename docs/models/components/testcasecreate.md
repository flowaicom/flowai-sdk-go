# TestCaseCreate

Schema for creating a test case, accepting nested objects


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Name`                                                           | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `Description`                                                    | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `ExpectedOutput`                                                 | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `Status`                                                         | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `IsActive`                                                       | **bool*                                                          | :heavy_minus_sign:                                               | N/A                                                              |
| `Trajectory`                                                     | [][components.Trajectory](../../models/components/trajectory.md) | :heavy_minus_sign:                                               | N/A                                                              |