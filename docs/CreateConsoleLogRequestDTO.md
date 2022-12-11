# CreateConsoleLogRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The log timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**Level** | [**LogLevel**](LogLevel.md) | The log level (Trace&#x3D;-1, Debug&#x3D;0, Info&#x3D;1, Warning&#x3D;2, Error&#x3D;3, Fatal&#x3D;4, Panic&#x3D;5) | 
**Message** | **string** | The log message | 
**Service** | **string** | The service which sent the log | 

## Methods

### NewCreateConsoleLogRequestDTO

`func NewCreateConsoleLogRequestDTO(level LogLevel, message string, service string, ) *CreateConsoleLogRequestDTO`

NewCreateConsoleLogRequestDTO instantiates a new CreateConsoleLogRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConsoleLogRequestDTOWithDefaults

`func NewCreateConsoleLogRequestDTOWithDefaults() *CreateConsoleLogRequestDTO`

NewCreateConsoleLogRequestDTOWithDefaults instantiates a new CreateConsoleLogRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreateConsoleLogRequestDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateConsoleLogRequestDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateConsoleLogRequestDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateConsoleLogRequestDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *CreateConsoleLogRequestDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *CreateConsoleLogRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *CreateConsoleLogRequestDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *CreateConsoleLogRequestDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *CreateConsoleLogRequestDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *CreateConsoleLogRequestDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *CreateConsoleLogRequestDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *CreateConsoleLogRequestDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetLevel

`func (o *CreateConsoleLogRequestDTO) GetLevel() LogLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CreateConsoleLogRequestDTO) GetLevelOk() (*LogLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CreateConsoleLogRequestDTO) SetLevel(v LogLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *CreateConsoleLogRequestDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateConsoleLogRequestDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateConsoleLogRequestDTO) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetService

`func (o *CreateConsoleLogRequestDTO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateConsoleLogRequestDTO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateConsoleLogRequestDTO) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


