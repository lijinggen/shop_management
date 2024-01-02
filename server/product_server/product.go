package product_server

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/product_dto"
	"github.com/shop_management/po/common_po"
	"github.com/shop_management/po/product_po"
	"github.com/shop_management/service"
	"github.com/shop_management/service/product_service"
	"github.com/shop_management/sm_error"
)

type ProductServer struct {
	productService service.ProductService
}

func NewProductServer() *ProductServer {
	return &ProductServer{
		productService: product_service.NewProductServiceImpl(),
	}
}

func (p ProductServer) Add(ctx *gin.Context) (interface{}, error) {
	dto := &product_po.Product{}
	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		return nil, sm_error.NewParamHttpError(err)
	}
	err = p.productService.Add(ctx, &product_dto.Product{
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
	})
	if err != nil {
		return nil, err
	}
	return &common_po.CommonResp{}, nil
}

func (p ProductServer) List(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}
