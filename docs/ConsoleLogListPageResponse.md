# ConsoleLogListPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | The actual page number | [optional] 
**Items** | Pointer to [**[]ConsoleLogDTO**](ConsoleLogDTO.md) | The items | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 
**TotalCount** | Pointer to **int32** | The total count of items | [optional] 
**TotalPages** | Pointer to **int32** | The total pages | [optional] 

## Methods

### NewConsoleLogListPageResponse

`func NewConsoleLogListPageResponse() *ConsoleLogListPageResponse`

NewConsoleLogListPageResponse instantiates a new ConsoleLogListPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleLogListPageResponseWithDefaults

`func NewConsoleLogListPageResponseWithDefaults() *ConsoleLogListPageResponse`

NewConsoleLogListPageResponseWithDefaults instantiates a new ConsoleLogListPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ConsoleLogListPageResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ConsoleLogListPageResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ConsoleLogListPageResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ConsoleLogListPageResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetItems

`func (o *ConsoleLogListPageResponse) GetItems() []ConsoleLogDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConsoleLogListPageResponse) GetItemsOk() (*[]ConsoleLogDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConsoleLogListPageResponse) SetItems(v []ConsoleLogDTO)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConsoleLogListPageResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPageSize

`func (o *ConsoleLogListPageResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ConsoleLogListPageResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ConsoleLogListPageResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ConsoleLogListPageResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ConsoleLogListPageResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ConsoleLogListPageResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ConsoleLogListPageResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ConsoleLogListPageResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ConsoleLogListPageResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ConsoleLogListPageResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ConsoleLogListPageResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ConsoleLogListPageResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


