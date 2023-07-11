package main

import (
	"peony/cmd"
	_ "peony/logic"
	"peony/utils"
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
}

func main() {
	cmd.NewServer()
}
