package v1

import (
	"github.com/freelifer/gin-plus"
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/output"
	"github.com/gravida/gcs/pkg/utils"
	"log"
)

type RoleController struct {
	ginplus.Controller
}

type Role struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Enable bool   `json:"enable"`
}

func (controller *RoleController) List(c *gin.Context) {
	page, pageSize := utils.DefaultQueryPage(c)
	roles, err := models.QueryAllRoles(page, pageSize)
	if err != nil {
		// 系统错误
		output.InternalErrorJSON(c, err.Error())
		return
	}

	total, _ := models.CountRoles()
	output.SuccessJSON1(c, gin.H{
		"pager": gin.H{"page": page, "pageSize": pageSize, "total": total},
		"data":  roles,
	})
}

func (controller *RoleController) Get(g *gin.Context) {
	id, err := utils.ParamFromID(g, "id")
	if err != nil {
		output.BadRequestJSON(g, err.Error())
		return
	}
	user, has, err := models.QueryRoleByID(id)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}
	if !has {
		output.NotFoundJSON(g, "role not found")
		return
	}
	output.SuccessJSON(g, user)
}

func (controller *RoleController) Post(g *gin.Context) {
	var remoteRole Role
	err := g.BindJSON(&remoteRole)
	if err != nil {
		output.BadRequestJSON(g, err.Error())
		return
	}

	if len(remoteRole.Name) == 0 {
		output.BadRequestJSON(g, "role name must not empty")
		return
	}

	has, err := models.ExistRoleByName(0, remoteRole.Name)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(g, "role name repeat")
		return
	}

	var role models.Role
	role.Name = remoteRole.Name
	role.Desc = remoteRole.Desc
	role.Enable = remoteRole.Enable

	err = models.AddRole(&role)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}

	g.JSON(200, gin.H{
		"data": role,
	})
}
func (controller *RoleController) Put(g *gin.Context) {
	id, err := utils.ParamFromID(g, "id")
	if err != nil {
		output.BadRequestJSON(g, err.Error())
		return
	}

	role, has, err := models.QueryRoleByID(id)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}
	if !has {
		output.BadRequestJSON(g, "role not found")
		return
	}

	var remoteRole Role
	err = g.BindJSON(&remoteRole)
	if err != nil {
		output.BadRequestJSON(g, err.Error())
		return
	}

	if remoteRole.Name != "" && remoteRole.Name != role.Name {
		has, err = models.ExistRoleByName(0, remoteRole.Name)
		if err != nil {
			output.InternalErrorJSON(g, err.Error())
			return
		}
		if has {
			output.BadRequestJSON(g, "role name repeat")
			return
		}
		role.Name = remoteRole.Name
	}

	if remoteRole.Desc != "" {
		role.Desc = remoteRole.Desc
	}

	role.Enable = remoteRole.Enable
	log.Println(role)

	err = models.UpdateRole(role)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}

	g.JSON(200, gin.H{
		"data": role,
	})
}
