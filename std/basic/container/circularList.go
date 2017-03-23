package main

import (
	"fmt"
	"container/ring"
)

func main() {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		fmt.Println("r",i)
		r = r.Next()
	}
	fmt.Println("========================")
	//moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring and returns that ring element
	//r = r.Move(12)        //6%10
	//fmt.Println(r.Value) //6

	//remove n % r.Len() elements from the ring starting at r.Next().
	//If n % r.Len() == 0 r remains unchanged
	fmt.Println("r.Next",r.Next().Value)
	r1 := r.Unlink(19)   //移除19%10=9个元素，
	fmt.Println("========================")

	for i := 0; i < r1.Len(); i++ {
		fmt.Println("r1",r1.Value)
		r1 = r1.Next()
	}
	fmt.Println("r len", r.Len())  //10-9=1
	fmt.Println("r1 len", r1.Len()) //9
	for i := 0; i < r.Len(); i++ {
		fmt.Println("r",i)
		r = r.Next()
	}
}
