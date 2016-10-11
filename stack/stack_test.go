package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	emptyStack := NewStack(0)
	if item := emptyStack.Pop(); item != nil {
		t.Error("Pop from empty stack not return nil")
	}
	fullStack := NewStack(2)
	if fullStack.Len() != 0 {
		t.Error("Initialize stack error")
	}
	fullStack.Push(1)
	fullStack.Push(2)
	value := fullStack.Pop()
	if value != 2 {
		t.Error("Value pop from stack not equal to which pushed into")
	}
}
