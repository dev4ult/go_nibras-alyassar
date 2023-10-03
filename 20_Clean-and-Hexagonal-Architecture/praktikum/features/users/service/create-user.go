package service

import (
	dto "clean_arch/features/users/dtos"
	entity "clean_arch/features/users/entity"
)

func(us *UserService) CreateUser(input dto.UserInput) (*dto.UserResponse, string) {
	user, err := us.repo.SelectByUsername(input.Username)

	if err != nil {
		return nil, err.Error()
	}

	if *user != (entity.UserEntity{}) {
		return nil, "User has already exist!"
	}

	// need to map dto to entity
	newUser, errCreate := us.repo.Insert(input)

	if errCreate != nil {
		return nil, err.Error()
	}

	// need to map entity to dto
	return newUser, "Success Created!"
}