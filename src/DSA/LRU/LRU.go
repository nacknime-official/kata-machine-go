package lru

type LRU[K comparable, V comparable] struct {
	length int
}

func (l *LRU[K, V]) Len() int { return l.length }

func NewLRU[K comparable, V comparable](capacity int) *LRU[K, V] {
}

func (l *LRU[K, V]) Update(key K, value V) {
}

func (l *LRU[K, V]) Get(key K) (V, bool) {
}
