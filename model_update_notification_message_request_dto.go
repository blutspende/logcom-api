/*
LogCom API

LogCom Swagger documentation

API version: 1.3.15
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// UpdateNotificationMessageRequestDTO struct for UpdateNotificationMessageRequestDTO
type UpdateNotificationMessageRequestDTO struct {
	// The status of the message (only the \"SEEN\" status is supported)
	Status *string `json:"status,omitempty"`
}

// NewUpdateNotificationMessageRequestDTO instantiates a new UpdateNotificationMessageRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNotificationMessageRequestDTO() *UpdateNotificationMessageRequestDTO {
	this := UpdateNotificationMessageRequestDTO{}
	return &this
}

// NewUpdateNotificationMessageRequestDTOWithDefaults instantiates a new UpdateNotificationMessageRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNotificationMessageRequestDTOWithDefaults() *UpdateNotificationMessageRequestDTO {
	this := UpdateNotificationMessageRequestDTO{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateNotificationMessageRequestDTO) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationMessageRequestDTO) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateNotificationMessageRequestDTO) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateNotificationMessageRequestDTO) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateNotificationMessageRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNotificationMessageRequestDTO struct {
	value *UpdateNotificationMessageRequestDTO
	isSet bool
}

func (v NullableUpdateNotificationMessageRequestDTO) Get() *UpdateNotificationMessageRequestDTO {
	return v.value
}

func (v *NullableUpdateNotificationMessageRequestDTO) Set(val *UpdateNotificationMessageRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNotificationMessageRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNotificationMessageRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNotificationMessageRequestDTO(val *UpdateNotificationMessageRequestDTO) *NullableUpdateNotificationMessageRequestDTO {
	return &NullableUpdateNotificationMessageRequestDTO{value: val, isSet: true}
}

func (v NullableUpdateNotificationMessageRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNotificationMessageRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


