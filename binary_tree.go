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

func (n *tNode) preOrderSearch(out chan<- string, term int) *tNode {
	if n != nil {
		out <- fmt.Sprintf("Passing through %d", n.val)
		if n.val == term {
			out <- fmt.Sprintf("term %d found", n.val)
			return n
		}

		l := n.left.preOrderSearch(out, term)
		if l != nil {
			return l
		}

		r := n.right.preOrderSearch(out, term)
		if r != nil {
			return r
		}
	}

	return nil
}

func (n *tNode) inOrderSearch(out chan<- string, term int) *tNode {
	if n != nil {
		l := n.left.inOrderSearch(out, term)
		if l != nil {
			return l
		}

		out <- fmt.Sprintf("Passing through %d", n.val)
		if n.val == term {
			out <- fmt.Sprintf("term %d found", n.val)
			return n
		}

		r := n.right.inOrderSearch(out, term)
		if r != nil {
			return r
		}
	}

	return nil
}

func (n *tNode) postOrderSearch(out chan<- string, term int) *tNode {
	if n != nil {
		l := n.left.postOrderSearch(out, term)
		if l != nil {
			return l
		}

		r := n.right.postOrderSearch(out, term)
		if r != nil {
			return r
		}

		out <- fmt.Sprintf("Passing through %d", n.val)
		if n.val == term {
			out <- fmt.Sprintf("term %d found", n.val)
			return n
		}
	}

	return nil
}

func TraverseBinaryTree(wg *sync.WaitGroup, out chan<- string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		bTree := initBTree()

		finished := "Finished traversing \n"
		searchTerm := 6

		out <- "Preorder Searching..."
		bTree.preOrderSearch(out, searchTerm)
		out <- finished

		out <- "InOrder Searching..."
		bTree.inOrderSearch(out, searchTerm)
		out <- finished

		out <- "PostOrder Searching..."
		bTree.postOrderSearch(out, searchTerm)
		out <- finished
	}()
}
