package product_repo

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/product_dto"
	"github.com/shop_management/model"
	"github.com/shop_management/repository"
	"gorm.io/gorm"
)

type productRepoImpl struct {
}

func NewProductRepoImpl() repository.ProductRepo {
	return &productRepoImpl{}
}

func (p *productRepoImpl) AddProduct(ctx *gin.Context, db *gorm.DB, dto *product_dto.Product) error {
	m := &model.Product{
		ID:               dto.ID,
		ImageURL:         dto.ImageURL,
		StorageCode:      dto.StorageCode,
		StoragePos:       dto.StoragePos,
		Name:             dto.Name,
		Color:            dto.Color,
		BasePrice:        dto.BasePrice,
		CostPrice:        dto.CostPrice,
		PurchasePrice:    dto.PurchasePrice,
		Factory:          dto.Factory,
		Stock:            dto.Stock,
		InProductionNums: dto.InProductionNums,
		InOrderNums:      dto.InOrderNums,
		CreateTime:       dto.CreateTime,
		ModifyTime:       dto.ModifyTime,
	}
	err := db.Create(m).Error
	if err != nil {
		return err
		//return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}
