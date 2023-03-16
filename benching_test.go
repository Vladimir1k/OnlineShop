package main

import (
	"testing"
)

var arr = []string{"4, 5, 3, 2, 5", "SDFSGSD", "SDFSDFE<", "ES", "AXZ", "ZXVWS"}

func BenchmarkСoncat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(arr)
	}
}

func BenchmarkСoncat2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat2(arr)
	}
}
