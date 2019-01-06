package v1

import (
	"github.com/freelifer/gin-plus"
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/output"
	"github.com/gravida/gcs/pkg/utils"
)

type PermissionController struct {
	ginplus.Controller
}

type OperationController struct {
	ginplus.Controller
}

type Operation struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Key    string `json:"key"`
	Enable bool   `json:"enable"`
}

func (controller *OperationController) List(c *gin.Context) {
	page, pageSize := utils.DefaultQueryPage(c)
	operations, err := models.QueryAllOperations(page, pageSize)
	if err != nil {
		// 系统错误
		output.InternalErrorJSON(c, err.Error())
		return
	}

	total, _ := models.CountOperation()
	output.SuccessJSON1(c, gin.H{
		"pager": gin.H{"page": page, "pageSize": pageSize, "total": total},
		"data":  operations,
	})

}

func (controller *OperationController) Post(c *gin.Context) {
	var op Operation
	err := c.BindJSON(&op)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if len(op.Name) == 0 || len(op.Key) == 0 {
		output.BadRequestJSON(c, "Operation name and key must not empty")
		return
	}

	has, err := models.ExistOperationByName(0, op.Name)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "Operation name repeat")
		return
	}

	has, err = models.ExistOperationByKey(0, op.Key)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "Operation key repeat")
		return
	}

	var operation models.Operation
	operation.Name = op.Name
	operation.Desc = op.Desc
	operation.Key = op.Key
	operation.Enable = op.Enable

	err = models.AddOperation(&operation)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": operation,
	})
}

func (controller *OperationController) Put(c *gin.Context) {
	id, err := utils.ParamFromID(c, "id")
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	op, has, err := models.QueryOperationByID(id)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}
	if !has {
		output.BadRequestJSON(c, "operation not found")
		return
	}

	var remoteOp Operation
	err = c.BindJSON(&remoteOp)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if remoteOp.Name != "" && remoteOp.Name != op.Name {
		has, err = models.ExistOperationByName(0, remoteOp.Name)
		if err != nil {
			output.InternalErrorJSON(c, err.Error())
			return
		}
		if has {
			output.BadRequestJSON(c, "Operation name repeat")
			return
		}
		op.Name = remoteOp.Name
	}

	if remoteOp.Key != "" && remoteOp.Key != op.Key {
		has, err = models.ExistOperationByKey(0, remoteOp.Key)
		if err != nil {
			output.InternalErrorJSON(c, err.Error())
			return
		}
		if has {
			output.BadRequestJSON(c, "Operation key repeat")
			return
		}
		op.Key = remoteOp.Key
	}

	if remoteOp.Desc != "" && remoteOp.Desc != op.Desc {
		op.Desc = remoteOp.Desc
	}

	err = models.UpdateOperation(op)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": op,
	})
}

func (controller *PermissionController) List(c *gin.Context) {
	page, pageSize := utils.DefaultQueryPage(c)
	permissionOps, err := models.QueryAllOperations(page, pageSize)
	if err != nil {
		// 系统错误
		output.InternalErrorJSON(c, err.Error())
		return
	}
	output.SuccessJSON(c, permissionOps)
}

func (controller *PermissionController) Get(g *gin.Context) {
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

func (controller *PermissionController) Post(c *gin.Context) {
	var op Operation
	err := c.BindJSON(&op)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if len(op.Name) == 0 || len(op.Key) == 0 {
		output.BadRequestJSON(c, "Operation name and key must not empty")
		return
	}

	has, err := models.ExistOperationByName(0, op.Name)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "Operation name repeat")
		return
	}

	has, err = models.ExistOperationByKey(0, op.Key)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "Operation key repeat")
		return
	}

	var operation models.Operation
	operation.Name = op.Name
	operation.Desc = op.Desc
	operation.Key = op.Key
	operation.Enable = op.Enable

	err = models.AddOperation(&operation)
	if err != nil {
		output.InternalErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": operation,
	})
}
func (controller *PermissionController) Put(g *gin.Context) {
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
