/*
LogCom API

LogCom Swagger documentation

API version: 1.3.4
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// NotificationMessageListPageResponse struct for NotificationMessageListPageResponse
type NotificationMessageListPageResponse struct {
	// The actual page number
	CurrentPage *int32 `json:"currentPage,omitempty"`
	// The items
	Items []NotificationMessageDTO `json:"items,omitempty"`
	// The number of items per page
	PageSize *int32 `json:"pageSize,omitempty"`
	// The total count of items
	TotalCount *int32 `json:"totalCount,omitempty"`
	// The total pages
	TotalPages *int32 `json:"totalPages,omitempty"`
}

// NewNotificationMessageListPageResponse instantiates a new NotificationMessageListPageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationMessageListPageResponse() *NotificationMessageListPageResponse {
	this := NotificationMessageListPageResponse{}
	return &this
}

// NewNotificationMessageListPageResponseWithDefaults instantiates a new NotificationMessageListPageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationMessageListPageResponseWithDefaults() *NotificationMessageListPageResponse {
	this := NotificationMessageListPageResponse{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *NotificationMessageListPageResponse) GetCurrentPage() int32 {
	if o == nil || isNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageListPageResponse) GetCurrentPageOk() (*int32, bool) {
	if o == nil || isNil(o.CurrentPage) {
    return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *NotificationMessageListPageResponse) HasCurrentPage() bool {
	if o != nil && !isNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *NotificationMessageListPageResponse) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NotificationMessageListPageResponse) GetItems() []NotificationMessageDTO {
	if o == nil || isNil(o.Items) {
		var ret []NotificationMessageDTO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageListPageResponse) GetItemsOk() ([]NotificationMessageDTO, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NotificationMessageListPageResponse) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NotificationMessageDTO and assigns it to the Items field.
func (o *NotificationMessageListPageResponse) SetItems(v []NotificationMessageDTO) {
	o.Items = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *NotificationMessageListPageResponse) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageListPageResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *NotificationMessageListPageResponse) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *NotificationMessageListPageResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *NotificationMessageListPageResponse) GetTotalCount() int32 {
	if o == nil || isNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageListPageResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *NotificationMessageListPageResponse) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *NotificationMessageListPageResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *NotificationMessageListPageResponse) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageListPageResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *NotificationMessageListPageResponse) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *NotificationMessageListPageResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o NotificationMessageListPageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !isNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationMessageListPageResponse struct {
	value *NotificationMessageListPageResponse
	isSet bool
}

func (v NullableNotificationMessageListPageResponse) Get() *NotificationMessageListPageResponse {
	return v.value
}

func (v *NullableNotificationMessageListPageResponse) Set(val *NotificationMessageListPageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMessageListPageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMessageListPageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMessageListPageResponse(val *NotificationMessageListPageResponse) *NullableNotificationMessageListPageResponse {
	return &NullableNotificationMessageListPageResponse{value: val, isSet: true}
}

func (v NullableNotificationMessageListPageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMessageListPageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


