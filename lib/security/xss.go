// @Author Ben.Zheng
// @DateTime 2022/7/18 2:28 PM

package security

import (
	"encoding/json"
	"html/template"

	"github.com/benz9527/idk-doc/lib/collections"
)

// URL requests with HTML or Javascript code will make browser
// does some unexpected operations. Then the browser will expose
// sensitive data to hacker.

func JsonXssFilter(unsafe ...any) (safe []string) {
	if len(unsafe) == 0 {
		return safe
	}

	safe = make([]string, len(unsafe))
	for _, u := range unsafe {
		// Go json encode will escape HTML, Javascript XSS characters
		res, err := json.Marshal(u)
		if err != nil {
			res = []byte{}
		}
		safe = append(safe, collections.Bytes2String(res))
	}
	return safe
}

func HtmlXssFilter(content string) (safe string) {
	if len(content) == 0 {
		return ""
	}

	safe = template.HTMLEscaper(content)
	return safe
}
