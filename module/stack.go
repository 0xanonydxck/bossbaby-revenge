package module

// --------------------------------------------

type Node[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Value T
}

// --------------------------------------------

type Stack[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

// --------------------------------------------

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// --------------------------------------------

func (s *Stack[T]) Push(v T) {
	node := &Node[T]{Value: v}
	if s.Head == nil {
		s.Head = node
		s.Tail = node
		return
	}

	s.Tail.Tail = node
	s.Tail = node
}

// --------------------------------------------

func (s *Stack[T]) Pop() (value T) {
	if s.Head == nil {
		return
	}

	value = s.Tail.Value
	if s.Head == s.Tail {
		s.Head = nil
		s.Tail = nil
		return
	}

	node := s.Head
	for node.Tail != s.Tail {
		node = node.Tail
	}

	node.Tail = nil
	s.Tail = node
	return
}

// --------------------------------------------

func (s *Stack[T]) Peek() (value T) {
	if s.Tail != nil {
		value = s.Tail.Value
		return
	}
	return
}
