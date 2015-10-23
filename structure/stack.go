package structure

import "errors"

type stack struct {
	data []interface{}
}

var ErrEmptyStack = errors.New("stack is empty")

func NewStack(size uint) *stack {
	return &stack{data: make([]interface{}, 0, size)}
}

//return the size of items in stack
func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

//pop the top item out, if stack is empty, will return ErrEmptyStack decleared above
func (s *stack) Pop() (interface{}, error) {
	if s.Len() > 0 {
		rect := s.data[s.Len()-1]
		s.data = s.data[:s.Len()-1]
		return rect, nil
	}
	return nil, ErrEmptyStack
}

//return and not pop the top item
func (s *stack) Peek() (interface{}, error) {
	if s.Len() > 0 {
		return s.data[s.Len()-1], nil
	}
	return nil, ErrEmptyStack
}
