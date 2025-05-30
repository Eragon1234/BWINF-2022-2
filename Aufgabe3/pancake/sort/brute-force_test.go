package sort

import (
	"BWINF/Aufgabe3/pancake"
	"testing"
)

func benchmarkBruteForce(input pancake.Stack, b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(input)
	}
}

func BenchmarkBruteForceWithLength14(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{11, 5, 6, 12, 1, 14, 9, 7, 3, 2, 8, 10, 13, 4}, b)
}

func BenchmarkBruteForceWithLength13(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{2, 8, 3, 9, 12, 13, 1, 6, 10, 5, 11, 4, 7}, b)
}

func BenchmarkBruteForceWithLength15(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{6, 10, 5, 9, 3, 11, 7, 15, 1, 2, 13, 12, 4, 8, 14}, b)
}

func BenchmarkBruteForceWithLength16(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{11, 16, 14, 1, 9, 12, 4, 2, 6, 13, 7, 3, 15, 10, 5, 8}, b)
}

func BenchmarkBruteForceWithLength11(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{6, 3, 7, 9, 2, 8, 4, 11, 1, 10, 5}, b)
}

func BenchmarkBruteForceWithLength8(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{2, 4, 6, 3, 5, 7, 1, 8}, b)
}

func BenchmarkBruteForceWithLength5(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{1, 5, 4, 2, 3}, b)
}

func BenchmarkBruteForceWithLength7(b *testing.B) {
	benchmarkBruteForce(pancake.Stack{5, 2, 4, 7, 1, 3, 6}, b)
}
