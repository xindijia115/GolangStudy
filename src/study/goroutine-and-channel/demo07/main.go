package main

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("SayHello", "hello world")
		time.Sleep(time.Second)
	}
}

func test() {
	// defer + recover
	defer func() {
		if err := recover(); err != nil {
			// assignment to entry in nil map
			fmt.Println(err)
			return
		}
	}()

	// 定义一个map不分配空间，这里会出现错误
	var m map[string]int
	m["age"] = 21
}

// defer + recover()捕获异常
func main() {
	// 开启两个协程
	go SayHello()
	// 出现异常 主程序不退出
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main", "Hello world")
		time.Sleep(time.Second)
	}

}
