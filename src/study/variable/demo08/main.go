package main

import (
	"fmt"
	"strconv"
)

// string类型转基本数据类型
func main() {
	var str string = "true"
	var b bool
	// 这个函数返回两个值 第二个值是我们不需要的可以用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Println(b)

	var str2 string = "12345"
	var n1 int64
	// 10 表示 十进制, 64表示int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Println(n1)

	var str3 string = "12.3456"
	var f float64
	f, _ = strconv.ParseFloat(str3, 64)
	fmt.Println(f)

}
