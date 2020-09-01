package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade,omitempty"`
}

func main() {
	stu1 := Student{
		Name:  "Tom",
		Age:   18,
		Grade: "middle school",
	}

	stu2 := Student{
		Name: "LiLy",
		Age:  19,
	}

	stuByts1, _ := json.Marshal(&stu1) //转化为json串

	stuByts2, _ := json.Marshal(&stu2)

	fmt.Println("stu1:", string(stuByts1))
	fmt.Println("stu2:", string(stuByts2))
}
