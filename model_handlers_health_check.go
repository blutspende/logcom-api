/*
LogCom API

LogCom Swagger documentation

API version: 1.3.13
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// HandlersHealthCheck struct for HandlersHealthCheck
type HandlersHealthCheck struct {
	// The API version
	ApiVersion []string `json:"apiVersion,omitempty"`
	// Docker build version
	BuildVersion *string `json:"buildVersion,omitempty"`
	MemStats *HandlersMemStats `json:"memStats,omitempty"`
	// The service name
	Service *string `json:"service,omitempty"`
	// The service status
	Status *string `json:"status,omitempty"`
}

// NewHandlersHealthCheck instantiates a new HandlersHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlersHealthCheck() *HandlersHealthCheck {
	this := HandlersHealthCheck{}
	return &this
}

// NewHandlersHealthCheckWithDefaults instantiates a new HandlersHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlersHealthCheckWithDefaults() *HandlersHealthCheck {
	this := HandlersHealthCheck{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HandlersHealthCheck) GetApiVersion() []string {
	if o == nil || isNil(o.ApiVersion) {
		var ret []string
		return ret
	}
	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersHealthCheck) GetApiVersionOk() ([]string, bool) {
	if o == nil || isNil(o.ApiVersion) {
    return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HandlersHealthCheck) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given []string and assigns it to the ApiVersion field.
func (o *HandlersHealthCheck) SetApiVersion(v []string) {
	o.ApiVersion = v
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *HandlersHealthCheck) GetBuildVersion() string {
	if o == nil || isNil(o.BuildVersion) {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersHealthCheck) GetBuildVersionOk() (*string, bool) {
	if o == nil || isNil(o.BuildVersion) {
    return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *HandlersHealthCheck) HasBuildVersion() bool {
	if o != nil && !isNil(o.BuildVersion) {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *HandlersHealthCheck) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetMemStats returns the MemStats field value if set, zero value otherwise.
func (o *HandlersHealthCheck) GetMemStats() HandlersMemStats {
	if o == nil || isNil(o.MemStats) {
		var ret HandlersMemStats
		return ret
	}
	return *o.MemStats
}

// GetMemStatsOk returns a tuple with the MemStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersHealthCheck) GetMemStatsOk() (*HandlersMemStats, bool) {
	if o == nil || isNil(o.MemStats) {
    return nil, false
	}
	return o.MemStats, true
}

// HasMemStats returns a boolean if a field has been set.
func (o *HandlersHealthCheck) HasMemStats() bool {
	if o != nil && !isNil(o.MemStats) {
		return true
	}

	return false
}

// SetMemStats gets a reference to the given HandlersMemStats and assigns it to the MemStats field.
func (o *HandlersHealthCheck) SetMemStats(v HandlersMemStats) {
	o.MemStats = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *HandlersHealthCheck) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersHealthCheck) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *HandlersHealthCheck) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *HandlersHealthCheck) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HandlersHealthCheck) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersHealthCheck) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HandlersHealthCheck) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HandlersHealthCheck) SetStatus(v string) {
	o.Status = &v
}

func (o HandlersHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.BuildVersion) {
		toSerialize["buildVersion"] = o.BuildVersion
	}
	if !isNil(o.MemStats) {
		toSerialize["memStats"] = o.MemStats
	}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableHandlersHealthCheck struct {
	value *HandlersHealthCheck
	isSet bool
}

func (v NullableHandlersHealthCheck) Get() *HandlersHealthCheck {
	return v.value
}

func (v *NullableHandlersHealthCheck) Set(val *HandlersHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlersHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlersHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlersHealthCheck(val *HandlersHealthCheck) *NullableHandlersHealthCheck {
	return &NullableHandlersHealthCheck{value: val, isSet: true}
}

func (v NullableHandlersHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlersHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


