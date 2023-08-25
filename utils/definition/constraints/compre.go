package constraints

import "golang.org/x/exp/constraints"

type CompareFunc[T any] func(T, T) bool

type Compare[T any] interface {
	Compare(T) bool
}

func LessFunc[T constraints.Ordered](t T, t2 T) bool {
	return t < t2
}

func GreaterFunc[T constraints.Ordered](t T, t2 T) bool {
	return t > t2
}

func EqualFunc[T constraints.Ordered](t T, t2 T) bool {
	return t > t2
}

// comparable 只能比较是否相等,不能比较大小
type OrderKey[T constraints.Ordered] interface {
	OrderKey() T
}

type CmpKey[T comparable] interface {
	CmpKey() T
}
