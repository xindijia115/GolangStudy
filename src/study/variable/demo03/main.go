package main

import "fmt"

// 一般用byte来保存单个字符(不是单个汉字)
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	// 输出的是对应的码值
	fmt.Println(c1, c2)
	//格式化输出
	fmt.Printf("%c\n", c1)
	fmt.Printf("%c\n", c2)

	// 汉字可以使用int 类型存储
	var c3 int = '东'
	fmt.Println(c3)
	fmt.Printf("%c\n", c3)
}
