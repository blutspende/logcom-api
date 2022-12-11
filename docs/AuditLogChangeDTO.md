# AuditLogChangeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the change | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The ID | [optional] 
**Message** | Pointer to **string** | The message | [optional] 
**NewValue** | Pointer to **string** | The new value | [optional] 
**OldValue** | Pointer to **string** | The old value | [optional] 
**Subject** | Pointer to **string** | The short description of the change | [optional] 
**SubjectName** | Pointer to **string** | The name of the subject | [optional] 
**SubjectPropertyName** | Pointer to **string** | The property name of the subject | [optional] 

## Methods

### NewAuditLogChangeDTO

`func NewAuditLogChangeDTO() *AuditLogChangeDTO`

NewAuditLogChangeDTO instantiates a new AuditLogChangeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogChangeDTOWithDefaults

`func NewAuditLogChangeDTOWithDefaults() *AuditLogChangeDTO`

NewAuditLogChangeDTOWithDefaults instantiates a new AuditLogChangeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AuditLogChangeDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditLogChangeDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditLogChangeDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AuditLogChangeDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *AuditLogChangeDTO) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogChangeDTO) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogChangeDTO) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogChangeDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AuditLogChangeDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogChangeDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogChangeDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditLogChangeDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewValue

`func (o *AuditLogChangeDTO) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AuditLogChangeDTO) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AuditLogChangeDTO) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AuditLogChangeDTO) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *AuditLogChangeDTO) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *AuditLogChangeDTO) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *AuditLogChangeDTO) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *AuditLogChangeDTO) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetSubject

`func (o *AuditLogChangeDTO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogChangeDTO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogChangeDTO) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogChangeDTO) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectName

`func (o *AuditLogChangeDTO) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *AuditLogChangeDTO) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *AuditLogChangeDTO) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *AuditLogChangeDTO) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSubjectPropertyName

`func (o *AuditLogChangeDTO) GetSubjectPropertyName() string`

GetSubjectPropertyName returns the SubjectPropertyName field if non-nil, zero value otherwise.

### GetSubjectPropertyNameOk

`func (o *AuditLogChangeDTO) GetSubjectPropertyNameOk() (*string, bool)`

GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPropertyName

`func (o *AuditLogChangeDTO) SetSubjectPropertyName(v string)`

SetSubjectPropertyName sets SubjectPropertyName field to given value.

### HasSubjectPropertyName

`func (o *AuditLogChangeDTO) HasSubjectPropertyName() bool`

HasSubjectPropertyName returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


