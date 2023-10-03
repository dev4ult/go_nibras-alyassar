package mocks

import (
	"clean_arch/features/user/entity"

	"github.com/mashingan/smapping"
)

type MockService struct{}

func (us *MockService) FetchAll() ([]dto.UserResponse, string) {
	users, err := us.repo.SelectAll()

	if err != nil {
		return nil, err.Error()
	}

	if len(users) == 0 {
		return nil, "No User Listed!"
	}

	var response []dto.UserResponse
	errMapping := smapping.FillStruct(&response, smapping.MapFields(&users))

	if errMapping != nil {
		return nil, "Cannot Mapping entity to DTO"
	}

	// need to map entity to dto
	return response, "Success Get Users!"
}

func (us *MockService) CreateUser(input dto.UserInput) (*dto.UserResponse, string) {
	user, err := us.repo.SelectByUsername(input.Username)

	if err != nil {
		return nil, err.Error()
	}

	if *user != (entity.UserEntity{}) {
		return nil, "User has already exist!"
	}

	var userEntity entity.UserEntity
	errMappingDto := smapping.FillStruct(&userEntity, smapping.MapFields(&input))

	if errMappingDto != nil {
		return nil, "Cannot Mapping DTO to entity"
	}

	// need to map dto to entity
	newUser, errCreate := us.repo.Insert(userEntity)

	if errCreate != nil {
		return nil, err.Error()
	}

	var userDTO dto.UserResponse
	errMappingEntity := smapping.FillStruct(&userDTO, smapping.MapFields(&newUser))

	if errMappingEntity != nil {
		return nil, "Cannot Mapping entity to DTO"
	}

	// need to map entity to dto
	return &userDTO, "Success Created!"
}