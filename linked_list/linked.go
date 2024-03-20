package llist

type LinkedList[T any] struct {
	value T
	next  *LinkedList[T]
}

func NewLinkedList[T any](v T) *LinkedList[T] {
	return &LinkedList[T]{value: v}
}

func (l *LinkedList[T]) Add(v T) *LinkedList[T] {
	if l.next == nil {
		l.next = &LinkedList[T]{value: v}
		return l
	}

	c := l.next

	for {
		if c.next == nil {
			c.next = &LinkedList[T]{value: v}
			return l
		}

		c = c.next
	}
}

func (l *LinkedList[T]) AddInTheMiddle(index int, v T) {
	var last *LinkedList[T]

	for i := 0; i < index-1; i++ {
		last = l.next
	}

	next := last.next

	last.next = &LinkedList[T]{
		value: v,
		next:  next,
	}
}

func (l *LinkedList[T]) Remove(index int) *LinkedList[T] {
	var itemBeforeTheItemThatWillBeRemove *LinkedList[T]
	for i := 0; i < index-2; i++ {
		itemBeforeTheItemThatWillBeRemove = l.next
	}

	itemThatWillBeRemoved := itemBeforeTheItemThatWillBeRemove.next
	itemBeforeTheItemThatWillBeRemove.next = itemThatWillBeRemoved.next

	return l
}
