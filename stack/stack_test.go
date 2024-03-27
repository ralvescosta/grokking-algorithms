package stack

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s.items) != 3 {
		t.Error("stack length is incorrect")
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item := s.Pop()
	if item == nil || *item != 3 || len(s.items) != 2 {
		t.Error("stack pop is incorrect")
	}

	item = s.Pop()
	if item == nil || *item != 2 || len(s.items) != 1 {
		t.Error("stack pop is incorrect")
	}

	item = s.Pop()
	if item == nil || *item != 1 || len(s.items) != 0 {
		t.Error("stack pop is incorrect")
	}

	item = s.Pop()
	if item != nil {
		t.Error("stack pop is incorrect")
	}
}
