package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
)

func Publish(c *gin.Context) {
	base := logic.NewBaseContext(c)
	var productReq v1.ProductReq
	err := base.Ctx.ShouldBind(&productReq)
	if err != nil {
		base.Response(http.StatusBadRequest, consts.Failed, err.Error())
		return
	}
	err = service.Product().Publish(base.Ctx, &productReq)
	if err != nil {
		base.Response(http.StatusInternalServerError, consts.Failed, consts.PublishFailed)
		return
	}
	base.To("/").Response(http.StatusOK, consts.Success, consts.PublishSuccess)
}
