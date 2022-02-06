package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := 67.09987
	value,_ := jsonNumber(&a)
	fmt.Println(value)
}

func jsonNumber(number *float64) (value string, err error) {
	var stringified []byte // json number (float64 in golang)
	if stringified, err = json.Marshal(number); err != nil {
		return
	}
	value = string(stringified)
	return
}
