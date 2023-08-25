package heap

import (
	_interface "github.com/hopeio/lemon/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

type MaxBaseHeap[T constraints.Ordered] []T

func NewMaxBaseHeap[T constraints.Ordered](l int) MaxBaseHeap[T] {
	maxHeap := make(MaxBaseHeap[T], 0, l)
	return maxHeap
}

func NewMaxBaseHeapFromArray[T constraints.Ordered](arr []T) MaxBaseHeap[T] {
	heap := newBaseHeapFromArray[T](arr, _interface.GreaterFunc[T])
	return MaxBaseHeap[T](heap)
}

func (h MaxBaseHeap[T]) Init() {
	BaseHeap[T](h).init(_interface.GreaterFunc[T])
}

func (h *MaxBaseHeap[T]) Push(x T) {
	(*BaseHeap[T])(h).push(x, _interface.GreaterFunc[T])
}

func (h *MaxBaseHeap[T]) Pop() T {
	return (*BaseHeap[T])(h).pop(_interface.GreaterFunc[T])
}

func (h *MaxBaseHeap[T]) Remove(i int) T {
	return (*BaseHeap[T])(h).remove(i, _interface.GreaterFunc[T])
}
