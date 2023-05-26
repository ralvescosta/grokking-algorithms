package stack

import "testing"

func TestPush(t *testing.T) {
	stack := New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	expectedValues := []int32{10, 20, 30}

	for i, v := range stack.vec {
		if v != expectedValues[i] {
			t.Errorf("expected: %v, actual: %v", expectedValues[i], v)
		}
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if v := stack.Pop(); v != 10 {
		t.Errorf("expected: %v, actual: %v", 10, v)
	}

	if v := stack.Pop(); v != 20 {
		t.Errorf("expected: %v, actual: %v", 20, v)
	}

	if v := stack.Pop(); v != 30 {
		t.Errorf("expected: %v, actual: %v", 30, v)
	}
}
