//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())//利用多核特性

	var c chan int = make(chan int)

	for i := 9; i >= 0; i-- {
		go myRoutine(c, i)
	}

	for i := 9; i >= 0; i-- {
		fmt.Println("Singal from Channel:",<-c)
	}

}

func myRoutine(c chan <- int, r int) {
	fmt.Println(r, "Routine GO GO!")
	c <- r
}