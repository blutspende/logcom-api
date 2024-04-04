# \AuditLogApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuditLogV1Int**](AuditLogApi.md#CreateAuditLogV1Int) | **Post** /v1/int/audit-logs | Create audit log
[**CreateAuditLogsV1Int**](AuditLogApi.md#CreateAuditLogsV1Int) | **Post** /v1/int/audit-logs/batch | Create audit logs
[**GetAuditLogByIDV1**](AuditLogApi.md#GetAuditLogByIDV1) | **Get** /v1/audit-logs/{id} | Get audit log by ID
[**GetAuditLogsV1**](AuditLogApi.md#GetAuditLogsV1) | **Get** /v1/audit-logs | Get audit logs



## CreateAuditLogV1Int

> CreateAuditLogV1Int(ctx).Model(model).Execute()

Create audit log



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
    model := *openapiclient.NewCreateAuditLogRequestDTO("Category_example", "ServiceCreated_example", "Subject_example") // CreateAuditLogRequestDTO | The audit log DTO

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.CreateAuditLogV1Int(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.CreateAuditLogV1Int``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuditLogV1IntRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**CreateAuditLogRequestDTO**](CreateAuditLogRequestDTO.md) | The audit log DTO | 

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


## CreateAuditLogsV1Int

> CreateAuditLogsV1Int(ctx).Model(model).Execute()

Create audit logs



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
    model := []openapiclient.CreateAuditLogRequestDTO{*openapiclient.NewCreateAuditLogRequestDTO("Category_example", "ServiceCreated_example", "Subject_example")} // []CreateAuditLogRequestDTO | The audit log DTOs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.CreateAuditLogsV1Int(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.CreateAuditLogsV1Int``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuditLogsV1IntRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**[]CreateAuditLogRequestDTO**](CreateAuditLogRequestDTO.md) | The audit log DTOs | 

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


## GetAuditLogByIDV1

> AuditLogDTO GetAuditLogByIDV1(ctx, id).Execute()

Get audit log by ID



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
    id := "id_example" // string | audit log ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.GetAuditLogByIDV1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetAuditLogByIDV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogByIDV1`: AuditLogDTO
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetAuditLogByIDV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | audit log ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogByIDV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLogDTO**](AuditLogDTO.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)


## GetAuditLogsV1

> AuditLogListPageResponse GetAuditLogsV1(ctx).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()

Get audit logs



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
    resp, r, err := apiClient.AuditLogApi.GetAuditLogsV1(context.Background()).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetAuditLogsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogsV1`: AuditLogListPageResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetAuditLogsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The desired page number | [default to 0]
 **pageSize** | **int32** | The desired number of items per page | [default to 25]
 **filter** | **string** | The search term | 
 **direction** | **string** | The sorting direction | 
 **sort** | **string** | The sorting parameter | 

### Return type

[**AuditLogListPageResponse**](AuditLogListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)

