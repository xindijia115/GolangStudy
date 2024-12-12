package main

import (
	"fmt"
	"time"
)

// 时间和日期
func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v type=%T\n", now, now)
	// 获取部分日期信息
	// 通过now获取年份月份
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期
	fmt.Printf("当前年月 %d %d %d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	str := fmt.Sprintf("当前年月 %d %d %d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(str)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
