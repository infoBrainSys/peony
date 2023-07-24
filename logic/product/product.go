package product

import (
	v1 "peony/api/v1"
	"peony/model"
	"peony/service"
	"peony/utils"
	"time"

	"github.com/gin-gonic/gin"
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
		Peony:   utils.NewHashPeony(email + time.Now().String()), // 使用（邮箱+时间戳）组合作为帖子唯一链接标识
	}
	// 发布帖子
	return utils.DB.Create(&product).Error
}
