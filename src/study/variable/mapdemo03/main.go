package main

import (
	"fmt"
)

// 查找
func main() {
	cities := make(map[string]string)
	cities["no1"] = "茂名"
	cities["no2"] = "广州"
	cities["no3"] = "深圳"

	// 根据key查找
	value, ok := cities["no1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在key=no1")
	}

	// for-range 遍历
	for key, value := range cities {
		fmt.Printf("key=%s, value=%s\n", key, value)
	}

	//使用for-range遍历一个较复杂的map
	studentMap := make(map[string]map[string]string)
	studentMap["no1"] = make(map[string]string)
	studentMap["no1"]["name"] = "XDJ"
	studentMap["no1"]["age"] = "21"

	studentMap["no2"] = make(map[string]string)
	studentMap["no2"]["name"] = "ZY"
	studentMap["no2"]["age"] = "21"
	//遍历
	for key, value := range studentMap {
		fmt.Println("学生的序号为:", key)
		for studentKey, studentValue := range value {
			fmt.Printf("%v=%v\t", studentKey, studentValue)
		}
		fmt.Println()
	}
	fmt.Println("studentMap的长度为: ", len(studentMap))
}
