package output

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
