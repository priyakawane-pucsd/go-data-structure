package stack

import (
	"errors"
)

type Stack struct {
	elements []any
	top      int
}

func New(size int) *Stack {
	return &Stack{
		elements: make([]any, size),
		top:      -1,
	}
}

func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Stack) Push(element any) error {
	if s.top == len(s.elements)-1 {
		return errors.New("Stack Overflow....")
	}
	s.top++
	s.elements[s.top] = element
	return nil
}

func (s *Stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		element := s.elements[s.top]
		s.top--
		return element, true
	}
}
