package main

import (
	"net"
	"fmt"
)

func main() {
	//defaultCidr := []string {
	//	"10.0.0.0/8",
	//	"172.16.0.0/12",
	//	"192.168.0.0/16",
	//}
    input := "222.40.30.30/16"

    inputIP,inputNet, _ := net.ParseCIDR(input)
    fmt.Println(inputIP.String(),inputNet.String())

	ip := net.ParseIP("172.16.1.100")
	first,cidr,_ := net.ParseCIDR("172.31.1.10/16")

	result1 := ip.Mask(cidr.Mask)
	result2 := first.Mask(cidr.Mask)

	fmt.Println(first.String())
	fmt.Println(result1.String(),result2.String())
	fmt.Println(result1.String()==result2.String())
}
