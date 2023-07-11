package v1

type RegisterReq struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterRes struct {
	Code    int
	Message string
}
