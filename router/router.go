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
		index.POST("/login", controller.Login)
		index.POST("/logout", middleware.AuthJwtToken, controller.Logout) // TODO 注销
		index.POST("/sendCode", controller.SendEmail)
	}

	// 用户组
	//user := app.Group("/user")
	//{
	//	//user.DELETE("/signOut", controller.SignOut) // TODO 删除账号
	//}

	// 帖子组
	post := app.Group("/post")
	post.Use(middleware.AuthJwtToken)
	{
		post.POST("/publish", controller.Publish)
	}

	// 搜索组，无需鉴权中间键
	search := app.Group("/search")
	{
		search.GET("/", controller.Search)
	}

	// 订单组

	order := app.Group("/order")
	{
		order.PUT("/check", middleware.AuthJwtToken, controller.Check)
		order.GET("/notify", controller.ParseNotify)
	}
}
