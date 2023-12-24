package common_assembly

import (
	"joineer-sms/dto/common_dto"
	"joineer-sms/po/common_po"
)

func ConvertPagerPoToDto(po *common_po.Pager) *common_dto.Pager {
	return &common_dto.Pager{
		Page:      po.Page,
		PageSize:  po.PageSize,
		TotalRows: po.TotalRows,
	}
}

func ConvertPagerDtoToPo(dto *common_dto.Pager) *common_po.Pager {
	totalPages := dto.TotalRows / dto.PageSize
	if dto.TotalRows != 0 && totalPages == 0 {
		totalPages++
	} else if dto.TotalRows%dto.PageSize != 0 {
		totalPages++
	}
	return &common_po.Pager{
		Page:       dto.Page,
		PageSize:   dto.PageSize,
		TotalRows:  dto.TotalRows,
		TotalPages: totalPages,
	}
}
