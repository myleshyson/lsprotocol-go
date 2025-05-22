package main

import (
	"encoding/json"
	"fmt"

	"github.com/myleshyson/lsprotocol-go/lsprotocol"
)

func main() {
	data := []byte("{\"code\": -2147483648}")
	var test lsprotocol.ResponseError

	if err := json.Unmarshal(data, &test); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", test)
}
