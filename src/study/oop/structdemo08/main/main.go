package main

import (
	"fmt"
	"study/oop/structdemo08/model"
)

func main() {
	p := model.GetInstance("DJ", 21)
	fmt.Println(*p)
	fmt.Println("name=", p.Name, "age=", p.Age)
}
