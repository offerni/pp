package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/offerni/pp"
)

func main() {
	fmt.Println("Creating test data...")
	start := time.Now()
	ints := pp.CreateLargeSliceMake(100000000)
	elapsed := time.Since(start)
	fmt.Printf("Created after %v milliseconds\n", elapsed.Milliseconds())

	var wg sync.WaitGroup
	resultsChannel := make(chan string)

	const searchTerm = 99999999

	fmt.Println("Searching...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		res := pp.IntBinarySearch(searchTerm, ints)
		elapsed := time.Since(start)
		resultsChannel <- fmt.Sprintf("Result: %v. Binary search Took %v microseconds\n", res, elapsed.Microseconds())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start2 := time.Now()
		res2 := pp.IntSerialSearch(searchTerm, ints)
		elapsed2 := time.Since(start2)
		resultsChannel <- fmt.Sprintf("Result: %v. Normal search Took %v microseconds\n", res2, elapsed2.Microseconds())
	}()

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	for result := range resultsChannel {
		fmt.Println(result)
	}
}
