package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 结构体定义方式
func main() {
	// 方式3
	//var person *Person = new(Person)
	var person *Person = &Person{}
	// 通过解引用的方式给属性赋值
	(*person).Name = "DJ"
	// 也可以这样
	person.Name = "ZY" // 编译器会优化成 (*person).Name
	(*person).Age = 21
	fmt.Println(*person)
	// {ZY 21}

}
