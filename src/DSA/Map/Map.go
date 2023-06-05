package mapp

type Map[K comparable, V comparable] struct {
	length int
}

func (m *Map[K, V]) Len() int { return m.length }

func NewMap[K comparable, V comparable]() *Map[K, V] {
}

func (m *Map[K, V]) Get(key K) V {
}

func (m *Map[K, V]) Set(key K, value V) {
}

func (m *Map[K, V]) Delete(key K) V {
}
