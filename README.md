# Cells API client

Rest API Client for Pydio Cells.

Current SDK has been update on **Oct. 29th 2018** with git commit **[29b9e7f](https://github.com/pydio/cells/commit/29b9e7f11286a8a07c4621018a6cf474434c9165)** for **Cells v1.2.1**.

## Overview

This API client was generated by the [go-swagger](https://github.com/go-swagger/go-swagger) project.

For more information, please visit [https://pydio.com](https://pydio.com)

## Installation

Put the package under your project folder and add the following in import:

```go
    "github.com/pydio/cells-sdk-go"
```

## Run SDK Unit Tests

- Insure you have a running instance of Cells with correct version
- Check that `config.json` and `config-s3.json` files are in `config` folder and contain valid info. (See samples files in same folder).
- Change RunEnvAwareTests flag to `true` (line 29 from `config/naming.go`). **WARNING: Do not commit this**.
- execute:

```sh
go test -v ./...
```