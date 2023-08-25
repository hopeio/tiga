package heap

import (
	_interface "github.com/hopeio/lemon/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type Heap[T _interface.OrderKey[V], V constraints.Ordered] []T

func NewHeap[T _interface.OrderKey[V], V constraints.Ordered](l int) Heap[T, V] {
	heap := make([]T, 0, l)
	return heap
}

func NewHeapFromArray[T _interface.OrderKey[V], V constraints.Ordered](arr []T, less _interface.CompareFunc[V]) Heap[T, V] {
	heap := Heap[T, V](arr)
	for i := 1; i < len(arr); i++ {
		heap.up(i, less)
	}
	return arr
}

func (h Heap[T, V]) init(less _interface.CompareFunc[V]) {
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n, less)
	}
}

func (h *Heap[T, V]) push(x T, less _interface.CompareFunc[V]) {
	hh := *h
	*h = append(hh, x)
	h.up(len(hh), less)
}

func (h *Heap[T, V]) pop(less _interface.CompareFunc[V]) T {
	hh := *h
	n := len(hh) - 1
	item := hh[0]
	hh[0], hh[n] = hh[n], hh[0]
	h.down(0, n, less)
	*h = hh[:n]
	return item
}

func (h *Heap[T, V]) remove(i int, less _interface.CompareFunc[V]) T {
	hh := *h
	n := len(hh) - 1
	item := hh[i]
	if n != i {
		hh[i], hh[n] = hh[n], hh[i]
		if !h.down(i, n, less) {
			h.up(i, less)
		}
	}
	*h = hh[:n]
	return item
}

func (h Heap[T, V]) down(i0, n int, less _interface.CompareFunc[V]) bool {
	return Down(h, i0, n, less)
}

func (h Heap[T, V]) up(j int, less _interface.CompareFunc[V]) {
	Up(h, j, less)
}

func (h Heap[T, V]) fix(i int, less _interface.CompareFunc[V]) {
	Fix(h, i, less)
}
