package pp

import (
	"fmt"
	"sync"
	"time"
)

func BinaryVsSerialSearch(wg *sync.WaitGroup, resultsChannel chan<- string) {
	start := time.Now()
	ints := CreateLargeSliceMake(100000000)
	elapsed := time.Since(start)
	fmt.Printf("Created after %v milliseconds\n", elapsed.Milliseconds())
	fmt.Println("Creating test data...")

	const searchTerm = 99999999

	fmt.Println("Searching...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		res := IntBinarySearch(searchTerm, ints)
		elapsed := time.Since(start)
		resultsChannel <- fmt.Sprintf("Result: %v. Binary search Took %v microseconds\n", res, elapsed.Microseconds())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start2 := time.Now()
		res2 := IntSerialSearch(searchTerm, ints)
		elapsed2 := time.Since(start2)
		resultsChannel <- fmt.Sprintf("Result: %v. Normal search Took %v microseconds\n", res2, elapsed2.Microseconds())
	}()
}

func IntBinarySearch(term int, data []int) bool {
	start := 0
	end := len(data) - 1

	for start <= end {
		mid := (start + end) / 2
		midVal := data[mid]

		if term == midVal {
			return true
		}

		if term < midVal {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
