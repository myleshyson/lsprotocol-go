package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Helper function that takes in a full jsonrpc message, and returns
// the corresponding message struct
func DecodeMessage(message []byte) (Message, error) {
	_, contentLength, content, err := SplitMessage(message)

	if err != nil {
		return nil, err
	}
	var temp map[string]any

	if err := json.Unmarshal(content[:contentLength], &temp); err != nil {
		return nil, err
	}

	method, exists := temp["method"]

	if !exists {
		return nil, fmt.Errorf("Missing method from request")
	}

	result, err := MessageRegistry[method.(string)](content)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Helper function that can be used as a split function in a scanner
func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	totalLength, contentLength, content, err := SplitMessage(data)

	if err != nil {
		return 0, nil, err
	}

	if contentLength < len(content) {
		return 0, nil, nil
	}

	return totalLength, data[:totalLength], nil
}

// Helper function that produces a ResponseError
func Error(code int32, err error) ResponseError {
	return ResponseError{
		Code:    code,
		Message: err.Error(),
	}
}

// Helper function that takes in a jsonrpc message, and returns
// the following:
// [
//
//	full length of the message,
//	the content length,
//	the content itself as an array of bytes,
//	an error if something went wrong
//
// ]
func SplitMessage(data []byte) (int, int, []byte, error) {
	headerPart, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return -1, -1, nil, fmt.Errorf("invalid message format")
	}

	var contentLength int
	for header := range bytes.SplitSeq(headerPart, []byte{'\r', '\n'}) {
		if bytes.HasPrefix(header, []byte("Content-Type: ")) {
			continue
		}

		length, err := strconv.Atoi(string(header[len("Content-Length: "):]))
		if err != nil {
			return -1, -1, nil, fmt.Errorf("invalid content length: %s", err)
		}
		contentLength = length
	}

	return len(data), contentLength, content, nil
}
