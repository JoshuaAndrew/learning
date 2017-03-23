package main

import (
	"fmt"
)

func f1() (result int) {
	defer func() { //return语句不是一条原子调用，return xxx其实是: 赋值＋RET指令
		result++  //defer语句被插入到 赋返回值 和 RET指令之间
		fmt.Println("inner:",&result)
	}()
	fmt.Println("outer:",&result)
	return 1
}

func f11() (result int) {
	result = 0
	fmt.Println(&result)
	func() {
		result++
		fmt.Println(&result)
		fmt.Println(result)
	}()
	return
}

func f2() (r int) {
	t := 5
	defer func() { //CALL runtime.deferproc
		t = t + 5
		fmt.Println(t)
	}()
	return t //CALL runtime.deferreturn
}

func f22() (r int) {
	t := 5
	r = t
	func() { //CALL runtime.deferproc
		t = t + 5
		fmt.Println(t)
	}()
	return //CALL runtime.deferreturn
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Println(r)
	}(r)
	return 1
}

func f33() (r int) {
	r = 1
	fmt.Println(&r)
	fmt.Println("================")
	func(r int) {  //注意这里是参数传递进来的,所以匿名函数中的变量r和外部函数中的r是不同的变量
		fmt.Println(&r)
		r = r + 5
		fmt.Println(r)
	}(r)
	return
}

func main() {
	w := f1()
	fmt.Println(w)
}
