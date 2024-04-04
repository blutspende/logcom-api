/*
LogCom API

LogCom Swagger documentation

API version: 1.3.6
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// CreateAuditLogRequestDTO struct for CreateAuditLogRequestDTO
type CreateAuditLogRequestDTO struct {
	// The category of the change
	Category string `json:"category"`
	// The creation timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The grouped changes if there are many related ones
	GroupedChanges []CreateAuditLogChangeDTO `json:"groupedChanges,omitempty"`
	// The message
	Message *string `json:"message,omitempty"`
	// The new value
	NewValue *string `json:"newValue,omitempty"`
	// The old value
	OldValue *string `json:"oldValue,omitempty"`
	// The service which the change affects
	ServiceAffected *string `json:"serviceAffected,omitempty"`
	// The service which created
	ServiceCreated string `json:"serviceCreated"`
	// The subject of the change
	Subject string `json:"subject"`
	// The name of the subject
	SubjectName *string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName *string `json:"subjectPropertyName,omitempty"`
}

// NewCreateAuditLogRequestDTO instantiates a new CreateAuditLogRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuditLogRequestDTO(category string, serviceCreated string, subject string) *CreateAuditLogRequestDTO {
	this := CreateAuditLogRequestDTO{}
	this.Category = category
	this.ServiceCreated = serviceCreated
	this.Subject = subject
	return &this
}

// NewCreateAuditLogRequestDTOWithDefaults instantiates a new CreateAuditLogRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuditLogRequestDTOWithDefaults() *CreateAuditLogRequestDTO {
	this := CreateAuditLogRequestDTO{}
	return &this
}

// GetCategory returns the Category field value
func (o *CreateAuditLogRequestDTO) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CreateAuditLogRequestDTO) SetCategory(v string) {
	o.Category = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateAuditLogRequestDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *CreateAuditLogRequestDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *CreateAuditLogRequestDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetGroupedChanges returns the GroupedChanges field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetGroupedChanges() []CreateAuditLogChangeDTO {
	if o == nil || isNil(o.GroupedChanges) {
		var ret []CreateAuditLogChangeDTO
		return ret
	}
	return o.GroupedChanges
}

// GetGroupedChangesOk returns a tuple with the GroupedChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetGroupedChangesOk() ([]CreateAuditLogChangeDTO, bool) {
	if o == nil || isNil(o.GroupedChanges) {
    return nil, false
	}
	return o.GroupedChanges, true
}

// HasGroupedChanges returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasGroupedChanges() bool {
	if o != nil && !isNil(o.GroupedChanges) {
		return true
	}

	return false
}

// SetGroupedChanges gets a reference to the given []CreateAuditLogChangeDTO and assigns it to the GroupedChanges field.
func (o *CreateAuditLogRequestDTO) SetGroupedChanges(v []CreateAuditLogChangeDTO) {
	o.GroupedChanges = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateAuditLogRequestDTO) SetMessage(v string) {
	o.Message = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetNewValue() string {
	if o == nil || isNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetNewValueOk() (*string, bool) {
	if o == nil || isNil(o.NewValue) {
    return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasNewValue() bool {
	if o != nil && !isNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *CreateAuditLogRequestDTO) SetNewValue(v string) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetOldValue() string {
	if o == nil || isNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetOldValueOk() (*string, bool) {
	if o == nil || isNil(o.OldValue) {
    return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasOldValue() bool {
	if o != nil && !isNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *CreateAuditLogRequestDTO) SetOldValue(v string) {
	o.OldValue = &v
}

// GetServiceAffected returns the ServiceAffected field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetServiceAffected() string {
	if o == nil || isNil(o.ServiceAffected) {
		var ret string
		return ret
	}
	return *o.ServiceAffected
}

// GetServiceAffectedOk returns a tuple with the ServiceAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetServiceAffectedOk() (*string, bool) {
	if o == nil || isNil(o.ServiceAffected) {
    return nil, false
	}
	return o.ServiceAffected, true
}

// HasServiceAffected returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasServiceAffected() bool {
	if o != nil && !isNil(o.ServiceAffected) {
		return true
	}

	return false
}

// SetServiceAffected gets a reference to the given string and assigns it to the ServiceAffected field.
func (o *CreateAuditLogRequestDTO) SetServiceAffected(v string) {
	o.ServiceAffected = &v
}

// GetServiceCreated returns the ServiceCreated field value
func (o *CreateAuditLogRequestDTO) GetServiceCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceCreated
}

// GetServiceCreatedOk returns a tuple with the ServiceCreated field value
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetServiceCreatedOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceCreated, true
}

// SetServiceCreated sets field value
func (o *CreateAuditLogRequestDTO) SetServiceCreated(v string) {
	o.ServiceCreated = v
}

// GetSubject returns the Subject field value
func (o *CreateAuditLogRequestDTO) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetSubjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CreateAuditLogRequestDTO) SetSubject(v string) {
	o.Subject = v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetSubjectName() string {
	if o == nil || isNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetSubjectNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectName) {
    return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasSubjectName() bool {
	if o != nil && !isNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *CreateAuditLogRequestDTO) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetSubjectPropertyName returns the SubjectPropertyName field value if set, zero value otherwise.
func (o *CreateAuditLogRequestDTO) GetSubjectPropertyName() string {
	if o == nil || isNil(o.SubjectPropertyName) {
		var ret string
		return ret
	}
	return *o.SubjectPropertyName
}

// GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogRequestDTO) GetSubjectPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPropertyName) {
    return nil, false
	}
	return o.SubjectPropertyName, true
}

// HasSubjectPropertyName returns a boolean if a field has been set.
func (o *CreateAuditLogRequestDTO) HasSubjectPropertyName() bool {
	if o != nil && !isNil(o.SubjectPropertyName) {
		return true
	}

	return false
}

// SetSubjectPropertyName gets a reference to the given string and assigns it to the SubjectPropertyName field.
func (o *CreateAuditLogRequestDTO) SetSubjectPropertyName(v string) {
	o.SubjectPropertyName = &v
}

func (o CreateAuditLogRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !isNil(o.GroupedChanges) {
		toSerialize["groupedChanges"] = o.GroupedChanges
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	if !isNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !isNil(o.ServiceAffected) {
		toSerialize["serviceAffected"] = o.ServiceAffected
	}
	if true {
		toSerialize["serviceCreated"] = o.ServiceCreated
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !isNil(o.SubjectPropertyName) {
		toSerialize["subjectPropertyName"] = o.SubjectPropertyName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuditLogRequestDTO struct {
	value *CreateAuditLogRequestDTO
	isSet bool
}

func (v NullableCreateAuditLogRequestDTO) Get() *CreateAuditLogRequestDTO {
	return v.value
}

func (v *NullableCreateAuditLogRequestDTO) Set(val *CreateAuditLogRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuditLogRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuditLogRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuditLogRequestDTO(val *CreateAuditLogRequestDTO) *NullableCreateAuditLogRequestDTO {
	return &NullableCreateAuditLogRequestDTO{value: val, isSet: true}
}

func (v NullableCreateAuditLogRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuditLogRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


