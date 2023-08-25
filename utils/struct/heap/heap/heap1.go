package heap

import (
	_interface "github.com/hopeio/lemon/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type Heap1[T _interface.OrderKey[V], V constraints.Ordered] struct {
	entity []T
	less   _interface.CompareFunc[V]
}

func NewHeap1[T _interface.OrderKey[V], V constraints.Ordered](l int, less _interface.CompareFunc[V]) *Heap1[T, V] {
	heap := make([]T, 0, l)
	return &Heap1[T, V]{heap, less}
}

func NewHeap1FromArray[T _interface.OrderKey[V], V constraints.Ordered](arr []T, less _interface.CompareFunc[V]) *Heap1[T, V] {
	heap := Heap1[T, V]{arr, less}
	for i := 1; i < len(arr); i++ {
		Up(heap.entity, i, less)
	}
	return &heap
}

func (heap *Heap1[T, V]) Init() {
	HeapInit(heap.entity, heap.less)
}

func (heap *Heap1[T, V]) Push(x T) {
	heap.entity = append(heap.entity, x)
	Up(heap.entity, len(heap.entity)-1, heap.less)
}

func (heap *Heap1[T, V]) Pop() T {
	n := len(heap.entity) - 1
	item := heap.entity[0]
	heap.entity[0], heap.entity[n] = heap.entity[n], heap.entity[0]
	Down(heap.entity, 0, n, heap.less)
	heap.entity = heap.entity[:n]
	return item
}

func (heap *Heap1[T, V]) First() T {
	return heap.entity[0]
}

func (heap *Heap1[T, V]) Last() T {
	return heap.entity[len(heap.entity)-1]
}

func (heap *Heap1[T, V]) Remove(i int) T {
	n := len(heap.entity) - 1
	item := heap.entity[i]
	if n != i {
		heap.entity[i], heap.entity[n] = heap.entity[n], heap.entity[i]
		if !Down(heap.entity, i, n, heap.less) {
			Up(heap.entity, i, heap.less)
		}
	}
	heap.entity = heap.entity[:n]
	return item
}
