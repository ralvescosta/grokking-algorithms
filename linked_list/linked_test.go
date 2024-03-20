package llist

import "testing"

func TestAdd(t *testing.T) {
	linkedList := NewLinkedList[int]()

	linkedList.Add(10).Add(5).Add(8).Add(7)

	if linkedList.list.value != 10 ||
		linkedList.list.next.value != 5 ||
		linkedList.list.next.next.value != 8 ||
		linkedList.list.next.next.next.value != 7 ||
		linkedList.list.next.next.next.next != nil {
		t.Error("add not work properly")
	}
}

func TestRemove(t *testing.T) {
	linkedList := NewLinkedList[int]()

	linkedList.Add(10).Add(5).Add(8).Add(7).Remove(3)

	if linkedList.list.value != 10 ||
		linkedList.list.next.value != 5 ||
		linkedList.list.next.next.value != 7 ||
		linkedList.list.next.next.next != nil {
		t.Error("add not work properly")
	}
}

func TestAddInTheMiddle(t *testing.T) {
	linkedList := NewLinkedList[int]()

	linkedList.Add(10).Add(5).Add(8).Add(7).AddInTheMiddle(2, 2)

	if linkedList.list.value != 10 ||
		linkedList.list.next.value != 5 ||
		linkedList.list.next.next.value != 2 ||
		linkedList.list.next.next.next.value != 8 ||
		linkedList.list.next.next.next.next.value != 7 ||
		linkedList.list.next.next.next.next.next != nil {
		t.Error("addInTheMiddle not work properly")
	}
}

// func BenchmarkAdd(b *testing.B) {
// 	list := NewLinkedList(10)

// 	for i := 0; i < b.N; i++ {
// 		list.Add(i)
// 	}
// }
