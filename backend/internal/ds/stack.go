package ds

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] { return &Stack[T]{data: make([]T, 0)} }

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 { return zero, false }
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 { return zero, false }
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Size() int { return len(s.data) }
