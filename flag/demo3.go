package main

import "flag"
import "fmt"
import "os"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("too less param")
		return
	}

	var ip = flag.Int("flagname", 1234, "help message for flagname")
	var real string
	flag.StringVar(&real, "dir", "this is a test", "help msg for dir")
	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(real)
}
