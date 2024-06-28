package pp

import (
	"fmt"
	"sync"
)

type node struct {
	data     int
	next     *node
	previous *node
}

type linkedList struct {
	head *node
}

func (list *linkedList) search(term int) *node {
	if list.head == nil {
		return nil
	}

	currentNode := list.head

	for currentNode != nil {
		if term == currentNode.data {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return nil
}

func (list *linkedList) searchWithDebug(term int, resultsChannel chan<- string) {
	node := list.search(term)

	if node != nil {
		resultsChannel <- fmt.Sprintf("Node found! The current node address is %p, the data is %d, the previous node address is %p and the next node address is %p", node, node.data, node.previous, node.next)
		return
	}

	resultsChannel <- fmt.Sprintf("Node %d does not exist in the current list", term)
}

func (list *linkedList) addNodeToEnd(data int) {
	newNode := &node{data, nil, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
		currentNode.previous = currentNode
	}
}

func (list *linkedList) removeFromEnd() {
	if list.head == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	currentNode := list.head
	var previousNode *node

	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = nil
}

func (list *linkedList) addNodeToBeginning(data int) {
	newNode := &node{data, nil, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		currentNode.previous = newNode
		newNode.next = currentNode
		list.head = newNode
	}
}

func (list *linkedList) removeFromBeginning() {
	if list.head == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	nextNode := list.head.next
	nextNode.previous = nil
	list.head = nextNode
}

func (list *linkedList) Print(resultsChannel chan<- string) {
	cn := list.head

	for cn != nil {
		resultsChannel <- fmt.Sprintf("The current node address is %p, the data is %d, the previous node address is %p and the next node address is %p", cn, cn.data, cn.previous, cn.next)
		cn = cn.next
	}
}

func ManipulateLinkedLists(wg *sync.WaitGroup, resultsChannel chan<- string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		var list linkedList

		list.addNodeToEnd(3)
		list.addNodeToEnd(4)
		list.addNodeToEnd(5)
		list.addNodeToEnd(6)
		list.addNodeToBeginning(2)
		list.addNodeToBeginning(1)
		list.addNodeToBeginning(0)
		list.addNodeToBeginning(-1)
		list.removeFromEnd()
		list.removeFromBeginning()
		list.searchWithDebug(5, resultsChannel)

		list.Print(resultsChannel)
	}()
}
