package common_dto

type Pager struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func (p *Pager) GetOffset() int64 {
	offset := (p.Page - 1) * p.PageSize

	if offset >= p.TotalRows {
		offset = (p.TotalRows/p.PageSize - 1) * p.PageSize
	}

	return offset
}
