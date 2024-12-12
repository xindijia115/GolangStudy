package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开一个文件 file对象 file指针 file 文件句柄
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// file 是一个指针
	fmt.Printf("file=%#v\n", file)
	// 当函数退出时及时关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}(file)

	// 创建一个*Reader 是带缓冲的
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			fmt.Println(str)
			break
		}
		// 输出内容
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束~~~")
}
