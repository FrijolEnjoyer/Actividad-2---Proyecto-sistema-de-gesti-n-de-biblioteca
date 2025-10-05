package ds

// Singly linked list for generic type T

type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

type List[T any] struct {
	head *ListNode[T]
	size int
}

func NewList[T any]() *List[T] { return &List[T]{} }

func (l *List[T]) InsertFront(v T) {
	n := &ListNode[T]{Value: v, Next: l.head}
	l.head = n
	l.size++
}

func (l *List[T]) ForEach(fn func(v T)) {
	for n := l.head; n != nil; n = n.Next {
		fn(n.Value)
	}
}

func (l *List[T]) Find(pred func(v T) bool) (T, bool) {
	for n := l.head; n != nil; n = n.Next {
		if pred(n.Value) {
			return n.Value, true
		}
	}
	var zero T
	return zero, false
}

func (l *List[T]) Size() int { return l.size }
