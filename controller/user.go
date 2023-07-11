package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
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
	// 登录成功逻辑在签发 IssueToken 后置中间键中执行
}
