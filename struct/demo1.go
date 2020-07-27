package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type Person struct {
	id   int
	name string
}

func main() {
	var s []*Person
	s1 := Person{
		id:   1,
		name: "lin",
	}
	s2 := Person{
		id:   2,
		name: "cwl",
	}
	s3 := Person{
		id:   2,
		name: "cwl",
	}

	s = append(s, &s1)
	s = append(s, &s2)
	s = append(s, &s3)
	f := excelize.NewFile()
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
	categories := map[string]string{"A1": "ID", "B1": "姓名"}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range s {
		//f.SetCellValue("Sheet1", k, v)
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(k + 2), (*v).id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(k + 2), (*v).name)
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%T - %v \n",s,s)
	//fmt.Printf("%p - %v \n",s[0],*s[0])
	//fmt.Printf("%p - %v \n",s[1],*s[1])
}
