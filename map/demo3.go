package main

import (
	"fmt"
	"image"
)

func main() {
	icons := make(map[string]image.Image)
	if icons != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}
