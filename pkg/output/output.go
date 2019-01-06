package output

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// const ( StatusContinue = 100 StatusSwitchingProtocols = 101 StatusOK = 200 StatusCreated = 201 StatusAccepted = 202 StatusNonAuthoritativeInfo = 203 StatusNoContent = 204 StatusResetContent = 205 StatusPartialContent = 206
// StatusMultipleChoices = 300 StatusMovedPermanently = 301 StatusFound = 302 StatusSeeOther = 303 StatusNotModified = 304 StatusUseProxy = 305 StatusTemporaryRedirect = 307
// StatusBadRequest = 400 StatusUnauthorized = 401 StatusPaymentRequired = 402 StatusForbidden = 403 StatusNotFound = 404 StatusMethodNotAllowed = 405 StatusNotAcceptable = 406 StatusProxyAuthRequired = 407 StatusRequestTimeout = 408 StatusConflict = 409 StatusGone = 410 StatusLengthRequired = 411 StatusPreconditionFailed = 412 StatusRequestEntityTooLarge = 413 StatusRequestURITooLong = 414 StatusUnsupportedMediaType = 415 StatusRequestedRangeNotSatisfiable = 416 StatusExpectationFailed = 417 StatusTeapot = 418
// StatusInternalServerError = 500 StatusNotImplemented = 501 StatusBadGateway = 502 StatusServiceUnavailable = 503 StatusGatewayTimeout = 504 StatusHTTPVersionNotSupported = 505
// )
// func ErrorJSON(c *gin.Context, err Err) {
// 	// log err.Error() 内部错误
// 	log.Println(err.Error())
// 	code := err.Code()
// 	errmsg := GetMsg(code) + ", [req_id:xxxxxx]"

// 	c.JSON(http.StatusOK, gin.H{
// 		"errcode": code,
// 		"errmsg":  errmsg,
// 		// "errmsg": err.Error(),
// 	})
// }

func SuccessJSON1(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// BadRequestJSON - 400 服务器无法理解请求的格式，客户端不应当尝试再次使用相同的内容发起请求
func BadRequestJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": msg,
	})
}

// UnauthorizedJSON - 401 请求未授权
func UnauthorizedJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
}

// ForbiddenJSON - 403 禁止访问
func ForbiddenJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": msg,
	})
}

// NotFoundJSON - 404 资源未找到
func NotFoundJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": msg,
	})
}

// 500  （服务器内部错误）  服务器遇到错误，无法完成请求。
func InternalErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": msg,
	})
}
