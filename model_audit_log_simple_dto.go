/*
LogCom API

LogCom Swagger documentation

API version: 1.3.12
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// AuditLogSimpleDTO struct for AuditLogSimpleDTO
type AuditLogSimpleDTO struct {
	// The category of the change
	Category *string `json:"category,omitempty"`
	// The creation timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The message
	Message *string `json:"message,omitempty"`
	// The service which the change affects
	ServiceAffected *string `json:"serviceAffected,omitempty"`
	// The service which created
	ServiceCreated *string `json:"serviceCreated,omitempty"`
	// The subject of the change
	Subject *string `json:"subject,omitempty"`
}

// NewAuditLogSimpleDTO instantiates a new AuditLogSimpleDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogSimpleDTO() *AuditLogSimpleDTO {
	this := AuditLogSimpleDTO{}
	return &this
}

// NewAuditLogSimpleDTOWithDefaults instantiates a new AuditLogSimpleDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogSimpleDTOWithDefaults() *AuditLogSimpleDTO {
	this := AuditLogSimpleDTO{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AuditLogSimpleDTO) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuditLogSimpleDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *AuditLogSimpleDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetId() uuid.UUID {
	if o == nil || isNil(o.Id) {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AuditLogSimpleDTO) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuditLogSimpleDTO) SetMessage(v string) {
	o.Message = &v
}

// GetServiceAffected returns the ServiceAffected field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetServiceAffected() string {
	if o == nil || isNil(o.ServiceAffected) {
		var ret string
		return ret
	}
	return *o.ServiceAffected
}

// GetServiceAffectedOk returns a tuple with the ServiceAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetServiceAffectedOk() (*string, bool) {
	if o == nil || isNil(o.ServiceAffected) {
    return nil, false
	}
	return o.ServiceAffected, true
}

// HasServiceAffected returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasServiceAffected() bool {
	if o != nil && !isNil(o.ServiceAffected) {
		return true
	}

	return false
}

// SetServiceAffected gets a reference to the given string and assigns it to the ServiceAffected field.
func (o *AuditLogSimpleDTO) SetServiceAffected(v string) {
	o.ServiceAffected = &v
}

// GetServiceCreated returns the ServiceCreated field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetServiceCreated() string {
	if o == nil || isNil(o.ServiceCreated) {
		var ret string
		return ret
	}
	return *o.ServiceCreated
}

// GetServiceCreatedOk returns a tuple with the ServiceCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetServiceCreatedOk() (*string, bool) {
	if o == nil || isNil(o.ServiceCreated) {
    return nil, false
	}
	return o.ServiceCreated, true
}

// HasServiceCreated returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasServiceCreated() bool {
	if o != nil && !isNil(o.ServiceCreated) {
		return true
	}

	return false
}

// SetServiceCreated gets a reference to the given string and assigns it to the ServiceCreated field.
func (o *AuditLogSimpleDTO) SetServiceCreated(v string) {
	o.ServiceCreated = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AuditLogSimpleDTO) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogSimpleDTO) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AuditLogSimpleDTO) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AuditLogSimpleDTO) SetSubject(v string) {
	o.Subject = &v
}

func (o AuditLogSimpleDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ServiceAffected) {
		toSerialize["serviceAffected"] = o.ServiceAffected
	}
	if !isNil(o.ServiceCreated) {
		toSerialize["serviceCreated"] = o.ServiceCreated
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLogSimpleDTO struct {
	value *AuditLogSimpleDTO
	isSet bool
}

func (v NullableAuditLogSimpleDTO) Get() *AuditLogSimpleDTO {
	return v.value
}

func (v *NullableAuditLogSimpleDTO) Set(val *AuditLogSimpleDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogSimpleDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogSimpleDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogSimpleDTO(val *AuditLogSimpleDTO) *NullableAuditLogSimpleDTO {
	return &NullableAuditLogSimpleDTO{value: val, isSet: true}
}

func (v NullableAuditLogSimpleDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogSimpleDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


