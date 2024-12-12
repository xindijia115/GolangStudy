package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println("手机开始工作")
}
func (p Phone) stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	name string
}

func (c Camera) start() {
	fmt.Println("相机开始工作")
}
func (c Camera) stop() {
	fmt.Println("相机停止工作")
}

func main() {
	// 定义一个接口数组 可以存放Phone 和 Camera结构体变量
	var usb [3]Usb
	usb[0] = Phone{"小米"}
	usb[1] = Phone{"vivo"}
	usb[2] = Camera{"索尼"}
	usb[0].start()
	usb[0].stop()
	fmt.Println(usb)

}
