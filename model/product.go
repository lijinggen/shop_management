package model

import "time"

type Product struct {
	BaseModel
	ID               string    `gorm:"type:varchar(36);primaryKey"`
	ImageURL         string    `gorm:"type:text"`
	StorageCode      string    `gorm:"type:varchar(255)"`
	StoragePos       string    `gorm:"type:varchar(255)"`
	Name             string    `gorm:"type:varchar(255)"`
	Color            string    `gorm:"type:varchar(255)"`
	BasePrice        float64   `gorm:"type:decimal(10,2)"`
	CostPrice        float64   `gorm:"type:decimal(10,2)"`
	PurchasePrice    float64   `gorm:"type:decimal(10,2)"`
	Factory          string    `gorm:"type:varchar(255)"`
	Stock            int       `gorm:"type:int"`
	InProductionNums int       `gorm:"type:int"`
	InOrderNums      int       `gorm:"type:int"`
	CreateTime       time.Time `gorm:"type:datetime"`
	ModifyTime       time.Time `gorm:"type:datetime"`
}

func (p *Product) TableName() string {
	return "product"
}
