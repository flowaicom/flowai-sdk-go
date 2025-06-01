# github.com/flowaicom/flowai-sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/flowaicom/flowai-sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/flowaicom/flowai-sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/flow-ai/main). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

server: FastAPI backend for Flow Expert
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/flowaicom/flowai-sdk-go](#githubcomflowaicomflowai-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/flowaicom/flowai-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [APIKeys](docs/sdks/apikeys/README.md)

* [List](docs/sdks/apikeys/README.md#list) - List API Keys for the authenticated user
* [Create](docs/sdks/apikeys/README.md#create) - Generate a new API Key for the authenticated user
* [Update](docs/sdks/apikeys/README.md#update) - Update an API Key
* [Revoke](docs/sdks/apikeys/README.md#revoke) - Revoke an API Key

### [Auth](docs/sdks/auth/README.md)

* [SDKLogin](docs/sdks/auth/README.md#sdklogin) - Sdk Login
* [CompleteValidatorSignup](docs/sdks/auth/README.md#completevalidatorsignup) - Complete Validator Signup

### [Authentication](docs/sdks/authentication/README.md)

* [VerifyValidatorAccess](docs/sdks/authentication/README.md#verifyvalidatoraccess) - Verify Validator Access

### [Batches](docs/sdks/batches/README.md)

* [ListMine](docs/sdks/batches/README.md#listmine) - List My Batches
* [Get](docs/sdks/batches/README.md#get) - Get Batch By Id
* [Delete](docs/sdks/batches/README.md#delete) - Delete Batch By Id
* [GetTestCases](docs/sdks/batches/README.md#gettestcases) - Get Batch Test Cases
* [GetValidations](docs/sdks/batches/README.md#getvalidations) - Get Validations For Batch
* [CreateValidationTask](docs/sdks/batches/README.md#createvalidationtask) - Create Validation Task For Batch Endpoint
* [GetValidationTask](docs/sdks/batches/README.md#getvalidationtask) - Get Validation Task For Batch Endpoint
* [GetStatus](docs/sdks/batches/README.md#getstatus) - Get Batch Status Endpoint

### [Datasets](docs/sdks/datasets/README.md)

* [GetItems](docs/sdks/datasets/README.md#getitems) - Get Dataset Items
* [Delete](docs/sdks/datasets/README.md#delete) - Delete Dataset


### [Health](docs/sdks/health/README.md)

* [Get](docs/sdks/health/README.md#get) - Get Service Health Status
* [Head](docs/sdks/health/README.md#head) - Check Service Health Status (HEAD)

### [Integrations](docs/sdks/integrations/README.md)

* [HandlePrefectWebhook](docs/sdks/integrations/README.md#handleprefectwebhook) - Prefect Webhook

### [Jobs](docs/sdks/jobs/README.md)

* [GetDetails](docs/sdks/jobs/README.md#getdetails) - Get Job Details
* [Cancel](docs/sdks/jobs/README.md#cancel) - Cancel Job
* [GenerateBatch](docs/sdks/jobs/README.md#generatebatch) - Generate Batch for Job
* [ListBatches](docs/sdks/jobs/README.md#listbatches) - List Batches for a Job
* [GenerateDataset](docs/sdks/jobs/README.md#generatedataset) - Generate Dataset for Job
* [TriggerPipeline](docs/sdks/jobs/README.md#triggerpipeline) - Trigger Pipeline From Stages

### [Projects](docs/sdks/projects/README.md)

* [List](docs/sdks/projects/README.md#list) - List Projects
* [Create](docs/sdks/projects/README.md#create) - Create Project
* [Get](docs/sdks/projects/README.md#get) - Get Project
* [Delete](docs/sdks/projects/README.md#delete) - Delete Project
* [Update](docs/sdks/projects/README.md#update) - Update Project
* [Archive](docs/sdks/projects/README.md#archive) - Archive Project
* [GetDataset](docs/sdks/projects/README.md#getdataset) - Get Project Dataset
* [ListValidators](docs/sdks/projects/README.md#listvalidators) - List Project Validators
* [AddValidator](docs/sdks/projects/README.md#addvalidator) - Add Project Validator
* [RemoveValidator](docs/sdks/projects/README.md#removevalidator) - Remove Project Validator
* [ListJobs](docs/sdks/projects/README.md#listjobs) - List Jobs for Project
* [CreateJob](docs/sdks/projects/README.md#createjob) - Create Job for Project

### [Root](docs/sdks/root/README.md)

* [Get](docs/sdks/root/README.md#get) - Root

### [Stages](docs/sdks/stages/README.md)

* [Ping](docs/sdks/stages/README.md#ping) - Ping Stages
* [List](docs/sdks/stages/README.md#list) - List Stages
* [Create](docs/sdks/stages/README.md#create) - Create Stages For Project
* [GetDetails](docs/sdks/stages/README.md#getdetails) - Get Stage Details
* [Update](docs/sdks/stages/README.md#update) - Update Stage
* [Delete](docs/sdks/stages/README.md#delete) - Delete Stage

### [TestCases](docs/sdks/testcases/README.md)

* [List](docs/sdks/testcases/README.md#list) - List Test Cases
* [Create](docs/sdks/testcases/README.md#create) - Create Test Case
* [Get](docs/sdks/testcases/README.md#get) - Get Test Case
* [Update](docs/sdks/testcases/README.md#update) - Update Test Case
* [Delete](docs/sdks/testcases/README.md#delete) - Delete Test Case
* [GetValidation](docs/sdks/testcases/README.md#getvalidation) - Get Single Validation From Batch

### [Users](docs/sdks/users/README.md)

* [GetMe](docs/sdks/users/README.md#getme) - Read Users Me
* [UpdateMe](docs/sdks/users/README.md#updateme) - Update User Me
* [UpdateRole](docs/sdks/users/README.md#updaterole) - Update User Role
* [GetBasicInfo](docs/sdks/users/README.md#getbasicinfo) - Get Current User Basic Info

### [Validations](docs/sdks/validations/README.md)

* [List](docs/sdks/validations/README.md#list) - List Validations
* [Create](docs/sdks/validations/README.md#create) - Create Test Case Validation
* [Get](docs/sdks/validations/README.md#get) - Get Validation
* [Update](docs/sdks/validations/README.md#update) - Update Test Case Validation
* [Delete](docs/sdks/validations/README.md#delete) - Delete Validation

### [ValidationTasks](docs/sdks/validationtasks/README.md)

* [GetStatus](docs/sdks/validationtasks/README.md#getstatus) - Check Validation Task Status

### [ValidatorTasks](docs/sdks/validatortasks/README.md)

* [SubmitValidation](docs/sdks/validatortasks/README.md#submitvalidation) - Submit Test Case Validation
* [LoadTestCases](docs/sdks/validatortasks/README.md#loadtestcases) - Load Test Cases For Validator
* [UpdateValidation](docs/sdks/validatortasks/README.md#updatevalidation) - Edit Test Case Validation
* [Submit](docs/sdks/validatortasks/README.md#submit) - Submit Validator Task

### [Webhooks](docs/sdks/webhooks/README.md)

* [Handle](docs/sdks/webhooks/README.md#handle) - Handle Clerk Webhook Events

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := flowaisdkgo.New()

	res, err := s.Health.Get(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseGetHealthStatus != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := flowaisdkgo.New(
		flowaisdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	res, err := s.Health.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseGetHealthStatus != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Handle` function may return the following errors:

| Error Type                    | Status Code | Content Type     |
| ----------------------------- | ----------- | ---------------- |
| apierrors.HTTPValidationError | 422         | application/json |
| apierrors.APIError            | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"github.com/flowaicom/flowai-sdk-go/models/apierrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := flowaisdkgo.New()

	res, err := s.Webhooks.Handle(ctx, nil, nil, nil)
	if err != nil {

		var e *apierrors.HTTPValidationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	flowaisdkgo "github.com/flowaicom/flowai-sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := flowaisdkgo.New(
		flowaisdkgo.WithServerURL("http://localhost:8000"),
	)

	res, err := s.Health.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseGetHealthStatus != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/flowaicom/flowai-sdk-go&utm_campaign=go)
