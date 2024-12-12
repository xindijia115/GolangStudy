package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)

	// 自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
