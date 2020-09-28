package main

import (
	"fmt"
	"time"
)

func main() {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds())
}
