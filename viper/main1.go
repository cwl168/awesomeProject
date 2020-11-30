package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watch, _ := fsnotify.NewWatcher()
	defer watch.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)

			}
		}
	}()
	//填写要监听的目录或者文件
	_ = watch.Add("/Users/caoweilin/go/src/learnDemo/viper/config.yaml")
	fmt.Println(<-done)
}
