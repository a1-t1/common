package httputils

import (
	"net/http"
	"strconv"

	"github.com/a1-t1/common/pkg/utils"
)

func OffsetAndLimitFromRequest(r *http.Request) (int64, int64) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("page_size"), 10, 64)
	if err != nil {
		limit = 10
	}

	offset := (page - 1) * limit
	return offset, limit
}

func ParsePaginationParams(r *http.Request) *utils.PaginationParams {
	offset, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("page_size"), 10, 64)
	if err != nil {
		limit = 10
	}
	return &utils.PaginationParams{
		Page:     offset,
		PageSize: limit,
	}
}
