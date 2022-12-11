# CreateAuditLogRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the change | 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**GroupedChanges** | Pointer to [**[]CreateAuditLogChangeDTO**](CreateAuditLogChangeDTO.md) | The grouped changes if there are many related ones | [optional] 
**Message** | Pointer to **string** | The message | [optional] 
**NewValue** | Pointer to **string** | The new value | [optional] 
**OldValue** | Pointer to **string** | The old value | [optional] 
**ServiceAffected** | Pointer to **string** | The service which the change affects | [optional] 
**ServiceCreated** | **string** | The service which created | 
**Subject** | **string** | The subject of the change | 
**SubjectName** | Pointer to **string** | The name of the subject | [optional] 
**SubjectPropertyName** | Pointer to **string** | The property name of the subject | [optional] 

## Methods

### NewCreateAuditLogRequestDTO

`func NewCreateAuditLogRequestDTO(category string, serviceCreated string, subject string, ) *CreateAuditLogRequestDTO`

NewCreateAuditLogRequestDTO instantiates a new CreateAuditLogRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuditLogRequestDTOWithDefaults

`func NewCreateAuditLogRequestDTOWithDefaults() *CreateAuditLogRequestDTO`

NewCreateAuditLogRequestDTOWithDefaults instantiates a new CreateAuditLogRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CreateAuditLogRequestDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateAuditLogRequestDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateAuditLogRequestDTO) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCreatedAt

`func (o *CreateAuditLogRequestDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateAuditLogRequestDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateAuditLogRequestDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateAuditLogRequestDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *CreateAuditLogRequestDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *CreateAuditLogRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *CreateAuditLogRequestDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *CreateAuditLogRequestDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *CreateAuditLogRequestDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *CreateAuditLogRequestDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *CreateAuditLogRequestDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *CreateAuditLogRequestDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetGroupedChanges

`func (o *CreateAuditLogRequestDTO) GetGroupedChanges() []CreateAuditLogChangeDTO`

GetGroupedChanges returns the GroupedChanges field if non-nil, zero value otherwise.

### GetGroupedChangesOk

`func (o *CreateAuditLogRequestDTO) GetGroupedChangesOk() (*[]CreateAuditLogChangeDTO, bool)`

GetGroupedChangesOk returns a tuple with the GroupedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedChanges

`func (o *CreateAuditLogRequestDTO) SetGroupedChanges(v []CreateAuditLogChangeDTO)`

SetGroupedChanges sets GroupedChanges field to given value.

### HasGroupedChanges

`func (o *CreateAuditLogRequestDTO) HasGroupedChanges() bool`

HasGroupedChanges returns a boolean if a field has been set.

### GetMessage

`func (o *CreateAuditLogRequestDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateAuditLogRequestDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateAuditLogRequestDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateAuditLogRequestDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewValue

`func (o *CreateAuditLogRequestDTO) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *CreateAuditLogRequestDTO) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *CreateAuditLogRequestDTO) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *CreateAuditLogRequestDTO) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *CreateAuditLogRequestDTO) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *CreateAuditLogRequestDTO) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *CreateAuditLogRequestDTO) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *CreateAuditLogRequestDTO) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetServiceAffected

`func (o *CreateAuditLogRequestDTO) GetServiceAffected() string`

GetServiceAffected returns the ServiceAffected field if non-nil, zero value otherwise.

### GetServiceAffectedOk

`func (o *CreateAuditLogRequestDTO) GetServiceAffectedOk() (*string, bool)`

GetServiceAffectedOk returns a tuple with the ServiceAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAffected

`func (o *CreateAuditLogRequestDTO) SetServiceAffected(v string)`

SetServiceAffected sets ServiceAffected field to given value.

### HasServiceAffected

`func (o *CreateAuditLogRequestDTO) HasServiceAffected() bool`

HasServiceAffected returns a boolean if a field has been set.

### GetServiceCreated

`func (o *CreateAuditLogRequestDTO) GetServiceCreated() string`

GetServiceCreated returns the ServiceCreated field if non-nil, zero value otherwise.

### GetServiceCreatedOk

`func (o *CreateAuditLogRequestDTO) GetServiceCreatedOk() (*string, bool)`

GetServiceCreatedOk returns a tuple with the ServiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCreated

`func (o *CreateAuditLogRequestDTO) SetServiceCreated(v string)`

SetServiceCreated sets ServiceCreated field to given value.


### GetSubject

`func (o *CreateAuditLogRequestDTO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateAuditLogRequestDTO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateAuditLogRequestDTO) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSubjectName

`func (o *CreateAuditLogRequestDTO) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *CreateAuditLogRequestDTO) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *CreateAuditLogRequestDTO) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *CreateAuditLogRequestDTO) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSubjectPropertyName

`func (o *CreateAuditLogRequestDTO) GetSubjectPropertyName() string`

GetSubjectPropertyName returns the SubjectPropertyName field if non-nil, zero value otherwise.

### GetSubjectPropertyNameOk

`func (o *CreateAuditLogRequestDTO) GetSubjectPropertyNameOk() (*string, bool)`

GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPropertyName

`func (o *CreateAuditLogRequestDTO) SetSubjectPropertyName(v string)`

SetSubjectPropertyName sets SubjectPropertyName field to given value.

### HasSubjectPropertyName

`func (o *CreateAuditLogRequestDTO) HasSubjectPropertyName() bool`

HasSubjectPropertyName returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


