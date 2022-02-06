package main

import (
	"fmt"
	"reflect"
)

/*

Array:

定义数组的格式: var <varName> [n]Type, n>=0
数组长度也是数组的一部分，因此具有不同长度的数组为不同类型  [2]int [3]int
可以使用new创建数组,此方法返回一个指向数组的指针
golang 支持多维数组

Slice:
其本身并不是数组，它指向底层的数组
作为变长数组的替代方案，可以关联底层数组的局部或全部
为引用类型
可以直接创建或从底层数组获取生成
如果多个slice指向相同底层数组，其中一个的值改变会影响全部

一般使用make()创建
make([]T, len, cap)  其中cap可以省略，则和len的值相同,len表示存数的元素个数，cap表示容量
使用len()获取元素个数，cap()获取容量

*/

func main() {

	//array()
	slice()

}

func array() {

	var a  [2]int
	var b  [2]int
	a = b
	fmt.Println(a)

	c := [2]string{"a", "b"}
	fmt.Println(c)

	d := [20]int{19:100}//利用下标初始化数组指定位置的数据
	fmt.Println(d)

	d1 := [...]int{3:20, 9:100}//利用下标初始化数组指定位置的数据
	fmt.Println(d1)

	var p *[10]int = &d1
	fmt.Println(p) //p是数组的地址
	fmt.Println(*p)

	e := [...]string{"a", "b", "c", "d"} //当数组的长度已知，可以使用...让编译器自己获取长度
	fmt.Println(e)

	f := [...]int{'a', 'b', 'c', 'd'}//当数组的长度已知，可以使用...让编译器自己获取长度
	fmt.Println(f)

	g := new([10]int8)  //可以使用new创建数组,此方法返回一个指向数组的指针
	fmt.Println(g)
	fmt.Println(*g)

	h := [2][3]int{{1, 2, 3}, {4, 5, 6}}//多维数据
	fmt.Println(h)
	fmt.Println(len(h))

	var i  [10]interface{}  //此示例说明数组是定常的
	for k := 0; k < 11; k++ {
		i[k] = "weiwei"
	}
	fmt.Println(i)

	//冒泡排序
	//数据结构	 数组
	//最差时间复杂度	 O(n^2)
	//最优时间复杂度	 O(n)
	//平均时间复杂度	 O(n^2)
	//最差空间复杂度	 总共O(n)，需要辅助空间O(1)
	var data = [...]int{90, 67, 98, 2, 56, 8, 50}
	for i := 0; i <= len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
			}
		}
	}
	fmt.Println(data)

}

func slice() {
	var a []interface{} //声明一个类型是slice变量a
	fmt.Println(a, len(a), cap(a))

	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := b[5:10] //注意此处从数组中截取出索引5到10的元素组成slice c的类型已经变为slice
	d := b[0:5]  //注意此处从数组中截取出索引0到5的元素组成slice c的类型已经变为slice
	e := b[:5]  //注意此处从数组中截取出索引0到5的元素组成slice c的类型已经变为slice
	x := b[4:]
	be := b[:]  //把b复制到be中

	brv := reflect.ValueOf(b)
	crv := reflect.ValueOf(c)
	fmt.Println(b, brv.Kind()) //[0 1 2 3 4 5 6 7 8 9] array
	fmt.Println(c, crv.Kind()) //[5 6 7 8 9] slice
	fmt.Println(d, e, x, be)

	f := make([]interface{}, 0, 20) //定义一个初始容量是20的空slice，每当slice长度达到容量限制时，容量自动加倍
	fmt.Println(f, len(f), cap(f))
	for i := 10; i < 52; i++ {
		//func append(slice []Type, elems ...Type) []Type
		f = append(f, i)
		fmt.Println(len(f), cap(f), f)
	}

	g := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}//如果多个slice指向相同底层数组，其中一个的值改变可能会影响很多slice
	fmt.Println(g,reflect.ValueOf(g).Kind())  //[0 1 2 3 4 5 6 7 8] array
	s1 := g[2:7]
	s2 := g[1:3]
	fmt.Println(s1, s2) //[2 3 4 5 6] [1 2]
	s1[0] = 100
	//s2[1] = 100
	fmt.Println(s1, s2) //[100 3 4 5 6] [1 100]
	fmt.Println(g,reflect.ValueOf(g).Kind()) //[0 1 100 3 4 5 6 7 8] array
	fmt.Println("====================Slice============================")

	h := [...]int{1, 2, 3, 4, 5, 6, 7}//如果多个slice指向相同底层数组，其中一个的值改变会影响全部slice
	fmt.Println(reflect.ValueOf(h).Kind())
	s11 := h[2:7]
	s22 := h[1:3]
	fmt.Println(s11, s22) //[3 4 5 6 7] [2 3]
	s22 = append(s22, 1, 1, 1, 1, 1, 1, 1)
	//如果append元素的个数超过了slice的容量,那么底层就会重新分配一个数组
	//此时修改其中一个slice的值，不会影响其他slice的值
	s11[0] = 100
	fmt.Println(s11, s22) //[100 4 5 6 7] [2 3 1 1 1 1 1 1 1]


	//The copy  function copies elements from a source slice into a destination slice
	//Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst)
	//func copy(dst, src []Type) int
	var i = []int{1, 2, 3, 4, 5, 6, 7}
	var j = []int{8, 9, 10}
	num := copy(i, j)  //copy直接覆盖destination slice 元素，并不会追加
	fmt.Println(num, i, j)

	var i1 = []int{1, 2, 3, 4, 5, 6, 7}
	var j1 = []int{8, 9, 10}
	num1 := copy(j1, i1[5:7])
	fmt.Println(num1, i1, j1)

	var i2 = []int{1, 2, 3, 4, 5, 6, 7}
	var j2 = []int{8, 9, 10}
	num2 := copy(j2, i2) //copy不会导致slice动态增长
	fmt.Println(num2, i2, j2)



}