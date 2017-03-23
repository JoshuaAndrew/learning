package main

import (
	"fmt"
	"time"
)

func main() {
	f5()
}

//错误的用法(1)---重复关闭channel
func f1() {
	ch := make(chan bool)
	close(ch)
	close(ch) // 这样会panic的，channel不能close两次
}

//错误的用法(2)---读取的时候channel提前关闭了channel
func f2() {
	ch := make(chan bool)
	close(ch)
	i := <-ch //不会panic, 但是i读取到的是零值(""/false/nil)
	fmt.Println(i)
}

//错误的用法(3)---向已经关闭的channel写数据
func f3() {
	ch := make(chan bool)
	close(ch)
	ch <- true //会panic
}

//技巧用法(1)---判断channel是否close
func f4() {
	ch := make(chan int)
	//ch <- 2345
	i, ok := <-ch //will set ok to false for a closed channel
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("channel closed")
	}
}

//技巧用法(2)---for循环读取channel
func f5() {
	ch := make(chan bool)
	for i := range ch {
		// ch关闭时，for循环会自动结束
		println(i)
	}
}

//技巧用法(3)---防止读取超时
func f6() {
	ch := make(chan bool)
	select {
	case <-time.After(time.Second * 2):
		println("read channel timeout")
	case i := <-ch:
		println(i)
	}
}

//技巧用法(4)---防止写入超时
func f7() {
	ch := make(chan bool)
	// 其实和读取超时很像
	select {
	case <-time.After(time.Second * 2):
		println("write channel timeout")
	case ch <- true:
		println("write ok")
	}
}
