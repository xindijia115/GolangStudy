package main

import "fmt"

func sum(a, b int) int {
	// defer 后面的语句会按顺序压栈 先进后出
	defer fmt.Println("a=", a) // 第三输出
	defer fmt.Println("b=", b) // 第二输出

	res := a + b
	fmt.Println("res=", res) // 第一输出
	return res
}

// defer 延迟机制
func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) // 第四输出
	//
}
