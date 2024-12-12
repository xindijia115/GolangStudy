package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B
	a.Num = 21
	b = B(a) // 结构体 A 和 B 中的变量类型和数目必须一样
	fmt.Println(a, b)
}
