package models

import (
	"log"
	helper "praktikum/helpers"
	"strconv"

	"gorm.io/gorm"
)

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

type IUserModel interface {
	FindUserAccount(user User) *User
	FindUser(paramId string) map[string]interface{}
	InsertUser(newUser User) *User
	SelectAllUser() []User
	UpdateUser(UserId int, newUser User) bool
	DeleteUser(UserId int) bool
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) IUserModel {
	return &UserModel {
		db: db,
	}
}

func (um *UserModel) FindUserAccount(user User) *User {
	if err := um.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func (um *UserModel) FindUser(paramId string) map[string]interface{} {
	var user User

	userId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := um.db.First(&user, userId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{}{
		"status": 200,
		"user":   user,
		"id":     userId,
	}
}

func (bm *UserModel) InsertUser(newUser User) *User {
	if err := bm.db.Create(&newUser).Error; err != nil {
		log.Fatal("Cannot Create New User! Err: " + err.Error())
		return nil
	} 

	return &newUser
}

func (bm *UserModel) SelectAllUser() []User {
	var users []User
	bm.db.Find(&users)

	if len(users) == 0 {
		log.Fatal("No User Found!")
		return nil
	}

	return users
}

func (bm *UserModel) UpdateUser(UserId int, newUser User) bool {
	if result := bm.db.Table("Users").Where("id", UserId).Updates(newUser); result.RowsAffected == 0 {
		log.Fatal("No User Updated!")
		return false
	}

	return true
}

func (bm *UserModel) DeleteUser(UserId int) bool {
	if result := bm.db.Table("Users").Where("id", UserId).Delete(&User{}, UserId); result.RowsAffected == 0 {
		log.Fatal("No User Updated!")
		return false
	}

	return true
}