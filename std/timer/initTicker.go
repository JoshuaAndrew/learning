package main

import (
	"time"
	"fmt"
)

func init()  {
	go run()
}

func run()  {
	for now := range time.NewTicker(2*time.Second).C {
		fmt.Printf("Node Ticker ticks at %s \n",now.Format("2006-01-02 15:04:05"))
	}
}

func main()  {
	fmt.Printf("Time: %s \n",time.Now().Format("2006-01-02 15:04:05"))
}
