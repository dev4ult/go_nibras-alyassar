package users

import (
	config "clean_arch/config"
	handler "clean_arch/features/users/handler"
	repo "clean_arch/features/users/repository"
	service "clean_arch/features/users/service"
)

func UserRouteHandler() handler.IUserController {
	db := config.InitDB() 
	repo := repo.NewUserRepo(db)
	service := service.NewUserService(repo)
	return handler.NewUserController(service)
}