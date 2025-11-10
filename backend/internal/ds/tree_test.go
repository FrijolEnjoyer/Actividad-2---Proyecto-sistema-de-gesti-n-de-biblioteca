package ds

import (
	"fmt"
	"testing"
)

func TestBSTPutGetAndDelete(t *testing.T) {
	bst := NewBST[string, int](stringsCompare)

	if _, replaced := bst.Put("b", 2); replaced {
		t.Fatalf("unexpected replace on first insert")
	}
	if _, replaced := bst.Put("a", 1); replaced {
		t.Fatalf("unexpected replace on insert a")
	}
	if prev, replaced := bst.Put("b", 3); !replaced || prev != 2 {
		t.Fatalf("expected replace of b, got replaced=%v prev=%d", replaced, prev)
	}

	if bst.Size() != 2 {
		t.Fatalf("expected size 2, got %d", bst.Size())
	}

	if !bst.Contains("a") {
		t.Fatalf("expected Contains(a) = true")
	}
	if _, ok := bst.Get("missing"); ok {
		t.Fatalf("missing key should not be found")
	}

	order := make([]string, 0)
	bst.TraverseInOrder(func(k string, _ int) { order = append(order, k) })
	expectedOrder := []string{"a", "b"}
	if !equalStrings(order, expectedOrder) {
		t.Fatalf("unexpected order: %v", order)
	}

	if removed, ok := bst.Delete("a"); !ok || removed != 1 {
		t.Fatalf("delete a failed, removed=%d, ok=%v", removed, ok)
	}
	if removed, ok := bst.Delete("missing"); ok || removed != 0 {
		t.Fatalf("delete missing should fail, removed=%d ok=%v", removed, ok)
	}
	if removed, ok := bst.Delete("b"); !ok || removed != 3 {
		t.Fatalf("delete b failed, removed=%d, ok=%v", removed, ok)
	}

	if !bst.IsEmpty() {
		t.Fatalf("tree should be empty after deletions")
	}
}

func TestBSTDeleteNodeWithTwoChildren(t *testing.T) {
	bst := NewBST[int, string](func(a, b int) int { return a - b })
	for i, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		bst.Put(v, fmt.Sprintf("node-%d", i))
	}

	removed, ok := bst.Delete(50)
	if !ok || removed != "node-0" {
		t.Fatalf("expected to remove root, ok=%v removed=%s", ok, removed)
	}

	order := make([]int, 0)
	bst.TraverseInOrder(func(k int, _ string) { order = append(order, k) })
	expected := []int{20, 30, 40, 60, 70, 80}
	if !equalInts(order, expected) {
		t.Fatalf("unexpected order after delete: %v", order)
	}
}

func stringsCompare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
