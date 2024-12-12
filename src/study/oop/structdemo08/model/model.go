package model

// 小写 所以只能在本包中使用
type person struct {
	Name string
	Age  int
}

// 定义一个获取实例的函数
func GetInstance(name string, age int) *person {
	return &person{name, age}
}
