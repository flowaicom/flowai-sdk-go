# FileUpload

Schema for file upload via base64 encoding


## Fields

| Field                       | Type                        | Required                    | Description                 |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `Filename`                  | *string*                    | :heavy_check_mark:          | N/A                         |
| `Content`                   | *string*                    | :heavy_check_mark:          | Base64 encoded file content |
| `ContentType`               | **string*                   | :heavy_minus_sign:          | MIME type of the file       |