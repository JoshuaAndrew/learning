package main

import (
	"fmt"
	"strings"
)

func main()  {
	cidr := "172/31/0.0/16"
	fmt.Println(strings.Replace(cidr,"/",".",-1))
}
