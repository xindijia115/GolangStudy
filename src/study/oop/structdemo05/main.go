package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"Name"` // `json:"Name"` 就是一个struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"` // 好像不写也可以？
}

func main() {
	person := Person{"DJ", 21, "coding"}
	// 使用json.Marshal进行序列化成json格式对象
	jsonStr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json", err)
	} else {
		fmt.Println(string(jsonStr))
		// {"Name":"DJ","age":21,"skill":"coding"}
	}
}
