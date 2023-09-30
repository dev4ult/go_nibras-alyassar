package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/try_middleware?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// initialMigration()

	return db
}

func initialMigration() {
	// DB.AutoMigrate(&model.Book{})
}

