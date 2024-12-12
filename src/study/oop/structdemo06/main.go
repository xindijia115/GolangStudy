package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 方法和结构体绑定，因为结构体是值类型传递，如果需要改变原结构体属性，可以传一个结构体指针
func (p *Person) walk() {
	fmt.Printf("%s is walking\n", p.Name)
}

func (p *Person) String() string {
	return fmt.Sprintf("Name is %s, Age is %d\n", p.Name, p.Age)
}

func main() {
	person := Person{"DJzz", 21}
	(&person).walk() // DJ is walking

	fmt.Println(&person)

}
