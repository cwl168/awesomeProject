package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
				l
			}
		}
	}()
}
