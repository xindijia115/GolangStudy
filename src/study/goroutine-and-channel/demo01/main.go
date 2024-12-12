package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("test() 协程执行")
		time.Sleep(500 * time.Millisecond)
	}
}

// 主线程 和 协程
func main() {
	// 开启一个协程
	go test()

	// 主线程也执行循环100次
	for i := 0; i < 100; i++ {
		fmt.Println("main() 主线程执行")
		time.Sleep(500 * time.Millisecond)
	}

	// 我们可以看到协程和主线程交替进行

}
