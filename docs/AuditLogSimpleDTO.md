# AuditLogSimpleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the change | [optional] 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp | [optional] 
**CreatedById** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) | The ID | [optional] 
**Message** | Pointer to **string** | The message | [optional] 
**ServiceAffected** | Pointer to **string** | The service which the change affects | [optional] 
**ServiceCreated** | Pointer to **string** | The service which created | [optional] 
**Subject** | Pointer to **string** | The subject of the change | [optional] 

## Methods

### NewAuditLogSimpleDTO

`func NewAuditLogSimpleDTO() *AuditLogSimpleDTO`

NewAuditLogSimpleDTO instantiates a new AuditLogSimpleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSimpleDTOWithDefaults

`func NewAuditLogSimpleDTOWithDefaults() *AuditLogSimpleDTO`

NewAuditLogSimpleDTOWithDefaults instantiates a new AuditLogSimpleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AuditLogSimpleDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditLogSimpleDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditLogSimpleDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AuditLogSimpleDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditLogSimpleDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLogSimpleDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLogSimpleDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditLogSimpleDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *AuditLogSimpleDTO) GetCreatedById() uuid.UUID`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AuditLogSimpleDTO) GetCreatedByIdOk() (*uuid.UUID, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AuditLogSimpleDTO) SetCreatedById(v uuid.UUID)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *AuditLogSimpleDTO) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetId

`func (o *AuditLogSimpleDTO) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogSimpleDTO) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogSimpleDTO) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogSimpleDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AuditLogSimpleDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogSimpleDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogSimpleDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditLogSimpleDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetServiceAffected

`func (o *AuditLogSimpleDTO) GetServiceAffected() string`

GetServiceAffected returns the ServiceAffected field if non-nil, zero value otherwise.

### GetServiceAffectedOk

`func (o *AuditLogSimpleDTO) GetServiceAffectedOk() (*string, bool)`

GetServiceAffectedOk returns a tuple with the ServiceAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAffected

`func (o *AuditLogSimpleDTO) SetServiceAffected(v string)`

SetServiceAffected sets ServiceAffected field to given value.

### HasServiceAffected

`func (o *AuditLogSimpleDTO) HasServiceAffected() bool`

HasServiceAffected returns a boolean if a field has been set.

### GetServiceCreated

`func (o *AuditLogSimpleDTO) GetServiceCreated() string`

GetServiceCreated returns the ServiceCreated field if non-nil, zero value otherwise.

### GetServiceCreatedOk

`func (o *AuditLogSimpleDTO) GetServiceCreatedOk() (*string, bool)`

GetServiceCreatedOk returns a tuple with the ServiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCreated

`func (o *AuditLogSimpleDTO) SetServiceCreated(v string)`

SetServiceCreated sets ServiceCreated field to given value.

### HasServiceCreated

`func (o *AuditLogSimpleDTO) HasServiceCreated() bool`

HasServiceCreated returns a boolean if a field has been set.

### GetSubject

`func (o *AuditLogSimpleDTO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogSimpleDTO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogSimpleDTO) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogSimpleDTO) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


