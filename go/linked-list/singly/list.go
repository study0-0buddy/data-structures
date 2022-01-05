package singly

import (
	"errors"
	"fmt"
)

// node of the list
type node struct {
	value int // use int for simplicity
	next  *node
}

func (n *node) GetValue() int {
	return n.value
}

// List defines singly linked list
type list struct {
	head *node
	size int
}

func NewList() *list {
	return &list{nil, 0}
}

// Size returns size of the list
func (l *list) Size() int {
	return l.size
}

// Display prints out all values of the list
func (l *list) Display() {
	// get copy of the head
	curr := l.head

	if curr == nil {
		fmt.Println("the list is empty")
	} else {
		// loop over the nodes and print values
		fmt.Print("List values: ")
		for curr != nil {
			fmt.Printf("%d ", curr.value)
			curr = curr.next
		}
		fmt.Println()
	}
}

// AppendBack adds a new node to end of the list
func (l *list) AppendBack(value int) {
	// create a new node
	n := &node{value, nil}

	// set the new node as head of the list
	// if there are no other nodes
	if l.head == nil {
		l.head = n
	} else {
		// loop over the nodes while
		// the next node of curr is not nil
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
	// increment size of the list
	l.size++
}

// AppendFront adds a new node with value
// to the front of the list
// the new node becomes a head of the list
func (l *list) AppendFront(value int) {
	// create a new node
	n := &node{value, nil}

	// new node's next link
	// is current head of the list
	n.next = l.head
	// new node becomes head of the list
	l.head = n

	// increment size of the list
	l.size++
}

// InsertAfter adds a new node to the list after specific node
// returns error if invalid node is provided
func (l *list) InsertAfter(n *node, value int) error {
	// validate the node
	if n == nil {
		return errors.New("'nil' is invalid value for a node")
	}

	// create a new node from the passed value
	// next node of the n must be next link of the newNode
	newNode := &node{value, n.next}

	// link the provided node to the new one
	n.next = newNode

	// increment size of the list
	l.size++

	return nil
}

// Delete removes the first element with passed value
// returns the deleted node
func (l *list) Delete(val int) *node {
	// return nil if list is empty
	if l.Size() == 0 {
		return nil
	}

	// if first element value equals given one
	// remove it and set a new front
	curr := l.head
	if l.head.value == val {
		l.head = curr.next
		// decrement size of the list
		l.size--
		return curr
	}

	// else look for a passed value in the list
	// loop over the list while the next element is no nil
	for curr != nil {
		// if next element's value equals given one
		// remove it and set current elements next
		if curr.next != nil && curr.next.value == val {
			removedElem := curr.next
			curr.next = curr.next.next
			l.size--
			return removedElem
		}
		curr = curr.next
	}

	// if not found return nil
	return nil
}

// Search searches for a node with passed value
// returns first node with matching value or nil if not found
func (l *list) Search(val int) *node {
	// temp variable to track current node
	curr := l.head

	// loop over the list
	for curr != nil {
		// if a node's value match val, return it
		if curr.value == val {
			return curr
		}
		curr = curr.next
	}

	// if no node with matching value found, return nil
	return nil
}
