// @Author Ben.Zheng
// @DateTime 2022/7/19 9:56 AM

package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get_allow_headers_as_one_line_without_end_sep(t *testing.T) {
	asserter := assert.New(t)
	expected := `Access-Control-Allow-Headers, Authorization, User-Agent, Keep-Alive, Content-Type, X-Requested-With, X-CSRF-Token, AccessToken, Token`
	actual := getAllowHeaders()
	asserter.Equal(expected, actual)
}
