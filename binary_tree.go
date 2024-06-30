package pp

import (
	"fmt"
	"sync"
)

type tNode struct {
	left  *tNode
	right *tNode
	val   int
}

// visual representation of this Binary Tree
/*

        10
       /  \
     15    30
    /  \     \
   3    6     2
  /           / \
 5           9   8

Preorder
Root - Left - Right

Inorder
Left - Root - Right

Postorder
Left - Right - Root

*/

func initBTree() *tNode {
	five := tNode{val: 5}
	three := tNode{val: 3, left: &five}
	six := tNode{val: 6}
	fifteen := tNode{val: 15, left: &three, right: &six}
	nine := tNode{val: 9}
	eight := tNode{val: 8}
	two := tNode{val: 2, left: &nine, right: &eight}
	thirty := tNode{val: 30, right: &two}
	return &tNode{val: 10, right: &thirty, left: &fifteen}
}

func (n *tNode) preOrderSearch(resultsChannel chan<- string, term int) *tNode {
	if n != nil {
		resultsChannel <- fmt.Sprintf("passing through %d", n.val)
		if n.val == term {
			resultsChannel <- fmt.Sprintf("term %d found", n.val)
			return n
		}

		l := n.left.preOrderSearch(resultsChannel, term)
		if l != nil {
			return l
		}

		r := n.right.preOrderSearch(resultsChannel, term)
		if r != nil {
			return r
		}
	}

	return nil
}

func (n *tNode) inOrderSearch(resultsChannel chan<- string, term int) *tNode {
	if n != nil {
		l := n.left.inOrderSearch(resultsChannel, term)
		if l != nil {
			return l
		}

		resultsChannel <- fmt.Sprintf("passing through %d", n.val)
		if n.val == term {
			resultsChannel <- fmt.Sprintf("term %d found", n.val)
			return n
		}

		r := n.right.inOrderSearch(resultsChannel, term)
		if r != nil {
			return r
		}
	}

	return nil
}

func (n *tNode) postOrderSearch(resultsChannel chan<- string, term int) *tNode {
	if n != nil {
		l := n.left.postOrderSearch(resultsChannel, term)
		if l != nil {
			return l
		}

		r := n.right.postOrderSearch(resultsChannel, term)
		if r != nil {
			return r
		}

		resultsChannel <- fmt.Sprintf("passing through %d", n.val)
		if n.val == term {
			resultsChannel <- fmt.Sprintf("term %d found", n.val)
			return n
		}
	}

	return nil
}

func TraverseBinaryTrees(wg *sync.WaitGroup, resultsChannel chan<- string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		bTree := initBTree()

		finished := "Finished traversing \n"
		searchTerm := 6

		resultsChannel <- "Preorder Searching..."
		bTree.preOrderSearch(resultsChannel, searchTerm)
		resultsChannel <- finished

		resultsChannel <- "InOrder Searching..."
		bTree.inOrderSearch(resultsChannel, searchTerm)
		resultsChannel <- finished

		resultsChannel <- "PostOrder Searching..."
		bTree.postOrderSearch(resultsChannel, searchTerm)
		resultsChannel <- finished
	}()
}
