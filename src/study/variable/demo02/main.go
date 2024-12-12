package main

import "fmt"

// 浮点类型
func main() {
	// 默认为float64 精度更高
	var salary float32 = 99.99
	fmt.Println(salary)

	// 十进制形式
	n1 := 5.20
	n2 := .520
	fmt.Println(n1, n2)

	//科学计数法
	n3 := 5.1234e2
	fmt.Println(n3)
	n4 := 5.1234e-2
	fmt.Println(n4)
}
