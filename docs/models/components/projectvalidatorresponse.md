# ProjectValidatorResponse

Schema for project validator response matching API spec.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | *string*                                                      | :heavy_check_mark:                                            | The ID of the validator record (project-specific assignment). |
| `ProjectID`                                                   | *string*                                                      | :heavy_check_mark:                                            | The ID of the project.                                        |
| `Email`                                                       | *string*                                                      | :heavy_check_mark:                                            | Email address of the validator.                               |
| `ClerkUserID`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Clerk User ID if user exists, otherwise null.                 |
| `CreatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | Timestamp of assignment creation.                             |