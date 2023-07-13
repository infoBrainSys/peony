package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
	"peony/utils/email"
	"time"
)

// Register 注册逻辑
func Register(c *gin.Context) {
	var registerReq v1.RegisterReq
	base := logic.NewBaseContext(c)

	err := base.Ctx.ShouldBind(&registerReq)
	if err != nil {
		base.Response(http.StatusBadRequest, consts.Failed, consts.RegisterFailed)
		return
	}
	if registerReq.Password != registerReq.ConfirmPassword {
		base.Response(http.StatusBadRequest, consts.Failed, consts.PasswordNotMatch)
		c.Abort()
		return
	}
	verifyCode, _ := utils.RDB.Get(base.Ctx, registerReq.Email+registerReq.Code).Result()
	if registerReq.Code != verifyCode {
		base.Response(http.StatusBadRequest, consts.Failed, consts.VerifyCodeWrong)
		c.Abort()
		return
	}

	err = service.User().Register(c, &registerReq)
	if err == nil {
		base.To("/").Response(http.StatusOK, consts.Success, consts.RegisterSuccess).Redirect()
		return
	} else {
		base.Response(http.StatusInternalServerError, consts.Failed, err.Error())
		return
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

// Logout 退出登录, 把 redis 中的 token 删除（加入黑名单）
func Logout(c *gin.Context) {
	base := logic.NewBaseContext(c)
	tokenStr := base.Ctx.Request.Header.Get("Authorization")[7:]
	getEmail, _ := service.JWT().GetEmail(tokenStr)

	utils.RDB.Del(base.Ctx, tokenStr+getEmail)
}

// SendEmail 发送邮件
func SendEmail(c *gin.Context) {
	base := logic.NewBaseContext(c)

	var sendEmailReq v1.SendEmailReq

	err := base.Ctx.ShouldBind(&sendEmailReq)
	if err != nil {
		base.Response(http.StatusBadRequest, consts.Failed, consts.EmailIsNotValid)
		return
	}

	code, files, err := service.Email().ParseFiles()
	if err != nil {
		base.Response(http.StatusInternalServerError, consts.Failed, consts.SendEmailFailed)
		return
	}

	err = email.SendEmail(sendEmailReq.Email, files)
	// redis 验证码键组合：邮箱+验证码
	defer utils.RDB.Set(base.Ctx,
		sendEmailReq.Email+code,
		code,
		time.Second*time.Duration(utils.V.GetInt64("email.codeExp")))

	if err != nil {
		base.Response(http.StatusInternalServerError, consts.Failed, consts.SendEmailFailed)
		return
	}
	base.Response(http.StatusOK, consts.Success, consts.SendEmailSuccess, code)
}
