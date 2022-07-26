// @Author Ben.Zheng
// @DateTime 2022/7/19 9:56 AM

package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get_allow_headers_as_one_line_without_end_sep(t *testing.T) {
	asserter := assert.New(t)
	_, actual := getAllowHeaders()
	expectedSize := 9
	asserter.Equal(expectedSize, actual)
}

func Test_get_allow_headers_as_one_line_without_end_sep_with_additional(t *testing.T) {
	asserter := assert.New(t)
	additional := `Access-Control-Allow-Headers, Authorization`
	_, actual := getAllowHeaders(additional)
	expectedSize := 9
	asserter.Equal(expectedSize, actual)
}
