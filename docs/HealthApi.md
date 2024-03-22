# \HealthApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthV1**](HealthApi.md#GetHealthV1) | **Get** /v1/health | Service health check



## GetHealthV1

> HandlersHealthCheck GetHealthV1(ctx).Execute()

Service health check



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthApi.GetHealthV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetHealthV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthV1`: HandlersHealthCheck
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetHealthV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthV1Request struct via the builder pattern


### Return type

[**HandlersHealthCheck**](HandlersHealthCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)

