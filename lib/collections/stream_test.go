// @Author Ben.Zheng
// @DateTime 2022/7/26 13:08

package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sliceStream_for_string_and_distinct(t *testing.T) {
	asserter := assert.New(t)
	src := []string{
		"a", "b", "c", "a", "d", "e", "c",
	}
	result := SliceStream[string](src).Distinct().ToSlice()
	expected := []string{
		"a", "b", "c", "d", "e",
	}
	asserter.Equal(len(expected), len(result))
}

func Test_DefaultStringSort(t *testing.T) {
	asserter := assert.New(t)
	t.Parallel()
	t.Run("single alpha string compare", func(tt *testing.T) {
		current := "a"
		standard := "A"
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
	t.Run("single ascii string compare", func(tt *testing.T) {
		current := "a"
		standard := "1"
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
	t.Run("single ascii string compare 2", func(tt *testing.T) {
		current := "a"
		standard := "{"
		expectedCurrentIsLower := -1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsLower, res)
	})
	t.Run("single ascii string compare 3", func(tt *testing.T) {
		current := "a"
		standard := "["
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
	t.Run("digit string compare", func(tt *testing.T) {
		current := "1.21"
		standard := "1.20"
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
	t.Run("abnormal digit string compare", func(tt *testing.T) {
		current := "1.2.1"
		standard := "1.2.#"
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
	t.Run("diff len string compare", func(tt *testing.T) {
		current := "1.2"
		standard := "1.2.#"
		expectedCurrentIsLower := -1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsLower, res)
	})
	t.Run("diff len string compare 2", func(tt *testing.T) {
		current := "2.1"
		standard := "1.2.#"
		expectedCurrentIsGreater := 1
		res := DefaultStringCompare[string](current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
}

func Test_using_generic_sort(t *testing.T) {
	asserter := assert.New(t)
	t.Parallel()
	t.Run("single alpha string compare with generic function", func(tt *testing.T) {
		var fn SortComparatorFn[string] = DefaultStringCompare[string]
		current := "a"
		standard := "A"
		expectedCurrentIsGreater := 1
		res := fn(current, standard)
		asserter.Equal(expectedCurrentIsGreater, res)
	})
}
