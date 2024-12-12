package main

import "fmt"

func main() {
	// 假如还有97天放假 问还有 xx星期xx天
	var n = 97
	var week = n / 7
	var day = n % 7
	fmt.Printf("还有%v个星期%v天\n", week, day)
	value := getValue(99.9)
	fmt.Println("摄氏温度为", value)
}

func getValue(a float64) float64 {
	return 5.0 / 9 * (a - 100)
}
