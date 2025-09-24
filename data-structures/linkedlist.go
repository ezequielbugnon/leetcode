package datastructures

type Node struct {
	property string
	next     *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(property string) {
	node := &Node{property: property}

	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func (l *LinkedList) Prepend(property string) {
	node := &Node{property: property, next: l.head}
	l.head = node
}

func (l *LinkedList) Find(property string) bool {
	current := l.head
	for current != nil {
		if current.property == property {
			return true
		}

		current = current.next
	}

	return false
}

func (l *LinkedList) Delete(property string) {

	if l.head == nil {
		return
	}

	if l.head.property == property {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.property != property {
		current = current.next
	}

	if current.next == nil {
		return
	}

	current.next = current.next.next
}
