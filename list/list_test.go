package list

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	l.PushFront(2)
	if e := l.First(); e.Value.(int) != 2 {
		t.Errorf("PushFront into doubly linked list error")
	}
	if s := l.Len(); s != 1 {
		t.Errorf("Get doubly linked list length error")
	}
}

func TestList(t *testing.T) {
	l := NewList()
	l.PushBack(3)
	l.PushBack(4)
	if e := l.Last(); e.Value.(int) != 4 {
		t.Errorf("PushBack into list error")
	}
	if l.Len() != 2 {
		t.Errorf("List len error")
	}
}
