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

func (list *linkedList) AddNodeToEnd(data int) {
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

func (list *linkedList) AddNodeToBeginning(data int) {
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

		list.AddNodeToEnd(3)
		list.AddNodeToEnd(4)
		list.AddNodeToEnd(5)
		list.AddNodeToEnd(6)
		list.AddNodeToBeginning(2)
		list.AddNodeToBeginning(1)
		list.AddNodeToBeginning(0)
		list.AddNodeToBeginning(-1)
		list.Print(resultsChannel)
	}()
}
