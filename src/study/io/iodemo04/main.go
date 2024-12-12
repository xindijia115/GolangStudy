package main

import "os"

func main() {
	source := "d:/abc.txt"
	target := "d:/kkk.txt"
	// 一次性全读取一个文件
	data, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}

	// 一次性写入
	err = os.WriteFile(target, data, 0666)
	if err != nil {
		panic(err)
	}
}
