package main

import (
	"fmt"
	"time"
)

// select 使用
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中的数据是:", num1)
	case num2 := <-ch2:
		fmt.Println("ch2中的数据是:", num2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout...")
	}
}
