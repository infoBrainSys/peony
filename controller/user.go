package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
)

// Register 注册逻辑
func Register(c *gin.Context) {
	var registerReq v1.RegisterReq
	base := logic.NewBaseContext(c)

	err := base.Ctx.ShouldBind(&registerReq)
	if err != nil {
		return
	}
	err = service.User().Register(c, &registerReq)
	if err == nil {
		base.To("/").Response(http.StatusOK, consts.Success, consts.RegisterSuccess).Redirect()
	} else {
		base.Response(http.StatusInternalServerError, consts.Failed, consts.RegisterFailed)
	}
}

// Login 登录逻辑
func Login(c *gin.Context) {
	base := logic.NewBaseContext(c)
	var loginReq v1.LoginReq
	err := base.Ctx.ShouldBind(&loginReq)
	if err != nil {
		base.Response(http.StatusBadRequest, consts.Failed, consts.LoginFailed)
		return
	}
	err = service.User().Login(c, &loginReq)
	if err != nil {
		base.Response(http.StatusInternalServerError, consts.Failed, consts.LoginFailed)
		return
	}
	// 登录成功逻辑在 IssueToken 后置中间键中执行
}

// Logout 退出登录, 把 jwt token 加入黑名单
func Logout(c *gin.Context) {
	base := logic.NewBaseContext(c)
	tokenStr := base.Ctx.Request.Header.Get("Authorization")[7:]
	email, _ := service.JWT().GetEmail(tokenStr)

	utils.RDB.Del(base.Ctx, tokenStr+email)
}
