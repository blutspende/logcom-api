/*
LogCom API

LogCom Swagger documentation

API version: 1.2.28
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"fmt"
)

// NotificationStatus the model 'NotificationStatus'
type NotificationStatus string

// List of NotificationStatus
const (
	NotificationError NotificationStatus = "ERROR"
	NotificationNew NotificationStatus = "NEW"
	NotificationSeen NotificationStatus = "SEEN"
	NotificationSent NotificationStatus = "SENT"
)

// All allowed values of NotificationStatus enum
var AllowedNotificationStatusEnumValues = []NotificationStatus{
	"ERROR",
	"NEW",
	"SEEN",
	"SENT",
}

func (v *NotificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationStatus(value)
	for _, existing := range AllowedNotificationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationStatus", value)
}

// NewNotificationStatusFromValue returns a pointer to a valid NotificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationStatusFromValue(v string) (*NotificationStatus, error) {
	ev := NotificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationStatus: valid values are %v", v, AllowedNotificationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationStatus) IsValid() bool {
	for _, existing := range AllowedNotificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationStatus value
func (v NotificationStatus) Ptr() *NotificationStatus {
	return &v
}

type NullableNotificationStatus struct {
	value *NotificationStatus
	isSet bool
}

func (v NullableNotificationStatus) Get() *NotificationStatus {
	return v.value
}

func (v *NullableNotificationStatus) Set(val *NotificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationStatus(val *NotificationStatus) *NullableNotificationStatus {
	return &NullableNotificationStatus{value: val, isSet: true}
}

func (v NullableNotificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

