package models

type User struct {
	Id       int    `json:"id" form:"id" gorm:"type:int(11)"`
	Username string `json:"username" form:"username" gorm:"type:varchar(100)"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(100)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255)"`
}

type UserResponse struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}