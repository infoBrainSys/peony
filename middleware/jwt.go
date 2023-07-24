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
	token := base.Ctx.Request.Header.Get("Authorization")[7:]
	email, _ := service.JWT().GetEmail(token)

	// 不存在则返回 err，如果 redis 中存在 token，则继续执行逻辑
	err := utils.RDB.Get(base.Ctx, email+"token").Err()
	if err != nil {
		base.To("/").AbortWithStatus(http.StatusUnauthorized, consts.Failed, consts.Unauthorized).Redirect()
		return
	}

	err = service.JWT().AuthJwtToken(token)
	if err != nil {
		base.To("/").AbortWithStatus(http.StatusUnauthorized, consts.Failed, consts.TokenIsNotValid).Redirect()
		return
	}
}
