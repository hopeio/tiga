package number

import (
	"github.com/hopeio/lemon/utils/definition/constraints"
	"strconv"
	"unsafe"
)

func FormatFloat(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func ToBytes[T constraints.Number](t T) []byte {
	size := unsafe.Sizeof(t)
	return unsafe.Slice((*byte)(unsafe.Pointer(&t)), size)
}
