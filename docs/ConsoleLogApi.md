# \ConsoleLogApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsoleLogV1**](ConsoleLogApi.md#CreateConsoleLogV1) | **Post** /v1/console-logs | Create console log
[**GetConsoleLogsV1**](ConsoleLogApi.md#GetConsoleLogsV1) | **Get** /v1/console-logs | Get console logs



## CreateConsoleLogV1

> CreateConsoleLogV1(ctx).Model(model).Execute()

Create console log



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
    model := *openapiclient.NewCreateConsoleLogRequestDTO(map[string][]openapiclient.LogLevel{ ... }, "Message_example", "Service_example") // CreateConsoleLogRequestDTO | The console log DTO

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsoleLogApi.CreateConsoleLogV1(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsoleLogApi.CreateConsoleLogV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsoleLogV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**CreateConsoleLogRequestDTO**](CreateConsoleLogRequestDTO.md) | The console log DTO | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)


## GetConsoleLogsV1

> ConsoleLogListPageResponse GetConsoleLogsV1(ctx).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()

Get console logs



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
    page := int32(56) // int32 | The desired page number (default to 0)
    pageSize := int32(56) // int32 | The desired number of items per page (default to 25)
    filter := "filter_example" // string | The search term (optional)
    direction := "direction_example" // string | The sorting direction (optional)
    sort := "sort_example" // string | The sorting parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsoleLogApi.GetConsoleLogsV1(context.Background()).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsoleLogApi.GetConsoleLogsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsoleLogsV1`: ConsoleLogListPageResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsoleLogApi.GetConsoleLogsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleLogsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The desired page number | [default to 0]
 **pageSize** | **int32** | The desired number of items per page | [default to 25]
 **filter** | **string** | The search term | 
 **direction** | **string** | The sorting direction | 
 **sort** | **string** | The sorting parameter | 

### Return type

[**ConsoleLogListPageResponse**](ConsoleLogListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)

