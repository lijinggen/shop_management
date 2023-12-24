package common_po

type CommonResp struct {
}

type Pager struct {
	Page       int64 `form:"pager.page" json:"page"`
	PageSize   int64 `form:"pager.page_size" json:"page_size"`
	TotalRows  int64 `form:"total_rows" json:"total_rows"`
	TotalPages int64 `form:"total_pages" json:"total_pages"`
}
