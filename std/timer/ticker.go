package main

import (
	"time"
	"reflect"
	"fmt"
)

func main() {

	 f2()

}


func f1(){
	/*
	 Ticker use a similar mechanism to timers:
	 a channel that is sent values. Here we`ll use the range builtin on the channel
	 to iterate over the values as they arrive every one second
	 */
	ticker := time.NewTicker(10*time.Second)

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

func f2(){
	//ticker := time.NewTicker(time.Second*2)
	for now := range time.NewTicker(time.Second*2).C {
		fmt.Printf("now: %v \n",now.Format("2006-01-02 15:04:05"))
	}
}