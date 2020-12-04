package b

import (
	"learnDemo/CircularReference/a"
)

func Result() int {
	sum2 := a.Calculate()
	return sum2
}
