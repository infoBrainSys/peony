package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "peony/api/v1"
	consts "peony/const"
	"peony/logic"
	"peony/service"
)

func Search(c *gin.Context) {
	base := logic.NewBaseContext(c)
	var searchReq v1.SearchReq
	err := base.Ctx.ShouldBind(&searchReq)
	if err != nil {
		base.Response(http.StatusBadRequest, consts.Failed, err.Error())
		return
	}
	result, err := service.Search().Search(base.Ctx, &searchReq)
	if err != nil {
		base.Response(http.StatusInternalServerError, consts.Failed, err.Error())
		return
	}
	base.Response(http.StatusOK, consts.Success, result)

}
