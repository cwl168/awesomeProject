package main

import (
	"reflect"
	"testing"
)

// 初始化map
func initMap() map[int]int {
	m := map[int]int{}
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
	return m
}

func getKeys1(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getKeys2(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 初始化默认
func getKeys3(m map[int]int) []int {
	// 注意：由于数组默认长度为0，后面append时，需要重新申请内存和拷贝，所以效率较低
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 使用反射
func getKeys4(m map[int]int) int {
	// 注意:虽然此写法简洁,但MapKeys函数内部操作复杂,效率极低
	keys := reflect.ValueOf(m).MapKeys()
	return len(keys)
}

func BenchmarkMapkeys1(b *testing.B) {
	// 初始化map
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getKeys1(m)
	}
}
func BenchmarkMapkeys2(b *testing.B) {
	// 初始化map
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getKeys2(m)
	}
}

func BenchmarkMapkeys3(b *testing.B) {
	// 初始化map
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getKeys3(m)
	}
}

func BenchmarkMapkeys4(b *testing.B) {
	// 初始化map
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getKeys4(m)
	}
}
