package utils

import "math"

type PaginatedResponse[T any] struct {
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	Data       *T    `json:"data"`
}

func (p PaginatedResponse[T]) Nothing() PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Page:       0,
		Size:       0,
		Total:      0,
		TotalPages: 0,
		Data:       nil,
	}
}

func (p PaginatedResponse[T]) WithData(data T, total int64, paginationParams *PaginationParams) *PaginatedResponse[T] {
	p.Data = &data
	if paginationParams == nil {
		paginationParams = &PaginationParams{
			Page:     1,
			PageSize: 10,
		}
	}
	p.Page = paginationParams.Page
	p.Size = paginationParams.PageSize
	p.TotalPages = int(math.Ceil(float64(total) / float64(paginationParams.PageSize)))
	p.Total = total
	return &p
}

type PaginationParams struct {
	Page     int64
	PageSize int64
}

func (p *PaginationParams) Offset() int64 {
	return (p.Page - 1) * p.PageSize
}
