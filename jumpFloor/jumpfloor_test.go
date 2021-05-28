package main

import "testing"

func BenchmarkJumpFloorNormal(b *testing.B) {
	num := 1000000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jumpFloorII(num)
	}
}

func BenchmarkJumpFloorAdvanced(b *testing.B) {
	num := 1000000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jumpFloorIIAdvanced(num)
	}
}
