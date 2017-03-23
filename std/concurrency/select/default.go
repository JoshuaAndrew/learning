//【golang select机制】
//检查每个case语句
//如果有任意一个chan是send or receive ready，那么就执行该block
//如果多个case是ready的，那么随机找1个并执行该block
//如果都没有ready，那么就block and wait
//如果有default block，而且其他的case都没有ready，就执行该default block
package main

import (
	"fmt"
	"time"
)

func main() {
	//SelectDefault()
	//channelIsFull()
	timeout1()
	//timeout2()
}

//此时因为 ch1 和 ch2 都为空，所以 case1 和 case2 都不会读取成功
//则 select 执行 default 语句
func SelectDefault() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	//ch1 <- 12345
	//ch2 <- 67890
	select {
	case msg1 := <-ch1:
		fmt.Println("received", msg1)
	case msg2 := <-ch2:
		fmt.Println("received", msg2)
	default:
		fmt.Println("default")
	}
}

//【用 select 语句来检测 chan 是否已经满】
//就是因为这个 default 特性，
//我们可以使用 select 语句来检测 chan 是否已经满了

//因为 ch 插入 1,2 的时候已经满了，
//当 ch 要插入 3 的时候，发现 ch 已经满了(case 阻塞住)
//则 select 执行 default 语句。 这样就可以实现对 channel 是否已满的检测，
//而不是一直等待。
func channelIsFull() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	select {
	case ch <- 3:
		fmt.Println("send 3 to channel successfully?")
	case ch <- 4:
		fmt.Println("send 4 to channel successfully! ")
	default:
		fmt.Println("channel is already full !")
	}

}

//【使用 select 实现 timeout 机制】
//当超时时间到的时候，case2 会操作成功。 所以 select 语句则会退出
//而不是一直阻塞在 ch 的读取操作上。 从而实现了对 ch 读取操作的超时设置
func timeout1() {
	timeout := make(chan bool, 1)
	go func() { //傻帽的定时方式
		time.Sleep(2 * time.Second) // sleep two second
		timeout <- true
	}()
	ch := make(chan int)
	//如果在2s内仍然没有从ch这个channel中读取到值就执行timeout分支

	//select-case一种超时技巧
	select {
	case <-ch:
	case <-timeout: //定时器
		fmt.Println("timeout!")
	}
}

//如果在5s内仍然没有从ch这个channel中读取到值就执行timeout分支
func timeout2() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case t := <-time.After(5 * time.Second):
		fmt.Println("Timeout!", t.String())
	}
}
