package main

import "fmt"

// 基本数据类型转换
func main() {
	// 变量i本身的数据类型并没有变化
	var i int32 = 100
	// 强制转化
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("n1 = %v; n2 = %v; n3 = %v\n", n1, n2, n3)
	// 大的数据类型转化为小的数据类型不会报错 但会溢出

}
