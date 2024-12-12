package main

import "fmt"

// map操作
func main() {
	// 增改
	a := make(map[string]string, 10)
	a["name"] = "DJ" // 增
	a["name"] = "ZY" // 修改

	// 删除 使用内置函数
	a["age"] = "21"
	delete(a, "age") // 删除, 如果key不存在 不进行操作 但也不会报错
	fmt.Println(a)

	// 假如我们想删除一个map中所有的数据，go没有提供一次性删除所有的方法，但我们可以遍历全部来进行删除
	// 或者 通过make 将原来的map指向一个空的map，原来的map就会被gc
	a = make(map[string]string)
	fmt.Println(a)
}
