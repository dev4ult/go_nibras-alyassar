package service

import (
	dto "clean_arch/features/users/dtos"
)

func(us *UserService) FetchAll() ([]dto.UserResponse, string) {
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