package stack

type Stack[T any] struct {
	Length int
}

func NewStack[T any]() *Stack[T] {

}

func (s *Stack[T]) Push(item T) {

}

func (s *Stack[T]) Pop() (T, bool) {

}

func (s *Stack[T]) Peek() (T, bool) {

}
