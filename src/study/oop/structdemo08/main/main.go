package main

import (
	"GoProject/src/study/oop/structdemo08/model"
	"fmt"
)

func main() {
	p := model.GetInstance("DJ", 21)
	fmt.Println(*p)
	fmt.Println("name=", p.Name, "age=", p.Age)
}
