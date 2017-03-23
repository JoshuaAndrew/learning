/*
Map
(1) key必须是支持==和!=比较运算的类型,不可以是func,map,slice
(2) map使用make创建  make([keyType]valueType,capacity),capacity容量可以省略
    超出容量后会自动扩容,但尽量提供一个合理的初始值,len(map)获取容量

*/
package main

import (
	"fmt"
)

func main() {
	var m1 map[int]string    //声明map
	m1 = make(map[int]string) //初始化map
	fmt.Println("map1", len(m1), m1)

	var m2 map[int]string
	m2 = map[int]string{} //初始化map
	fmt.Println("map2", len(m2), m2)

	var m3 map[int]string = make(map[int]string)
	fmt.Println("map3", len(m3), m3)

	m := make(map[int]string, 10)
	m[0] = "OK" //向map中添加数据
	m[1] = "ERROR"
	fmt.Println("map4", len(m), m)
	delete(m, 1) //删除map中数据
	fmt.Println("map5", len(m), m)

	mm := make(map[int]map[int]string)
	mm[1] = make(map[int]string)
	mm[1][1] = "wOk"
	for k, v := range mm {
		println("range", k, len(v))
	}

	a, ok := mm[2][1] //a标示值,ok标示当前值是否已经初始化
	if !ok {
		//这里ok标示
		mm[2] = make(map[int]string)
	}
	mm[2][1] = "ERROR"
	a = mm[2][1]
	println(a, ok)
}
