package pp

import (
	"fmt"
	"sync"
)

// visual representation of this Binary Tree
/*

        9
       / \
      5   15
     /    /  \
    3    10   30
   /           \
  2             50
                 \
                  51

*/

func initBSTree(out chan<- string) *tNode {
	var bt *tNode

	// The first number will be the Root.
	values := []int{9, 5, 3, 2, 15, 10, 30, 50, 51}
	for _, val := range values {
		bt = bt.insert(out, val)
	}

	return bt
}

func (n *tNode) insert(out chan<- string, val int) *tNode {
	newNode := &tNode{val: val}
	if n == nil {
		out <- fmt.Sprintf("No nodes found, creating Root on value %d", newNode.val)
		return newNode
	}

	out <- fmt.Sprintf("Passing through %d", n.val)
	if newNode.val < n.val {
		if n.left == nil {
			out <- fmt.Sprintf("<- Creating %d on the left", val)
			n.left = newNode
		} else {
			n.left = n.left.insert(out, val)
		}
	} else {
		if n.right == nil {
			out <- fmt.Sprintf("-> Creating %d on the right", val)
			n.right = newNode
		} else {
			n.right = n.right.insert(out, val)
		}
	}

	return n
}

func (n *tNode) search(out chan<- string, term int) *tNode {
	if n == nil {
		out <- fmt.Sprintf("Term %d not found", term)
		return nil
	}

	out <- fmt.Sprintf("Passing through %d", n.val)

	if n.val == term {
		out <- fmt.Sprintf("Term %d found!", term)
		return n
	}

	if term < n.val {
		return n.left.search(out, term)
	}

	return n.right.search(out, term)
}

func TraverseBinarySearchTree(wg *sync.WaitGroup, resultsChannel chan<- string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		resultsChannel <- "Initializing Binary Search Tree...\n"
		bst := initBSTree(resultsChannel)
		resultsChannel <- "Finished Initializing. \n"

		searchTerm := 51
		resultsChannel <- "Traversing Binary Search Tree...\n"
		bst.search(resultsChannel, searchTerm)
		resultsChannel <- "Finished traversing \n"
	}()
}
