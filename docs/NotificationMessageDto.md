# NotificationMessageDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | The notification timestamp | [optional] [default to null]
**CreatedById** | [***uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] [default to null]
**CreatedByName** | **string** | The user&#39;s name who created | [optional] [default to null]
**Id** | [***uuid.UUID**](uuid.UUID.md) | The ID | [optional] [default to null]
**Message** | **string** | The notification message | [optional] [default to null]
**Service** | **string** | The service which sent the notification | [optional] [default to null]
**Status** | [***NotificationMessageDtoStatus**](NotificationMessageDTO_status.md) |  | [optional] [default to null]

[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


