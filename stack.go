// Package stack implements the Stack data structure

package stack

type Stack struct {
	Stack  []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.Stack = append(s.Stack, value)
}

func (s *Stack) Pop() interface{} {
	length := len(s.Stack)
	poppedValue := s.Stack[length-1]
	s.Stack = s.Stack[:length-1]

	return poppedValue
}

func (s Stack) Peek() interface{} {
	length := len(s.Stack)
	return s.Stack[length-1]

}

func (s Stack) IsEmpty() bool {
	return s.Size() <= 0
}

func (s Stack) Size() int {
	return len(s.Stack)
}
