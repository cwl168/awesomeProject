package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bytes)
}
