package ds

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1); s.Push(2)
	if s.Size() != 2 { t.Fatalf("exp 2") }
	v, ok := s.Pop(); if !ok || v != 2 { t.Fatalf("exp 2") }
}

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1); q.Enqueue(2)
	if q.Size() != 2 { t.Fatalf("exp 2") }
	v, ok := q.Dequeue(); if !ok || v != 1 { t.Fatalf("exp 1") }
}

func TestListFind(t *testing.T) {
	l := NewList[int]()
	l.InsertFront(1); l.InsertFront(5)
	v, ok := l.Find(func(x int) bool { return x == 1 })
	if !ok || v != 1 { t.Fatalf("not found") }
}

func TestArray(t *testing.T) {
	a := NewArray[int](3)
	if !a.Set(1, 9) { t.Fatalf("set") }
	v, ok := a.Get(1); if !ok || v != 9 { t.Fatalf("get") }
}
