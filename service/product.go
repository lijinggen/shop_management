package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/product_dto"
)

type ProductService interface {
	Add(ctx *gin.Context, dto *product_dto.Product) error
	List(ctx *gin.Context, dto *product_dto.ProductListReq) (*product_dto.ProductListResp, error)
}
