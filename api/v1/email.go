package v1

type SendEmailReq struct {
	Email string `form:"email" json:"email" binding:"required"`
}
