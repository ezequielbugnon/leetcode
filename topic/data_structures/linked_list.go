package datastructures

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddNode(value int) bool {
	newNode := &Node{
		value: value,
	}

	if l.head == nil {
		l.head = newNode
		return true
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode

	return true
}

func (l *LinkedList) Print() {
	current := l.head
	for current.next != nil {
		fmt.Printf("linkedlist: %v\n", current.value)
		current = current.next
	}
	fmt.Printf("linkedlist: %v\n", current.value)
}

func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil {
		return
	}

	var prev *Node
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev

}

func (l *LinkedList) RemoveElement(value int) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
		return true
	}

	current := l.head

	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
		}
		current = current.next
	}

	return false

}
