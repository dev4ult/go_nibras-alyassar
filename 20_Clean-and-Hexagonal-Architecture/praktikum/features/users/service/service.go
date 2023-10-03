package service

import (
	dto "clean_arch/features/users/dtos"
	entity "clean_arch/features/users/entity"
)

func (us *UserService) FetchAll() ([]dto.UserResponse, string) {
	users, err := us.repo.SelectAll()

	if err != nil {
		return nil, err.Error()
	}

	if len(users) == 0 {
		return nil, "No User Listed!"
	}

	// need to map entity to dto
	return users, "Success Get Users!"
}

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