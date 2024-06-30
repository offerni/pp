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
     / \  / \
    3   6 10 30
   /     \    \
  2       8    50

*/

func initBSTree() *tNode {
	two := &tNode{val: 2}
	three := &tNode{val: 3, left: two}
	eight := &tNode{val: 8}
	six := &tNode{val: 6, right: eight}
	five := &tNode{val: 5, left: three, right: six}
	ten := &tNode{val: 10}
	fifty := &tNode{val: 50}
	thirty := &tNode{val: 30, right: fifty}
	fifteen := &tNode{val: 15, left: ten, right: thirty}
	return &tNode{val: 9, left: five, right: fifteen}
}

func (n *tNode) search(out chan<- string, term int) *tNode {
	if n == nil {
		out <- fmt.Sprintf("Term %d not found", term)
		return nil
	}

	out <- fmt.Sprintf("Passing through %d", n.val)

	if n.val == term {
		out <- fmt.Sprintf("Term %d found", term)
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

		finished := "Finished traversing \n"
		searchTerm := 50

		resultsChannel <- "Traversing Binary Search Tree..."
		bst := initBSTree()
		bst.search(resultsChannel, searchTerm)

		resultsChannel <- finished
	}()
}
