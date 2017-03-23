package main

import (
	"fmt"
)

//地址
type Address struct {
	nation,
	state,
	city,
	region,
	street string
}

type Contact struct {
	phone, email, qq string
}

type Person struct {
	Name    string
	Age     int
	address Address
	contact Contact
}

func main() {
	a1 := &Address{"China", "Beijing", "Beijing", "HaiDian", "中关村南大街"}
	c1 := &Contact{"18500986197", "ww@isoftstone.com", "457129605"}
	p1 := Person{"ww", 30, *a1, *c1}

	fmt.Println(p1)

}
