package routers

import (
	// "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/apis/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(gzip.Gzip(gzip.BestCompression))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Group := r.Group("/v1")
	{
		// easyapi.Router(v1, new(version1.ProjectApi))
		v1Group.GET("/musics", v1.Musics)
		v1Group.GET("/musics/:id", v1.GetMusic)
		v1Group.POST("/musics", v1.PostMusic)
		// v1.GET("/projects/:id", project.Get)
		// v1.POST("/projects", project.Post)
		// v1.PUT("/projects", project.Put)
		// v1.DELETE("/projects/:id", project.Del)
	}
	return r
}
