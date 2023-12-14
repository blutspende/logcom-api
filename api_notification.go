/*
LogCom API

LogCom Swagger documentation

API version: 1.2.27
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
	"github.com/google/uuid"
	"strings"
)


// NotificationApiService NotificationApi service
type NotificationApiService service

type ApiCreateNotificationV1Request struct {
	ctx context.Context
	ApiService *NotificationApiService
	model *CreateNotificationRequestDTO
}

// The notification DTO
func (r ApiCreateNotificationV1Request) Model(model CreateNotificationRequestDTO) ApiCreateNotificationV1Request {
	r.model = &model
	return r
}

func (r ApiCreateNotificationV1Request) Execute() (*http.Response, error) {
	return r.ApiService.CreateNotificationV1Execute(r)
}

/*
CreateNotificationV1 Create notification

Creates a new notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNotificationV1Request
*/
func (a *NotificationApiService) CreateNotificationV1(ctx context.Context) ApiCreateNotificationV1Request {
	return ApiCreateNotificationV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationApiService) CreateNotificationV1Execute(r ApiCreateNotificationV1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.CreateNotificationV1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/notifications"

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

type ApiGetNotificationMessagesV1Request struct {
	ctx context.Context
	ApiService *NotificationApiService
	page *int32
	pageSize *int32
	filter *string
	direction *string
	sort *string
}

// The desired page number
func (r ApiGetNotificationMessagesV1Request) Page(page int32) ApiGetNotificationMessagesV1Request {
	r.page = &page
	return r
}

// The desired number of items per page
func (r ApiGetNotificationMessagesV1Request) PageSize(pageSize int32) ApiGetNotificationMessagesV1Request {
	r.pageSize = &pageSize
	return r
}

// The search term
func (r ApiGetNotificationMessagesV1Request) Filter(filter string) ApiGetNotificationMessagesV1Request {
	r.filter = &filter
	return r
}

// The sorting direction
func (r ApiGetNotificationMessagesV1Request) Direction(direction string) ApiGetNotificationMessagesV1Request {
	r.direction = &direction
	return r
}

// The sorting parameter
func (r ApiGetNotificationMessagesV1Request) Sort(sort string) ApiGetNotificationMessagesV1Request {
	r.sort = &sort
	return r
}

func (r ApiGetNotificationMessagesV1Request) Execute() (*NotificationMessageListPageResponse, *http.Response, error) {
	return r.ApiService.GetNotificationMessagesV1Execute(r)
}

/*
GetNotificationMessagesV1 Get notification messages

Gets all notification message

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNotificationMessagesV1Request
*/
func (a *NotificationApiService) GetNotificationMessagesV1(ctx context.Context) ApiGetNotificationMessagesV1Request {
	return ApiGetNotificationMessagesV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotificationMessageListPageResponse
func (a *NotificationApiService) GetNotificationMessagesV1Execute(r ApiGetNotificationMessagesV1Request) (*NotificationMessageListPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationMessageListPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.GetNotificationMessagesV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/notification-messages"

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

type ApiPeekNotificationMessagesV1Request struct {
	ctx context.Context
	ApiService *NotificationApiService
}

func (r ApiPeekNotificationMessagesV1Request) Execute() (*NotificationMessageListPageResponse, *http.Response, error) {
	return r.ApiService.PeekNotificationMessagesV1Execute(r)
}

/*
PeekNotificationMessagesV1 Peek notification messages

Gets portion of the notification messages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPeekNotificationMessagesV1Request
*/
func (a *NotificationApiService) PeekNotificationMessagesV1(ctx context.Context) ApiPeekNotificationMessagesV1Request {
	return ApiPeekNotificationMessagesV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotificationMessageListPageResponse
func (a *NotificationApiService) PeekNotificationMessagesV1Execute(r ApiPeekNotificationMessagesV1Request) (*NotificationMessageListPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationMessageListPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.PeekNotificationMessagesV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/notification-messages/peek"

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

type ApiUpdateNotificationMessageV1Request struct {
	ctx context.Context
	ApiService *NotificationApiService
	id uuid.UUID
	model *UpdateNotificationMessageRequestDTO
}

// The updated notification DTO
func (r ApiUpdateNotificationMessageV1Request) Model(model UpdateNotificationMessageRequestDTO) ApiUpdateNotificationMessageV1Request {
	r.model = &model
	return r
}

func (r ApiUpdateNotificationMessageV1Request) Execute() (*http.Response, error) {
	return r.ApiService.UpdateNotificationMessageV1Execute(r)
}

/*
UpdateNotificationMessageV1 Update notification message

Updates a notification message by the optionally provided properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Message ID
 @return ApiUpdateNotificationMessageV1Request
*/
func (a *NotificationApiService) UpdateNotificationMessageV1(ctx context.Context, id uuid.UUID) ApiUpdateNotificationMessageV1Request {
	return ApiUpdateNotificationMessageV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *NotificationApiService) UpdateNotificationMessageV1Execute(r ApiUpdateNotificationMessageV1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.UpdateNotificationMessageV1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/notification-messages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
