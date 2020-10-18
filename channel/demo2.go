package main

func main() {
	/*worklist := make(chan []string)
	worklist <- os.Args[1:]

	for list := range worklist {
		fmt.Println(list)
	}*/
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"

	//fmt.Println("xxxx")
}
