// Package stack implements the Stack data structure

package stack

type Stack struct {
	Length int
	Stack  []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.Stack = append(s.Stack, value)
	s.Length += 1
}

func (s *Stack) Pop() interface{} {
	poppedValue := s.Stack[s.Length-1]
	s.Stack = s.Stack[:s.Length-1]

	s.Length -= 1
	return poppedValue
}

func (s Stack) Peek() interface{} {
	return s.Stack[s.Length-1]

}

func (s Stack) IsEmpty() bool {
	return s.Length <= 0
}

func (s Stack) Size() int {
	return s.Length
}
