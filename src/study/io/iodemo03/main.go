package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件写入
func main() {
	// 创建一个文件 写入内容
	filePath := "d:/abc.txt"
	// 仅写 或不存在时创建
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	// 关闭句柄
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file err:", err)
			return
		}
	}(file)
	str := "hello world\r\n"
	// 使用*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush err:", err)
		return
	}
}
