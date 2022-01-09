package main_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stuartmscott/foamsort"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestFoamSort(t *testing.T) {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	main.FoamSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	})
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slice)
}

func TestBubbleSort(t *testing.T) {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	main.BubbleSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	})
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slice)
}

const EXPENSIVE_COMPARISON = true

// Best Case: Sorted Integers

func BenchmarkFoamSort_Best_1K(b *testing.B) {
	benchmarkFoamSort_Best(b, 1000)
}

func BenchmarkBubbleSort_Best_1K(b *testing.B) {
	benchmarkBubbleSort_Best(b, 1000)
}

func BenchmarkStandardSort_Best_1K(b *testing.B) {
	benchmarkStandardSort_Best(b, 1000)
}

// Worst Case: Reverse-Sorted Integers

func BenchmarkFoamSort_Worst_1K(b *testing.B) {
	benchmarkFoamSort_Worst(b, 1000)
}

func BenchmarkBubbleSort_Worst_1K(b *testing.B) {
	benchmarkBubbleSort_Worst(b, 1000)
}

func BenchmarkStandardSort_Worst_1K(b *testing.B) {
	benchmarkStandardSort_Worst(b, 1000)
}

// Random Case: Random Integers

func BenchmarkFoamSort_Random_1K(b *testing.B) {
	benchmarkFoamSort_Random(b, 1000)
}

func BenchmarkBubbleSort_Random_1K(b *testing.B) {
	benchmarkBubbleSort_Random(b, 1000)
}

func BenchmarkStandardSort_Random_1K(b *testing.B) {
	benchmarkStandardSort_Random(b, 1000)
}

func benchmarkFoamSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateBestSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateBestSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateBestSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkFoamSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateWorstSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateWorstSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateWorstSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkFoamSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateRandomSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateRandomSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := generateRandomSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func foamSort(slice []int) {
	main.FoamSort(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			time.Sleep(100 * time.Nanosecond)
		}
		return slice[a] < slice[b]
	})
}

func bubbleSort(slice []int) {
	main.BubbleSort(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			time.Sleep(100 * time.Nanosecond)
		}
		return slice[a] < slice[b]
	})
}

func standardSort(slice []int) {
	sort.Slice(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			time.Sleep(100 * time.Nanosecond)
		}
		return slice[a] < slice[b]
	})
}

func generateBestSlice(c int) []int {
	s := make([]int, c, c)
	for i := 0; i < c; i++ {
		s[i] = i
	}
	return s
}

func generateWorstSlice(c int) []int {
	s := make([]int, c, c)
	for i := 0; i < c; i++ {
		s[i] = c - i
	}
	return s
}

func generateRandomSlice(c int) []int {
	s := make([]int, c, c)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < c; i++ {
		s[i] = rand.Int()
	}
	return s
}

func sorted(s []int) bool {
	p := s[0]
	for i := 1; i < len(s); i++ {
		n := s[i]
		if n < p {
			return false
		}
		p = n
	}
	return true
}
