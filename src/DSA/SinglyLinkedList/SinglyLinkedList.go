package singlylinkedlist

type SinglyLinkedList[T comparable] struct {
	length int
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {

}

func (l *SinglyLinkedList[T]) Prepend(item T) {

}

func (l *SinglyLinkedList[T]) InsertAt(item T, idx int) {

}

func (l *SinglyLinkedList[T]) Append(item T) {

}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {

}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {

}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {

}
