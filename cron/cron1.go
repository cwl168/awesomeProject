package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

var db int

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/5 * * * * ?", print5)
	c.AddFunc("*/15 * * * * ?", print15)
	c.Start()
	println(db)
	select {}
}
func print5() {
	db = 5
	log.Println("Run 5s cron")
	println(db)
}
func print15() {
	db = 15
	log.Println("Run 15s cron")
	println(db)
}
