package main

import "os"

//变量f的作用域只有在if语句内，因此后面的语句将无法引入它，这将导致编译错误。
func main() {
	fname := "1.txt"
	if f, err := os.Open(fname); err != nil {
		return
	}
	buf := make([]byte, 4*1024)
	f.Read(buf)
	f.Close()
}
