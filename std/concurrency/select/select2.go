//【golang select机制】
//检查每个case语句
//如果有任意一个chan是send or receive ready，那么就执行该block
//如果多个case是ready的，那么随机找1个并执行该block
//如果都没有ready，那么就block and wait
//如果有default block，而且其他的case都没有ready，就执行该default block
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		//sub routine不停的从channel中读
		for value := range c {
			fmt.Println(value)
		}
	}()

	for {
		//main routine不停的向channel中写
		select {
		case c <- 0:
		case c <- 1:
		case c <- 2:
		}
	}
}
