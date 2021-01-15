package routes

import (
	"net/http"
	"web_app/controllers"
	_ "web_app/docs"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetUp 开启gin
func SetUp() *gin.Engine {
	r := gin.Default()
	// middleware.JwtAuth()
	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.JwtAuth())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/luffy", func(c *gin.Context) {
		c.String(http.StatusOK, "luffy")
	})
	r.POST("/login", controllers.UserLogin)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/save", controllers.AddUser)
		userRouter.POST("/getlist", controllers.GetUsers)
		userRouter.POST("/del", controllers.DelUsers)
	}
	RoleRouter := r.Group("/role")
	{
		RoleRouter.POST("/save", controllers.AddRole)
		RoleRouter.POST("/getlist", controllers.GetRoles)
		RoleRouter.POST("/del", controllers.DelRoles)
	}
	menuRouter := r.Group("/menu")
	{
		menuRouter.POST("/save", controllers.AddMenu)
		menuRouter.POST("/get_tree", controllers.GetMenuTree)
		// RoleRouter.POST("/del", controllers.DelRoles)
	}

	return r
}
