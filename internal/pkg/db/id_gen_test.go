// @Author Ben.Zheng
// @DateTime 2022/8/22 13:48

package db

// References:
// https://github.com/ai/nanoid
// https://github.com/jaevor/go-nanoid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_standard_nanoid_gen(t *testing.T) {
	asserter := assert.New(t)

	t.Run("general", func(tt *testing.T) {
		idGen, err := StandardNanoId(21)
		asserter.NoError(err)
		id := idGen()
		asserter.Len(id, 21)
	})

	t.Run("len as negative", func(tt *testing.T) {
		_, err := StandardNanoId(-1)
		asserter.Error(err)
	})

	t.Run("max than right index boundary", func(tt *testing.T) {
		_, err := StandardNanoId(256)
		asserter.Error(err)
	})

	t.Run("min than left index boundary", func(tt *testing.T) {
		_, err := StandardNanoId(1)
		asserter.Error(err)
	})

	t.Run("diff", func(tt *testing.T) {
		idGen, err := StandardNanoId(21)
		asserter.NoError(err)
		id1 := idGen()
		id2 := idGen()
		asserter.NotEqual(id1, id2)
	})
}

// go test -run none -bench Benchmark_standard_nanoid_gen -benchtime 3s -benchmem
func Benchmark_standard_nanoid_gen(b *testing.B) {
	asserter := assert.New(b)

	idGen, err := StandardNanoId(21)
	asserter.NoError(err)
	for i := 0; i < b.N; i++ {
		_ = idGen()
	}
}
