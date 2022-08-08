// @Author Ben.Zheng
// @DateTime 2022/8/8 13:53

package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_read_an_exists_dir_without_filename_with_backslash(t *testing.T) {
	asserter := assert.New(t)

	onlyPath := ".\\test"
	fn := func() {
		NewConfigurationReader(onlyPath)
	}
	asserter.Panics(fn)
}
