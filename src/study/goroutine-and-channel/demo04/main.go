package main

import "fmt"

// 演示 channel 的读写
func main() {
	// 创建一个可以存放3个int的管道
	intChan := make(chan int, 3)

	// 看看intChan是什么
	fmt.Printf("intChan = %v, *intChain = %v\n", intChan, &intChan)

	// 向管道写入数据
	intChan <- 10
	intChan <- 20
	intChan <- 30
	// 写入管道的数据不能超过其容量 取出后可以继续写入
	<-intChan
	intChan <- 40

	// 查看管道长度和容量
	fmt.Printf("intChan的长度为%v, intChain的容量为%v\n", len(intChan), cap(intChan))

	// 从管道中读取数据
	n1 := <-intChan
	fmt.Println("n=", n1)
	fmt.Printf("intChan的长度为%v, intChain的容量为%v\n", len(intChan), cap(intChan))

	n2 := <-intChan
	fmt.Println("n=", n2)
	n3 := <-intChan
	fmt.Println("n=", n3)

	// 如果管道的数据全部取出 继续取的话会出现错误 deadlock
	//n4 := <-intChan
	//fmt.Println("n=", n4)

}
