package main

import "fmt"

// map的使用
func main() {
	// 第一种方式
	// 声明一个map
	var a map[string]string
	// 分配一个空间 只有分配了一个空间才能使用
	a = make(map[string]string, 10)
	a["name"] = "XDJ"
	a["age"] = "21"
	a["name"] = "ZY" // key相同时覆盖前面的值
	fmt.Println(a)
	// map中的key是无序的

	// 第二种方式
	var b = make(map[string]string, 10)
	b["name"] = "DJ"
	b["age"] = "21"
	fmt.Println(b)

	// 第三种方式
	c := map[string]string{
		"name":  "DJ",
		"age":   "21",
		"hobby": "rap 打篮球",
	}
	fmt.Println(c)
}
