package b

import (
	"learnDemo/import_cycle/a"
)

func Ride(a, b int) int {
	return a * b
}

func Result() int {
	sum2 := a.Calculate()
	return sum2
}
