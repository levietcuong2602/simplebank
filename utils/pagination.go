package utils

import (
	"math"
)

type Pager struct {
	Offset         int32       `json:"offset"`
	Limit          int32       `json:"limit"`
	CurrentPageNum int32       `json:"current_page_num"`
	TotalCount     int32       `json:"total_count"`
	HasPrev        bool        `json:"has_prev"`
	HasNext        bool        `json:"has_next"`
	PrevPageNum    interface{} `json:"prev"`
	NextPageNum    interface{} `json:"next_page_num"`
	LastPageNum    int32       `json:"last_page_num"`
}

type Pagination struct {
	Pager Pager         `json:"pager"`
	Data  []interface{} `json:"data"`
}

type NewPaginationParams struct {
	PageNum    int32 `json:"page_num"`
	Limit      int32 `json:"limit"`
	TotalCount int32 `json:"total_count"`
	Data       []interface{}
}

func NewPagination(arg NewPaginationParams) Pagination {
	totalPage := int32(math.Ceil(float64(arg.TotalCount / arg.Limit)))
	currentPageNum := int32(0)
	if totalPage >= arg.PageNum {
		currentPageNum = int32(arg.PageNum)
	}
	hasPrev := currentPageNum > 1
	hasNext := currentPageNum < int32(totalPage)
	offset := int32(0)
	if currentPageNum > 0 {
		offset = int32((currentPageNum - 1) * int32(arg.Limit))
	}
	var prevPageNum interface{}
	if hasPrev {
		prevPageNum = int32(currentPageNum - 1)
	}
	var nextPageNum interface{}
	if hasPrev {
		nextPageNum = int32(currentPageNum + 1)
	}

	pager := Pager{
		Offset:         offset,
		Limit:          arg.Limit,
		CurrentPageNum: int32(currentPageNum),
		TotalCount:     arg.TotalCount,
		HasPrev:        hasPrev,
		HasNext:        hasNext,
		PrevPageNum:    prevPageNum,
		NextPageNum:    nextPageNum,
	}
	pagination := Pagination{
		Pager: pager,
		Data:  arg.Data,
	}

	return pagination
}
