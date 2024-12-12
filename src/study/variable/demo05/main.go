package main

import "fmt"

// string 类型
func main() {
	var address string = "大学城广东外语外贸大学"
	fmt.Println(address)

	//字符串一旦赋值了就不可以被修改
	var s = "hello"
	// s[0] = 'a' 这样是错误的
	s = "dj"
	fmt.Println(s)

	// 使用反引号能够将字符串按照原来的格式输出
	s1 := `
		package main
		
		import "fmt"
		
		// string 类型
		func main() {
			var address string = "大学城广东外语外贸大学"
			fmt.Println(address)
			
			//字符串一旦赋值了就不可以被修改
			var s = "hello"
			// s[0] = 'a' 这样是错误的
			s = "dj"
			fmt.Println(s)
		}
		`
	fmt.Println(s1)

	// 字符串拼接
	var str = "hello" + "world"
	str += "dj"
	fmt.Println(str)

	//拼接操作很长时 注意加号要留在行尾
	str4 := "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world"
	fmt.Println(str4)
}
