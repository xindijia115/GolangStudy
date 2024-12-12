package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 自定义一个文件复制函数
func CopyFile(source string, dest string) (written int64, err error) {
	// 打开源文件
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()
	// 创建reader
	reader := bufio.NewReader(sourceFile)

	// 创建目的文件
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()
	// 创建Writer
	writer := bufio.NewWriter(destFile)
	return io.Copy(writer, reader)
}

func main() {
	source := "d:/a.jpg"
	dest := "d:/b.jpg"
	_, err := CopyFile(source, dest)
	if err != nil {
		fmt.Println("拷贝失败")
		return
	}
	fmt.Println("拷贝完成")
}
