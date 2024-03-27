package stack

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}

	lastItemIndex := len(s.items) - 1

	item := s.items[lastItemIndex]
	s.items = s.items[0:lastItemIndex]

	return &item
}
