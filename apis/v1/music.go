package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/pkg/utils"
	"net/http"
)

func Musics(c *gin.Context) {

	page := utils.DefaultQueryForInt(c, "page", 1)
	pageSize := utils.DefaultQueryForInt(c, "pageSize", 10)
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data":    gin.H{"page": page, "pageSize": pageSize},
	})
}

func GetMusic(c *gin.Context) {

	c.String(200, "GetMusic, ...")
}

func PostMusic(c *gin.Context) {

	c.String(200, "PostMusic, ...")
}
