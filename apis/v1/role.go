package v1

import (
	"github.com/freelifer/gin-plus"
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/output"
	"github.com/gravida/gcs/pkg/utils"
)

type RoleController struct {
	ginplus.Controller
}

type Role struct {
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

func (r *RoleController) List(g *gin.Context) {
	page, pageSize := utils.DefaultQueryPage(g)
	roles, err := models.QueryAllRoles(page, pageSize)
	if err != nil {
		// 系统错误
		output.InternalErrorJSON(g, err.Error())
		return
	}
	output.SuccessJSON(g, roles)
}

func (r *RoleController) Get(g *gin.Context) {
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
func (r *RoleController) Post(g *gin.Context) {
	var role Role
	err := g.BindJSON(&role)
	if err != nil {
		output.BadRequestJSON(g, err.Error())
		return
	}

	if len(role.Name) == 0 {
		output.BadRequestJSON(g, "role name must not empty")
		return
	}

	has, err := models.ExistRoleByName(0, role.Name)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(g, "role name repeat")
		return
	}

	var rr models.Role
	rr.Name = role.Name
	rr.Enable = role.Enable

	err = models.AddRole(&rr)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}

	g.JSON(200, gin.H{
		"data": rr,
	})
}
func (r *RoleController) Put(g *gin.Context) {
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

	if remoteRole.Name != "" {
		role.Name = remoteRole.Name

		has, err = models.ExistRoleByName(0, role.Name)
		if err != nil {
			output.InternalErrorJSON(g, err.Error())
			return
		}
		if has {
			output.BadRequestJSON(g, "role name repeat")
			return
		}
	}

	role.Enable = remoteRole.Enable

	err = models.UpdateRole(role)
	if err != nil {
		output.InternalErrorJSON(g, err.Error())
		return
	}

	g.JSON(200, gin.H{
		"data": role,
	})
}
