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

## Utilities
This package provides a limited number of utility functions for dealing with jsonrpc messages.

**DecodeMessage**
This is a helper that takes in a full jsonrpc message and returns a message object.

```golang
var content = []byte("Content-Length: 20\r\n\r\n{...etc}")
message, err := protocol.DecodeMessage(content)
```

**SplitMessage**
This is a helper that takes in a full jsonrpc message and returns information about the message.

```golang
var content = []byte("Content-Length: 20\r\n\r\n{...etc}")
messageLengthIncludingTheHeader, contentLength, content, err := protocol.SplitMessage(content)
```

**Split**
This is a helper you can use with scanner to read jsonrpc messages from a stream.

```golang
var scanner = bufio.NewScanner(reader)
scanner.Split(protocol.Split)
for scanner.Scan() {
	msg := scanner.Bytes()
	// handle message
}
```

**Error**
This is a tiny helper that takes in an error coe and an `error` instance, and returns a `ResponseError`

```golang
responseError := protocol.Error(someErrorCode, err)
```

## Interfaces
The following interfaces are provided by this package:

- `Request` - represents a jsonrpc request
- `Response` - represents a jsonrpc response
- `Notification` - represents a jsonrpc notification
- `Message` - represents a jsonrpc message
- `IncomingMessage` - notification or request. not in the spec but super helpful.
