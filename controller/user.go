package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
	"time"
)

// Register 注册逻辑
func Register(c *gin.Context) {
	var registerReq v1.RegisterReq
	base := logic.NewBaseContext(c)

	err := base.Ctx.ShouldBind(&registerReq)
	if err != nil {
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, err)
		return
	}
	if registerReq.Password != registerReq.ConfirmPassword {
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, errors.New(consts.PasswordNotMatch))
		return
	}
	verifyCode, _ := utils.RDB.Get(base.Ctx, registerReq.Email+registerReq.Code).Result()
	if registerReq.Code != verifyCode {
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, errors.New(consts.VerifyCodeErr))
		return
	}

	err = service.User().Register(c, &registerReq)
	if err == nil {
		base.To("/login").Response(http.StatusOK, consts.Success, consts.RegisterSuccess).Redirect()
		return
	} else {
		base.AbortWithStatus(http.StatusUnprocessableEntity, consts.Failed, errors.New(consts.UserExist))
		return
	}
}

// Login 登录逻辑
func Login(c *gin.Context) {
	base := logic.NewBaseContext(c)
	var loginReq v1.LoginReq
	emailCh := make(chan string, 1)

	// 绑定用户输入
	if err := base.Ctx.ShouldBind(&loginReq); err != nil {
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, err)
		return
	}

	// 校验登录
	if err := service.User().Login(base.Ctx, &loginReq, emailCh); err != nil {
		base.AbortWithStatus(http.StatusUnauthorized, consts.Failed, err)
		return
	}

	// 签发token
	email := <-emailCh
	tokenStr, err := service.JWT().IssueToken(email)
	token := "Bearer " + tokenStr

	defer utils.RDB.Set(base.Ctx, email+"token", email+token, utils.V.GetDuration("jwt.exp")*time.Second)
	if err != nil {
		base.AbortWithStatus(http.StatusInternalServerError, consts.Failed, err)
		return
	}

	base.To("/").Response(http.StatusOK, consts.Success, token, consts.LoginSuccess).Redirect()
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
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, consts.SendEmailFailed)
		return
	}

	code, files, err := service.Email().ParseFiles()
	if err != nil {
		base.AbortWithStatus(http.StatusInternalServerError, consts.Failed, consts.SendEmailFailed)
		return
	}

	err = utils.SendEmail(sendEmailReq.Email, files)
	// redis 验证码键组合：邮箱+验证码
	defer utils.RDB.Set(base.Ctx,
		sendEmailReq.Email+code,
		code,
		time.Second*time.Duration(utils.V.GetInt64("email.codeExp")))

	if err != nil {
		base.AbortWithStatus(http.StatusBadRequest, consts.Failed, consts.SendEmailFailed)
		return
	}
	base.Response(http.StatusOK, consts.Success, consts.SendEmailSuccess, code)
}
