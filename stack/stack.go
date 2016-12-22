package stack

type stack struct {
	data []interface{}
}

// NewStack return a new stack
// The stack storage capacity will auto increase because the underground
// storage is slice.
func NewStack(size uint) *stack {
	return &stack{data: make([]interface{}, 0, size)}
}

//Len return the size of items in stack
func (s *stack) Len() int {
	return len(s.data)
}

// IsEmpty return if the stack is empty
func (s *stack) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	} else {
		return false
	}
}

// Push item to stack
func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

// Pop the top item out, if stack is empty, will return nil
func (s *stack) Pop() interface{} {
	if s.Len() > 0 {
		rect := s.data[s.Len() - 1]
		s.data = s.data[:s.Len() - 1]
		return rect
	}
	return nil
}

// Peek return and not pop the top item
func (s *stack) Peek() interface{} {
	if s.Len() > 0 {
		return s.data[s.Len() - 1]
	}
	return nil
}
