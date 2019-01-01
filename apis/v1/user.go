package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/output"
	"github.com/gravida/gcs/pkg/utils"
)

type Email struct {
	Email string `json:"email"`
}

type User struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
}

// 注册第一步 验证邮箱
// curl -H "Content-Type:application/json" -X POST -d '{"email": ""}' http://localhost:54321/v1/register
func Register(c *gin.Context) {
	var email Email
	err := c.BindJSON(&email)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if len(email.Email) == 0 {
		output.BadRequestJSON(c, "Register email must not empty")
		return
	}

	if !utils.VerifyEmailFormat(email.Email) {
		output.BadRequestJSON(c, "Register must a real email")
		return
	}

	exist, err := models.ExistUserByEmail(0, email.Email)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if exist {
		output.BadRequestJSON(c, "Register must a new email")
		return
	}

	user := &models.User{Email: email.Email, Status: 0}
	err = models.AddUser(user)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"msg": user,
	})

}

func Login(c *gin.Context) {

}

// curl http://localhost:54321/v1/validate?code=1k2lj432k4j23lrh9314ut-341gu-1bvhervn
func ValidateEmail(c *gin.Context) {
	code := c.Query("code")

	if len(code) == 0 {
		output.BadRequestJSON(c, "ValidateEmail code must not empty")
		return
	}

	codeInfo, has, err := models.QueryCodeByCode(code)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if !has {
		output.BadRequestJSON(c, "ValidateEmail code is not exist")
		return
	}

	// 验证过期

	c.JSON(200, gin.H{
		"data": codeInfo,
	})
}

// curl http://localhost:54321/v1/users
func Users(c *gin.Context) {
	page, pageSize := utils.DefaultQueryPage(c)
	users, err := models.QueryAllUsers(page, pageSize)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	output.SuccessJSON(c, users)

}

// curl http://localhost:54321/v1/users/1
func GetUser(c *gin.Context) {
	id, err := utils.ParamFromID(c, "id")
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	user, has, err := models.QueryUserByID(id)
	if err != nil {
		output.NotFoundJSON(c, err.Error())
		return
	}
	if !has {
		output.NotFoundJSON(c, "user not found")
		return
	}
	output.SuccessJSON(c, user)
}

// curl -H "Content-Type:application/json" -X POST -d '{"name": "aaa", "passwd":"123456", "email":"q@q.com", "phone":"111"}' http://localhost:54321/v1/users
func PostUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if len(user.Name) == 0 || len(user.Passwd) == 0 || len(user.Email) == 0 || len(user.Phone) == 0 {
		output.BadRequestJSON(c, "user name/passwd/email/phone must not empty")
		return
	}

	if !utils.VerifyEmailFormat(user.Email) {
		output.BadRequestJSON(c, "user email must a real email")
		return
	}

	if !utils.VerifyPhoneFormat(user.Phone) {
		output.BadRequestJSON(c, "user phone must a real phone")
		return
	}

	has, err := models.ExistUserByName(0, user.Name)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "user name repeat")
		return
	}

	has, err = models.ExistUserByEmail(0, user.Email)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "user email repeat")
		return
	}

	has, err = models.ExistUserByPhone(0, user.Phone)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if has {
		output.BadRequestJSON(c, "user phone repeat")
		return
	}

	var u models.User
	u.Name = user.Name
	u.Passwd = user.Passwd
	u.Email = user.Email
	u.Phone = user.Phone
	if user.Status == 0 {
		user.Status = 1
	}
	u.Status = user.Status

	if err = u.GetUserSaltAndEncodePasswd(); err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	err = models.AddUser(&u)
	if err != nil {
		output.NotFoundJSON(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"data": u,
	})
}

// curl -H "Content-Type:application/json" -X PUT -d '{"name": "aaa"}' http://localhost:54321/v1/users
func PutUser(c *gin.Context) {
	id, err := utils.ParamFromID(c, "id")
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	u, has, err := models.QueryUserByID(id)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}
	if !has {
		output.BadRequestJSON(c, "user not found")
		return
	}

	var user User
	err = c.BindJSON(&user)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	if user.Name != "" {
		u.Name = user.Name

		has, err = models.ExistUserByName(0, user.Name)
		if err != nil {
			output.BadRequestJSON(c, err.Error())
			return
		}
		if has {
			output.BadRequestJSON(c, "user name repeat")
			return
		}
	}

	if user.Email != "" {
		u.Email = user.Email
	}

	if user.Phone != "" {
		u.Phone = user.Phone
	}

	if user.Status != 0 {
		u.Status = user.Status
	}

	err = models.UpdateUser(u)
	if err != nil {
		output.BadRequestJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": u,
	})

}
