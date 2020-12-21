package main

import "log"

func main() {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("recover:%v", e)
			}
		}()

		panic("cwl168")
	}()
	log.Println("go lear")
}
