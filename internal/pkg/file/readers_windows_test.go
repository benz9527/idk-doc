// @Author Ben.Zheng
// @DateTime 2022/8/8 13:06

package file

import (
	"path/filepath"
	"regexp"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_filepath_dir(t *testing.T) {
	asserter := assert.New(t)
	t.Parallel()
	t.Run("win slash", func(tt *testing.T) {
		tt.Parallel()
		src := ".\\conf\\idk.yaml"
		expectedDir, expectedFilename, expectedExt := "conf", "idk", ".yaml"
		dir, ext := filepath.Dir(src), filepath.Ext(src)
		filename := src[3+len(dir):]
		filename = filename[:len(filename)-len(ext)]
		asserter.Equal(expectedDir, dir)
		asserter.Equal(expectedExt, ext)
		asserter.Equal(expectedFilename, filename)
	})

	t.Run("win regexp", func(tt *testing.T) {
		tt.Parallel()
		src := "C:\\abc\\123.txt"
		matched, err := regexp.MatchString(`^[A-Za-z]:\\`, src)
		asserter.Nil(err)
		asserter.True(matched)
	})

	t.Run("Linux slash", func(tt *testing.T) {
		tt.Parallel()
		src := "./conf/idk.yaml"
		expectedDir, expectedFilename, expectedExt := "conf", "idk", ".yaml"
		dir, ext := filepath.Dir(src), filepath.Ext(src)
		filename := src[3+len(dir):]
		filename = filename[:len(filename)-len(ext)]
		asserter.Equal(expectedDir, dir)
		asserter.Equal(expectedExt, ext)
		asserter.Equal(expectedFilename, filename)
	})

}

func Test_read_a_not_exists_config(t *testing.T) {
	asserter := assert.New(t)
	v := viper.New()
	notFoundFilepath := "./test/not_found"
	fn := func() {
		NewConfigurationReader(v, notFoundFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_but_not_support_config(t *testing.T) {
	asserter := assert.New(t)
	v := viper.New()
	notSupportFilepath := "./test/reader.properties"
	fn := func() {
		NewConfigurationReader(v, notSupportFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_but_not_support_config_with_backslash(t *testing.T) {
	asserter := assert.New(t)
	v := viper.New()
	notSupportFilepath := ".\\test\\reader.properties"
	fn := func() {
		NewConfigurationReader(v, notSupportFilepath)
	}
	asserter.Panics(fn)
}

func Test_read_an_exists_dir_without_filename(t *testing.T) {
	asserter := assert.New(t)
	v := viper.New()
	onlyPath := "./test"
	fn := func() {
		NewConfigurationReader(v, onlyPath)
	}
	asserter.Panics(fn)
}
