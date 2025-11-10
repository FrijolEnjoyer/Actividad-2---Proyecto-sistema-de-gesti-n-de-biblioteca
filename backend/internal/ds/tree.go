package ds

// BST representa un árbol de búsqueda binaria genérico. Se alimenta con un comparador
// provisto por el consumidor para decidir cómo ordenar las claves de tipo K y así
// asociarlas a valores de tipo V.
type BST[K any, V any] struct {
	root *bstNode[K, V]
	cmp  func(a, b K) int
	size int
}

// bstNode almacena un par clave/valor junto con referencias a sus hijos.
// Es interna a la implementación del árbol y no se exporta.
type bstNode[K any, V any] struct {
	key         K
	value       V
	left, right *bstNode[K, V]
}

// NewBST crea un árbol vacío configurado con el comparador recibido. Si olvidamos pasar
// el comparador, fallamos de inmediato para evitar árboles inconsistentes.
func NewBST[K any, V any](cmp func(a, b K) int) *BST[K, V] {
	if cmp == nil {
		panic("nil comparator")
	}
	return &BST[K, V]{cmp: cmp}
}

// Size devuelve cuántos elementos viven actualmente en el árbol.
func (t *BST[K, V]) Size() int { return t.size }

// IsEmpty permite preguntar en una sola línea si el árbol está vacío.
// Esta función es útil para determinar rápidamente si el árbol contiene elementos.
func (t *BST[K, V]) IsEmpty() bool { return t.size == 0 }

func (t *BST[K, V]) Put(key K, value V) (V, bool) {
	if t.root == nil {
		t.root = &bstNode[K, V]{key: key, value: value}
		t.size++
		var zero V
		return zero, false
	}

	current := t.root
	var zero V
	for {
		comparison := t.cmp(key, current.key)
		if comparison == 0 {
			previous := current.value
			current.value = value
			return previous, true
		}
		if comparison < 0 {
			if current.left == nil {
				current.left = &bstNode[K, V]{key: key, value: value}
				t.size++
				return zero, false
			}
			current = current.left
			continue
		}

		if current.right == nil {
			current.right = &bstNode[K, V]{key: key, value: value}
			t.size++
			return zero, false
		}
		current = current.right
	}
}

func (t *BST[K, V]) Get(key K) (V, bool) {
	n := t.root
	for n != nil {
		cmp := t.cmp(key, n.key)
		if cmp == 0 {
			return n.value, true
		}
		if cmp < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	var zero V
	return zero, false
}

func (t *BST[K, V]) Contains(key K) bool {
	_, ok := t.Get(key)
	return ok
}

func (t *BST[K, V]) Delete(key K) (V, bool) {
	var removed V
	var deleted bool
	t.root, removed, deleted = deleteNode(t.root, key, t.cmp)
	if deleted {
		t.size--
	}
	return removed, deleted
}

func deleteNode[K any, V any](node *bstNode[K, V], key K, cmp func(a, b K) int) (*bstNode[K, V], V, bool) {
	if node == nil {
		var zero V
		return nil, zero, false
	}

	comparison := cmp(key, node.key)
	switch {
	case comparison < 0:
		newLeft, value, deleted := deleteNode(node.left, key, cmp)
		node.left = newLeft
		return node, value, deleted
	case comparison > 0:
		newRight, value, deleted := deleteNode(node.right, key, cmp)
		node.right = newRight
		return node, value, deleted
	default:
		removed := node.value
		if node.left == nil {
			return node.right, removed, true
		}
		if node.right == nil {
			return node.left, removed, true
		}
		successor := minNode(node.right)
		replacement := &bstNode[K, V]{
			key:   successor.key,
			value: successor.value,
			left:  node.left,
			right: deleteMin(node.right),
		}
		return replacement, removed, true
	}
}

func minNode[K any, V any](node *bstNode[K, V]) *bstNode[K, V] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func deleteMin[K any, V any](node *bstNode[K, V]) *bstNode[K, V] {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}

func (t *BST[K, V]) TraverseInOrder(fn func(key K, value V)) {
	if fn == nil {
		return
	}
	traverseInOrder(t.root, fn)
}

func traverseInOrder[K any, V any](node *bstNode[K, V], fn func(key K, value V)) {
	if node == nil {
		return
	}
	traverseInOrder(node.left, fn)
	fn(node.key, node.value)
	traverseInOrder(node.right, fn)
}
