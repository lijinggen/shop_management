package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
}

func (t BaseModel) BeforeCreate(tx *gorm.DB) (e error) {
	now := time.Now()
	tx.Statement.SetColumn("modify_time", now)
	tx.Statement.SetColumn("create_time", now)
	tx.Statement.SetColumn("id", uuid.NewString())
	return
}

func (t BaseModel) BeforeUpdate(tx *gorm.DB) (e error) {
	now := time.Now()
	tx.Statement.SetColumn("modify_time", now)
	return
}

func (t BaseModel) BeforeDelete(tx *gorm.DB) (e error) {
	now := time.Now()
	tx.Statement.SetColumn("deleted_time", now)
	return
}
