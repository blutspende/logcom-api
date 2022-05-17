# CreateAuditLogRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the change | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The creation timestamp | [optional] [default to null]
**CreatedById** | [***uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] [default to null]
**CreatedByName** | **string** | The user&#39;s name who created | [optional] [default to null]
**GroupedChanges** | [**[]GroupedAuditLogChangesDto**](GroupedAuditLogChangesDTO.md) | The grouped changes if there are many related ones | [optional] [default to null]
**Message** | **string** | The message | [optional] [default to null]
**NewValue** | **string** | The new value | [optional] [default to null]
**OldValue** | **string** | The old value | [optional] [default to null]
**ServiceAffected** | **string** | The service which the change affects | [optional] [default to null]
**ServiceCreated** | **string** | The service which created | [default to null]
**Subject** | **string** | The subject of the change | [default to null]
**SubjectName** | **string** | The name of the subject | [optional] [default to null]
**SubjectPropertyName** | **string** | The property name of the subject | [optional] [default to null]

[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


