package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/file_dto"
)

type FileServiceInterface interface {
	UploadFile(ctx *gin.Context, req *file_dto.UploadReq) (string, error)
	ParseExcel(ctx *gin.Context, url string, dealData func(*gin.Context, [][]string) error) error
}
