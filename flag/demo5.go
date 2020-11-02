package main

import (
	"flag"
	"log"
)

//go run flag/demo5.go go -name=as
var name string

//子命令
func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")

	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name:%s", name)
}
