package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 记录开始时间
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test方法耗时:%v秒\n", end-start)
}
