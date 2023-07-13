package v1

type RegisterReq struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
	Code            string `form:"code" json:"code" binding:"required"` // 邮箱校验码
}

type RegisterRes struct {
	Code    int // 状态码
	Message string
}
