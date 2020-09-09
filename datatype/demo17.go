package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sept", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	endlessSummer := summer[:7]
	fmt.Println(endlessSummer)
	//fmt.Println(summer[:20]) //越界panic: out of range ，超过7的容量

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	a := []int{0, 1, 2, 3, 4, 5}
	s := []int{0, 1, 2, 3, 4, 5}
	if equal(a, s) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	v := make([]int, 4)
	y := make([]int, 4, 7)
	fmt.Println(len(v))
	fmt.Println(len(y))

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
