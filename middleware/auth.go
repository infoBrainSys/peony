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
	utils.RDB.Set(base.Ctx, "email", email, 0)
	c.Next()
}
