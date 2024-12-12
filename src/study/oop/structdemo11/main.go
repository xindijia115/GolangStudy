package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

// Monster 实现了两个接口的全部方法 也就实现了接口
type Monster struct {
}

func (m *Monster) Say() {
	fmt.Println("Say...")
}
func (m *Monster) Hello() {
	fmt.Println("Hello...")
}

func main() {
	// Monster 实现了AInterface , BInterface
	var monster Monster
	var a AInterface = &monster
	var b BInterface = &monster

	a.Say()
	b.Hello()

}
