package main

import (
	"fmt"
	"unsafe"
)

// 定义全局变量
var hobby = "唱 跳 rap"

// 多个变量时 可以这样写
var (
	a = "a"
	b = "b"
	c = "c"
)

func main() {
	// 先定义后赋值
	var i int
	i = 10
	fmt.Println("i=", i)
	// 定义并赋值
	var name string = "xdj"
	fmt.Println("name=", name)
	// 不声明变量类型 自动推导
	var sex = "男"
	fmt.Println("sex=", sex)

	// := 方式
	age := 21
	fmt.Println("age=", age)

	//多变量声名
	var n1, n2, n3 int = 1, 2, 3
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)

	fmt.Println("hobby=", hobby)

	fmt.Println(a, b, c)

	// 查看数据类型
	var n4 = 100
	fmt.Printf("n1的数据类型是%T", n4)
	fmt.Println()
	//查看占用内存大小
	fmt.Printf("n4占用字节数为%d", unsafe.Sizeof(n4))

}
