package routers

import (
	// "github.com/gin-contrib/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/apis/v1"
	"net/http"
	"strings"
)

////// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		fmt.Println(origin)
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", origin)                                    // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	// r.Use(gzip.Gzip(gzip.BestCompression))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Group := r.Group("/v1")
	{
		v1Group.POST("/register", v1.Register)
		v1Group.POST("/login", v1.Login)
		v1Group.GET("/validate", v1.ValidateEmail)

		v1Group.GET("/users", v1.Users)
		v1Group.GET("/users/:id", v1.GetUser)
		v1Group.POST("/users", v1.PostUser)
		v1Group.PUT("/users/:id", v1.PutUser)

		// easyapi.Router(v1, new(version1.ProjectApi))
		v1Group.GET("/musics", v1.Musics)
		v1Group.GET("/musics/:id", v1.GetMusic)
		v1Group.POST("/musics", v1.PostMusic)
		v1Group.PUT("/musics/:id", v1.PutMusic)
		// v1.GET("/projects/:id", project.Get)
		// v1.POST("/projects", project.Post)
		// v1.PUT("/projects", project.Put)
		// v1.DELETE("/projects/:id", project.Del)
	}
	return r
}
