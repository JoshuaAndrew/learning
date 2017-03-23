//【golang select机制】
//检查每个case语句
//如果有任意一个chan是send or receive ready，那么就执行该block
//如果多个case是ready的，那么随机找1个并执行该block
//如果都没有ready，那么就block and wait
//如果有default block，而且其他的case都没有ready，就执行该default block
//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	"reflect"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	singal := make(chan bool)
	go func() {
		for { //在一个routine中无限循环等待,c1,c2中接受数据
			select {
			case v, ok := <-c1: //will set ok to false for a closed channel
				fmt.Println("type:", reflect.TypeOf(!ok))
				fmt.Println("c1", !ok)
				if !ok {
					singal <- true
					fmt.Println("=============================?")
					break
				}
				fmt.Println("c1", v) //注意此处必须放到上面的if语句下面
				fmt.Println("=============================")
			case v, ok := <-c2: //will set ok to false for a closed channel
				if !ok {
					singal <- true
					fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&?")
					break
				}
				fmt.Println("c2", v) //注意此处必须放到上面的if语句下面
				fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
			}
		}
	}()

	c1 <- 2
	c2 <- "weiweiu"

	c1 <- 345
	c2 <- "isoftstone"

	close(c1) //关闭channel后,channel的状态标示false
	close(c2) //关闭channel后,channel的状态标示false

	msg := <-singal //此处必须阻塞main routine直到c1,c2中任意一个关闭,然后执行后面的代码

	fmt.Println("go ahead!", msg)
}
