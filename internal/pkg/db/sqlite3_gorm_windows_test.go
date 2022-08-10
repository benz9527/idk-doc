// @Author Ben.Zheng
// @DateTime 2022/8/10 13:27

package db

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"regexp"
	"testing"
)

func Test_multiple_upper_dir_validated_by_regx(t *testing.T) {
	asserter := assert.New(t)

	dbPath := ".\\db\\idk.db"
	match, err := regexp.MatchString(`^(\.\.\\?)+`, dbPath)
	asserter.Nil(err)
	asserter.False(match)
}

func Test_if_db_path_is_abs(t *testing.T) {
	asserter := assert.New(t)

	winDBPath1 := ".\\db\\idk.db"
	winDBPath2 := "..\\db\\idk.db"
	linuxDBPath1 := "./db/idk.db"
	linuxDBPath2 := "../db/idk.db"

	asserter.False(filepath.IsAbs(winDBPath1))
	asserter.False(filepath.IsAbs(winDBPath2))
	asserter.False(filepath.IsAbs(linuxDBPath1))
	asserter.False(filepath.IsAbs(linuxDBPath2))
}
