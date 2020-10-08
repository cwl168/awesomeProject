package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"
)

type Tags []Tag

func (t Tags) Len() int {
	return len(t)
}

func (t Tags) Less(i, j int) bool {
	ti := t[i].Commit.CommittedDate
	tj := t[j].Commit.CommittedDate
	return ti.Before(tj) // 按照 committed_date 进行升序
}

func (t Tags) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func main() {
	var tag Tags
	err := json.Unmarshal([]byte(sourceDatas), &tag)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	sort.Sort(tag)
	fmt.Printf("排序后的的数据 %s", tag)
}

type Tag struct {
	Name   string `json:"name"`
	Commit struct {
		CommittedDate time.Time `json:"committed_date"`
	} `json:"commit"`
}

// 需要排序的json
var sourceDatas = `[
    {
        "name":"v0.9.4",
        "commit":{
            "committed_date":"2020-02-28T20:14:27+08:00"
        }
    },
    {
        "name":"v0.11.5",
        "commit":{
            "committed_date":"2020-07-01T17:15:53+08:00"
        }
    },
    {
        "name":"v0.14.6",
        "commit":{
            "committed_date":"2020-08-02T18:20:22+08:00"
        }
    },
    {
        "name":"v0.9.7",
        "commit":{
            "committed_date":"2020-03-25T17:04:17+08:00"
        }
    }
]`
