package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	consts "peony/const"
	"peony/logic"
	"peony/service"
	"peony/utils"
)

// AuthUser 用户对保护资源进行操作时需要鉴别用户是否存在，PS：由于JwtToke 的无状态特性，防止用户在删除账号后仍然可使用 Token 进行敏感操作（TODO：待优化）
func AuthUser(c *gin.Context) {
	base := logic.NewBaseContext(c)
	tkn := base.Ctx.Request.Header.Get("Authorization")[7:]
	fmt.Println(tkn)
	email, err := service.JWT().GetEmail(tkn)
	if err != nil {
		base.Response(http.StatusUnauthorized, consts.Failed, err.Error()).Redirect()
		c.Abort()
		return
	}
	if ok := service.User().UserExist(email); !ok {
		base.Response(http.StatusUnauthorized, consts.Failed, consts.UserNotExist).Redirect()
		c.Abort()
		return
	}
	utils.RDB.Set(base.Ctx, "email", email, 300)
	c.Next()
}
