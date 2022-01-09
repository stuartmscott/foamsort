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
			if less(a, b) {
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
	// Adds indeces of non-overlapping pairs to job queue.
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
			if !less(i, j) {
				slice[i], slice[j] = slice[j], slice[i]
				swapped = true
			}
			i = j
		}
	}
}
