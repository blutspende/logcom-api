# Go API client for LogCom
LogCom client and Swagger documentation

## Overview
The API documentation can be found [**here**](api/README.md)

## Installation
Put the package under your project folder and add the following in import:
```go
import "github.com/DRK-Blutspende-BaWueHe/logcom-api"
```

## Usage

### Init the client
Call the `logcom.Init` function after the logger initialization.
```go
logcom.Init(logcom.Configuration{
    ServiceName: "Your Service (Project) Name",
    LogComURL:   configuration.LogComUrl,
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

If a log is not intended to be sent to LogCom, use the default `zerolog` methods:
- `Send()`
- `Msg(...)`
- `Msgf(...)`

### Audit logging


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
