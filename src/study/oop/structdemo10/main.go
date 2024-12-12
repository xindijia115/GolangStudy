package main

import "fmt"

type A struct {
	name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A say ok", a.name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.name)
}

type B struct {
	A
	name string
}

func (b *B) SayOk() {
	fmt.Println("B say ok", b.name)
}

func main() {
	var b B
	b.name = "DJ" // B 本身的name属性
	b.age = 21    // 使用继承的A的属性
	b.SayOk()     // 使用B本省的方法
	b.A.name = "ZZ"
	b.hello() // 使用继承的A的方法
	fmt.Println(b)
}
