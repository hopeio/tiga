package heap

import (
	_interface "github.com/hopeio/tiga/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type BaseHeap[T constraints.Ordered] []T

func NewBaseHeap[T constraints.Ordered](l int) BaseHeap[T] {
	heap := make([]T, 0, l)
	return heap
}

func NewBaseHeapFromArray[T constraints.Ordered](arr []T, less _interface.CompareFunc[T]) BaseHeap[T] {
	heap := BaseHeap[T](arr)
	for i := 1; i < len(arr); i++ {
		heap.up(i, less)
	}
	return arr
}

func (h BaseHeap[T]) init(less _interface.CompareFunc[T]) {
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n, less)
	}
}

func (h *BaseHeap[T]) push(x T, less _interface.CompareFunc[T]) {
	hh := *h
	*h = append(hh, x)
	h.up(len(hh), less)
}

func (h *BaseHeap[T]) pop(less _interface.CompareFunc[T]) T {
	hh := *h
	n := len(hh) - 1
	item := hh[0]
	hh[0], hh[n] = hh[n], hh[0]
	h.down(0, n, less)
	*h = hh[:n]
	return item
}

func (h *BaseHeap[T]) remove(i int, less _interface.CompareFunc[T]) T {
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

func (h BaseHeap[T]) down(i0, n int, less _interface.CompareFunc[T]) bool {
	return DownForBase(h, i0, n, less)
}

func (h BaseHeap[T]) up(j int, less _interface.CompareFunc[T]) {
	UpForBase(h, j, less)
}

func (h BaseHeap[T]) fix(i int, less _interface.CompareFunc[T]) {
	FixForBase(h, i, less)
}
