/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:router.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

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

		// 获取通知
		v1.GET("inbox/list", api.InboxList)
		// 获取是否有未读邮件
		v1.GET("inbox/list/unread", api.InboxListUnread)
		// 读取邮件（消除未读标记）
		v1.GET("inbox/looked", api.InboxLooked)

		// 认证1：游客根据地址和学生身份证号和姓名进行信息查询
		v1.GET("certificate/address", api.CertificateWithAddress)
		// 认证2：游客根据文件和编号查询真伪
		v1.GET("certificate/file", api.CertificateWithFile)

		// 检查是否登陆
		v1.GET("me", api.Me)

		// 需要登录保护的
		authUser := v1.Group("")
		authUser.Use(middleware.AuthUserRequired())
		{
			// User Routing
			authUser.GET("user/me", api.UserMe)
			authUser.DELETE("user/logout", api.UserLogout)
			authUser.PUT("user/modify/user", api.UserModify)
			authUser.POST("user/identity/auth", api.UserAuth)
			authUser.GET("user/certification/list", api.Certification)
			authUser.GET("user/certification/getInfo", api.CertificationInfo)
			authUser.GET("user/certification/file/:filename", api.CertificationFile)
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
			authAdmin.GET("admin/msg/list", api.MsgList)
			authAdmin.PUT("admin/authenticated/user", api.AdminAuthUser)
			authAdmin.PUT("admin/academic/certification", api.AdminACStudent)
			authAdmin.GET("admin/userInfo", api.AdminGetUser)
			authAdmin.GET("admin/msg", api.GetAMsg)

		}
		authUniversity := v1.Group("")
		authUniversity.Use(middleware.AuthUniversityRequired())
		{
			// University Routing
			authUniversity.GET("university/me", api.UniversityMe)
			authUniversity.DELETE("university/logout", api.UniversityLogout)
			authUniversity.PUT("university/modify/university", api.UniversityModify)
			authUniversity.POST("university/studentAuth", api.StudentAuth)
			authUniversity.GET("university/studentAuth/list", api.StudentAuthList)
			authUniversity.GET("university/studentAuth/list/msg", api.StudentAuthMsg)
		}
	}
	return r
}
