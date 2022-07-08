package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (model *Model) BeforeCreate(db *gorm.DB) (err error) {
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now().UnixMilli()
	model.UpdatedAt = time.Now().UnixMilli()
	return
}
