package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "peony/logic/jwt"
	_ "peony/logic/product"
	_ "peony/logic/search"
	_ "peony/logic/user"
)

type BaseContext struct {
	Ctx  *gin.Context
	Path string
}

func NewBaseContext(ctx *gin.Context) *BaseContext {
	return &BaseContext{Ctx: ctx}
}

// Response 返回 code:{0:成功, 1:失败}，data:任意消息
func (b *BaseContext) Response(status int, code int, data ...interface{}) *BaseContext {
	b.Ctx.JSON(status, gin.H{
		"code": code,
		"data": data,
	})
	return b
}

// To 前往的路径
func (b *BaseContext) To(path string) *BaseContext {
	b.Path = path
	return b
}

// Redirect 跳转
func (b *BaseContext) Redirect() {
	b.Ctx.Redirect(http.StatusMovedPermanently, b.Path)
}
