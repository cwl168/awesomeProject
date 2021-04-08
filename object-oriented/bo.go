package object_oriented

import "fmt"

type St struct {
	name string //相当于成员变量
}

//成员方法
func (t *St) get_name() string {
	return t.name + "test"
}

type St1 struct {
	St
}

//多态
type Si interface {
	inter() string
}
type T1 struct {
}
type T2 struct {
}

func (t1 T1) inter() string {
	return "T1"
}
func (t1 T2) inter() string {
	return "T2"
}

func GetName() {
	//封装
	t := St{name: "cwl"} //相当于new对象
	fmt.Println(t.get_name())

	//继承
	t1 := St1{St{name: "lin"}}
	fmt.Println(t1.get_name())

	//多态
	var si Si
	si = T1{}
	si = T2{}
	fmt.Println(si.inter())
}
