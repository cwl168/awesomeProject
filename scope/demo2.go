package main

import "os"

func main() {
	fname := "1.txt"
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	buf := make([]byte, 4*1024)
	f.Read(buf)
	f.Close()
}
