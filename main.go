package main

import (
	"log"
	"runtime"
)

func main() {
	slice := []int{1, 5, 2, 6, 3, 7, 4, 9, 8, 0}
	log.Println(slice)
	FoamSort(slice, func(a, b int) bool {
		return slice[a] < slice[b]
	})
	log.Println(slice)
}

func FoamSort(slice []int, less func(int, int) bool) {
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
	}
	// Closes queues so worker threads can terminate.
	close(jobs)
	close(results)
}

func BubbleSort(slice []int, less func(int, int) bool) {
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
	}
}

// RedditSort is another variant of a parallel Bubble Sort intended to improve
// performance by dividing the workload into larger chunks for each worker and
// therefore make better use of the cache hierarchy.
func RedditSort(slice []int, less func(int, int) bool) {
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
	}
	// Closes queues so worker threads can terminate.
	close(jobs)
	close(results)
}
