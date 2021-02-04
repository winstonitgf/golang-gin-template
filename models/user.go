package models

type User struct {
	Common
	Account  string `json:"account" form:"account" gorm:"size:100;index:ix_account_type_status"`
	Password string `json:"-" gorm:"size:200"`
	Nickname string `json:"nickname" form:"nickname"  gorm:"size:50"`
	Email    string `json:"email" form:"email" gorm:"size:100"`
	Phone    string `json:"phone" form:"phone" gorm:"size:20"`
	Token    string `json:"token" form:"token" gorm:"size:1000"`
	RoleID   uint   `json:"role_id" form:"role_id"`
}
