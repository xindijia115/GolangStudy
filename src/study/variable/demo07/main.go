package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型转化为string类型
func main() {
	var num1 int = 99
	var num2 float64 = 99.9999
	var b bool = true

	var str string
	//方式1
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)
	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)
	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

	// 方式2
	var num3 int = 99
	var num4 float64 = 99.9999
	var b1 bool = true

	// 10 表示十进制
	str = strconv.FormatInt(int64(num3), 10)
	str = strconv.Itoa(num3)
	fmt.Println(str)
	// 'f'字符串格式， 10表示小数点保留10位，64表示是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Println(str)
	str = strconv.FormatBool(b1)
	fmt.Println(str)
}
