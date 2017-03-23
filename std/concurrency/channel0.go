//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
	"time"
)

//no channel
//二货阻塞main routine的办法
func main() {
	go f()
	time.Sleep(2 * time.Second)
}

func f() {
	fmt.Println("GO GO!")
}
