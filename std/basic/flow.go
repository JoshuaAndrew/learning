/*

for  range   if else  switch case default fallthrough
跳轉語句:break continue goto



 */
package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {

	//for 第一种情况：单条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(i)

	//for 第二种情况：初始化,条件
	for j := 5; j <= 10; j++ {
		if j % 2 == 0 {
			fmt.Println(j, "is even")
		} else {
			fmt.Println(j, "is odd")
		}
	}

	//for 第三种情况：
	for {
		//无限循环
		fmt.Println("loop")
		if num := 9; num < 0 {
			fmt.Println(num, "is negative")
		} else if num < 10 {
			fmt.Println(num, "has 1 digit")
		} else {
			fmt.Println(num, "has multiple digits")
		}
		break
	}

	// 这里我们使用range来计算一个切片的所有元素和
	// 这种方法对数组也适用
	nums := []int{2, 3, 4}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("index:",index)
	}
	fmt.Println("sum:", sum)


	vs := []interface{}{
		[]int{1, 2, 3},                //slice 切片
		[]int{1, 2, 3}[:],            //切片再切还是切片
		[3]int{1, 2, 3},              //array 数组，确定数组长度
		//[3]int{1, 2, 3}[:],         //数组不能使用切片操作符
		[...]int{1, 2, 3},            //array 数组，由编译器自动计算数组长度。
	}
	for i, v := range(vs) {
		rv := reflect.ValueOf(v)    //进入疯狂的reflect世界
		fmt.Println(i, rv.Kind())
	}

	// range函数用来遍历字符串时，返回Unicode代码点。
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "gopher" {
		fmt.Println(i, string(c))
	}


	/*
	switch case default fallthrough

	 可以使用任何类型或表达式作为条件语句
	 不需要写break，一旦条件符合就自动跳出switch语句语句
	 如果希望继续执行下一个case，需要使用fallthrough
         支持初始化表达式,右侧需要跟 ";"
	*/


	/*
	 you can use commas to separate multiple expression in the same case statement
	 we use the optional default case in this example as well
	 */
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is Weekend!")
	default:
		fmt.Println("Today is Weekday")
	}
        /*
         switch without an expression is an alternate way to express if/else logic.
         Here we also show how the case expressions can be no-constants.
         */
	now := time.Now().Hour();
	fmt.Println(now)
	switch {      //注意这里不能有表达式，本质就是if else
	case now < 12:
		fmt.Println("AM")
	case now >= 12 && now < 18:
		fmt.Println("PM")
	default:
		fmt.Println("Night")
	}


	a :=1
	switch a {//不需要写break，一旦条件符合就自动跳出switch语句语句
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	switch b := 1; { //支持初始化表达式,右侧需要跟 ";"
	case b<0:
		fmt.Println(b, "b<0")
		fallthrough   //如果希望继续执行下一个case，需要使用fallthrough
	case b>0:
		fmt.Println(b, "b>0")
		fallthrough
	case b<10:
		fmt.Println(b,"b<10")
	default:
		fmt.Println("b=0")
	}



        /*
        跳轉語句:break continue goto

        都可以配合标签使用跳出多层循环
        标签的名称区分大小写，若定义标签却未使用会导致编译错误
        break continue配合标签可以跳出多层循环,但是不会再次进入循环
        goto是调整程序执行的位置,与其他两个语句配合标签的效果并不相同

        */
	break_LABEL:
	for i:=0;i<10;i++{
		if i> 5{
			break break_LABEL
		}else{
			fmt.Println("break index:",i)
		}
	}

	for i:=0;i<10;i++{
		for{
			counter :=0
			if counter > 5 {
				break
			}else{
				fmt.Println("counter index:",counter)
			}
			counter++
		}
		fmt.Println("outer break index:",i)
	}

	continue_LABEL:
	for i:=0;i<10;i++{
		if i> 5{
			continue continue_LABEL
		}else{
			fmt.Println("continue index:",i)
		}
	}













}
