package ds

type Queue[T any] struct {
	data []T
	head int
}

func NewQueue[T any]() *Queue[T] { return &Queue[T]{data: make([]T, 0)} }

func (q *Queue[T]) Enqueue(v T) { q.data = append(q.data, v) }

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.head >= len(q.data) { return zero, false }
	v := q.data[q.head]
	q.head++
	// compact occasionally
	if q.head > 32 && q.head*2 >= len(q.data) {
		q.data = append([]T(nil), q.data[q.head:]...)
		q.head = 0
	}
	return v, true
}

func (q *Queue[T]) Size() int { return len(q.data) - q.head }
