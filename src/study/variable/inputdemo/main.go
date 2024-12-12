package main

import "fmt"

func main() {
	// 键盘获取输入
	// 方式1
	var name string
	var age byte
	var sal float64
	var isPass bool

	fmt.Println("请输入姓名 ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("请输入年龄 ")
	_, err = fmt.Scanln(&age)
	if err != nil {
		return
	}

	fmt.Println("请输入薪水 ")
	_, err = fmt.Scanln(&sal)
	if err != nil {
		return
	}

	fmt.Println("请输入是否通过考试  ")
	_, err = fmt.Scanln(&isPass)
	if err != nil {
		return
	}

	fmt.Printf("name=%v age=%v sal=%v isPass=%v\n", name, age, sal, isPass)

	//方式2 空格隔开
	fmt.Println("请输入学生信息 ")
	_, err = fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	if err != nil {
		return
	}
	fmt.Printf("name=%v age=%v sal=%v isPass=%v\n", name, age, sal, isPass)
}
