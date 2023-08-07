package main

import "testing"

func BenchmarkSortOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := []int{33, 25, 97, 14, 21, 44, 9, 15, 16, 88}
		sort1(list)
	}
}

func BenchmarkSortTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := []int{33, 25, 97, 14, 21, 44, 9, 15, 16, 88}
		sort2(list)
	}
}
