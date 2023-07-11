package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
)

// AuthJwtToken 校验 Token
func AuthJwtToken(c *gin.Context) {
	base := logic.NewBaseContext(c)
	token := base.Ctx.Request.Header.Get("Authorization")
	err := service.JWT().AuthJwtToken(token)
	if err != nil {
		base.To("/").Response(http.StatusUnauthorized, consts.Failed, err.Error()).Redirect()
		c.Abort()
	}
	c.Next()
}

// IssueToken 签发 Token 中间键
func IssueToken(c *gin.Context) {
	c.Next()
	base := logic.NewBaseContext(c)

	info, err := utils.RDB.Get(base.Ctx, "err").Result()
	defer utils.RDB.Del(base.Ctx, "email", "err")
	if err != nil || info != "" {
		base.To(base.Ctx.Request.RequestURI).
			Response(http.StatusUnauthorized, consts.Failed, info).
			Redirect()
		c.Abort()
		return
	}

	email, err := utils.RDB.Get(base.Ctx, "email").Result()
	if err != nil {
		return
	}
	tokenStr, err := service.JWT().IssueToken(email)
	if err != nil {
		base.To(base.Ctx.Request.RequestURI).
			Response(http.StatusInternalServerError, consts.Failed, err.Error()).
			Redirect()
		c.Abort()
		return
	}
	base.To("/").Response(http.StatusOK, consts.Success, tokenStr, consts.LoginSuccess).Redirect()
}
