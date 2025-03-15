package httputils

type ListResponse struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

func MakeListResponse(items interface{}, total int64) ListResponse {
	return ListResponse{Items: items, Total: total}
}
