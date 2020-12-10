package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02"))
	/*go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	select {}*/
}
