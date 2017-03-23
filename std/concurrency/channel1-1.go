//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
)

func main() {
	//通过channel来实现main进程和routine之间的通信
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO!")
		c <- true
	}()
	<-c  //直到子routine跑完, 取到消息. main在此阻塞
}
