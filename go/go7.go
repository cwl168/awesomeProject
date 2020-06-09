package main

func main() {
	ch := make(chan string)
	go func() {
		ch <- "send"
	}()
	time.Sleep(time.Second * 3)
}
