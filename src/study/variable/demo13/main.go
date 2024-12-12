package main

import "fmt"

// 数组初始化
func main() {
	var a [4]int
	a[0] = 1
	fmt.Println(a)
	// 方式一
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1", arr1)

	// 方式2 [...]会自动推导数组的大小
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr2", arr2)

	// 方式3
	var arr3 = [3]int{1, 2, 3}
	fmt.Println("arr3", arr3)

	// 方式4 指定下标初始化 左边的为下标 以最大的下标+1作为数组大小 没有赋值的下标赋默认值
	var arr4 = [...]int{1: 100, 0: 99, 2: 8888}
	fmt.Println("arr4", arr4)

	f := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(f)

	e := [5]int{4: 100} // [0 0 0 0 100]
	fmt.Println(e)

	//类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05", strArr05)
}
