package main

import "fmt"

func main() {
	//获取变量的地址
	var i int = 10
	fmt.Printf("i的地址为%v\n", &i)

	//指针
	var p *int
	p = &i
	// *p表示引用
	fmt.Printf("变量地址为%v; 值大小为%v", p, *p)
}
