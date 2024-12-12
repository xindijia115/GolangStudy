package main

import "fmt"

// slice 切片
func main() {
	// 声名一个数组 通过切片引用数组
	var intArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// 引用
	sliceArr := intArr[1:7]            // 取值左闭右开
	fmt.Println("sliceArr:", sliceArr) //[2 3 4 5 6 7]
	fmt.Println("len:", len(sliceArr)) // 元素个数
	fmt.Println("cap:", cap(sliceArr)) // 切片容量 len(intArr) - start

	var slice = make([]float64, 5, 10)
	slice[0] = 1.0
	fmt.Println("slice:", slice)    // slice: [1 0 0 0 0]
	fmt.Println("len:", len(slice)) // 5
	fmt.Println("cap:", cap(slice)) // 10
}
