package heap

import (
	_interface "github.com/hopeio/tiga/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type Heap1[T _interface.OrderKey[V], V constraints.Ordered] struct {
	entry []T
	less  _interface.CompareFunc[V]
}

func NewHeap1[T _interface.OrderKey[V], V constraints.Ordered](l int, less _interface.CompareFunc[V]) *Heap1[T, V] {
	heap := make([]T, 0, l)
	return &Heap1[T, V]{heap, less}
}

func NewHeap1FromArray[T _interface.OrderKey[V], V constraints.Ordered](arr []T, less _interface.CompareFunc[V]) *Heap1[T, V] {
	heap := Heap1[T, V]{arr, less}
	for i := 1; i < len(arr); i++ {
		Up(heap.entry, i, less)
	}
	return &heap
}

func (heap *Heap1[T, V]) Init() {
	HeapInit(heap.entry, heap.less)
}

func (heap *Heap1[T, V]) Push(x T) {
	heap.entry = append(heap.entry, x)
	Up(heap.entry, len(heap.entry)-1, heap.less)
}

func (heap *Heap1[T, V]) Pop() T {
	n := len(heap.entry) - 1
	item := heap.entry[0]
	heap.entry[0], heap.entry[n] = heap.entry[n], heap.entry[0]
	Down(heap.entry, 0, n, heap.less)
	heap.entry = heap.entry[:n]
	return item
}

func (heap *Heap1[T, V]) First() T {
	return heap.entry[0]
}

func (heap *Heap1[T, V]) Last() T {
	return heap.entry[len(heap.entry)-1]
}

func (heap *Heap1[T, V]) Remove(i int) T {
	n := len(heap.entry) - 1
	item := heap.entry[i]
	if n != i {
		heap.entry[i], heap.entry[n] = heap.entry[n], heap.entry[i]
		if !Down(heap.entry, i, n, heap.less) {
			Up(heap.entry, i, heap.less)
		}
	}
	heap.entry = heap.entry[:n]
	return item
}
