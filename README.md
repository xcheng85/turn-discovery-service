# turn-discovery-service

## Content
- [Installation](#installation)
- [Golang Programming](#golang)

## installation

```shell
go mod init github.com/xcheng85/turn-discovery-service
# structural logging
go get -u go.uber.org/zap
# global configuration
go get -u github.com/ilyakaznacheev/cleanenv
```
## build
```shell
go build .
```

## Run locally
```shell
CONFIG_PATH="config.yaml" SECRET_PATH="secret" ELB_EXTERNAL_IP="0.0.0.0" go run .
```

## golang
This golang application uses the following techniques:

- custom struct with yaml/json mapping
- spread string for function argument to indicate variable length of string
- multiple return for function
- length of array, slice-like strings
- os open file, copy bytes between files
- defer for auto resource clean in the scope of function
- http middleware, chain of responsiblity pattern (multiple middlewares)
- usage of time package
- define function type
- use interface{} to represent any-type (any in typescript)
- struct implement interface
- usage of strings package: string joins
- usage of fmt package: string and int joins
- golang style enum
- switch
- internal function in a package
- string to byte slice
- slice of string (difference vs typescript)
- append slice with slics
- for range loop
- initialize struct instance