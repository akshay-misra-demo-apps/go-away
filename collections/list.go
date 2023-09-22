package structs

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (list *DoublyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}

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
func (list *DoublyLinkedList[T]) RemoveByIndex(index int) (bool, error) {
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

func (list *DoublyLinkedList[T]) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
