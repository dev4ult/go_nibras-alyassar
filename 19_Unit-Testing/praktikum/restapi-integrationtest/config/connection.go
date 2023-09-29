package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "praktikum/models"
)

  var DB *gorm.DB

  func ConnectDB() {
	dsn := "root@tcp(127.0.0.1:3306)/try_middleware?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	initialMigration()
  }

  func initialMigration() {
	DB.AutoMigrate(&model.Book{})
  }

