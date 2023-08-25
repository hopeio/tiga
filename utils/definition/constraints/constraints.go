package constraints

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Callback[T any] interface {
	~func() | ~func() error | ~func(T) | ~func(T) error
}

type ID interface {
	constraints.Integer | ~string | ~[]byte | ~[8]byte | ~[16]byte
}

type Basic struct {
}

type Range interface {
	constraints.Ordered | time.Time | ~*time.Time | ~string
}

func SignedConvert[T, V constraints.Signed](v V) T {
	return T(v)
}

func FloatConvert[T, V constraints.Float](v V) T {
	return T(v)
}

func UnsignedConvert[T, V constraints.Unsigned](v V) T {
	return T(v)
}

func IntegerConvert[T, V constraints.Integer](v V) T {
	return T(v)
}

func NumberConvert[T, V Number](v V) T {
	return T(v)
}
