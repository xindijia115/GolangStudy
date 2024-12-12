package main

import (
	"errors"
	"fmt"
	_ "time"
)

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func test02() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_ = readConf("config2.ini")
	fmt.Println("test02()继续执行....")
}

func main() {
	//测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码...")
}
