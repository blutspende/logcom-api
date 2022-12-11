# AuditLogDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the change | [optional] 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**CreatedByName** | Pointer to **string** | The user&#39;s name who created | [optional] 
**GroupedChanges** | Pointer to [**[]AuditLogChangeDTO**](AuditLogChangeDTO.md) | The grouped changes | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The ID | [optional] 
**Message** | Pointer to **string** | The message | [optional] 
**NewValue** | Pointer to **string** | The new value | [optional] 
**OldValue** | Pointer to **string** | The old value | [optional] 
**RequestId** | Pointer to **string** | The request ID making dependent logs trackable | [optional] 
**ServiceAffected** | Pointer to **string** | The service which the change affects | [optional] 
**ServiceCreated** | Pointer to **string** | The service which created | [optional] 
**Subject** | Pointer to **string** | The subject of the change | [optional] 
**SubjectName** | Pointer to **string** | The name of the subject | [optional] 
**SubjectPropertyName** | Pointer to **string** | The property name of the subject | [optional] 

## Methods

### NewAuditLogDTO

`func NewAuditLogDTO() *AuditLogDTO`

NewAuditLogDTO instantiates a new AuditLogDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogDTOWithDefaults

`func NewAuditLogDTOWithDefaults() *AuditLogDTO`

NewAuditLogDTOWithDefaults instantiates a new AuditLogDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AuditLogDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditLogDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditLogDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AuditLogDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditLogDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLogDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLogDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditLogDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *AuditLogDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AuditLogDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AuditLogDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *AuditLogDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *AuditLogDTO) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *AuditLogDTO) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *AuditLogDTO) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *AuditLogDTO) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetGroupedChanges

`func (o *AuditLogDTO) GetGroupedChanges() []AuditLogChangeDTO`

GetGroupedChanges returns the GroupedChanges field if non-nil, zero value otherwise.

### GetGroupedChangesOk

`func (o *AuditLogDTO) GetGroupedChangesOk() (*[]AuditLogChangeDTO, bool)`

GetGroupedChangesOk returns a tuple with the GroupedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedChanges

`func (o *AuditLogDTO) SetGroupedChanges(v []AuditLogChangeDTO)`

SetGroupedChanges sets GroupedChanges field to given value.

### HasGroupedChanges

`func (o *AuditLogDTO) HasGroupedChanges() bool`

HasGroupedChanges returns a boolean if a field has been set.

### GetId

`func (o *AuditLogDTO) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogDTO) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogDTO) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AuditLogDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditLogDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewValue

`func (o *AuditLogDTO) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AuditLogDTO) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AuditLogDTO) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AuditLogDTO) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *AuditLogDTO) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *AuditLogDTO) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *AuditLogDTO) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *AuditLogDTO) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetRequestId

`func (o *AuditLogDTO) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AuditLogDTO) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AuditLogDTO) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AuditLogDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetServiceAffected

`func (o *AuditLogDTO) GetServiceAffected() string`

GetServiceAffected returns the ServiceAffected field if non-nil, zero value otherwise.

### GetServiceAffectedOk

`func (o *AuditLogDTO) GetServiceAffectedOk() (*string, bool)`

GetServiceAffectedOk returns a tuple with the ServiceAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAffected

`func (o *AuditLogDTO) SetServiceAffected(v string)`

SetServiceAffected sets ServiceAffected field to given value.

### HasServiceAffected

`func (o *AuditLogDTO) HasServiceAffected() bool`

HasServiceAffected returns a boolean if a field has been set.

### GetServiceCreated

`func (o *AuditLogDTO) GetServiceCreated() string`

GetServiceCreated returns the ServiceCreated field if non-nil, zero value otherwise.

### GetServiceCreatedOk

`func (o *AuditLogDTO) GetServiceCreatedOk() (*string, bool)`

GetServiceCreatedOk returns a tuple with the ServiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCreated

`func (o *AuditLogDTO) SetServiceCreated(v string)`

SetServiceCreated sets ServiceCreated field to given value.

### HasServiceCreated

`func (o *AuditLogDTO) HasServiceCreated() bool`

HasServiceCreated returns a boolean if a field has been set.

### GetSubject

`func (o *AuditLogDTO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogDTO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogDTO) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogDTO) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectName

`func (o *AuditLogDTO) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *AuditLogDTO) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *AuditLogDTO) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *AuditLogDTO) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSubjectPropertyName

`func (o *AuditLogDTO) GetSubjectPropertyName() string`

GetSubjectPropertyName returns the SubjectPropertyName field if non-nil, zero value otherwise.

### GetSubjectPropertyNameOk

`func (o *AuditLogDTO) GetSubjectPropertyNameOk() (*string, bool)`

GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPropertyName

`func (o *AuditLogDTO) SetSubjectPropertyName(v string)`

SetSubjectPropertyName sets SubjectPropertyName field to given value.

### HasSubjectPropertyName

`func (o *AuditLogDTO) HasSubjectPropertyName() bool`

HasSubjectPropertyName returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


