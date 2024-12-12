package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用os.Readfile一次性将文件读取到位
	file := "d:/test.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file error:", err)
	}
	fmt.Println(string(content))
	// 我们没有显式Open文件 也不需要显式Close文件
	// open and close 被封装到readFile里面了
}
