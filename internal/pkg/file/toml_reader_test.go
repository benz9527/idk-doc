// @Author Ben.Zheng
// @DateTime 2022/8/8 9:06

package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tomlReader_read_from_toml(t *testing.T) {
	asserter := assert.New(t)

	path, err := filepath.Abs("./test")
	asserter.Nil(err)
	_, err = os.Stat(path + "\\reader.toml")
	asserter.Nil(err)
	reader := newTomlReader(path, "reader", "toml")

	a, err := reader.GetAny("reader")
	t.Log(a)

	s, err := reader.GetSlice("reader.myslice")
	asserter.Nil(err)
	expectedSlice := []any{
		map[string]any{
			"days": int64(100),
			"name": "ben1",
			"work": "idk-doc1",
		},
		map[string]any{
			"days": int64(200),
			"name": "ben2",
			"work": "idk-doc2",
		}}
	asserter.Equal(expectedSlice, s)

	str, err := reader.GetString("reader.mystring")
	asserter.Nil(err)
	expectedString := "This is a string"
	asserter.Equal(expectedString, str)

	_int64, err := reader.GetInt64("reader.myint64")
	asserter.Nil(err)
	expectedInt64 := int64(12345)
	asserter.Equal(expectedInt64, _int64)

	_float64, err := reader.GetFloat64("reader.myfloat64")
	asserter.Nil(err)
	expectedFloat64 := float64(0.12345)
	asserter.Equal(expectedFloat64, _float64)

	_bool, err := reader.GetBoolean("reader.mybool")
	asserter.Nil(err)
	asserter.True(_bool)

	obj, err := reader.GetMap("reader.myobj")
	asserter.Nil(err)
	expectedMap := map[string]any{
		"name":  "ben",
		"title": "gopher",
	}
	asserter.Equal(expectedMap, obj)

	asserter.NotNil(reader)
}
