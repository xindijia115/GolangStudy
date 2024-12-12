package main

import "fmt"

type Goods struct {
	name  string
	price int
}

type Books struct {
	Goods  // 嵌套匿名结构体
	writer string
}

func main() {
	var book Books
	//结构体可以使用匿名结构体的所有字段和方法 包括大写和小写的
	book.name = "天龙八部"
	book.price = 10
	book.writer = "DJ"
	fmt.Println(book)
}
