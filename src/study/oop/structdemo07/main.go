package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test01() {
	p.Name = "DJ"
	fmt.Println("test01's name=", p.Name) // DJ
}

func (p *Person) test02() {
	(*p).Name = "ZY"
	fmt.Println("test02's name=", (*p).Name) // ZY
}

func main() {
	person := Person{"tom"}
	person.test01()
	fmt.Println(person.Name) // tom
	(&person).test01()       // 形式上传入的是地址，但本质还是值拷贝，不会影响原来的值
	fmt.Println(person.Name) // tom

	(&person).test02()
	fmt.Println(person.Name) // ZY
	person.test02()          // 形式上是传入的是值类型，但本质是传入的是地址，会修改原来的值
	fmt.Println(person.Name) // ZY
}
