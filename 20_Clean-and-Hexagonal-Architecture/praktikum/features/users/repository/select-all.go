package repository

import entity "clean_arch/features/users/entity"

func (ur *UserRepo) SelectAll() ([]entity.UserEntity, error) {
	var users []entity.UserEntity
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}