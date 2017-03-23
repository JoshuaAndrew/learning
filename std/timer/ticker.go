package main

import (
	"time"
	"fmt"
	"reflect"
)

func main() {

	/*
	 Ticker use a similar mechanism to timers:
	 a channel that is sent values. Here we`ll use the range builtin on the channel
	 to iterate over the values as they arrive every one second
	 */
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			ty := reflect.TypeOf(t)
			fmt.Println(ty, "Ticker at", t)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}