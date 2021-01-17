package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("login", api.Login)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 需要登录保护的
		authUser := v1.Group("")
		authUser.Use(middleware.AuthUserRequired())
		{
			// User Routing
			authUser.GET("user/me", api.UserMe)
			authUser.DELETE("user/logout", api.UserLogout)
		}
		authAdmin := v1.Group("")
		authAdmin.Use(middleware.AuthAdminRequired())
		{
			// Admin Routing
			authUser.GET("admin/me", api.UserMe)
			authUser.DELETE("admin/logout", api.UserLogout)
			authUser.POST("university/register", api.UserRegister)
		}
		authUniversity := v1.Group("")
		authUniversity.Use(middleware.AuthUniversityRequired())
		{
			// University Routing
			authUser.GET("university/me", api.UserMe)
			authUser.DELETE("university/logout", api.UserLogout)
		}
	}
	return r
}
