/*

二元操作符：
优先级    操作符                                  备注
5        *, /, %, <<(左移), >>(右移), & ,&^      &^位清理操作符
4        +, –, |,  ^(两个操作数)                  ^异或(xor)
3        ==, !=, <, <=, >, >=
2        &&
1        ||

一元操作符包括：&, !, *, +, –, ^,++，--  (外加用于通信的<-)

一元操作符^是求补码/反码操作

 */
package main

import (
	"fmt"
	"math"
)

const s string = "Running"
/*
二进制运算符
二元运算:
 6: 0110
11: 1011
---------
 &: 0010 = 2  (与:一假即假)
 |: 1111 = 15 (或:一真即真)
 ^: 1101 = 13 (异或: 同假异真)
&^: 0100 = 4  (位清理: 第二个操作数是1就把第一个操作数置0)

一元运算:
 6: 0110
---------
 ^:



*/
func binary() {
	fmt.Println(1 << 10)   //左移
	fmt.Println(1 << 10 << 10 >> 10) //右移
	var (
		a = 6
		b = 11
	)
	// &(与), |(或), ^(异或) , &^(位清理), ^(取补码)
	fmt.Println(a & b, a | b, a ^ b, a &^ b, ^a)

	//+,-,*,/(整除),%(取模)
	fmt.Println(a + b, a - b, a * b, a / b, a % b)

}
/*
golang 中++ ,--不是作为表达式，而是作为语句的,因此其不能作为赋值语句出现在等号的右边
golang 中++ ,--不能放到变量的前面。
*/
func plusAndplus(){
	a:=1
	a++
	a++
	a--
	fmt.Println("a=",a)
}

func main() {
	binary()
	plusAndplus()
	//values
	fmt.Println("go" + "lang-" + "v1.7.4")
	fmt.Println("1+1=", 1 + 1)
	fmt.Println("7.0/3.0=", 7.0 / 3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	//variables
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 3, 4
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := 'c'
	fmt.Println(f)
	fmt.Println(string(f))

	//constants
	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const m = 3e20 / n  //浮点数
	fmt.Println(m)
	fmt.Println(3e20)

	fmt.Println(int64(m))
	fmt.Println(math.Sin(n))

}
