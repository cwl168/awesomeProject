package person

import "fmt"

type Test struct {
	Name string
}

//定义不能导出的结构体
type person struct {
	Name string
	//首字母小写,不能导出了,其它包不能使用,类似 private
	age int
	sal float64
}

//定义工厂模式的函数 首字母大写 类似构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//提供一个Set方法 设置年龄
func (user *person) Setage(age int) {
	if age > 0 && age < 150 {
		user.age = age
	} else {
		fmt.Println("年龄数值不对！")
	}
}

//获取年龄
func (user *person) Getage() int {
	return user.age
}

//更改年龄
func (user *person) Updateage(age int) int {
	user.age = age
	return user.age
}

//更改姓名
func (user *person) Updatename(name string) string {
	user.Name = name
	return user.Name
}
