package common

import "time"

type Pagination struct {
	Take int `json:"take" form:"take"`
	Page int `json:"page" form:"page"`
}

type Model struct {
	Id         int        `gorm:"primary_key" json:"id"`
	CreatedAt  int        `json:"created_at"`
	ModifiedAt int        `json:"modified_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type ErrorModel struct {
	Error     error `json:"error" form:"error"`
	ErrorCode int   `json:"error_code" form:"error_code"`
}