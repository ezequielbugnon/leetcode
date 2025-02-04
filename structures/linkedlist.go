package structures

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{head: nil}
}

func (l *LinkedList) Add(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (l *LinkedList) Remove(value int) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
		return true
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		return false
	}

	current.next = current.next.next
	return true
}
