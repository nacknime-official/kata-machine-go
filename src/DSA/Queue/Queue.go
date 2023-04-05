package queue

type Queue[T any] struct {
	Length int
}

func NewQueue[T any]() *Queue[T] {

}

func (q *Queue[T]) Enqueue(item T) {

}

func (q *Queue[T]) Deque() (T, bool) {

}

func (q *Queue[T]) Peek() (T, bool) {

}
