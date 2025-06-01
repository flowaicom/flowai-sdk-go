<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := flowaisdkgo.New()

	res, err := s.Health.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseGetHealthStatus != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->