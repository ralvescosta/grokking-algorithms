package llist

type (
	List[T any] struct {
		value T
		next  *List[T]
	}

	LinkedList[T any] struct {
		list *List[T]
		last *List[T]
	}
)

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(v T) *LinkedList[T] {
	if l.list == nil {
		l.list = &List[T]{value: v}
		l.last = l.list
		return l
	}

	l.last.next = &List[T]{value: v}
	l.last = l.last.next

	return l
}

func (l *LinkedList[T]) AddInTheMiddle(index int, v T) *LinkedList[T] {
	last := l.list

	for i := 0; i < index-1; i++ {
		last = l.list.next
	}

	next := last.next

	last.next = &List[T]{
		value: v,
		next:  next,
	}

	return l
}

func (l *LinkedList[T]) Remove(index int) *LinkedList[T] {
	var itemBeforeTheItemThatWillBeRemove *List[T]
	for i := 0; i < index-2; i++ {
		itemBeforeTheItemThatWillBeRemove = l.list.next
	}

	itemThatWillBeRemoved := itemBeforeTheItemThatWillBeRemove.next
	itemBeforeTheItemThatWillBeRemove.next = itemThatWillBeRemoved.next

	return l
}

func (l *LinkedList[T]) Revert() *LinkedList[T] {
	var current, next, prev *List[T]
	current = l.list

	// Runtime Example
	//l.list = &List { 1, next: &List{ 2, next: &List{ 3, next: nil } } }
	for current != nil {
		//1
		// current = &List { 1, next: &List{ 2, next: &List{ 3, next: nil } } }
		// next = nil
		// prev = nil
		//
		//2
		// current = &List{ 2, next: List{ 3, next: nil } }
		// next = &List{ 2, next: &List{ 3, next: nil } }
		// prev =  &List { 1, next: nil }
		//
		//3
		// current = &List{ 3, next: nil }
		// next = &List{ 3, next: nil }
		// prev = &List{ 2, next: List { 1, next: nil } }

		next = current.next
		//1
		// current = &List { 1, next: &List{ 2, next: &List{ 3, next: nil } } }
		// next = &List{ 2, next: &List{ 3, next: nil } }
		// prev = nil
		//
		//2
		// current = &List{ 2, next: List{ 3, next: nil } }
		// next = &List{ 3, next: nil }
		// prev =  &List { 1, next: nil }
		//
		//3
		// current = &List{ 3, next: nil }
		// next = nil
		// prev = &List{ 2, next: List { 1, next: nil } }

		current.next = prev
		//1
		// current = &List { 1, next: nil }
		// next = &List{ 2, next: &List{ 3, next: nil } }
		// prev = nil
		//
		//2
		// current = &List{ 2, next: &List { 1, next: nil } }
		// next = &List{ 3, next: nil }
		// prev =  &List { 1, next: nil }
		//
		//3
		// current = &List{ 3, next: &List{ 2, next: &List { 1, next: nil } } }
		// next = nil
		// prev = &List{ 2, next: List { 1, next: nil } }

		prev = current
		//1
		// current = &List { 1, next: nil }
		// next = &List{ 2, next: &List{ 3, next: nil } }
		// prev = &List { 1, next: nil }
		//
		//2
		// current = &List{ 2, next: &List { 1, next: nil } }
		// next = &List{ 3, next: nil }
		// prev = &List{ 2, next: List { 1, next: nil } }
		//
		//3
		// current = &List{ 3, next: &List{ 2, next: &List { 1, next: nil } } }
		// next = nil
		// prev =  &List{ 3, next: &List{ 2, next: &List { 1, next: nil } } }

		current = next
		//1
		// current = &List{ 2, next: List{ 3, next: nil } }
		// next = &List{ 2, next: &List{ 3, next: nil } }
		// prev =  &List { 1, next: nil }
		//
		//2
		// current = &List{ 3, next: nil }
		// next = &List{ 3, next: nil }
		// prev = &List{ 2, next: List { 1, next: nil } }
		//
		//3
		// current = nil
		// next = nil
		// prev =  &List{ 3, next: &List{ 2, next: &List { 1, next: nil } } }
	}

	l.list = prev
	l.last = l.list

	return l
}
