package heap

import (
	_interface "github.com/hopeio/tiga/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type MaxHeap[T _interface.OrderKey[V], V constraints.Ordered] Heap[T, V]

func NewMaxHeap[T _interface.OrderKey[V], V constraints.Ordered](l int) MaxHeap[T, V] {
	maxHeap := make([]T, 0, l)
	return MaxHeap[T, V]{
		entry: maxHeap,
		less:  _interface.GreaterFunc[V],
	}
}

func NewMaxHeapFromArray[T _interface.OrderKey[V], V constraints.Ordered](arr []T) MaxHeap[T, V] {
	heap := NewHeapFromArray[T, V](arr, _interface.GreaterFunc[V])
	return MaxHeap[T, V](heap)
}

func (h *MaxHeap[T, V]) Init() {
	(*Heap[T, V])(h).Init()
}

func (h *MaxHeap[T, V]) Push(x T) {
	(*Heap[T, V])(h).Push(x)
}

func (h *MaxHeap[T, V]) Pop() T {
	return (*Heap[T, V])(h).Pop()
}

func (h *MaxHeap[T, V]) Remove(i int) T {

	return (*Heap[T, V])(h).Remove(i)
}
