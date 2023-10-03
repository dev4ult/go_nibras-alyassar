package mocks

import "clean_arch/features/user/entity"

type MockRepo struct{}

func (mr *MockRepo) Insert(input entity.UserEntity) (*entity.UserEntity, error) {
	if err := mr.db.Create(input).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func (mr *MockRepo) SelectAll() ([]entity.UserEntity, error) {
	var users []entity.UserEntity
	if err := mr.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (mr *MockRepo) SelectById(userId int) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := mr.db.First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (mr *MockRepo) SelectByUsername(username string) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := mr.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}