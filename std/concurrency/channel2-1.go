//GoSublime double ctrl+.
//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
)

//该程序创建f0->f9共10个routine,一旦f9这个routine执行结束,那么整个程序就结束
//问题是f9这个routine不一定最后被系统调度执行
func main() {
	var c chan bool = make(chan bool)
	for i := 0; i < 10; i++ {
		go f(c, i)
	}
	<-c //阻塞main routine
}

func f(c chan bool, index int) {
	fmt.Printf("Index=%d\n", index)
	if index == 9 { //问题:index=9的routine不一定最后被系统调度执行
		c <- true //当最后一个routine退出时发送channel
	}
}
