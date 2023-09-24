package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `json:"id" form:"id" gorm:"type:int(11)"`
	Name  string `json:"name" form:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" form:"email" gorm:"type:varchar(255)"`
	Blog []Blog
}