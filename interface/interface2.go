package main

import "fmt"

func main() {
	/*values := map[string]interface{}{
		"modified_by": 4,
		"state":       5,
	}
	// map 为引用类型
	update(values)
	fmt.Println(values)*/
	i := 1
	update2(i)
	fmt.Println(i)

	ps := new([]string)
	fmt.Println(*ps)
	*ps = append(*ps, "aa")
	fmt.Println(*ps)

	a := make([]int, 20)
	b := make([]int, 42)
	a = append(a, b...)
	fmt.Println(len(a), cap(a))

	e := []string{"my", "name", "is"}
	e = append(e, "cwl")

	fmt.Println(len(e), cap(e))

}
func update(vv map[string]interface{}) {
	//v['state'] = 2323

	for key, value := range vv {
		switch val := value.(type) {
		case string:
			fmt.Printf("string:[%v],key:[%s]\n", val, key)
		case int:
			vv["modified_by"] = 44444
			fmt.Printf("int:[%v],key:[%s]\n", val, key)
		}
	}
}

func update2(value interface{}) {

	switch val := value.(type) {
	case int:
		fmt.Printf("int:[%v]\n", val)
		value = 44444
	}
}
