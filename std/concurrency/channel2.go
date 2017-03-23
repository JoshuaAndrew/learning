//GoSublime double ctrl+.
//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	//"runtime"
)

//fix Bug channel3.go    channel
//fix Bug waitGroup.go   waitGroup

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU()) //问题 runtime.GOMAXPROCS(1)
	var c chan bool = make(chan bool)    //不带缓冲区的channel
	for i := 0; i < 10; i++ {
		go func(index int) {
			var sum int
			for i := 0; i <= 100; i++ {
				sum += i
			}
			fmt.Printf("Index=%d,Sum=%d\n", index, sum)
			if index == 9 { //问题: index是9的routine不一定最后被系统调度执行
				c <- true //当最后一个routine退出时发送channel
				//因此总共只发送了一次消息
			}

		}(i)
	}
	<-c //阻塞main routine
	//因此这里只接受一次消息
}
