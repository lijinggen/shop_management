package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shop_management/vars"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDB() (*gorm.DB, error) {
	dbName := os.Getenv("SM_DB_NAME")
	if dbName == "" {
		dbName = "shop_management_dev"
	}
	dsn := fmt.Sprintf("root:ljg098098@tcp(120.24.169.86:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Config.CreateBatchSize = 100
	// 设置创建时的插件
	return db.Debug(), nil
}

func GetDBFromContext(ctx *gin.Context) *gorm.DB {
	value, _ := ctx.Get(vars.DbMetadataName)
	return value.(*gorm.DB).Debug()
}
