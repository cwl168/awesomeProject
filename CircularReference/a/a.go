package a

import (
	"learnDemo/CircularReference/c"
)

func Calculate() int {
	sum := c.Ride(2, 2) + c.Add(1, 2)
	return sum
}
