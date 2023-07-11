package v1

type LoginReq struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginRes struct {
	Code    int
	Message string
}
