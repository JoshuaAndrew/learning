/*
   json.Marshal   other --> json
   json.Unmarshal json  --> other
 */

package main

import (
	"fmt"
	//"os"
	"encoding/json"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int        `json:"page"`
	Fruits []string   `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.344567)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple":5, "lettuce":7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	r1 := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	r1B, _ := json.Marshal(r1)
	fmt.Println("r1",string(r1B))

	r2 := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	r2B, _ := json.Marshal(r2)
	fmt.Println("r2",string(r2B))



}
