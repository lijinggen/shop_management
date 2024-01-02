package product_dto

import (
	"github.com/shop_management/dto/common_dto"
	"time"
)

type Product struct {
	ID               string
	ImageURL         string
	StorageCode      string
	StoragePos       string
	Name             string
	Color            string
	BasePrice        float64
	CostPrice        float64
	PurchasePrice    float64
	Factory          string
	Stock            int
	InProductionNums int
	InOrderNums      int
	CreateTime       time.Time
	ModifyTime       time.Time
}

type ProductListReq struct {
	Pager *common_dto.Pager `json:"pager"`
}

type ProductListResp struct {
	Pager *common_dto.Pager `json:"pager"`
	Data  []*Product        `json:"data"`
}
