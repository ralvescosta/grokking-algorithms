package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		q := NewQueue[int]()

		first, second, third := 1, 2, 3
		q.Push(first)
		q.Push(second)
		q.Push(third)

		if v := q.Pop(); *v != first {
			t.Errorf("Expected 1, got %d", *v)
		}

		if v := q.Pop(); *v != second {
			t.Errorf("Expected 2, got %d", *v)
		}

		if v := q.Pop(); *v != third {
			t.Errorf("Expected 3, got %d", *v)
		}
	})

	t.Run("Empty Queue", func(t *testing.T) {
		q := NewQueue[int]()
		if v := q.Pop(); v != nil {
			t.Errorf("Expected nil, got %d", v)
		}
	})
}
