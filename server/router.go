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
		//v1.POST("aaa",api.AdminRegister)

		// 文件上传
		v1.POST("file/upload", api.FileUpload)
		// 文件下载
		v1.GET("file/download/:filename", api.FileDownload)

		// 需要登录保护的
		authUser := v1.Group("")
		authUser.Use(middleware.AuthUserRequired())
		{
			// User Routing
			authUser.GET("user/me", api.UserMe)
			authUser.DELETE("user/logout", api.UserLogout)
			authUser.PUT("user/modify/user", api.UserModify)
		}
		authAdmin := v1.Group("")
		authAdmin.Use(middleware.AuthAdminRequired())
		{
			// Admin Routing
			authAdmin.GET("admin/me", api.AdminMe)
			authAdmin.DELETE("admin/logout", api.AdminLogout)
			authAdmin.POST("university/register", api.UniversityRegister)
			authAdmin.PUT("admin/modify/admin", api.AdminModify)
			authAdmin.PUT("admin/modify/user", api.AdminModifyUser)
			authAdmin.PUT("admin/modify/university", api.AdminModifyUniversity)
		}
		authUniversity := v1.Group("")
		authUniversity.Use(middleware.AuthUniversityRequired())
		{
			// University Routing
			authUniversity.GET("university/me", api.UniversityMe)
			authUniversity.DELETE("university/logout", api.UniversityLogout)
			authUniversity.PUT("university/modify/university", api.UniversityModify)
		}
	}
	return r
}
