package main

import "fmt"

func cal(n1 float64, n2 float64, operator byte) float64 {
	switch operator {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	default:
		return -1
	}
}

func main() {
	var n1 float64
	var n2 float64
	var operator byte
	fmt.Println("请输入两个数和运算符:")
	_, err := fmt.Scanf("%f %f %c", &n1, &n2, &operator)
	if err != nil {
		return
	}

	//var n1 float64 = 1.2
	//var n2 float64 = 2.3
	//var operator byte = '+'

	res := cal(n1, n2, operator)
	fmt.Println("result:", res)
}
