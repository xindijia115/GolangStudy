package main

import "fmt"

// for 循环
func main() {
	var n = 10
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	// 另一种写法
	j := 0
	for j < n {
		fmt.Println(j)
		j++
	}

	////死循环
	//for {
	//	fmt.Println("dj")
	//	// break
	//}
	//等价
	/*
		for ; ; {

		}
	*/

	//遍历字符串 传统方式
	var str string = "hello world"
	for i := 0; i < len(str); i++ { // 注意使用这种方式中文会出现问题
		fmt.Println(string(str[i]))
	}
	// for-range
	for index, value := range str {
		fmt.Println(index, string(value))
	}
}
