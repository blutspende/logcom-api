/*
LogCom API

LogCom Swagger documentation

API version: 1.3.10
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"github.com/google/uuid"
)

// AuditLogChangeDTO struct for AuditLogChangeDTO
type AuditLogChangeDTO struct {
	// The category of the change
	Category *string `json:"category,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The message
	Message *string `json:"message,omitempty"`
	// The new value
	NewValue *string `json:"newValue,omitempty"`
	// The old value
	OldValue *string `json:"oldValue,omitempty"`
	// The short description of the change
	Subject *string `json:"subject,omitempty"`
	// The name of the subject
	SubjectName *string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName *string `json:"subjectPropertyName,omitempty"`
}

// NewAuditLogChangeDTO instantiates a new AuditLogChangeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogChangeDTO() *AuditLogChangeDTO {
	this := AuditLogChangeDTO{}
	return &this
}

// NewAuditLogChangeDTOWithDefaults instantiates a new AuditLogChangeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogChangeDTOWithDefaults() *AuditLogChangeDTO {
	this := AuditLogChangeDTO{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AuditLogChangeDTO) SetCategory(v string) {
	o.Category = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetId() uuid.UUID {
	if o == nil || isNil(o.Id) {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AuditLogChangeDTO) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuditLogChangeDTO) SetMessage(v string) {
	o.Message = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetNewValue() string {
	if o == nil || isNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetNewValueOk() (*string, bool) {
	if o == nil || isNil(o.NewValue) {
    return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasNewValue() bool {
	if o != nil && !isNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *AuditLogChangeDTO) SetNewValue(v string) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetOldValue() string {
	if o == nil || isNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetOldValueOk() (*string, bool) {
	if o == nil || isNil(o.OldValue) {
    return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasOldValue() bool {
	if o != nil && !isNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *AuditLogChangeDTO) SetOldValue(v string) {
	o.OldValue = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AuditLogChangeDTO) SetSubject(v string) {
	o.Subject = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetSubjectName() string {
	if o == nil || isNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetSubjectNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectName) {
    return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasSubjectName() bool {
	if o != nil && !isNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *AuditLogChangeDTO) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetSubjectPropertyName returns the SubjectPropertyName field value if set, zero value otherwise.
func (o *AuditLogChangeDTO) GetSubjectPropertyName() string {
	if o == nil || isNil(o.SubjectPropertyName) {
		var ret string
		return ret
	}
	return *o.SubjectPropertyName
}

// GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogChangeDTO) GetSubjectPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPropertyName) {
    return nil, false
	}
	return o.SubjectPropertyName, true
}

// HasSubjectPropertyName returns a boolean if a field has been set.
func (o *AuditLogChangeDTO) HasSubjectPropertyName() bool {
	if o != nil && !isNil(o.SubjectPropertyName) {
		return true
	}

	return false
}

// SetSubjectPropertyName gets a reference to the given string and assigns it to the SubjectPropertyName field.
func (o *AuditLogChangeDTO) SetSubjectPropertyName(v string) {
	o.SubjectPropertyName = &v
}

func (o AuditLogChangeDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !isNil(o.Subject) {
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

type NullableAuditLogChangeDTO struct {
	value *AuditLogChangeDTO
	isSet bool
}

func (v NullableAuditLogChangeDTO) Get() *AuditLogChangeDTO {
	return v.value
}

func (v *NullableAuditLogChangeDTO) Set(val *AuditLogChangeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogChangeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogChangeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogChangeDTO(val *AuditLogChangeDTO) *NullableAuditLogChangeDTO {
	return &NullableAuditLogChangeDTO{value: val, isSet: true}
}

func (v NullableAuditLogChangeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogChangeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


