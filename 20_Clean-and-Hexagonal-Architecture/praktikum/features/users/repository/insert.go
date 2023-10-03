package repository

import entity "clean_arch/features/users/entity"

func (ur *UserRepo) Insert(input entity.UserEntity) (*entity.UserEntity, error) {
	if err := ur.db.Create(input).Error ; err != nil {
		return nil, err
	}

	return &input, nil
}
