package router

import (
	"github.com/gin-gonic/gin"
	"peony/controller"
	"peony/middleware"
)

func RegisterRouter(app *gin.Engine) {
	// 注册登录组
	index := app.Group("/")
	{
		index.POST("/register", controller.Register)
		index.POST("/login", middleware.IssueToken, controller.Login)
		index.POST("/logout", middleware.AuthUser, middleware.AuthJwtToken, controller.Logout) // TODO 注销
		index.POST("/sendCode", controller.SendEmail)
	}

	// 用户组
	//user := app.Group("/user")
	//{
	//	//user.DELETE("/signOut", controller.SignOut) // TODO 删除账号
	//}

	// 帖子组
	post := app.Group("/post")
	post.Use(middleware.AuthUser, middleware.AuthJwtToken)
	{
		post.POST("/publish", controller.Publish)
		post.GET("/search", controller.Search)
	}

	// 订单组
	order := app.Group("/order")
	post.Use(middleware.AuthUser, middleware.AuthJwtToken)
	{
		order.GET("/:peony", controller.Search)
		//order.PUT("/check", controller.Check) // TODO 支付订单
	}
}
