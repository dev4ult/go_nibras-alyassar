package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Id      int    `json:"id" form:"id" gorm:"type:int(11)"`
	UserId  int    `json:"user_id" form:"user_id" gorm:"type:int(11)"`
	Title   string `json:"title" form:"title" gorm:"type:varchar(100)"`
	Content string `json:"content" form:"content" gorm:"type:varchar(255)"`
}