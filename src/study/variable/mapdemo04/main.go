package main

import "fmt"

// map 切片
func main() {
	/*
		要求：使用一个map来记录student的信息 name 和 age, 也就是说一个
		student对应一个map,并且学生的个数可以动态的增加=>map切片
	*/

	student := make([]map[string]string, 2)
	student[0] = make(map[string]string, 2)
	student[0]["name"] = "XDJ"
	student[0]["age"] = "21"

	student[1] = make(map[string]string, 2)
	student[1]["name"] = "ZY"
	student[1]["age"] = "21"

	// 新增
	newStudent := make(map[string]string, 2)
	newStudent["name"] = "ZY"
	newStudent["age"] = "21"

	student = append(student, newStudent)
	fmt.Println(student)
}
