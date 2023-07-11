package product

import (
	"github.com/gin-gonic/gin"
	"peony/api/v1"
	"peony/model"
	"peony/service"
	"peony/utils"
)

type sProduct struct {
}

func init() {
	service.RegisterProduct(New())
}

func New() service.IProduct {
	return &sProduct{}
}

// Publish 发布帖子
func (p *sProduct) Publish(c *gin.Context, in *v1.ProductReq) error {
	// 从 token 中获取发布者 email
	email, _ := utils.RDB.Get(c, "email").Result()
	defer utils.RDB.Del(c, "email")
	var product = model.Product{
		UUID:    utils.UUID.String(),
		Title:   in.Title,
		Content: in.Content,
		Price:   in.Price,
		Status:  1,
		Peony:   utils.NewHashPeony(email + utils.UUID.String()),
	}
	// 发布帖子
	return utils.DB.Create(&product).Error
}
