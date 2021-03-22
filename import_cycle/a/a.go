package a

import "learnDemo/import_cycle/b"

func Add(a, b int) int {
	return a + b
}

func Calculate() int {
	sum := b.Ride(2, 2) + Add(1, 2)
	return sum
}
