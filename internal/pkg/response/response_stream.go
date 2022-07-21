// @Author Ben.Zheng
// @DateTime 2022/7/19 10:16 AM

package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Stream struct {
	context     *fiber.Ctx
	jsonContent string
	isJson      bool
	*appResponse
}

func NewResponseStream(ctx *fiber.Ctx) *Stream {
	return &Stream{
		context: ctx,
		appResponse: &appResponse{
			coreResponse: &coreResponse{
				httpCode:       http.StatusOK,
				appStatusCode:  1,
				errorDetails:   "",
				additionalData: nil,
			},
			requestUrl: "",
		},
		jsonContent: "",
		isJson:      false,
	}
}

type coreResponse struct {
	appStatusCode  int
	httpCode       int
	errorDetails   string
	additionalData any
}

type appResponse struct {
	*coreResponse
	requestUrl string
}

func (r *appResponse) ErrorDetails() string {
	if len(r.errorDetails) == 0 {
		return ""
	}

	if isHttpStatusCodeAbnormal(r.httpCode) && r.isAppCodeAbnormal() {

	}
	return ""
}

func (r *appResponse) isAppCodeAbnormal() bool {
	return true
}

func isHttpStatusCodeAbnormal(code int) bool {
	return code >= 400 && code <= 599
}
