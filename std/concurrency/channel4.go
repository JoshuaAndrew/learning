package main

import (
	"runtime"
	"sync"
)

//WaitGroup是golang standard library 中提供的routine之间同步的机制
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done() //每次减一
		defer println("A.defer")
		func() {
			defer println("B.defer")
			runtime.Goexit()
			println("B") // no
		}()
		println("A") // no
	}()
	wg.Wait()
}
