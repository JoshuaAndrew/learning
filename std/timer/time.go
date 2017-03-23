package main

import (
	"time"
	"fmt"
)

func main() {

	time := time.Now()
	fmt.Println(time)
	fmt.Println(time.Format("2006-01-02 15:04:05"))
}
