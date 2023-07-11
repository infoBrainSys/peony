package v1

type ProductReq struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Price   string `form:"price" json:"price" binding:"required"`
}
