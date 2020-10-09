package main

import (
	"fmt"
	"sort"
)

type Article struct {
	Title   string
	Desc    string
	Visited int
}
type byArticle []*Article

func main() {
	var groups byArticle
	groups = []*Article{
		{"Golang", "Hi, golang.", 13},
		{"Python", "Oh, python.", 20},
		{"Java", "SHit, java.", 19},
		{"Python", "Oh, python, It is good.", 43},
		{"Java", "SHit, java, it is bad.", 21},
		{"Golang", "Hi, golang, my love.", 13},
		{"Ruby", "Fuck, ruby.", 25}}
	sort.Sort(groups)
	for _, g := range groups {
		fmt.Printf("%s\t%d\t%s\n", g.Title, g.Visited, g.Desc)
	}
}

func (s byArticle) Len() int {
	return len(s)
}
func (s byArticle) Less(i, j int) bool {
	if s[i].Title != s[j].Title {
		return s[i].Title < s[j].Title
	}
	if s[i].Visited != s[j].Visited {
		return s[i].Visited < s[j].Visited
	}
	if len(s[i].Desc) != len(s[j].Desc) {
		return len(s[i].Desc) < len(s[j].Desc)
	}
	return false
}
func (s byArticle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
