package main

import "fmt"

func main() {
	originSlice := []int{1, 2, 3, 4, 7, 6, 1, 4, 7}
	sliceUnique := make([]int, 0, len(originSlice))
	mapExist := make(map[int]interface{}, 0)

	for _, v := range originSlice {
		if _, ok := mapExist[v]; !ok {
			mapExist[v] = true
			sliceUnique = append(sliceUnique, v)
		}
	}
	fmt.Println(sliceUnique, len(sliceUnique), cap(sliceUnique))
}
