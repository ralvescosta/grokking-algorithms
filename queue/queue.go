package queue

type (
	Queue[T any] struct {
		values []T
	}
)

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{values: make([]T, 0)}
}

func (q *Queue[T]) Push(v T) *Queue[T] {
	q.values = append(q.values, v)
	return q
}

func (q *Queue[T]) Pop() *T {
	if len(q.values) == 0 {
		return nil
	}

	first := q.values[0]
	q.values = q.values[1:len(q.values)]
	return &first
}
