package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/output"
	"github.com/gravida/gcs/pkg/utils"
)

type Music struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Url     string `json:"url"`
	Source  string `json:"source"`
}

// curl http://localhost:8080/v1/musics
func Musics(c *gin.Context) {
	page := utils.DefaultQueryForInt(c, "page", 1)
	pageSize := utils.DefaultQueryForInt(c, "pageSize", 10)
	musics, err := models.GetMusics(page, pageSize)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
	}
	output.SuccessJSON(c, musics)
}

// curl http://localhost:8080/v1/musics/1
func GetMusic(c *gin.Context) {
	id, err := utils.ParamFromID(c, "id")
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	music, err := models.GetMusicByID(id)
	if err != nil {
		output.NotFoundJSON(c, err.Error())
		return
	}
	output.SuccessJSON(c, music)
}

// curl -H "Content-Type:application/json" -X POST -d '{"name": "aaa", "picture":"picture", "url":"url", "source":"github"}' http://localhost:8080/v1/musics
func PostMusic(c *gin.Context) {
	var music models.Music
	err := c.BindJSON(&music)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if len(music.Name) == 0 || len(music.Picture) == 0 || len(music.Url) == 0 || len(music.Source) == 0 {
		output.BadRequestJSON(c, "music name/picture/url/source must not empty")
		return
	}

	_, has, err := models.ExistMusicByName(music.Name)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "music name/picture/url/source must not empty")
		return
	}
	err = models.AddMusic(&music)
	if err != nil {
		output.NotFoundJSON(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"id": music.Id,
	})
}

func PutMusic(c *gin.Context) {
	id, err := utils.ParamFromID(c, "id")
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	music, err := models.GetMusicByID(id)
	if err != nil {
		output.NotFoundJSON(c, err.Error())
		return
	}
	var m Music

	err = c.BindJSON(&m)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if m.Name != "" {
		music.Name = m.Name
	}

	if m.Picture != "" {
		music.Picture = m.Picture
	}

	if m.Url != "" {
		music.Url = m.Url
	}

	if m.Source != "" {
		music.Source = m.Source
	}

	err = models.UpdateMusic(music)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	output.SuccessJSON(c, music)
}
