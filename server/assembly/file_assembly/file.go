package file_assembly

import (
	"github.com/shop_management/dto/file_dto"
	"github.com/shop_management/po/file_po"
)

func ConvertURPoToDto(req *file_po.UploadReq) *file_dto.UploadReq {
	return &file_dto.UploadReq{
		File:    req.File,
		MaxSize: req.MaxSize,
	}
}
