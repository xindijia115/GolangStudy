package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

type MyFunType func(int, int) int

func myFun(varFun MyFunType, n1 int, n2 int) int {
	return varFun(n1, n2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 函数也是一种变量 可以把函数赋值给一个变量 使用变量进行函数调用
func main() {
	a := getSum
	res := a(1, 2)
	fmt.Println(res)

	// 既然函数作为一种数据类型 那么函数也可以作为调用函数的形参
	res2 := myFun(getSum, 10, 20)
	fmt.Println(res2)

	sum, sub := getSumAndSub(10, 5)
	fmt.Println(sum, sub)

}
