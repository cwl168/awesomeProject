package mergeSort

import (
	"fmt"
	"testing"

	"github.com/fortytw2/leaktest"
)

func TestMergeSort1(t *testing.T) {
	slice := []int{8, 4, 5, 7, 1, 3, 2, 6}

	fmt.Println(mergeSort1(slice))
}

func TestMergeSort2(t *testing.T) {
	defer leaktest.Check(t)()
	mergeSort2()
	Gen()
}
