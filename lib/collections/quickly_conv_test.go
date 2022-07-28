// @Author Ben.Zheng
// @DateTime 2022/7/18 4:14 PM

package collections

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func Test_no_mem_allocate_string_to_bytes(t *testing.T) {
	asserter := assert.New(t)

	var source = `this is a source string`
	var expected = []byte{
		't', 'h', 'i', 's', ' ',
		'i', 's', ' ',
		'a', ' ',
		's', 'o', 'u', 'r', 'c', 'e', ' ',
		's', 't', 'r', 'i', 'n', 'g',
	}
	var res = String2Bytes(source)
	asserter.Equal(expected, res)
}

func Test_unknown_if_mem_allocated_string_to_bytes(t *testing.T) {
	asserter := assert.New(t)

	var fnAddr string
	converter := func(src string) []byte {
		fnAddr = fmt.Sprintf("%p", &src)
		return *(*[]byte)(unsafe.Pointer(&struct {
			Data string
			Cap  int
			Len  int
		}{
			Data: src,
			Cap:  len(src),
			Len:  len(src),
		}))
	}

	var source = `this is a source string`
	var expected = []byte{
		't', 'h', 'i', 's', ' ',
		'i', 's', ' ',
		'a', ' ',
		's', 'o', 'u', 'r', 'c', 'e', ' ',
		's', 't', 'r', 'i', 'n', 'g',
	}
	var expectedAddr = fmt.Sprintf("%p", &source)
	var res = converter(source)
	asserter.NotEqual(expectedAddr, fnAddr)
	asserter.Equal(expected, res)
}

func Test_no_mem_allocate_bytes_to_string(t *testing.T) {
	asserter := assert.New(t)
	var expected = `this is a source string`
	var src = []byte{
		't', 'h', 'i', 's', ' ',
		'i', 's', ' ',
		'a', ' ',
		's', 'o', 'u', 'r', 'c', 'e', ' ',
		's', 't', 'r', 'i', 'n', 'g',
	}
	var res = Bytes2String(src)
	asserter.Equal(expected, res)
}
