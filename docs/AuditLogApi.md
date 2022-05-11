# \AuditLogApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuditLogV1Int**](AuditLogApi.md#CreateAuditLogV1Int) | **Post** /v1/int/audit-logs | Create audit log
[**GetAuditLogsV1**](AuditLogApi.md#GetAuditLogsV1) | **Get** /v1/audit-logs | Get audit logs


# **CreateAuditLogV1Int**
> CreateAuditLogV1Int(ctx, model)
Create audit log

Creates a new audit log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**CreateAuditLogRequestDto**](CreateAuditLogRequestDto.md)| The audit log DTO | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditLogsV1**
> AuditLogListPageResponse GetAuditLogsV1(ctx, page, pageSize, optional)
Get audit logs

Gets all audit log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| The desired page number | [default to 0]
  **pageSize** | **int32**| The desired number of items per page | [default to 25]
 **optional** | ***AuditLogApiGetAuditLogsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditLogApiGetAuditLogsV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchTerm** | **optional.String**| The search term | 
 **direction** | **optional.String**| The sorting direction | 
 **sort** | **optional.String**| The sorting parameter | 

### Return type

[**AuditLogListPageResponse**](AuditLogListPageResponse.md)

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

