package model

type Paginated[T interface{}] struct {
	Page         int   `json:"page"`
	PageSize     int   `json:"pageSize"`
	TotalRecords int64 `json:"totalRecords"`
	Data         []T   `json:"data"`
}
