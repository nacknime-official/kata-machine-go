package ringbuffer

type RingBuffer[T any] struct {
	Length int
}

func NewRingBuffer[T any]() *RingBuffer[T] {

}

func (r *RingBuffer[T]) Push(item T) {

}

func (r *RingBuffer[T]) Get(idx int) (T, bool) {

}

func (r *RingBuffer[T]) Pop() (T, bool) {

}
