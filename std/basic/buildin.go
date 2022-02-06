/*
 golang 内置25个关键字:

break continue switch case default if else  for import return goto

package(定义包名) const(定义常量) var(定义变量) func(定义函数)  type(定义类型) interface(定义接口) struct(定义结构)

chan go select defer map fallthrough range

*/

/*
1.单个变量的声明:
变量的声明格式:  var <变量名称> <变量类型>
变量的赋值格式:  <变量名称> =  <表达式>
声明的同时赋值:  var <变量名称> [变量类型] =  <表达式> /  <变量名称>  :=  <表达式>
2.多个变量的声明:
   var (
      a = 234
      b = "sdf"
      e,f = "ddd", 9876
   )

   var c,d int
   c,d = 1,2

3.变量的类型转换:
golang中不存在隐式类型转换,所有类型转换必须显示声明
类型转换必须发生在两种相互兼容的类型之间
类型转换的格式: <ValueA> [:]= <TypeOfValueA> <ValueB>
	var b float = 1.1
	a := int(b)
4. 常量和枚举
  定义常量组时如果不提供初始值则使用上行的表达式,使用相同的表达式不代表具有相同的值
  iota是常量计数器,从0开始,常量组中每定义一个常量计数器递增1
  每遇到一个const关键字,iota就会重置为0
  通过初始化规则和iota可以达到枚举的效果
const(
   a = "A"
   b
   c = iota
   d
)

const (
   Monday = iota
   Tuesday
   Wednesday
   Thursday
   Friday
   Saturday
   Sunday
)

利用iota和<<运算符实现计算机中存储单位的枚举
const (
   _          = iota
   KB float64 = 1 << (iota * 10 )
   MB
   GB
   TB
   PB
   EB
   ZB
   YB

)

5. golang 15个内置函数
func append(slice []Type, elems ...Type) []Type
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
,len,cap,make,new,complex,real,imag,close,panic,recover,print,println

6. 指针
   Golang 保留了指针,但是与其他语言不同的是
   golang中指针不支持运算,不使用->操作符来引用对象,而是使用"."选择符来操作指针目标对象的的成员
   指针的默认值是nil
   操作符"&"表示取变量的地址,  "*" 间接访问目标对象
*/

package main

import (
	"fmt"
)


const (
	_         = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	_          = iota
	KB float64 = 1 << (iota * 10 )
	MB
	GB
	TB
	PB
	EB
	ZB
	YB

)

const(
	a = "A"
	b
	c = iota
	d
)

type myStruct struct {
	name   bool
	userid int64
}
//全局变量
var structZero myStruct

var intZero int
var int8Zero int8
var int32Zero int32
var runeZero rune
var int64Zero int64

var uintZero uint
var uint8Zero uint8
var byteZero byte
var uint32Zero uint32
var uint64Zero uint64

var boolZero bool

var float32Zero float32
var float64Zero float64

var stringZero string
var funcZero func(int) int
var byteArrayZero [5]byte
var boolArrayZero [5]bool
var byteSliceZero []byte
var boolSliceZero []bool
var mapZero map[string]bool
var interfaceZero interface{}
var chanZero chan int
var pointerZero *int

func main() {

	fmt.Println("week:",Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday)
	fmt.Println("Capaicty:",KB,MB,GB,TB,PB,EB,ZB,YB)
	fmt.Println("Value:",a,b,c,d)
	fmt.Println("structZero: ", structZero)
	fmt.Println("intZero: ", intZero)
	fmt.Println("int8Zero:", int8Zero)
	fmt.Println("int32Zero: ", int32Zero)
	fmt.Println("runeZero: ", runeZero)
	fmt.Println("int64Zero: ", int64Zero)
	fmt.Println("uintZero: ", uintZero)
	fmt.Println("uint8Zero: ", uint8Zero)
	fmt.Println("uint32Zero: ", uint32Zero)
	fmt.Println("uint64Zero: ", uint64Zero)
	fmt.Println("byteZero: ", byteZero)
	fmt.Println("boolZero: ", boolZero)
	fmt.Println("float32Zero: ", float32Zero)
	fmt.Println("float64Zero: ", float64Zero)
	fmt.Println("stringZero: ", stringZero)
	fmt.Println("funcZero: ", funcZero)
	fmt.Println("funcZero == nil?", funcZero == nil)
	fmt.Println("byteArrayZero: ", byteArrayZero)
	fmt.Println("boolArrayZero: ", boolArrayZero)
	fmt.Println("byteSliceZero: ", byteSliceZero)
	fmt.Println("byteSliceZero's len?", len(byteSliceZero))
	fmt.Println("byteSliceZero's cap?", cap(byteSliceZero))
	fmt.Println("byteSliceZero == nil?", byteSliceZero == nil)
	fmt.Println("boolSliceZero: ", boolSliceZero)
	fmt.Println("mapZero: ", mapZero)
	fmt.Println("mapZero's len?", len(mapZero))
	fmt.Println("mapZero == nil?", mapZero == nil)
	fmt.Println("interfaceZero: ", interfaceZero)
	fmt.Println("interfaceZero == nil?", interfaceZero == nil)
	fmt.Println("chanZero: ", chanZero)
	fmt.Println("chanZero == nil?", chanZero == nil)
	fmt.Println("pointerZero: ", pointerZero)
	fmt.Println("pointerZero == nil?", pointerZero == nil)
	//fmt.Println("nil == nil", nil == nil)

	varTest()
	pointer()
}

func varTest() {
	var (  //变量组
		a = "a"
		b = "b"
	)
	fmt.Println(a, b)
}

func pointer(){

	a := "34567"
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

}