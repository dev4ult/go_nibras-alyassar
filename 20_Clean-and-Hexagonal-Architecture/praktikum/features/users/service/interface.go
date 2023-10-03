package service

import (
	dto "clean_arch/features/users/dtos"
	repo "clean_arch/features/users/repository"
)

type IUserService interface {
	FetchAll() ([]dto.UserResponse, string)
	CreateUser(input dto.UserInput) (*dto.UserResponse, string)
}

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) IUserService {
	return &UserService{
		repo: repo,
	}
}
