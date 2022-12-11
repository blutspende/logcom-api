# ConsoleLogDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The log timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The ID | [optional] 
**Level** | Pointer to [**LogLevel**](LogLevel.md) | The log level (Trace&#x3D;-1, Debug&#x3D;0, Info&#x3D;1, Warning&#x3D;2, Error&#x3D;3, Fatal&#x3D;4, Panic&#x3D;5) | [optional] 
**Message** | Pointer to **string** | The log message | [optional] 
**RequestId** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The request ID making dependent logs trackable | [optional] 
**Service** | Pointer to **string** | The service which sent the log | [optional] 

## Methods

### NewConsoleLogDTO

`func NewConsoleLogDTO() *ConsoleLogDTO`

NewConsoleLogDTO instantiates a new ConsoleLogDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleLogDTOWithDefaults

`func NewConsoleLogDTOWithDefaults() *ConsoleLogDTO`

NewConsoleLogDTOWithDefaults instantiates a new ConsoleLogDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ConsoleLogDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsoleLogDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsoleLogDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsoleLogDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *ConsoleLogDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConsoleLogDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConsoleLogDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *ConsoleLogDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *ConsoleLogDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *ConsoleLogDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *ConsoleLogDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *ConsoleLogDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetId

`func (o *ConsoleLogDTO) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleLogDTO) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleLogDTO) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleLogDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *ConsoleLogDTO) GetLevel() LogLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ConsoleLogDTO) GetLevelOk() (*LogLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ConsoleLogDTO) SetLevel(v LogLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ConsoleLogDTO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *ConsoleLogDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConsoleLogDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConsoleLogDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConsoleLogDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRequestId

`func (o *ConsoleLogDTO) GetRequestId() uuid.UUID`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ConsoleLogDTO) GetRequestIdOk() (*uuid.UUID, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ConsoleLogDTO) SetRequestId(v uuid.UUID)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ConsoleLogDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetService

`func (o *ConsoleLogDTO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ConsoleLogDTO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ConsoleLogDTO) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ConsoleLogDTO) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


