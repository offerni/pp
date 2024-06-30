package main

import (
	"fmt"
	"sync"

	"github.com/offerni/pp"
)

func main() {
	var wg sync.WaitGroup
	resultsChannel := make(chan string)

	// pp.BinaryVsSerialSearch(&wg, resultsChannel)
	// pp.ManipulateLinkedLists(&wg, resultsChannel)
	// pp.TraverseBinaryTree(&wg, resultsChannel)
	pp.TraverseBinarySearchTree(&wg, resultsChannel)

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	for result := range resultsChannel {
		fmt.Println(result)
	}
}
