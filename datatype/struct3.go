package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Computer struct {
	Cpu       string `json:"cup"`
	Ram       int    `json:"ram"`
	Rom       int    `json:"rom"`
	CreatedAt string `json:"created_at"`
}

/*type Computer struct {
	cpu string `json:"cup"`
	ram int    `json:"ram"`
	rom int    `json:"rom"`
}*/

func main() {
	str := "Hello World"
	fmt.Println(strings.Count(str, "")) //程序输出 1

	hp := &Computer{"i7", 2, 250, time.Now().Format(time.RFC3339)} //{"cup":"i7","ram":2,"rom":250}

	res, err := json.Marshal(hp)
	if err == nil {
		fmt.Println(reflect.TypeOf(res))
		fmt.Println(string(res))
	}
}
