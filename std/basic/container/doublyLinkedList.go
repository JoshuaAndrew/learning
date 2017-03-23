package main

import (
	"fmt"
	"container/list"
)

func main() {
	//doubly linked list  双向链表
	var doublyLinkedList = list.New()
	for i := 0; i < 10; i++ {
		doublyLinkedList.PushBack(i)
	}

	for e := doublyLinkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println(doublyLinkedList.Front().Value)
	fmt.Println(doublyLinkedList.Back().Value)
	fmt.Println("==================")
	doublyLinkedList.InsertAfter(99, doublyLinkedList.Back())
	doublyLinkedList.PushBack(100)
	for e := doublyLinkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//fmt.Println(doublyLinkedList.Init())

	fmt.Println(doublyLinkedList.Len())

	//
	doublyLinkedList.PushBack(1)
	fmt.Println(doublyLinkedList.Front().Next().Value) //1
	fmt.Println(doublyLinkedList.Front().Next().Prev().Value) //0
	value := doublyLinkedList.Remove(doublyLinkedList.Front())
	fmt.Println("value removed:", value)            //1
	e := doublyLinkedList.Front();
	if e != nil { //获取Element时主要注意判空
		value1 := doublyLinkedList.Remove(e) //panic: runtime error: invalid memory address or nil pointer dereference
		fmt.Println("first element is:",value1)
	}else{
		fmt.Println("List is empty")
	}

}
