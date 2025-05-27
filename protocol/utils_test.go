package protocol

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDecodeNotification(t *testing.T) {
	content := []byte("{\"jsonrpc\": \"2.0\", \"id\": 1, \"method\": \"workspace/codeLens/refresh\", \"params\": null}")
	requestMessage := []byte("Content-Length: " + strconv.Itoa(len(content)) + "\r\n\r\n" + string(content))

	message, err := DecodeMessage(requestMessage)

	if err != nil {
		t.Fatal(err)
	}

	if _, check := message.(CodeLensRefreshRequest); !check {
		fmt.Printf("Unexpected message type: %T", message)
		t.Fatal("Expected CodeLensRefreshRequest")
	}
}

func TestDecodeRequest(t *testing.T) {
	content := []byte("{\"jsonrpc\": \"2.0\", \"id\": 1, \"method\": \"workspace/diagnostic\", \"params\": {\"previousResultIds\": []}}")
	requestMessage := []byte("Content-Length: " + strconv.Itoa(len(content)) + "\r\n\r\n" + string(content))

	message, err := DecodeMessage(requestMessage)

	if err != nil {
		t.Fatal(err)
	}

	if _, check := message.(WorkspaceDiagnosticRequest); !check {
		fmt.Printf("Unexpected message type: %T", message)
		t.Fatal("Expected WorkspaceDiagnosticRequest")
	}
}
