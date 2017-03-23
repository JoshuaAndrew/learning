package main

import (
	"fmt"
	"math/rand"
	"time"
)

func r(name string, delay time.Duration) {
	t0 := time.Now()
	fmt.Println(name, " start at ", t0)
	time.Sleep(delay)
	t1 := time.Now()
	fmt.Println(name, " end at ", t1)
	fmt.Println(name, " lasted ", t1.Sub(t0))
}

func main() {

	var elapsed = time.Now().Unix()
	fmt.Println("elapsed : %d ", elapsed)
	rand.Seed(elapsed) //生成随机种子

	var name string
	var delay time.Duration
	for i := 0; i < 5; i++ {
		name = fmt.Sprintf("go_%02d", i) //生成ID
		//生成随机等待时间，从0-10秒
		delay = time.Duration(rand.Intn(10)) * time.Second
		go r(name, delay)
	}

	//让主进程停住，不然主进程退了，goroutine也就退了
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Println("done")

}
