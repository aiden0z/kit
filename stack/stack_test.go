package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	emptyStack := NewStack(0)
	if _, err := emptyStack.Pop(); err != ErrEmptyStack {
		t.Errorf("Pop from empty stack not return error %v", err)
	}
	fullStack := NewStack(2)
	if fullStack.Len() != 0 {
		t.Errorf("Initialize stack error")
	}
	fullStack.Push(1)
	fullStack.Push(2)
	value, err := fullStack.Pop()
	if err != nil {
		t.Errorf("Pop from stack return error %v", err)
	}
	if value != 2 {
		t.Errorf("Value pop from stack not equal to which pushed into")
	}
}
