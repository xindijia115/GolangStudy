package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个全局变量map，各个写协程共享
var (
	myMap = make(map[int]uint64)
	// 声明一个互斥锁 : Mutex 互斥
	lock sync.Mutex
)

// 定义一个函数来求n的阶乘
func test(n int) {
	sum := uint64(1)
	for i := 1; i <= n; i++ {
		sum *= uint64(i)
	}
	// 加锁
	lock.Lock()
	myMap[n] = sum
	// 用完释放
	lock.Unlock()
}

// 求各个数的阶乘 并放到一个map中
func main() {
	// 开启多个协程来并发求阶乘
	for i := 1; i <= 10; i++ {
		go test(i)
	}

	// 主线程遍历map
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("%d的阶乘为：%d\n", k, v)
	}
	lock.Unlock()

	// 休眠一小会 防止主线程退出导致协程销毁
	time.Sleep(10 * time.Second)

	// 当我们运行时 会出现异常 ： concurrent map writes
	// 原因是 协程和主线程并发执行 协程在对全局变量map进行写 而主线程在对map进行读
	// 产生了资源利用问题（同一个资源不能同时被多个线程使用）
	// 解决方法：可以对共享资源进行加锁 使得某一时刻只能一个线程访问

}
