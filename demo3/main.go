package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	log.Println(path)
	fmt.Println(os.Getenv("TERM"))
}

/**
➜  demo3 git:(master) ✗ ./demo3
2020/11/11 11:33:58 /Users/caoweilin/go/src/learnDemo/demo3/demo3
➜  demo3 git:(master) ✗ go run main.go
2020/11/11 11:34:15 /var/folders/kj/87srbqwn3tq1tzv0p4wknz0r0000gp/T/go-build503078963/b001/exe/main
➜  demo3 git:(master) ✗ go run main.go
2020/11/11 11:38:35 /var/folders/kj/87srbqwn3tq1tzv0p4wknz0r0000gp/T/go-build627995643/b001/exe/main
➜  demo3 git:(master) ✗ go run main.go
2020/11/11 11:38:44 /var/folders/kj/87srbqwn3tq1tzv0p4wknz0r0000gp/T/go-build541432292/b001/exe/main
➜  demo3 git:(master) ✗ go run main.go
2020/11/11 11:38:46 /var/folders/kj/87srbqwn3tq1tzv0p4wknz0r0000gp/T/go-build115641852/b001/exe/main
➜  demo3 git:(master) ✗ go run main.go
2020/11/11 11:38:48 /var/folders/kj/87srbqwn3tq1tzv0p4wknz0r0000gp/T/go-build099015284/b001/exe/main



*/
