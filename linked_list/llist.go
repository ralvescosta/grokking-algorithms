package llist

type Node struct {
	ID   int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(ID int) *LinkedList {
	if l.head == nil {
		l.head = &Node{ID: ID}
		return l
	}

	current := l.head.next
	for current.next != nil {
		current = current.next
	}

	current.next = &Node{ID: ID}

	return l
}

func (l *LinkedList) Remove(ID int) {
	if l.head == nil {
		return
	}

	if l.head.ID == ID {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.ID != ID {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}
