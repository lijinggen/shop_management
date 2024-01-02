package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/product_dto"
	"gorm.io/gorm"
)

type ProductRepo interface {
	AddProduct(ctx *gin.Context, db *gorm.DB, dto *product_dto.Product) error
}
