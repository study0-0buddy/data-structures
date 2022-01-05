package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	// should return a valid singly linked list
	l := NewList()

	assert.NotNil(t, l)
	// at init the list size must be 0 and head is nil
	assert.Equal(t, 0, l.size)
	assert.Nil(t, l.head)
}

func TestSize(t *testing.T) {
	l := NewList()

	// the list is empty, so size must be 0
	assert.Equal(t, 0, l.Size())

	// use all add/insert operations
	l.AppendBack(1)
	l.InsertAfter(l.head, 2)
	l.AppendFront(3)

	// the list size must be 3 as three items were added
	assert.Equal(t, 3, l.Size())
}

func TestAppendBack(t *testing.T) {
	l := NewList()

	l.AppendBack(1)
	// the list's head shouldn't be nil
	assert.NotNil(t, l.head)
	// the head's value should be 1
	assert.Equal(t, 1, l.head.value)
	// the head's next link should be nil
	assert.Nil(t, l.head.next)

	l.AppendBack(2)
	// the head's value should be 1
	assert.Equal(t, 1, l.head.value)
	// the head's next node's value should be 2
	assert.Equal(t, 2, l.head.next.value)
}

func TestAppendFront(t *testing.T) {
	l := NewList()

	l.AppendFront(1)
	// the list's head shouldn't be nil
	assert.NotNil(t, l.head)
	// the head's value should be 1
	assert.Equal(t, 1, l.head.value)
	// the head's next link should be nil
	assert.Nil(t, l.head.next)

	l.AppendFront(2)
	// the head's value should be 2
	assert.Equal(t, 2, l.head.value)
	// the head's next node's value should be 1
	assert.Equal(t, 1, l.head.next.value)
}

func TestInsertAfter(t *testing.T) {
	l := NewList()

	// should return error as the list is empty
	assert.Error(t, l.InsertAfter(l.head, 1))

	// append new items
	l.AppendBack(1)
	l.AppendBack(3)

	// insert a new node after the head
	l.InsertAfter(l.head, 2)
	// the head should be 1
	assert.Equal(t, 1, l.head.value)
	// the head's next node's value should be 2
	assert.Equal(t, 2, l.head.next.value)
	// the newly added node's next value should be 3
	assert.Equal(t, 3, l.head.next.next.value)
}

func TestDelete(t *testing.T) {
	l := NewList()

	// should be nil, as the list is empty
	assert.Nil(t, l.Delete(1))

	// 1 -> 2 -> 3
	l.AppendBack(1)
	l.AppendBack(2)
	l.AppendBack(3)

	// should return the node with value 1 and set new head
	n := l.Delete(1)
	assert.NotNil(t, n)
	// removed node's value should be 1
	assert.Equal(t, 1, n.GetValue())
	// new head should be the node with value 2
	assert.Equal(t, 2, l.head.GetValue())
	// size should be decreased
	assert.Equal(t, 2, l.Size())

	// remove last node
	n = l.Delete(3)
	assert.NotNil(t, n)
	// removed node's value should be 1
	assert.Equal(t, 3, n.GetValue())
	// size should be decreased
	assert.Equal(t, 1, l.Size())

	// 2 -> 3 -> 2
	l.AppendBack(3)
	l.AppendBack(2)

	// should remove first matching value
	n = l.Delete(2)
	assert.NotNil(t, n)
	// removed node's value should be 2
	assert.Equal(t, 2, n.GetValue())
	// new head should be the node with value 3
	assert.Equal(t, 3, l.head.GetValue())
	// head's next node's value should be 2
	assert.Equal(t, 2, l.head.next.value)
	// size should be decreased
	assert.Equal(t, 2, l.Size())
}

func TestSearch(t *testing.T) {
	l := NewList()

	// should be nil, as the list is empty
	assert.Nil(t, l.Search(1))

	l.AppendBack(1)
	l.AppendBack(2)
	// should be nil, as the list doesn't have a node with value 3
	assert.Nil(t, l.Search(3))

	l.AppendBack(3)
	// should return valid node, as 3 is found
	res := l.Search(3)
	assert.Equal(t, 3, res.GetValue())
	assert.NotNil(t, l.Search(3))
}
