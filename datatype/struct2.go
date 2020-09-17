package main

type E interface {
	Start()
	Stop()
	Renewal(x int)
}

type W interface {
	Open()
	Close()
	Renewal()
}

type Car struct {
	E
	W
}

func main() {
	var car Car
	car.Start()
	car.Stop()
	car.Open()
	car.Close()
	car.E.Renewal(5)
	car.W.Renewal()
	//car.Renewal()

}
