package arraylist

type ArrayList[T comparable] struct {
	length int
}

func (l *ArrayList[T]) Len() int { return l.length }

func NewArrayList[T comparable](capacity int) *ArrayList[T] {

}

func (l *ArrayList[T]) Prepend(item T) {

}

func (l *ArrayList[T]) InsertAt(item T, idx int) {

}

func (l *ArrayList[T]) Append(item T) {

}

func (l *ArrayList[T]) Remove(item T) (T, bool) {

}

func (l *ArrayList[T]) Get(idx int) (T, bool) {

}

func (l *ArrayList[T]) RemoveAt(idx int) (T, bool) {

}
