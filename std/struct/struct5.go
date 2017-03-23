/**

golang struct 支持多继承

 */
package main

import "fmt"

type Father struct {
	name string
}

func (this *Father) Say() string {
	return "Hello I am " + this.name
}

type Mother struct {
	name string
}

func (this *Mother) Say() string {
	return "Hello I am " + this.name
}

type Child struct {
	name string
	Father
	Mother
}

func (this *Child) Say() string {
	return "Hello I am " + this.name
}

func main() {
	c := new(Child)
	c.Father.name = "牛魔王"
	c.Mother.name = "铁扇公主"
	c.name = "红孩儿"
	fmt.Println(c.Father.Say())
	fmt.Println(c.Mother.Say())
	fmt.Println(c.Say())
}