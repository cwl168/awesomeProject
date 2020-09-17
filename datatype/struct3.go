package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Computer struct {
	Cpu       string    `json:"cup"`
	Ram       int       `json:"ram"`
	Rom       int       `json:"rom"`
	CreatedAt time.Time `json:"created_at"`
}

/*type Computer struct {
	cpu string `json:"cup"`
	ram int    `json:"ram"`
	rom int    `json:"rom"`
}*/

func main() {
	hp := &Computer{"i7", 2, 250} //{"cup":"i7","ram":2,"rom":250}

	res, err := json.Marshal(hp)
	if err == nil {
		fmt.Println(reflect.TypeOf(res))
		fmt.Println(string(res))
	}
}
