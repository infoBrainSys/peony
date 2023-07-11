package service

import (
	"github.com/gin-gonic/gin"
	"peony/api/v1"
)

type IProduct interface {
	Publish(c *gin.Context, in *v1.ProductReq) error
}

var (
	localProduct IProduct
)

func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

func RegisterProduct(i IProduct) {
	localProduct = i
}
