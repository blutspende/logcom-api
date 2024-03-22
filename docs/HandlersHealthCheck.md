# HandlersHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **[]string** | The API version | [optional] 
**BuildVersion** | Pointer to **string** | Docker build version | [optional] 
**MemStats** | Pointer to [**HandlersMemStats**](HandlersMemStats.md) |  | [optional] 
**Service** | Pointer to **string** | The service name | [optional] 
**Status** | Pointer to **string** | The service status | [optional] 

## Methods

### NewHandlersHealthCheck

`func NewHandlersHealthCheck() *HandlersHealthCheck`

NewHandlersHealthCheck instantiates a new HandlersHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersHealthCheckWithDefaults

`func NewHandlersHealthCheckWithDefaults() *HandlersHealthCheck`

NewHandlersHealthCheckWithDefaults instantiates a new HandlersHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HandlersHealthCheck) GetApiVersion() []string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HandlersHealthCheck) GetApiVersionOk() (*[]string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HandlersHealthCheck) SetApiVersion(v []string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HandlersHealthCheck) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBuildVersion

`func (o *HandlersHealthCheck) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *HandlersHealthCheck) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *HandlersHealthCheck) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *HandlersHealthCheck) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetMemStats

`func (o *HandlersHealthCheck) GetMemStats() HandlersMemStats`

GetMemStats returns the MemStats field if non-nil, zero value otherwise.

### GetMemStatsOk

`func (o *HandlersHealthCheck) GetMemStatsOk() (*HandlersMemStats, bool)`

GetMemStatsOk returns a tuple with the MemStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStats

`func (o *HandlersHealthCheck) SetMemStats(v HandlersMemStats)`

SetMemStats sets MemStats field to given value.

### HasMemStats

`func (o *HandlersHealthCheck) HasMemStats() bool`

HasMemStats returns a boolean if a field has been set.

### GetService

`func (o *HandlersHealthCheck) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HandlersHealthCheck) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HandlersHealthCheck) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HandlersHealthCheck) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersHealthCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersHealthCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersHealthCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersHealthCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


