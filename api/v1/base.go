package v1

type DeleteRes string

type Page struct {
	Total     int64 `json:"total" dc:"总数"`
	PageNum   int64 `json:"pageNum" dc:"第几页"`
	PageSize  int   `json:"pageSize" dc:"每页的数量"`
	PageTotal int   `json:"pageTotal" dc:"总共多少页"`
}

type PageReq struct {
	PageNum  int64 `json:"pageNum" dc:"第几页"`
	PageSize int   `json:"pageSize" dc:"每页的数量"`
}
