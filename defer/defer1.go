package main
import "fmt"

func main() {
	name := "go"
	defer fmt.Println(name)

	name = "python"
	fmt.Println(name)
}
