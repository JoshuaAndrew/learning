package main

/*
3. 通过浏览器访问

http://localhost:6060/debug/pprof/

能够查看到程序的overview

4.你也可以通过终端命令查看

Then use the pprof tool to look at the heap profile:

>> go tool pprof http://localhost:6060/debug/pprof/heap

Or to look at a 30-second CPU profile:

>> go tool pprof http://localhost:6060/debug/pprof/profile

Or to look at the goroutine blocking profile:

>> go tool pprof http://localhost:6060/debug/pprof/goroutine
>> go tool pprof http://localhost:6060/debug/pprof/block


*/

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	flag.Parse()

	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		time.Sleep(time.Millisecond * 100)
		counter++

	}
	wg.Done()
}
