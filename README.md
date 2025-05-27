# lsprotocol generator for go

This is a generator plugin for go using https://github.com/microsoft/lsprotocol.

## Quickstart for using in a go project
```shell
go get github.com/myleshyson/lsprotocol-go@latest
```
```golang
import "github.com/myleshyson/lsprotocol-go/protocol"

var request protocol.InitializeRequest
```

All lsprotocol types are in the `protocol` package.

## TODO
- Currently only types are provided. Would be nice to also provide utility functions for creating a request/response.
- the microsoft tests only test for proper unmarshalling. probably should add tests to make sure structs marshal correctly as well.
- I'm not a golang expert. There may be more efficient ways to do things like validation or tagged unions.
