package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	UploadRouter "gitee.com/under-my-umbrella/cloud/router/upload"
	UsersRouter "gitee.com/under-my-umbrella/cloud/router/users"
	// MiddleWare "gitee.com/under-my-umbrella/cloud/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 设置静态文件路由
	r.Static("/files/avatar", "./files/avatar")
	// 设置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 允许所有来源访问
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
	// 创建路由组
	api := r.Group("/api")
	// 设置不同模块的路由
	UploadRouter.SetupUploadRoutes(api)
	UsersRouter.SetupUsersRoutes(api)

	return r
}