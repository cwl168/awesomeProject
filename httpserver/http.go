package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const goroutinueNum = 2000

func main() {
	var wg sync.WaitGroup

	for i := 0; i < goroutinueNum; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			//resp, err := http.NewRequest("GET","http://127.0.0.1:8000/voiceapp/host/config?ip=2",nil)
			//resp, err := http.Get("http://127.0.0.1:8000/voiceapp/host/config?ip=2")
			resp, err := http.Get("http://127.0.0.1:8000/voiceapp/host/config?ip=2")
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("resp:", string(body))
		}()
		//time.Sleep(time.Millisecond*10)
	}
	wg.Wait()
	fmt.Println("done")
}
