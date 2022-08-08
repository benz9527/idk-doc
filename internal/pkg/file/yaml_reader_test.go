// @Author Ben.Zheng
// @DateTime 2022/8/5 16:07

package file

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_viper_read_from_io_as_yaml(t *testing.T) {
	asserter := assert.New(t)

	v := viper.New()
	v.SetConfigType("yaml")

	var yamlDemo = []byte(`
reader:
  myslice:
    - {
        name: "ben1",
        work: "idk-doc1",
        days: 100
    }
    - {
        name: "ben2",
        work: "idk-doc2",
        days: 200
    }
  mystring: "This is a string"
  myint64: 12345
  myfloat64: 0.12345
  mybool: true
  myobj: {
    name: "ben",
    title: "gopher"
  }`)
	err := v.ReadConfig(bytes.NewBuffer(yamlDemo))
	asserter.Nil(err)
	result := v.Get("reader.myslice")
	expectedSlice := []any{
		map[string]any{
			"days": 100,
			"name": "ben1",
			"work": "idk-doc1",
		},
		map[string]any{
			"days": 200,
			"name": "ben2",
			"work": "idk-doc2",
		}}
	asserter.Equal(expectedSlice, result)
}

func Test_yamlReader_read_from_yaml(t *testing.T) {
	asserter := assert.New(t)
	v := viper.New()
	path, err := filepath.Abs("./test")
	asserter.Nil(err)
	_, err = os.Stat(path + "\\reader.yaml")
	asserter.Nil(err)
	reader := newYamlReader(v, path, "reader", "yaml")

	s, err := reader.GetSlice("reader.myslice")
	asserter.Nil(err)
	expectedSlice := []any{
		map[string]any{
			"days": 100,
			"name": "ben1",
			"work": "idk-doc1",
		},
		map[string]any{
			"days": 200,
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
