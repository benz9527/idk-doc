// @Author Ben.Zheng
// @DateTime 2022/7/18 3:01 PM

package security

import (
	"encoding/json"
	"testing"

	"idk-doc/lib/collections"

	"github.com/stretchr/testify/assert"
)

func Test_HTML_XSS_filter(t *testing.T) {
	asserter := assert.New(t)

	html := `<h1>This is a XSS H1 tag</h1>`
	html = HtmlXssFilter(html)
	var expected = `&lt;h1&gt;This is a XSS H1 tag&lt;/h1&gt;`
	asserter.Equal(expected, html)
}

func Test_HTML_XSS_filter_JSON(t *testing.T) {
	asserter := assert.New(t)

	obj := struct {
		Name string
		Age  int
	}{
		Name: `<p>obj name is p tag</p>`,
		Age:  22,
	}

	bytes, err := json.Marshal(obj)
	asserter.NoError(err)
	res := collections.Bytes2String(bytes)
	t.Log(res)
}

func Test_Json_XSS_filter(t *testing.T) {
	obj := struct {
		Name string
		Age  int
		Js   string
	}{
		Name: `<p>obj name is p tag</p>`,
		Age:  22,
		Js:   `<script>alert('you have been pwned')</script>`,
	}

	res := JsonXssFilter(obj)
	t.Log(res)
}
