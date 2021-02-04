package models

import (
	"gorm.io/gorm"
)

type Common struct {
	ID        int            `json:"id" form:"id"`
	CreatedAt int64          `json:"created_at" form:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updated_at" form:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
