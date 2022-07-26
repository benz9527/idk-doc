// @Author Ben.Zheng
// @DateTime 2022/7/26 11:28

package collections

import "math"

type SortComparatorFn[T comparable] func(a, b T) int

type sliceStream[T comparable] struct {
	source []T
}

func (s *sliceStream[T]) Distinct() *sliceStream[T] {
	var res = make(map[T]string, len(s.source))
	for _, v := range s.source {
		if _, ok := res[v]; !ok {
			res[v] = ""
		}
	}
	s.source = make([]T, len(res))
	i := 0
	for k := range res {
		s.source[i] = k
		i++
	}
	return s
}

func (s *sliceStream[T]) Sort(fn SortComparatorFn[T]) *sliceStream[T] {
	// TODO(Ben) Implement the sort details.
	return s
}

func (s *sliceStream[T]) ToSlice() []T {
	return s.source
}

func SliceStream[T comparable](src []T) *sliceStream[T] {
	return &sliceStream[T]{
		source: src,
	}
}

// DefaultStringCompare
// Internal compare as the ASCII character index.
// Generally, the digit is lt letter, and upper case letter is lt lower case one.
// digit: 48-57
// upper case letter: 65-90
// lower case letter: 97-122
// @return -1: current lt standard; 0: current eq standard; 1: current gt standard
func DefaultStringCompare[T string](current, standard T) int {
	minLen := int(math.Min(float64(len(current)), float64(len(standard))))
	for i := 0; i < minLen; i++ {
		if current[i] == standard[i] {
			continue
		} else if current[i] > standard[i] {
			return 1
		} else if current[i] < standard[i] {
			return -1
		}
	}
	isCurrentLenEq := len(current) == len(standard)
	if isCurrentLenEq {
		return 0
	}

	isCurrentShorter := len(current) < len(standard)
	if isCurrentShorter {
		return -1
	}
	return 1
}
