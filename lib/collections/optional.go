// @Author Ben.Zheng
// @DateTime 2022/7/18 5:29 PM

package collections

type optionalSlice[T any] struct {
	isEmpty bool
	result  []T
}

func (o optionalSlice[T]) OrElse(e []T) []T {
	if o.isEmpty {
		return e
	}
	return o.result
}

func OfSliceNullable[T any](s []T) optionalSlice[T] {
	o := optionalSlice[T]{}
	if s == nil || len(s) == 0 {
		o.isEmpty = true
	} else {
		o.result = s
	}
	return o
}

type optionalMap[K comparable, V any] struct {
	isEmpty bool
	result  map[K]V
}

func (o optionalMap[K, V]) OrElse(e map[K]V) map[K]V {
	if o.isEmpty {
		return e
	}

	return o.result
}

func OfMapNullable[K comparable, V any](m map[K]V) optionalMap[K, V] {
	o := optionalMap[K, V]{}
	if m == nil || len(m) == 0 {
		o.isEmpty = true
	} else {
		o.result = m
	}
	return o
}
