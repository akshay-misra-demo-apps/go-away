package collections

import (
	"fmt"
)

// Node for Singly or Doubly Linked List
type node[T any] struct {
	Value T
	Next  *node[T]
	Prev  *node[T]
}

// doublyLinkedList struct
type doublyLinkedList[T any] struct {
	Head *node[T]
	Tail *node[T]
}

// NewDoublyLinkedList public func to create doubly linked list struct
func NewDoublyLinkedList[T any]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{}
}

// Add a new entry to the doubly linked list
func (list *doublyLinkedList[T]) Add(value T) {
	newNode := &node[T]{Value: value}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	newNode.Prev = list.Tail
	list.Tail.Next = newNode
	list.Tail = newNode
}

// RemoveByIndex remove a node by index
func (list *doublyLinkedList[T]) RemoveByIndex(index int) (bool, error) {
	current := list.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return false, fmt.Errorf("index %v out of range", index)
		}
		current = current.Next
	}

	if current == nil {
		return false, fmt.Errorf("index %v out of range", index)
	}

	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		list.Head = current.Next
	}

	if current.Next != nil {
		current.Next.Prev = current.Prev
	} else {
		list.Tail = current.Prev
	}

	return true, nil
}

func (list *doublyLinkedList[T]) String() {
	current := list.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
