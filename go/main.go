package main

import (
	"data-structures/linked-list/singly"
	"fmt"
	"log"
)

func main() {
	// create a new singly linked list
	sl := singly.NewList()

	// check the list size
	fmt.Printf("List size: %d\n", sl.Size())

	// append a new values
	sl.AppendBack(1)
	sl.AppendBack(2)
	sl.AppendBack(3)
	sl.AppendBack(5)
	sl.Display()
	// search for a node with value 4
	fmt.Printf("%d is found: %v\n", 4, sl.Search(4) != nil)
	// insert a new node with value 4 after the node with value 3
	if err := sl.InsertAfter(sl.Search(3), 4); err != nil {
		log.Fatal(err)
	}
	sl.AppendBack(6)
	sl.AppendBack(7)

	// display all the list values
	sl.Display()

	// search for a node with value 4
	fmt.Printf("%d is found: %v\n", 4, sl.Search(4) != nil)

	// check the list size
	fmt.Printf("List size: %d\n", sl.Size())
	sl.AppendFront(9)
	sl.AppendFront(10)
	sl.AppendFront(11)
	sl.Display()
}
