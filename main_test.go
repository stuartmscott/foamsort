package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stuartmscott/foamsort"
	"sort"
	"testing"
)

func TestFoamSort(t *testing.T) {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	main.FoamSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	}, nil)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slice)
}

func TestBubbleSort(t *testing.T) {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	main.BubbleSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	}, nil)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slice)
}

func TestRedditSort(t *testing.T) {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	main.RedditSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	}, nil)
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

func BenchmarkRedditSort_Best_1K(b *testing.B) {
	benchmarkRedditSort_Best(b, 1000)
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

func BenchmarkRedditSort_Worst_1K(b *testing.B) {
	benchmarkRedditSort_Worst(b, 1000)
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

func BenchmarkRedditSort_Random_1K(b *testing.B) {
	benchmarkRedditSort_Random(b, 1000)
}

func BenchmarkStandardSort_Random_1K(b *testing.B) {
	benchmarkStandardSort_Random(b, 1000)
}

// Helper Functions

func benchmarkFoamSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateBestSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateBestSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkRedditSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateBestSlice(c)
		redditSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Best(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateBestSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkFoamSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateWorstSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateWorstSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkRedditSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateWorstSlice(c)
		redditSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Worst(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateWorstSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkFoamSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateRandomSlice(c)
		foamSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkBubbleSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateRandomSlice(c)
		bubbleSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkRedditSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateRandomSlice(c)
		redditSort(slice)
		assert.True(b, sorted(slice))
	}
}

func benchmarkStandardSort_Random(b *testing.B, c int) {
	for n := 0; n < b.N; n++ {
		slice := main.GenerateRandomSlice(c)
		standardSort(slice)
		assert.True(b, sorted(slice))
	}
}

func foamSort(slice []int) {
	main.FoamSort(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			busyWork()
		}
		return slice[a] < slice[b]
	}, nil)
}

func bubbleSort(slice []int) {
	main.BubbleSort(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			busyWork()
		}
		return slice[a] < slice[b]
	}, nil)
}

func redditSort(slice []int) {
	main.RedditSort(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			busyWork()
		}
		return slice[a] < slice[b]
	}, nil)
}

func standardSort(slice []int) {
	sort.Slice(slice, func(a, b int) bool {
		if EXPENSIVE_COMPARISON {
			busyWork()
		}
		return slice[a] < slice[b]
	})
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

func busyWork() {
	var sum int
	for i := 0; i < 1000; i++ {
		sum += i
	}
	expected := 499500
	if sum != expected {
		panic(fmt.Sprintf("Expected %d, got %d", expected, sum))
	}
}
