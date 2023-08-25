package slices

import (
	_interface "github.com/hopeio/lemon/utils/definition/constraints"
	"golang.org/x/exp/constraints"
)

// 没有泛型，范例，实际需根据不同类型各写一遍,用CmpKey，基本类型又用不了，go需要能给基本类型实现方法不能给外部类型实现方法
// 1.20以后字段均是comparable的结构体也是comparable的
func IsCoincide[S ~[]T, T comparable](s1, s2 S) bool {
	for i := range s1 {
		for j := range s2 {
			if s1[i] == s2[j] {
				return true
			}
		}
	}
	return false
}

func IsCoincideByKey[S ~[]_interface.CmpKey[T], T comparable](s1, s2 S) bool {
	for i := range s1 {
		for j := range s2 {
			if s1[i].CmpKey() == s2[j].CmpKey() {
				return true
			}
		}
	}
	return false
}

func RemoveDuplicates[S ~[]T, T comparable](s S) S {
	var m = make(map[T]struct{})
	for _, i := range s {
		m[i] = struct{}{}
	}
	s = s[:0]
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

func RemoveDuplicatesByKey[S ~[]_interface.CmpKey[T], T comparable](s S) S {
	var m = make(map[T]_interface.CmpKey[T])
	for _, i := range s {
		m[i.CmpKey()] = i
	}
	s = s[:0]
	for _, i := range m {
		s = append(s, i)
	}
	return s
}

// 取并集
func Intersection[S ~[]T, T comparable](a S, b S) S {
	if len(a) < SmallArrayLen && len(b) < SmallArrayLen {
		if len(a) > len(b) {
			return intersection(a, b)
		}
		return intersection(b, a)
	}
	panic("TODO:大数组利用map取并集")
}

func intersection[S ~[]T, T comparable](a S, b S) S {
	var ret S
	for _, x := range a {
		if In(x, b) {
			ret = append(ret, x)
		}
	}
	return ret
}

func IntersectionByKey[S ~[]_interface.CmpKey[T], T comparable](a S, b S) S {
	if len(a) < SmallArrayLen && len(b) < SmallArrayLen {
		if len(a) > len(b) {
			return intersectionByKey(a, b)
		}
		return intersectionByKey(b, a)
	}
	panic("TODO:大数组利用map取并集")
}

func intersectionByKey[S ~[]_interface.CmpKey[T], T comparable](a S, b S) S {
	var ret S
	for _, x := range a {
		if InByKey(x.CmpKey(), b) {
			ret = append(ret, x)
		}
	}
	return ret
}

// 有序数组取交集
func OrderedArrayIntersection[S ~[]T, T constraints.Ordered](a S, b S) S {
	var ret S
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	var idx int
	for _, x := range a {
		if x > b[len(b)-1] {
			return ret
		}
		for j := idx; idx < len(b)-1; j++ {
			if a[len(a)-1] < b[idx] {
				return ret
			}
			if x == b[idx] {
				ret = append(ret, x)
				idx = j
			}
		}
	}
	return ret
}

// 取并集
func Union[S ~[]T, T comparable](a S, b S) S {
	var m = make(map[T]struct{}, len(a)+len(b))
	for _, x := range a {
		m[x] = struct{}{}
	}
	for _, x := range b {
		m[x] = struct{}{}
	}
	var ret = make(S, len(m))
	for k, _ := range m {
		ret = append(ret, k)
	}
	return ret
}

func UnionByKey[S ~[]_interface.CmpKey[T], T comparable](a S, b S) S {
	var m = make(map[T]_interface.CmpKey[T], len(a)+len(b))
	for _, x := range a {
		m[x.CmpKey()] = x
	}
	for _, x := range b {
		m[x.CmpKey()] = x
	}
	var ret = make(S, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

// 取差集,采用递归
func Difference[S ~[]T, T comparable](a S, b S) S {
	if len(a) < SmallArrayLen && len(b) < SmallArrayLen {
		if len(a) > len(b) {
			return difference(a, b)
		}
		return difference(b, a)
	}
	panic("TODO:大数组利用map取差集")
}

func difference[S ~[]T, T comparable](a S, b S) S {
	var ret S
	for _, x := range a {
		if !In(x, b) {
			ret = append(ret, x)
		}
	}
	return ret
}

func DifferenceByKey[S ~[]_interface.CmpKey[T], T comparable](a S, b S) S {
	if len(a) < SmallArrayLen && len(b) < SmallArrayLen {
		if len(a) > len(b) {
			return differenceByKey(a, b)
		}
		return differenceByKey(b, a)
	}
	panic("TODO:大数组利用map取差集")
}

func differenceByKey[S ~[]_interface.CmpKey[T], T comparable](a S, b S) S {
	var ret S
	for _, x := range a {
		if !InByKey(x.CmpKey(), b) {
			ret = append(ret, x)
		}
	}
	return ret
}

// 取差集，采用循环
func DifferenceByLoop[S ~[]T, T comparable](a, b S) S {
	if len(a) > len(b) {
		a, b = b, a
	}
	var diff S
Loop:
	for _, i := range b {
		for _, j := range a {
			if i == j {
				continue Loop
			}
		}
		diff = append(diff, i)
	}
	return diff
}

// 取差集，通过循环比较key
func DifferenceByKeyByLoop[S ~[]_interface.CmpKey[T], T comparable](a, b S) S {
	if len(a) > len(b) {
		a, b = b, a
	}
	var diff S
Loop:
	for _, i := range b {
		for _, j := range a {
			if i.CmpKey() == j.CmpKey() {
				continue Loop
			}
		}
		diff = append(diff, i)
	}
	return diff
}
