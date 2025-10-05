package ds

// Fixed-capacity array wrapper demonstrating basic set/get operations.

type Array[T any] struct {
	data []T
}

func NewArray[T any](cap int) *Array[T] {
	return &Array[T]{data: make([]T, cap)}
}

func (a *Array[T]) Set(i int, v T) bool {
	if i < 0 || i >= len(a.data) { return false }
	a.data[i] = v
	return true
}

func (a *Array[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= len(a.data) { return zero, false }
	return a.data[i], true
}

func (a *Array[T]) Len() int { return len(a.data) }
