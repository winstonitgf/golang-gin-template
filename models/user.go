package models

type User struct {
	Id        uint   `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	Account   string `json:"account" form:"account"`
	Password  string `json:"-" form:"password"`
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	AvatarUrl string `json:"avatar_url" form:"avatar_url"`
	Token     string `json:"token" form:"token"`
	Balance   uint   `json:"balance" form:"balance"`
	Status    uint   `json:"status" form:"status"`
	CreatedAt uint   `json:"created_at"`
	UpdatedAt uint   `json:"updated_at" gorm:"type:int"`
	DeletedAt *uint  `json:"deleted_at" redis:"-"`
}
