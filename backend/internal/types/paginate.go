package types

import (
	"strconv"
)

type Paginate struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Offset    int    `json:"-"`
	Limit     int    `json:"-"`
	SortBy    string `json:"sort_by,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

type PaginateContextKey string

var PaginateCtxKey = PaginateContextKey("paginate")

func DefaultPaginate() *Paginate {
	return &Paginate{
		Page:      1,
		Size:      10,
		SortBy:    "created_at",
		SortOrder: "desc",
	}
}

func (p *Paginate) CalculateOffsetLimit() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Size < 1 {
		p.Size = 10
	}
	if p.Size > 100 {
		p.Size = 100
	}

	p.Offset = (p.Page - 1) * p.Size
	p.Limit = p.Size
}

func ParsePaginateFromQuery(pageStr, sizeStr, sortBy, sortOrder string) *Paginate {
	paginate := DefaultPaginate()

	if pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			paginate.Page = page
		}
	}

	if sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil && size > 0 {
			paginate.Size = size
		}
	}

	if sortBy != "" {
		paginate.SortBy = sortBy
	}

	if sortOrder != "" {
		if sortOrder == "asc" || sortOrder == "desc" {
			paginate.SortOrder = sortOrder
		}
	}

	paginate.CalculateOffsetLimit()
	return paginate
}
