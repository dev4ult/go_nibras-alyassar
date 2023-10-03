package repository

import (
	"gorm.io/gorm"

	entity "clean_arch/features/users/entity"
)

type IUserRepo interface {
	Insert(input entity.UserEntity) (*entity.UserEntity, error)
	SelectAll() ([]entity.UserEntity, error)
	SelectById(userId int) (*entity.UserEntity, error)
	SelectByUsername(username string) (*entity.UserEntity, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo {
		db: db,
	}
}