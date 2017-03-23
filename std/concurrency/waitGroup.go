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
)

//注意输出结果无序,因为此处利用到了cpu的多核特性
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10) //此处add的数量必须和goroutine的数量一致
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, index int) {
			var a int
			for i := 0; i < 1000000; i++ {
				a += i
			}
			fmt.Printf("Index=%d,Value=%d\n", index, a)
			wg.Done() //每个routine结束时对waitgroup中的数字减一
		}(&wg, i)
	}
	wg.Wait()
}
