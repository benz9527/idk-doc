// @Author Ben.Zheng
// @DateTime 8/20/22 1:09 PM

package db

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_linux_path_with_db_ext_validation(t *testing.T) {
	asserter := assert.New(t)

	case1 := "/tmp/idk_test.db"
	case2 := "/tmp/idk_testdb"
	_, err := regexp.MatchString(`.*/.*\.db$`, case1)
	asserter.Nil(err)
	_, err = regexp.MatchString(`.*/.*\.db$`, case2)
	asserter.Nil(err)
}
