package stack

type Stack[T any] struct {
	items []T
}

func New[T any](items ...T) Stack[T] {
	return Stack[T]{items: items[:]}
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var t T
		return t, false
	}
	t := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return t, true
}
