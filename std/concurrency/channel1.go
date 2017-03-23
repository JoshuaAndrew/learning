//GoSublime double ctrl+.
//GoSublime ctrl+.,ctrl+g
//GoSublime ctrl+.,ctrl+h
package main

import (
	"fmt"
)

func main() {
	var c chan bool = make(chan bool)
	//fmt.Printf("c1 %p\n", &c)
	go func() {
		fmt.Println("GO GO!")
		c <- true
		//fmt.Printf("c2 %p\n", &c)
		close(c) //channel关闭之后无法再向channel中存数据，但是可以继续从channel中读取数据
		//此处关闭了channel之后,下面的for循环就循环一次就结束
	}()
	//fmt.Printf("c3 %p\n", &c)
	//for v := range c { //无限循环等待从c中取数据,如果routine中的channel不关闭,此处会一直等待
	//	fmt.Println(v)
	//}
	fmt.Println(<-c)
}

