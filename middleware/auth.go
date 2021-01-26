package middleware

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		utype := session.Get("user_type")
		uid := session.Get("user_id")
		if uid != nil && utype != nil {
			user, err := model.GetUserWithType(utype.(string), uid)
			if err == nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}

//// AuthRequired 需要登录
//func AuthRequired() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if user, _ := c.Get("user"); user != nil {
//			if _, ok := user.(*model.User); ok {
//				c.Next()
//				return
//			}
//		}
//
//		c.JSON(200, serializer.CheckLogin())
//		c.Abort()
//	}
//}

// AuthUserRequired 需要普通用户进行登录
func AuthUserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.User); ok {
				if u.Status == model.Suspend {
					c.JSON(200, serializer.CheckSuspend())
				}
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// AuthAdminRequired 需要管理员进行登录
func AuthAdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.Admin); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// AuthUniversityRequired 需要学校用户进行登录
func AuthUniversityRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.University); ok {
				if u.Status == model.Suspend {
					c.JSON(200, serializer.CheckSuspend())
				}
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
