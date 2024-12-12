package main

import "fmt"

// 结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var cat Cat = Cat{"ZY", 18, "粉红色", "爱睡觉"}
	fmt.Println(cat)

	var cat1 Cat
	cat1.Name = "DJ"
	cat1.Age = 21
	cat1.Color = "黑色"
	cat1.Hobby = "吃"
	fmt.Println(cat1)
}
