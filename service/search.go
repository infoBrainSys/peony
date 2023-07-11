package service

import (
	"context"
	v1 "peony/api/v1"
)

type ISearch interface {
	Search(c context.Context, in *v1.SearchReq) (result *v1.SearchResp, err error)
}

var (
	localSearch ISearch
)

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for interface ISearch, forgot register?")
	}
	return localSearch
}

func RegisterSearch(i ISearch) {
	localSearch = i
}
