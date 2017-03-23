/*
Method Expressions:
MethodExpr    = ReceiverType "." MethodName
ReceiverType  = TypeName | "(" "*" TypeName ")" | "(" ReceiverType ")"

*/
package main

import (
	"fmt"
)

// 先定义结构体
type Rectangle struct {
	width  int
	height int
}

//矩形的面积
func (r *Rectangle) Area() int { /* area() whose receiver is of type *T */
	r.modify(r)
	return r.height * r.width
}

func (r *Rectangle) Area() int { /* area() whose receiver is of type *T */
	return r.height * r.width
}

//矩形的周长
func (r *Rectangle) Perimeter() int {
	return 2 * (r.height + r.width)
}

func (r *Rectangle) modify(rect *Rectangle) int {
	if rect.height < 50 {
		r.height = 50
	}
	return r.height * r.width
}

//初始化struct的方法:
//【1】 得到Rectangle对象的地址
//rect :=new(Rectangle)
//rect :=&Rectangle{}
//rect :=&Rectangle{10,20}
//rect :=&Rectangle{width:10,height:20}
//定义 + 初始化同时进行
//rect := &struct{width int, height int}{10, 20}  //匿名结构体
//rect := &struct{}  //匿名空结构体
//【2】 得到Rectangle对象
//rect := Rectangle{}
//rect := Rectangle{10,20}
//rect := Rectangle{width:10,height:20}
//rect := struct{width int, height int}{10, 20}  //匿名结构体
//rect := struct{}  //匿名空结构体
func main() {
	rect := new(Rectangle)
	//rect := &Rectangle{height: 20, width: 20}
	rect.height = 20
	rect.width = 20
	area := rect.area()
	perimeter := rect.perimeter()
	modify := rect.modify(rect) //如果高小于50就重置为50
	fmt.Println(area)
	fmt.Println(perimeter)
	fmt.Println(modify)
}
