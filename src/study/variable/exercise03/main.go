package main

import "fmt"

func main() {
	// 三个班 每个班五个同学 统计每个班的平均成绩 和三个班的平均成绩

	var class = 3
	var students = 5
	var total = 0.0
	var score float64

	for i := 1; i <= class; i++ {
		var sum = 0.0
		for j := 1; j <= students; j++ {
			fmt.Printf("请输入%d班第%d位同学的成绩\n", i, j)
			_, err := fmt.Scanln(&score)
			if err != nil {
				return
			}
			sum += score
		}
		total += sum
		fmt.Printf("%d班的平均成绩为：%f\n", i, sum/float64(students))
	}
	fmt.Printf("三个班的平均成绩为%f\n", total/float64(class*students))
}
