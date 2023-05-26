package llist

import "testing"

func TestAdd(t *testing.T) {
	list := New().Add(1).Add(2).Add(3)

	node := list.head
	for i := 1; i <= 3; i++ {
		if node == nil {
			break
		}

		if node.ID != i {
			t.Errorf("expected: %v, actual: %v", i, node.ID)
		}

		node = node.next
	}
}

func TestRemove(t *testing.T) {
	list := New().Add(1).Add(4).Add(2).Add(7)
	list.Remove(7)

	unexpectedID := 7

	node := list.head
	for {
		if node == nil {
			break
		}

		if node.ID == unexpectedID {
			t.Errorf("ID: %v, should have been removed", unexpectedID)
		}

		node = node.next
	}
}
