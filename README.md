# Go API client for LogCom
LogCom client and Swagger documentation

## Overview
The API documentation can be found [**here**](docs/README.md)

## Installation
```
go get -v github.com/DRK-Blutspende-BaWueHe/logcom-api
```

## Usage
It is **important to pass a context value** (e.g. `gin.Context`) containing the necessary information from the incoming, actual request.

### Init the client
Call the `logcomapi.Init` function after the logger initialization.
```go
logcomapi.Init(logcomapi.Configuration{
    ServiceName: "Your Service (Project) Name",
    LogComURL:   "http://the.logcom.url",
    HeaderProvider: func(ctx context.Context) http.Header {
        if ginCtx, ok := ctx.(*gin.Context); ok {
            return ginCtx.Request.Header
        }
        return http.Header{}
    },
}, &log.Logger)
```

### Console logging
For sending console logs to LogCom too, use one of these methods:
- `SendContext()`
  ```go
  log.Error().(err).SendContext(c)
  ```
- `MsgContext(...)`
  ```go
  log.Debug().MsgContext(c, "Success")
  ```
- `MsgfContext(...)`
  ```go
  log.Warn().MsgfContext(c, "Something went wrong: %v", err)
  ```

For sending custom console log only to LogCom, use one of these methods:
- For simple message: `SendConsoleLog(...)`
  ```go
  SendConsoleLog(ctx, zerolog.ErrorLevel, "Something went wrong")
  ```

- For more customized message: `SendConsoleLogWithModel(...)`
  ```go
  SendConsoleLogWithModel(ctx, logcomapi.CreateConsoleLogRequestDto{
      CreatedAt:     nil,
      CreatedById:   nil,
      CreatedByName: "",
      Level:         0,
      Message:       "",
      Service:       "",
  })
  ```

If a log is not intended to be sent to LogCom, use the default `zerolog` methods:
- `Send()`
- `Msg(...)`
- `Msgf(...)`

### Audit logging

#### Creation

Function to use: `SendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{})`
- `ctx`: The gin context.
- `subject`: The subject what is created.
- `subjectName`: The unique identifier or name of the subject.
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important.

Example:
```go
err := logcomapi.SendAuditLogWithCreation(ctx, "MATERIAL", "ANTIG", newMaterialDTO)
if err != nil {
    _ = txConn.Rollback()
    return errors.New("audit logging error: " + err.Error())
}
```
The translated result in Bloodlab-UI:
```
ANTIG material was created successfully
```

#### Modification

Function to use: `SendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{})`
- `ctx`: The gin context.
- `subject`: The subject what is created.
- `subjectName`: The unique identifier or name of the subject.
- `oldValue`: The original object.
- `newValue`: The modified object.

Example:
```go
err := logcomapi.SendAuditLogWithModification(ctx, "MATERIAL", "ANTIG", originalMaterialDTO, updatedMaterialDTO)
if err != nil {
    _ = txConn.Rollback()
    return errors.New("audit logging error: " + err.Error())
}
```
The translated result in Bloodlab-UI:
- When has only 1 change
  ```
  ANTIG material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```
- When has more than 1 changes
  ```
  3 properties of ANTIG material were modified successfully

  Material code was modified successfully from "ANTIG" to "ANTI-G"
  Material enabled state was modified successfully from "false" to "true"
  Material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```

**When performance matters**

Use the following solution instead of the code above which uses reflection to get the changes between the models.

1. Implement the `logcomapi.GetChanges(...)` interface for the model. An example for the material model:
   ```go
   func (m Material) GetChanges(oldModel interface{}, ignoredProperties map[string]interface{}) []logcomapi.ModelChange {
       changedPropertyNames := make([]logcomapi.ModelChange, 0)
       other := oldModel.(Material)

       if ignoredProperties == nil {
           ignoredProperties = make(map[string]interface{}, 0)
       }

       if isNotIgnored("code", ignoredProperties) && m.Code != other.Code {
           changedPropertyNames = append(changedPropertyNames, logcomapi.ModelChange{
               PropertyName: "code", 
               OldValue:     other.Code,
               NewValue:     m.Code,
           })
       }
       if isNotIgnored("name", ignoredProperties) && m.Name != other.Name {
           changedPropertyNames = append(changedPropertyNames, logcomapi.ModelChange{
               PropertyName: "name",
               OldValue:     other.Name,
               NewValue:     m.Name,
           })
       }
       if isNotIgnored("enabled", ignoredProperties) && m.Enabled != other.Enabled {
           changedPropertyNames = append(changedPropertyNames, logcomapi.ModelChange{
               PropertyName: "enabled",
               OldValue:     other.Enabled,
               NewValue:     m.Enabled,
           })
       }

       return changedPropertyNames
   }

   func isNotIgnored(key string, ignoredProperties map[string]interface{}) bool {
       _, ok := ignoredProperties[key]
       return !ok
   }
   ```
2. Use `SendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange)` function. An example:
   ```go
   changedProperties := updatedMaterial.GetChanges(originalMaterial, nil)
   err = logcomapi.SendAuditLogWithModificationModelChanges(ctx, "MATERIAL", "ANTIG", changedProperties)
   if err != nil {
       _ = txConn.Rollback()
       return errors.New("audit logging error: " + err.Error())
   }
   ```

#### Deletion

Function to use: `SendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{})`
- `ctx`: The gin context.
- `subject`: The subject what is deleted.
- `subjectName`: The unique identifier or name of the subject.
- `oldValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important.

Example:
```go
err := logcomapi.SendAuditLogWithDeletion(ctx, "MATERIAL", "ANTIG", deletedMaterialDTO)
if err != nil {
    _ = txConn.Rollback()
    return errors.New("audit logging error: " + err.Error())
}
```
The translated result in Bloodlab-UI:
```
ANTIG material was removed successfully
```

## Author
laborit@blutspende.de

## Disclaimer
By making use of any information, content and source code in this repository, You agree to the following:

**NO WARRANTIES**

All the information, content and source code provided in this repository is provided "**AS-IS**" and with **NO WARRANTIES**.

**DISCLAIMER OF LIABILITY**

*DRK Blutspendedienst BaWü u. Hessen gGmbH* specifically DISCLAIMS LIABILITY FOR INCIDENTAL OR CONSEQUENTIAL DAMAGES and assumes no responsibility or liability for any loss or damage suffered by any person as a result of the use or misuse of any of the information, content or source code in this repository.

*DRK Blutspendedienst BaWü u. Hessen gGmbH* assumes or undertakes NO LIABILITY for any loss or damage suffered as a result of the use, misuse or reliance on the information, content and source code in this repository.

**USE AT YOUR OWN RISK**

This repository is for *DRK Blutspendedienst BaWü u. Hessen gGmbH* only.
