// Author : weiweiu@isoftstone.com
// Date   : 2016-05-20
// GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const ROUTINE_NUM = 10

//注意输出结果无序,因为此处利用到了cpu的多核特性
func main() {
	t0 := time.Now()
    wg0()
	t1 := time.Now()
	fmt.Println("======================",t1.Sub(t0))
    wg1()
	t2 := time.Now()
    fmt.Println("======================",t2.Sub(t1))
    wg2()
	t3 := time.Now()
	fmt.Println("======================",t3.Sub(t2))
}

func wg0()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(ROUTINE_NUM) //此处add的数量必须和goroutine的数量一致
	for i := 0; i < ROUTINE_NUM; i++ {
		go func(wg *sync.WaitGroup, index int) {
			var a int
			for i := 0; i < 100; i++ {
				a += i
			}
			time.Sleep(1 * time.Second)
			fmt.Printf("Index=%d,Value=%d\n", index, a)
			wg.Done() //每个routine结束时对waitgroup中的数字减一
		}(&wg, i)
	}
	wg.Wait()
}

func wg1()  {
	wg := sync.WaitGroup{}
	wg.Add(ROUTINE_NUM) //此处add的数量必须和goroutine的数量一致
	for i := 0; i < ROUTINE_NUM; i++ {
		go func(wg *sync.WaitGroup, index int) {
			var a int
			for i := 0; i < 100; i++ {
				a += i
			}
			fmt.Printf("Index=%d,Value=%d\n", index, a)
			wg.Done() //每个routine结束时对waitgroup中的数字减一
		}(&wg, i)
	}
	wg.Wait()
}

func wg2()  {
	wg := sync.WaitGroup{}
	wg.Add(ROUTINE_NUM) //此处add的数量必须和goroutine的数量一致
	for j := 0; j < ROUTINE_NUM; j++ {
		go func() {
			var a int
			for i := 0; i < 100; i++ {
				a += i
			}
			fmt.Printf("Index=%d,Value=%d\n", j, a)
			wg.Done() //每个routine结束时对waitgroup中的数字减一
		}()
	}
	wg.Wait()
}
