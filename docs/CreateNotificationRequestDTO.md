# CreateNotificationRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The log timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**EventCategory** | Pointer to [**NotificationEventCategory**](NotificationEventCategory.md) | The notification event category | [optional] 
**Message** | **string** | The log message | 
**Service** | **string** | The service which sent the log | 
**Targets** | Pointer to **map[string][]string** | The targets who will receive the notification | [optional] 

## Methods

### NewCreateNotificationRequestDTO

`func NewCreateNotificationRequestDTO(message string, service string, ) *CreateNotificationRequestDTO`

NewCreateNotificationRequestDTO instantiates a new CreateNotificationRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRequestDTOWithDefaults

`func NewCreateNotificationRequestDTOWithDefaults() *CreateNotificationRequestDTO`

NewCreateNotificationRequestDTOWithDefaults instantiates a new CreateNotificationRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreateNotificationRequestDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateNotificationRequestDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateNotificationRequestDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateNotificationRequestDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *CreateNotificationRequestDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *CreateNotificationRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *CreateNotificationRequestDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *CreateNotificationRequestDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *CreateNotificationRequestDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *CreateNotificationRequestDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *CreateNotificationRequestDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *CreateNotificationRequestDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetEventCategory

`func (o *CreateNotificationRequestDTO) GetEventCategory() NotificationEventCategory`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *CreateNotificationRequestDTO) GetEventCategoryOk() (*NotificationEventCategory, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *CreateNotificationRequestDTO) SetEventCategory(v NotificationEventCategory)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *CreateNotificationRequestDTO) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### GetMessage

`func (o *CreateNotificationRequestDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateNotificationRequestDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateNotificationRequestDTO) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetService

`func (o *CreateNotificationRequestDTO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateNotificationRequestDTO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateNotificationRequestDTO) SetService(v string)`

SetService sets Service field to given value.


### GetTargets

`func (o *CreateNotificationRequestDTO) GetTargets() map[string][]string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CreateNotificationRequestDTO) GetTargetsOk() (*map[string][]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CreateNotificationRequestDTO) SetTargets(v map[string][]string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CreateNotificationRequestDTO) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


