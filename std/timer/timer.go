package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second * 5)

	<-timer1.C
	fmt.Println("Timer 1 expired ")


	/*
	  you can cancel the timer before it expires
	 */
	timer2 := time.NewTimer(time.Second * 5)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop := timer2.Stop() // cannel the timer
	if stop {
		fmt.Println("Timer 2 stopped!")
	}

}