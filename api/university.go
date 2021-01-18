package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"singo/service"
)

// UniversityRegister 学校用户注册接口
func UniversityRegister(c *gin.Context) {
	var service service.UniversityRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UniversityMe 学校用户详情
func UniversityMe(c *gin.Context) {
	if university, ok := CurrentUser(c).(*model.University); ok {
		res := serializer.BuildUniversityResponse(*university)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CheckLogin())
	}
}

// UniversityLogout 学校用户登出
func UniversityLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
