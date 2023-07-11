package v1

type SearchReq struct {
	Peony string `form:"peony" json:"peony" binding:"required"`
}

type SearchResp struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Price   string `json:"price"`
}
