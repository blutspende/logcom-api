# \NotificationApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationV1**](NotificationApi.md#CreateNotificationV1) | **Post** /v1/notifications | Create notification
[**GetNotificationMessagesV1**](NotificationApi.md#GetNotificationMessagesV1) | **Get** /v1/notification-messages | Get notification messages
[**UpdateNotificationMessageV1**](NotificationApi.md#UpdateNotificationMessageV1) | **Patch** /v1/notification-messages/{id} | Update notification message


# **CreateNotificationV1**
> CreateNotificationV1(ctx, model)
Create notification

Creates a new notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**CreateNotificationRequestDto**](CreateNotificationRequestDto.md)| The notification DTO | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationMessagesV1**
> NotificationMessageListPageResponse GetNotificationMessagesV1(ctx, page, pageSize, optional)
Get notification messages

Gets all notification message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| The desired page number | [default to 0]
  **pageSize** | **int32**| The desired number of items per page | [default to 25]
 **optional** | ***NotificationApiGetNotificationMessagesV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationApiGetNotificationMessagesV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**| The search term | 
 **direction** | **optional.String**| The sorting direction | 
 **sort** | **optional.String**| The sorting parameter | 

### Return type

[**NotificationMessageListPageResponse**](NotificationMessageListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationMessageV1**
> UpdateNotificationMessageV1(ctx, id, model)
Update notification message

Updates a notification message by the optionally provided properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**uuid.UUID**](.md)| Message ID | 
  **model** | [**UpdateNotificationMessageRequestDto**](UpdateNotificationMessageRequestDto.md)| The updated notification DTO | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

