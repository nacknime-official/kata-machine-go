package doublylinkedlist

type DoublyLinkedList[T comparable] struct {
	length int
}

func (l *DoublyLinkedList[T]) Len() int { return l.length }

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {

}

func (l *DoublyLinkedList[T]) Prepend(item T) {

}

func (l *DoublyLinkedList[T]) InsertAt(item T, idx int) {

}

func (l *DoublyLinkedList[T]) Append(item T) {

}

func (l *DoublyLinkedList[T]) Remove(item T) (T, bool) {

}

func (l *DoublyLinkedList[T]) Get(idx int) (T, bool) {

}

func (l *DoublyLinkedList[T]) RemoveAt(idx int) (T, bool) {

}
