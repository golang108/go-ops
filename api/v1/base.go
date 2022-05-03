package v1

type DeleteRes string

type Page struct {
	Total     int `json:"total" dc:"总数"`
	PageNum   int `json:"pageNum" dc:"第几页"`
	PageSize  int `json:"pageSize" dc:"每页的数量"`
	PageTotal int `json:"pageTotal" dc:"总共多少页"`
}

type PageReq struct {
	PageNum  int `json:"pageNum" dc:"第几页" d:"1"  v:"min:0#分页号码错误"`
	PageSize int `json:"pageSize" dc:"每页的数量" d:"10" v:"max:50#分页数量最大50条"`
}

type KVString struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}
