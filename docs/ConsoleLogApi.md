# \ConsoleLogApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsoleLogV1Int**](ConsoleLogApi.md#CreateConsoleLogV1Int) | **Post** /v1/int/console-logs | Create console log
[**GetConsoleLogsV1**](ConsoleLogApi.md#GetConsoleLogsV1) | **Get** /v1/console-logs | Get console logs


# **CreateConsoleLogV1Int**
> CreateConsoleLogV1Int(ctx, model)
Create console log

Creates a new console log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**CreateConsoleLogRequestDto**](CreateConsoleLogRequestDto.md)| The console log DTO | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsoleLogsV1**
> ConsoleLogListPageResponse GetConsoleLogsV1(ctx, page, pageSize, optional)
Get console logs

Gets all console log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| The desired page number | [default to 0]
  **pageSize** | **int32**| The desired number of items per page | [default to 25]
 **optional** | ***ConsoleLogApiGetConsoleLogsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsoleLogApiGetConsoleLogsV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchTerm** | **optional.String**| The search term | 
 **direction** | **optional.String**| The sorting direction | 
 **sort** | **optional.String**| The sorting parameter | 

### Return type

[**ConsoleLogListPageResponse**](ConsoleLogListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

