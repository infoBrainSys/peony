package search

import (
	"context"
	"peony/api/v1"
	"peony/model"
	"peony/service"
	"peony/utils"
)

type sSearch struct {
}

func init() {
	service.RegisterSearch(New())
}

func New() service.ISearch {
	return &sSearch{}
}

func (s *sSearch) Search(c context.Context, in *v1.SearchReq) (result *v1.SearchResp, err error) {
	var search v1.SearchResp
	tx := utils.DB.Model(model.Product{}).Where(in).Scan(&search)
	return &search, tx.Error
}
