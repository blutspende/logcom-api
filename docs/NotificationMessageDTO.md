# NotificationMessageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The notification timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The ID | [optional] 
**Message** | Pointer to **string** | The notification message | [optional] 
**Service** | Pointer to **string** | The service which sent the notification | [optional] 
**Status** | Pointer to [**NotificationStatus**](NotificationStatus.md) | The status of the message | [optional] 

## Methods

### NewNotificationMessageDTO

`func NewNotificationMessageDTO() *NotificationMessageDTO`

NewNotificationMessageDTO instantiates a new NotificationMessageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessageDTOWithDefaults

`func NewNotificationMessageDTOWithDefaults() *NotificationMessageDTO`

NewNotificationMessageDTOWithDefaults instantiates a new NotificationMessageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NotificationMessageDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationMessageDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationMessageDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationMessageDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *NotificationMessageDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *NotificationMessageDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *NotificationMessageDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *NotificationMessageDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *NotificationMessageDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *NotificationMessageDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *NotificationMessageDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *NotificationMessageDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetId

`func (o *NotificationMessageDTO) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationMessageDTO) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationMessageDTO) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationMessageDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *NotificationMessageDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationMessageDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationMessageDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationMessageDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetService

`func (o *NotificationMessageDTO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NotificationMessageDTO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NotificationMessageDTO) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *NotificationMessageDTO) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationMessageDTO) GetStatus() NotificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationMessageDTO) GetStatusOk() (*NotificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationMessageDTO) SetStatus(v NotificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationMessageDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


