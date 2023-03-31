# \NotificationApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationV1**](NotificationApi.md#CreateNotificationV1) | **Post** /v1/notifications | Create notification
[**GetNotificationMessagesV1**](NotificationApi.md#GetNotificationMessagesV1) | **Get** /v1/notification-messages | Get notification messages
[**PeekNotificationMessagesV1**](NotificationApi.md#PeekNotificationMessagesV1) | **Get** /v1/notification-messages/peek | Peek notification messages
[**UpdateNotificationMessageV1**](NotificationApi.md#UpdateNotificationMessageV1) | **Patch** /v1/notification-messages/{id} | Update notification message



## CreateNotificationV1

> CreateNotificationV1(ctx).Model(model).Execute()

Create notification



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
    model := *openapiclient.NewCreateNotificationRequestDTO("Message_example", "Service_example") // CreateNotificationRequestDTO | The notification DTO

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationApi.CreateNotificationV1(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.CreateNotificationV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**CreateNotificationRequestDTO**](CreateNotificationRequestDTO.md) | The notification DTO | 

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


## GetNotificationMessagesV1

> NotificationMessageListPageResponse GetNotificationMessagesV1(ctx).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()

Get notification messages



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
    resp, r, err := apiClient.NotificationApi.GetNotificationMessagesV1(context.Background()).Page(page).PageSize(pageSize).Filter(filter).Direction(direction).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.GetNotificationMessagesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationMessagesV1`: NotificationMessageListPageResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.GetNotificationMessagesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationMessagesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The desired page number | [default to 0]
 **pageSize** | **int32** | The desired number of items per page | [default to 25]
 **filter** | **string** | The search term | 
 **direction** | **string** | The sorting direction | 
 **sort** | **string** | The sorting parameter | 

### Return type

[**NotificationMessageListPageResponse**](NotificationMessageListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)


## PeekNotificationMessagesV1

> NotificationMessageListPageResponse PeekNotificationMessagesV1(ctx).Execute()

Peek notification messages



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
    resp, r, err := apiClient.NotificationApi.PeekNotificationMessagesV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.PeekNotificationMessagesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PeekNotificationMessagesV1`: NotificationMessageListPageResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.PeekNotificationMessagesV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPeekNotificationMessagesV1Request struct via the builder pattern


### Return type

[**NotificationMessageListPageResponse**](NotificationMessageListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints)
[[Back to Model list]](README.md#documentation-for-models)
[[Back to README]](README.md)


## UpdateNotificationMessageV1

> UpdateNotificationMessageV1(ctx, id).Model(model).Execute()

Update notification message



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Message ID
    model := *openapiclient.NewUpdateNotificationMessageRequestDTO() // UpdateNotificationMessageRequestDTO | The updated notification DTO

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationApi.UpdateNotificationMessageV1(context.Background(), id).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.UpdateNotificationMessageV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **uuid.UUID** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationMessageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**UpdateNotificationMessageRequestDTO**](UpdateNotificationMessageRequestDTO.md) | The updated notification DTO | 

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

