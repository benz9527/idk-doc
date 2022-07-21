// @Author Ben.Zheng
// @DateTime 2022/7/18 6:54 PM

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_collection_of_slice_nullable_and_or_else_for_nil(t *testing.T) {
	asserter := assert.New(t)
	expectedSlice := []struct {
		Name string
	}{
		{Name: "abc"},
	}
	res := OfSliceNullable[struct{ Name string }](nil).
		OrElse(expectedSlice)
	asserter.Equal(expectedSlice, res)
}

func Test_collection_of_slice_nullable_and_or_else_for_empty(t *testing.T) {
	asserter := assert.New(t)
	expectedSlice := []struct {
		Name string
	}{
		{Name: "abc"},
	}
	res := OfSliceNullable[struct{ Name string }]([]struct{ Name string }{}).
		OrElse(expectedSlice)
	asserter.Equal(expectedSlice, res)
}

func Test_collection_of_slice_nullable_and_or_else_for_contains_elements(t *testing.T) {
	asserter := assert.New(t)
	expectedSlice := []struct {
		Name string
	}{
		{Name: "abc"},
	}
	res := OfSliceNullable[struct{ Name string }](expectedSlice).
		OrElse([]struct{ Name string }{})
	asserter.Equal(expectedSlice, res)
}

func Test_collection_of_map_nullable_and_or_else_for_nil(t *testing.T) {
	asserter := assert.New(t)
	expectedMap := map[string]struct {
		Name string
	}{
		"foo": {Name: "abc"},
	}
	res := OfMapNullable[string, struct{ Name string }](nil).
		OrElse(expectedMap)
	asserter.Equal(expectedMap, res)
}

func Test_collection_of_map_nullable_and_or_else_for_empty(t *testing.T) {
	asserter := assert.New(t)
	expectedMap := map[string]struct {
		Name string
	}{
		"foo": {Name: "abc"},
	}

	empty := make(map[string]struct{ Name string }, 0)
	res := OfMapNullable[string, struct{ Name string }](empty).
		OrElse(expectedMap)
	asserter.Equal(expectedMap, res)
}

func Test_collection_of_map_nullable_and_or_else_for_empty_with_len_init(t *testing.T) {
	asserter := assert.New(t)
	expectedMap := map[string]struct {
		Name string
	}{
		"foo": {Name: "abc"},
	}

	empty := make(map[string]struct{ Name string }, 1)
	expectedLen := 0
	asserter.Equal(expectedLen, len(empty))

	res := OfMapNullable[string, struct{ Name string }](empty).
		OrElse(expectedMap)
	asserter.Equal(expectedMap, res)
}

func Test_collection_of_map_nullable_and_or_else_for_contains_elements(t *testing.T) {
	asserter := assert.New(t)
	expectedMap := map[string]struct {
		Name string
	}{
		"foo": {Name: "abc"},
	}

	empty := make(map[string]struct{ Name string }, 1)
	res := OfMapNullable[string, struct{ Name string }](expectedMap).
		OrElse(empty)
	asserter.Equal(expectedMap, res)
}
