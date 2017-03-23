package main

import (
	"fmt"
)

func main() {
	sliceCopy2()
	rangeMap()
}

func sliceCopy() {
	sm := make([]map[int]string, 5) //声明一个容量为5的slice
	for _, v := range sm {
		//v是sm的拷贝
		v = make(map[int]string, 1) //声明一个容量为1的map
		v[0] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
}

func sliceCopy2() {
	sm := make([]map[int]string, 5) //声明一个容量为5的slice
	for i := range sm {
		//v是sm的拷贝
		sm[i] = make(map[int]string, 1) //声明一个容量为1的map
		sm[i][i] = "OK"                 //给slice中第i个map赋值
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}

func rangeMap() {
	// 使用range来遍历字典的时候，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

}