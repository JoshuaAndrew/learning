/*
  利用channel阻塞主routine
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Doing something…")
		c <- 234
		close(c)
	}()
	fmt.Println("I am Undone!")

	r, ok := <-c
	fmt.Println("Done!",r,ok)
}
