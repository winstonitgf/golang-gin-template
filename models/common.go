package models

type Common struct {
	ID        int   `json:"id" form:"id" redis:"id"`
	CreatedAt int64 `json:"created_at" form:"created_at" gorm:"autoCreateTime" redis:"created_at"`
	UpdatedAt int64 `json:"updated_at" form:"updated_at" gorm:"autoCreateTime" redis:"updated_at"`
	DeletedAt int64 `json:"deleted_at" form:"deleted_at" redis:"deleted_at"`
}
