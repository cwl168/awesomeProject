package main

import (
	"fmt"
	"sort"
)

func main() {
	/*names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)*/
	/*names := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Strings(names)*/
	names := []int{3, 4, 5, 2, 1}
	sort.Ints(names)
	/*names := sort.IntSlice{3, 4, 5, 2, 1}
	sort.Sort(names)*/
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%v\n", v)
	}
}
