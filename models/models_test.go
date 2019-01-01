package models_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/settings"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	settings.DatabaseCfg.Type = "sqlite3"
	settings.DatabaseCfg.Path = "data.db"
	Setup()
	Convey("根据用户名判断用户是否存在", t, func() {
		isok, err := ExistUserByName(0, "pig")
		// So(err, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(isok, ShouldEqual, false)
	})

	Convey("添加邮箱", t, func() {
		user := &User{Email: "pig1"}
		err := AddUser(user)
		// So(err, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("查询用户列表", t, func() {
		users, err := QueryAllUsers(1, 20)
		// So(err, ShouldNotBeNil)
		So(err, ShouldBeNil)
		t.Log(gin.H{
			"data": users,
		})
		// So(err, ShouldBeNil)
	})

}

func Add(a, b int) int {
	return a + b
}
