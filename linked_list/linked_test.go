package llist

import "testing"

func TestAdd(t *testing.T) {
	list := NewLinkedList(10)

	list.Add(5).Add(8).Add(7)

	if list.value != 10 ||
		list.next.value != 5 ||
		list.next.next.value != 8 ||
		list.next.next.next.value != 7 ||
		list.next.next.next.next != nil {
		t.Error("add not work properly")
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList(10)

	list.Add(5).Add(8).Add(7).Remove(3)

	if list.value != 10 ||
		list.next.value != 5 ||
		list.next.next.value != 7 ||
		list.next.next.next != nil {
		t.Error("add not work properly")
	}
}
