package main

import "fmt"

type Retriever struct {
	Contents string
}
type Retriever interface {
	Get(url string) string
}

type InterfaceRetriever interface {
	Get(url string) string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
func main() {
	var r = Retriever{"test"}
	fmt.Println(r.Get("ttttt"))
}
