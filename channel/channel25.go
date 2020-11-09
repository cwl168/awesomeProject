package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var abort = make(chan struct{})

func cancelled() bool {
	select {
	case <-abort:
		return true
	default:
		return false
	}
}

type entry struct {
	url    string
	nbytes int64
}

func main() {
	errors := make(chan error, 2)
	entries := make(chan entry, 2)
	var n sync.WaitGroup

	n.Add(1)
	go getBaidu(&n, entries, errors)

	n.Add(1)
	go getOther(&n, entries, errors)

	go func() {
		n.Wait()
		close(entries)
	}()

loop:
	for {
		select {
		case err := <-errors:
			fmt.Printf("%v\n", err)
			close(abort)
		case item, ok := <-entries:
			if !ok {
				break loop
			}
			fmt.Printf("%d\t%s\n", item.nbytes, item.url)
		}
	}

	fmt.Println("end")
}

func getBaidu(n *sync.WaitGroup, item chan<- entry, outErr chan<- error) {
	defer n.Done()
	time.Sleep(3 * time.Second)
	if cancelled() {
		fmt.Println("baidu cancelled")
		return
	}
	url := "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		outErr <- fmt.Errorf("fetch url err: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		outErr <- fmt.Errorf("status code err: %v", resp.StatusCode)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		outErr <- fmt.Errorf("io.Copy err: %v", err)
		return
	}
	item <- entry{
		url:    url,
		nbytes: nbytes,
	}
	return
}

func getOther(n *sync.WaitGroup, item chan<- entry, outErr chan<- error) {
	defer n.Done()
	if cancelled() {
		fmt.Println("other cancelled")
		return
	}
	url := "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		outErr <- fmt.Errorf("fetch url err: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		outErr <- fmt.Errorf("status code err: %v", resp.StatusCode)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		outErr <- fmt.Errorf("io.Copy err: %v", err)
		return
	}
	item <- entry{
		url:    url,
		nbytes: nbytes,
	}
	return
}
