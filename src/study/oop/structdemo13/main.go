package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var p Point = Point{1, 2}
	var a interface{}

	a = p // 向下转型

	var b Point
	// b = a // 这是不可以的
	b = a.(Point) // 类型断言
	fmt.Println(b)
}
