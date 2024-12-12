package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

// 结构体是值类型
func main() {
	var monster1 Monster
	monster1.Name = "DJ"
	monster1.Age = 21

	monster2 := monster1
	monster2.Name = "Zy"

	fmt.Println(monster1)
	fmt.Println(monster2)
	//{DJ 21}
	//{Zy 21}
}
