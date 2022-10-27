# CreateNotificationRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | The log timestamp | [optional] [default to null]
**CreatedById** | [***uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] [default to null]
**CreatedByName** | **string** | The user&#39;s name who created | [optional] [default to null]
**EventCategory** | **string** | The notification event category | [optional] [default to null]
**Message** | **string** | The log message | [default to null]
**Service** | **string** | The service which sent the log | [default to null]
**Targets** | [**map[string][]string**](array.md) | The targets who will receive the notification | [optional] [default to null]

[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


