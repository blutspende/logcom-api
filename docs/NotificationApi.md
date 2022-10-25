# \NotificationApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationV1**](NotificationApi.md#CreateNotificationV1) | **Post** /v1/notifications | Create notification
[**GetNotificationsV1**](NotificationApi.md#GetNotificationsV1) | **Get** /v1/notifications | Get notifications


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

# **GetNotificationsV1**
> NotificationListPageResponse GetNotificationsV1(ctx, page, pageSize, optional)
Get notifications

Gets all notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| The desired page number | [default to 0]
  **pageSize** | **int32**| The desired number of items per page | [default to 25]
 **optional** | ***NotificationApiGetNotificationsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationApiGetNotificationsV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchTerm** | **optional.String**| The search term | 
 **direction** | **optional.String**| The sorting direction | 
 **sort** | **optional.String**| The sorting parameter | 

### Return type

[**NotificationListPageResponse**](NotificationListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

