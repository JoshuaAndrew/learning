package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	value []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

//放入元素
func (s *Stack)Push(v ...interface{}) {
	s.value = append(s.value, v...)
}

//返回下一个元素,并从Stack移除元素
func (s *Stack)Pop() (err error) {
	if s.Size() > 0 {
		s.value = s.value[:s.Size() - 1]
		return nil
	}
	return errors.New(" Stack is Empty!")

}

//返回栈顶元素
func (s *Stack) Top() (v interface{}) {
	if s.Size() > 0 {
		return s.value[s.Size() - 1]
	}
	return nil
}


func (s *Stack) Size() int {
	return len(s.value)
}

//是否为空
func (s *Stack)Empty() (bool) {
	if s.Size() == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	s := NewStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
		s.Push("v")
	}
	fmt.Println(s.Size(),s.Empty(),s.value)
	s.Pop()
	fmt.Println(s.Size(),s.Empty(),s.value)
}
