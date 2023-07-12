package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
	"time"
)

// AuthJwtToken 校验 Token
func AuthJwtToken(c *gin.Context) {
	base := logic.NewBaseContext(c)
	token := base.Ctx.Request.Header.Get("Authorization")[7:]
	email, _ := service.JWT().GetEmail(token)

	// 不存在则返回 err，如果 redis 中存在 token，则继续执行签发逻辑
	err := utils.RDB.Get(base.Ctx, email+token).Err()
	if err != nil {
		base.To("/").Response(http.StatusInternalServerError, consts.Failed, err).Redirect()
		c.Abort()
		return
	}

	err = service.JWT().AuthJwtToken(token)
	if err != nil {
		base.To("/").Response(http.StatusUnauthorized, consts.Failed, err).Redirect()
		c.Abort()
		return
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

	// 将 token 写入 redis，判断 token 是否存在 redis 中，如果存在则这个 token 有效
	defer utils.RDB.Set(
		base.Ctx,
		email+tokenStr, // token key 构造方式： email+token
		tokenStr,
		time.Second*time.Duration(utils.V.GetInt64("jwt.exp")),
	)
	if err != nil {
		base.To(base.Ctx.Request.RequestURI).
			Response(http.StatusInternalServerError, consts.Failed, err.Error()).
			Redirect()
		c.Abort()
		return
	}
	base.To("/").Response(http.StatusOK, consts.Success, tokenStr, consts.LoginSuccess).Redirect()
}
