package main

import "fmt"

type fileError struct {
	s string
}

func (fe *fileError) Error() string {
	return fe.s + "3333"
}

//只是模拟一个错误
func openFile() ([]byte, error) {
	return nil, &fileError{"文件错误，自定义"}
}
func main() {
	conent, err := openFile()
	fmt.Printf("%T\n", err)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(conent))
	}
}
