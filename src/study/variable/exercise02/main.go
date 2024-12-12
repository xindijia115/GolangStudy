package main

import "fmt"

func main() {
	var maxNum int = 100
	var cnt int
	var sum int

	for i := 0; i < maxNum; i++ {
		if i%9 == 0 {
			sum += i
			cnt++
		}
	}
	fmt.Printf("cnt is %d\n", cnt)
	fmt.Printf("sum is %d\n", sum)
}
