package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Birthday time.Time `xml:"birthday" json:"birthday" form:"birthday" time_format:"2006-01-02"`
}

//parsing time ""2019-03-04"" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "T"
func main() {
	s := []byte(`{"birthday": "2019-03-04"}`)
	p := Person{}
	fmt.Println(json.Unmarshal(s, &p))

}
