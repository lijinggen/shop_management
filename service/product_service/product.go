package product_service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/product_dto"
	"github.com/shop_management/repository"
	"github.com/shop_management/repository/product_repo"
	"github.com/shop_management/service"
	"github.com/shop_management/util"
)

type productServiceImpl struct {
	productRepo repository.ProductRepo
}

func NewProductServiceImpl() service.ProductService {
	return &productServiceImpl{
		productRepo: product_repo.NewProductRepoImpl(),
	}
}

func (p *productServiceImpl) Add(ctx *gin.Context, dto *product_dto.Product) error {
	err := p.productRepo.AddProduct(ctx, util.GetDBFromContext(ctx), dto)
	if err != nil {
		return err
	}
	return nil
}

func (p *productServiceImpl) List(ctx *gin.Context, dto *product_dto.ProductListReq) (*product_dto.ProductListResp, error) {
	return nil, nil
}
