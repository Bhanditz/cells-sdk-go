# Cells API client

**Warning, this is a Work In Progress**.

Rest API Client for Pydio Cells.

Current SDK has been update on **Sept. 12th 2018** with git commit **[da01e73](https://github.com/pydio/cells/commit/da01e73fb123177d0748513b272d0d9d0053a31e)**.

## Overview

This API client was generated by the [go-swagger](https://github.com/go-swagger/go-swagger) project.

For more information, please visit [https://pydio.com](https://pydio.com)

## Installation

Put the package under your project folder and add the following in import:

```go
    "github.com/pydio/cells-sdk-go"
```

## Update SDK

_This is for the maintainers of cells-sdk-go project only._

To update the Cells go sdk, you must follow the steps below:

_First time only_:

```sh
# go to the roots of this directory, typically:
cd $GOPATH/src/github.com/pydio/cells-sdk-go

# If necessary, retrieve swagger binary, rename it and give execution permission

## for linux
wget https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_linux_amd64

## for Mac OS
wget https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_darwin_amd64

mv swagger_linux_amd64 swagger
chmod u+x swagger
```

_After each API Spec modification_:

```sh

# Clean existing generated code and previous swagger
## Vorsicht :)
rm -r client models rest.swagger.json

# Retrieve latest spec file
wget https://raw.githubusercontent.com/pydio/cells/master/common/proto/rest/rest.swagger.json

# Simply generate updated code
./swagger generate client --skip-validation -f rest.swagger.json
```

```sh
# Apply the tweak to workaround int64 serialisation issue between protobuf and swagger
go run build/main.go tweak-model
```

You should also update version information at the top of this page.

_**NOTE**: we use the --skip-validation flag to avoid circular issues with object that make reference to same type of objects, typically activities and jobs. See [issue #957 in go-swagger repository](https://github.com/go-swagger/go-swagger/issues/957) for more details._

### Run SDK Unit Tests

- Insure you have a running instance of Cells with correct version
- Check that `config.json` and `config-s3.json` files are in `config` folder and contain valid info. (See samples files in same folder).
- Change RunEnvAwareTests flag to `true` (line 29 from `config/naming.go`). **WARNING: Do not commit this**.
- execute:

```sh
go test -v ./...
```