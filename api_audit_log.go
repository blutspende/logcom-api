/*
LogCom API

LogCom Swagger documentation

API version: 1.3.6
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AuditLogApiService AuditLogApi service
type AuditLogApiService service

type ApiCreateAuditLogV1IntRequest struct {
	ctx context.Context
	ApiService *AuditLogApiService
	model *CreateAuditLogRequestDTO
}

// The audit log DTO
func (r ApiCreateAuditLogV1IntRequest) Model(model CreateAuditLogRequestDTO) ApiCreateAuditLogV1IntRequest {
	r.model = &model
	return r
}

func (r ApiCreateAuditLogV1IntRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateAuditLogV1IntExecute(r)
}

/*
CreateAuditLogV1Int Create audit log

Creates a new audit log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAuditLogV1IntRequest
*/
func (a *AuditLogApiService) CreateAuditLogV1Int(ctx context.Context) ApiCreateAuditLogV1IntRequest {
	return ApiCreateAuditLogV1IntRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuditLogApiService) CreateAuditLogV1IntExecute(r ApiCreateAuditLogV1IntRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogApiService.CreateAuditLogV1Int")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/int/audit-logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateAuditLogsV1IntRequest struct {
	ctx context.Context
	ApiService *AuditLogApiService
	model *[]CreateAuditLogRequestDTO
}

// The audit log DTOs
func (r ApiCreateAuditLogsV1IntRequest) Model(model []CreateAuditLogRequestDTO) ApiCreateAuditLogsV1IntRequest {
	r.model = &model
	return r
}

func (r ApiCreateAuditLogsV1IntRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateAuditLogsV1IntExecute(r)
}

/*
CreateAuditLogsV1Int Create audit logs

Creates new audit logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAuditLogsV1IntRequest
*/
func (a *AuditLogApiService) CreateAuditLogsV1Int(ctx context.Context) ApiCreateAuditLogsV1IntRequest {
	return ApiCreateAuditLogsV1IntRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuditLogApiService) CreateAuditLogsV1IntExecute(r ApiCreateAuditLogsV1IntRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogApiService.CreateAuditLogsV1Int")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/int/audit-logs/batch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAuditLogByIDV1Request struct {
	ctx context.Context
	ApiService *AuditLogApiService
	id string
}

func (r ApiGetAuditLogByIDV1Request) Execute() (*AuditLogDTO, *http.Response, error) {
	return r.ApiService.GetAuditLogByIDV1Execute(r)
}

/*
GetAuditLogByIDV1 Get audit log by ID

Gets an audit log by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id audit log ID
 @return ApiGetAuditLogByIDV1Request
*/
func (a *AuditLogApiService) GetAuditLogByIDV1(ctx context.Context, id string) ApiGetAuditLogByIDV1Request {
	return ApiGetAuditLogByIDV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuditLogDTO
func (a *AuditLogApiService) GetAuditLogByIDV1Execute(r ApiGetAuditLogByIDV1Request) (*AuditLogDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditLogDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogApiService.GetAuditLogByIDV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/audit-logs/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAuditLogsV1Request struct {
	ctx context.Context
	ApiService *AuditLogApiService
	page *int32
	pageSize *int32
	filter *string
	direction *string
	sort *string
}

// The desired page number
func (r ApiGetAuditLogsV1Request) Page(page int32) ApiGetAuditLogsV1Request {
	r.page = &page
	return r
}

// The desired number of items per page
func (r ApiGetAuditLogsV1Request) PageSize(pageSize int32) ApiGetAuditLogsV1Request {
	r.pageSize = &pageSize
	return r
}

// The search term
func (r ApiGetAuditLogsV1Request) Filter(filter string) ApiGetAuditLogsV1Request {
	r.filter = &filter
	return r
}

// The sorting direction
func (r ApiGetAuditLogsV1Request) Direction(direction string) ApiGetAuditLogsV1Request {
	r.direction = &direction
	return r
}

// The sorting parameter
func (r ApiGetAuditLogsV1Request) Sort(sort string) ApiGetAuditLogsV1Request {
	r.sort = &sort
	return r
}

func (r ApiGetAuditLogsV1Request) Execute() (*AuditLogListPageResponse, *http.Response, error) {
	return r.ApiService.GetAuditLogsV1Execute(r)
}

/*
GetAuditLogsV1 Get audit logs

Gets all audit log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAuditLogsV1Request
*/
func (a *AuditLogApiService) GetAuditLogsV1(ctx context.Context) ApiGetAuditLogsV1Request {
	return ApiGetAuditLogsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditLogListPageResponse
func (a *AuditLogApiService) GetAuditLogsV1Execute(r ApiGetAuditLogsV1Request) (*AuditLogListPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditLogListPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogApiService.GetAuditLogsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/audit-logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 0 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 0")
	}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
