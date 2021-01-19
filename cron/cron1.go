package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "0/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
func print5() {
	log.Println("Run 5s cron")
}
func print15() {
	log.Println("Run 15s cron")
}
