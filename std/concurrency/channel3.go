// Author : weiweiu@isoftstone.com
// Date   : 2016-05-20
// GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	"time"
	"runtime"
)

//fix Bug channel2.go channel2-1.go
//注意输出结果无序,因为此处利用到了cpu的多核特性
func main() {
	//fmt.Printf("runtime.NumCPU()=%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	t1 := time.Now().Nanosecond()
	var c chan bool = make(chan bool,10) //带缓冲区的channel
	for i := 0; i < 10; i++ {
		go func(index int) {
			var sum int
			for i := 0; i <= 100000; i++ {
				sum += i
			}
			fmt.Printf("Index=%d,Sum=%d\n", index, sum)
			time.Sleep(time.Second)
			c <- true //总共发送了10次消息
		}(i)
	}

	for i := 0; i < 10; i++ { //总共接收了10次消息
		<-c
	}
	t2 := time.Now().Nanosecond()
	fmt.Println("Time Elapsed:",t2-t1)
}
