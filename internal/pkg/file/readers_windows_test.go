// @Author Ben.Zheng
// @DateTime 2022/8/8 13:06

package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_read_a_not_exists_config(t *testing.T) {
	asserter := assert.New(t)

	notFoundFilepath := "./test/not_found"
	fn := func() {
		NewConfigurationReader(notFoundFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_but_not_support_config(t *testing.T) {
	asserter := assert.New(t)

	notSupportFilepath := "./test/reader.properties"
	fn := func() {
		NewConfigurationReader(notSupportFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_but_not_support_config_with_backslash(t *testing.T) {
	asserter := assert.New(t)

	notSupportFilepath := ".\\test\\reader.properties"
	fn := func() {
		NewConfigurationReader(notSupportFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_dir_without_filename(t *testing.T) {
	asserter := assert.New(t)

	onlyPath := "./test"
	fn := func() {
		NewConfigurationReader(onlyPath)
	}
	asserter.Panics(fn)
}
