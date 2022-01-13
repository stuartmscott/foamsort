package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var slice []int
	less := func(a, b int) bool {
		return slice[a] < slice[b]
	}

	// Best
	slice = GenerateBestSlice(1000)
	if err := CreateGif("foam_best", slice, FoamSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateBestSlice(1000)
	if err := CreateGif("bubble_best", slice, BubbleSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateBestSlice(1000)
	if err := CreateGif("reddit_best", slice, RedditSort, less); err != nil {
		log.Fatal(err)
	}

	// Worst
	slice = GenerateWorstSlice(1000)
	if err := CreateGif("foam_worst", slice, FoamSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateWorstSlice(1000)
	if err := CreateGif("bubble_worst", slice, BubbleSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateWorstSlice(1000)
	if err := CreateGif("reddit_worst", slice, RedditSort, less); err != nil {
		log.Fatal(err)
	}

	// Random
	slice = GenerateRandomSlice(1000)
	if err := CreateGif("foam_random", slice, FoamSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateRandomSlice(1000)
	if err := CreateGif("bubble_random", slice, BubbleSort, less); err != nil {
		log.Fatal(err)
	}
	slice = GenerateRandomSlice(1000)
	if err := CreateGif("reddit_random", slice, RedditSort, less); err != nil {
		log.Fatal(err)
	}
}

func FoamSort(slice []int, less func(int, int) bool, callback func()) {
	limit := len(slice)
	jobs := make(chan int, limit/2)
	results := make(chan bool, limit/2)
	// Repeatedly takes an index from job queue.
	// Compares, and optionally swaps, with adjacent index.
	work := func() {
		for j := range jobs {
			a := j
			b := j + 1
			if slice[a] == slice[b] || less(a, b) {
				results <- false
			} else {
				slice[a], slice[b] = slice[b], slice[a]
				results <- true
			}
		}
	}
	// Spawns a worker thread for each available process.
	for w := 0; w < runtime.GOMAXPROCS(0); w++ {
		go work()
	}
	// Adds indices of non-overlapping pairs to job queue.
	// Waits for all results.
	sort := func(start int) (result bool) {
		for i := start; i < limit-1; i += 2 {
			jobs <- i
		}
		for i := start; i < limit-1; i += 2 {
			result = <-results || result
		}
		return
	}
	// Sorts even, then odd, pairs.
	// Repeats until no swaps occur.
	for swapped := true; swapped; {
		swapped = sort(0)
		swapped = sort(1) || swapped
		if c := callback; c != nil {
			c()
		}
	}
	// Closes queues so worker threads can terminate.
	close(jobs)
	close(results)
}

func BubbleSort(slice []int, less func(int, int) bool, callback func()) {
	limit := len(slice)
	// Repeats until no swaps occur.
	for swapped := true; swapped; {
		swapped = false
		// Iterate through all pairs.
		for i := 0; i < limit-1; {
			j := i + 1
			// Compare index with adjacent.
			if slice[i] != slice[j] && !less(i, j) {
				slice[i], slice[j] = slice[j], slice[i]
				swapped = true
			}
			i = j
		}
		if c := callback; c != nil {
			c()
		}
	}
}

// RedditSort is another variant of a parallel Bubble Sort intended to improve
// performance by dividing the workload into larger chunks for each worker and
// therefore make better use of the cache hierarchy.
func RedditSort(slice []int, less func(int, int) bool, callback func()) {
	limit := len(slice)
	workers := runtime.GOMAXPROCS(0)
	batch := limit / workers
	cache := 16 // Number of integers in a 64 Byte cache line
	// Align batch size to a multiple of cache line size
	for ; batch%cache != 0; batch++ {
	}

	jobs := make(chan int, workers)
	results := make(chan bool, workers)
	// Repeatedly takes an index from job queue.
	// For each element in the batch starting at the index.
	// Compares, and optionally swaps, with adjacent element.
	work := func() {
		// Add 1 to batch size so that it overlaps into next batch
		batch = batch + 1
		for j := range jobs {
			swapped := false
			for i := 0; i < batch && j < limit-1; i++ {
				a := j
				b := j + 1
				if slice[a] != slice[b] && !less(a, b) {
					slice[a], slice[b] = slice[b], slice[a]
					swapped = true
				}
				j = b
			}
			results <- swapped
		}
	}
	// Spawns a worker thread for each available process.
	for w := 0; w < workers; w++ {
		go work()
	}
	// Repeatedly tell each worker to process their batch until no swaps occur.
	for swapped := true; swapped; {
		swapped = false
		for w := 0; w < workers; w++ {
			jobs <- w * batch
		}
		for w := 0; w < workers; w++ {
			swapped = <-results || swapped
		}
		if c := callback; c != nil {
			c()
		}
	}
	// Closes queues so worker threads can terminate.
	close(jobs)
	close(results)
}

func GenerateBestSlice(c int) []int {
	s := make([]int, c, c)
	for i := 0; i < c; i++ {
		s[i] = i
	}
	return s
}

func GenerateWorstSlice(c int) []int {
	s := make([]int, c, c)
	for i := 0; i < c; i++ {
		s[i] = c - i
	}
	return s
}

func GenerateRandomSlice(c int) []int {
	s := make([]int, c, c)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < c; i++ {
		s[i] = rand.Intn(c)
	}
	return s
}
