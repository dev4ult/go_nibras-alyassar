package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id        int    `json:"id" form:"id" gorm:"type:int(11)"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(100)"`
	Author    string `json:"author" form:"author" gorm:"type:varchar(255)"`
	Publisher string `json:"publisher" form:"publisher" gorm:"type:varchar(255)"`
}