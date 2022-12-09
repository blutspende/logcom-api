# ConsoleLogDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | The log timestamp | [optional] [default to null]
**CreatedById** | [***uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] [default to null]
**CreatedByName** | **string** | The user&#39;s name who created | [optional] [default to null]
**Id** | [***uuid.UUID**](uuid.UUID.md) | The ID | [optional] [default to null]
**Level** | [***ConsoleLogDtoLevel**](ConsoleLogDTO_level.md) |  | [optional] [default to null]
**Message** | **string** | The log message | [optional] [default to null]
**RequestId** | [***uuid.UUID**](uuid.UUID.md) | The request ID making dependent logs trackable | [optional] [default to null]
**Service** | **string** | The service which sent the log | [optional] [default to null]

[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


