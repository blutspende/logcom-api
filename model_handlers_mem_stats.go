/*
LogCom API

LogCom Swagger documentation

API version: 1.3.18
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// HandlersMemStats struct for HandlersMemStats
type HandlersMemStats struct {
	Alloc *string `json:"alloc,omitempty"`
	HeapAlloc *string `json:"heapAlloc,omitempty"`
	HeapInUse *string `json:"heapInUse,omitempty"`
	NumberOfGoRoutines *int32 `json:"numberOfGoRoutines,omitempty"`
	StackInUse *string `json:"stackInUse,omitempty"`
	Sys *string `json:"sys,omitempty"`
	TotalAlloc *string `json:"totalAlloc,omitempty"`
}

// NewHandlersMemStats instantiates a new HandlersMemStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlersMemStats() *HandlersMemStats {
	this := HandlersMemStats{}
	return &this
}

// NewHandlersMemStatsWithDefaults instantiates a new HandlersMemStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlersMemStatsWithDefaults() *HandlersMemStats {
	this := HandlersMemStats{}
	return &this
}

// GetAlloc returns the Alloc field value if set, zero value otherwise.
func (o *HandlersMemStats) GetAlloc() string {
	if o == nil || isNil(o.Alloc) {
		var ret string
		return ret
	}
	return *o.Alloc
}

// GetAllocOk returns a tuple with the Alloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetAllocOk() (*string, bool) {
	if o == nil || isNil(o.Alloc) {
    return nil, false
	}
	return o.Alloc, true
}

// HasAlloc returns a boolean if a field has been set.
func (o *HandlersMemStats) HasAlloc() bool {
	if o != nil && !isNil(o.Alloc) {
		return true
	}

	return false
}

// SetAlloc gets a reference to the given string and assigns it to the Alloc field.
func (o *HandlersMemStats) SetAlloc(v string) {
	o.Alloc = &v
}

// GetHeapAlloc returns the HeapAlloc field value if set, zero value otherwise.
func (o *HandlersMemStats) GetHeapAlloc() string {
	if o == nil || isNil(o.HeapAlloc) {
		var ret string
		return ret
	}
	return *o.HeapAlloc
}

// GetHeapAllocOk returns a tuple with the HeapAlloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetHeapAllocOk() (*string, bool) {
	if o == nil || isNil(o.HeapAlloc) {
    return nil, false
	}
	return o.HeapAlloc, true
}

// HasHeapAlloc returns a boolean if a field has been set.
func (o *HandlersMemStats) HasHeapAlloc() bool {
	if o != nil && !isNil(o.HeapAlloc) {
		return true
	}

	return false
}

// SetHeapAlloc gets a reference to the given string and assigns it to the HeapAlloc field.
func (o *HandlersMemStats) SetHeapAlloc(v string) {
	o.HeapAlloc = &v
}

// GetHeapInUse returns the HeapInUse field value if set, zero value otherwise.
func (o *HandlersMemStats) GetHeapInUse() string {
	if o == nil || isNil(o.HeapInUse) {
		var ret string
		return ret
	}
	return *o.HeapInUse
}

// GetHeapInUseOk returns a tuple with the HeapInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetHeapInUseOk() (*string, bool) {
	if o == nil || isNil(o.HeapInUse) {
    return nil, false
	}
	return o.HeapInUse, true
}

// HasHeapInUse returns a boolean if a field has been set.
func (o *HandlersMemStats) HasHeapInUse() bool {
	if o != nil && !isNil(o.HeapInUse) {
		return true
	}

	return false
}

// SetHeapInUse gets a reference to the given string and assigns it to the HeapInUse field.
func (o *HandlersMemStats) SetHeapInUse(v string) {
	o.HeapInUse = &v
}

// GetNumberOfGoRoutines returns the NumberOfGoRoutines field value if set, zero value otherwise.
func (o *HandlersMemStats) GetNumberOfGoRoutines() int32 {
	if o == nil || isNil(o.NumberOfGoRoutines) {
		var ret int32
		return ret
	}
	return *o.NumberOfGoRoutines
}

// GetNumberOfGoRoutinesOk returns a tuple with the NumberOfGoRoutines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetNumberOfGoRoutinesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfGoRoutines) {
    return nil, false
	}
	return o.NumberOfGoRoutines, true
}

// HasNumberOfGoRoutines returns a boolean if a field has been set.
func (o *HandlersMemStats) HasNumberOfGoRoutines() bool {
	if o != nil && !isNil(o.NumberOfGoRoutines) {
		return true
	}

	return false
}

// SetNumberOfGoRoutines gets a reference to the given int32 and assigns it to the NumberOfGoRoutines field.
func (o *HandlersMemStats) SetNumberOfGoRoutines(v int32) {
	o.NumberOfGoRoutines = &v
}

// GetStackInUse returns the StackInUse field value if set, zero value otherwise.
func (o *HandlersMemStats) GetStackInUse() string {
	if o == nil || isNil(o.StackInUse) {
		var ret string
		return ret
	}
	return *o.StackInUse
}

// GetStackInUseOk returns a tuple with the StackInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetStackInUseOk() (*string, bool) {
	if o == nil || isNil(o.StackInUse) {
    return nil, false
	}
	return o.StackInUse, true
}

// HasStackInUse returns a boolean if a field has been set.
func (o *HandlersMemStats) HasStackInUse() bool {
	if o != nil && !isNil(o.StackInUse) {
		return true
	}

	return false
}

// SetStackInUse gets a reference to the given string and assigns it to the StackInUse field.
func (o *HandlersMemStats) SetStackInUse(v string) {
	o.StackInUse = &v
}

// GetSys returns the Sys field value if set, zero value otherwise.
func (o *HandlersMemStats) GetSys() string {
	if o == nil || isNil(o.Sys) {
		var ret string
		return ret
	}
	return *o.Sys
}

// GetSysOk returns a tuple with the Sys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetSysOk() (*string, bool) {
	if o == nil || isNil(o.Sys) {
    return nil, false
	}
	return o.Sys, true
}

// HasSys returns a boolean if a field has been set.
func (o *HandlersMemStats) HasSys() bool {
	if o != nil && !isNil(o.Sys) {
		return true
	}

	return false
}

// SetSys gets a reference to the given string and assigns it to the Sys field.
func (o *HandlersMemStats) SetSys(v string) {
	o.Sys = &v
}

// GetTotalAlloc returns the TotalAlloc field value if set, zero value otherwise.
func (o *HandlersMemStats) GetTotalAlloc() string {
	if o == nil || isNil(o.TotalAlloc) {
		var ret string
		return ret
	}
	return *o.TotalAlloc
}

// GetTotalAllocOk returns a tuple with the TotalAlloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersMemStats) GetTotalAllocOk() (*string, bool) {
	if o == nil || isNil(o.TotalAlloc) {
    return nil, false
	}
	return o.TotalAlloc, true
}

// HasTotalAlloc returns a boolean if a field has been set.
func (o *HandlersMemStats) HasTotalAlloc() bool {
	if o != nil && !isNil(o.TotalAlloc) {
		return true
	}

	return false
}

// SetTotalAlloc gets a reference to the given string and assigns it to the TotalAlloc field.
func (o *HandlersMemStats) SetTotalAlloc(v string) {
	o.TotalAlloc = &v
}

func (o HandlersMemStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alloc) {
		toSerialize["alloc"] = o.Alloc
	}
	if !isNil(o.HeapAlloc) {
		toSerialize["heapAlloc"] = o.HeapAlloc
	}
	if !isNil(o.HeapInUse) {
		toSerialize["heapInUse"] = o.HeapInUse
	}
	if !isNil(o.NumberOfGoRoutines) {
		toSerialize["numberOfGoRoutines"] = o.NumberOfGoRoutines
	}
	if !isNil(o.StackInUse) {
		toSerialize["stackInUse"] = o.StackInUse
	}
	if !isNil(o.Sys) {
		toSerialize["sys"] = o.Sys
	}
	if !isNil(o.TotalAlloc) {
		toSerialize["totalAlloc"] = o.TotalAlloc
	}
	return json.Marshal(toSerialize)
}

type NullableHandlersMemStats struct {
	value *HandlersMemStats
	isSet bool
}

func (v NullableHandlersMemStats) Get() *HandlersMemStats {
	return v.value
}

func (v *NullableHandlersMemStats) Set(val *HandlersMemStats) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlersMemStats) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlersMemStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlersMemStats(val *HandlersMemStats) *NullableHandlersMemStats {
	return &NullableHandlersMemStats{value: val, isSet: true}
}

func (v NullableHandlersMemStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlersMemStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


