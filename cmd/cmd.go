package cmd

import (
	"github.com/gin-gonic/gin"
	"peony/router"
	"peony/utils"
	"peony/utils/email"
)

func init() {
	// 初始化viper
	utils.InitViper()

	// 初始化数据库
	utils.InitGorm()

	// 初始化redis
	utils.InitRedis()

	// 初始化 UUID
	utils.InitUUID()

	// 初始化支付
	utils.InitPay()

	// 初始化邮件服务
	email.InitEmail()
}

func NewServer() {
	app := gin.Default()
	app.LoadHTMLGlob("utils/email/*.html")
	router.RegisterRouter(app)
	app.Run()
}
