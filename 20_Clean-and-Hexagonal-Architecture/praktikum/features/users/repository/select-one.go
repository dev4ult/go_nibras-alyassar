package repository

import entity "clean_arch/features/users/entity"

func (ur *UserRepo) SelectById(userId int) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := ur.db.First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) SelectByUsername(username string) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}