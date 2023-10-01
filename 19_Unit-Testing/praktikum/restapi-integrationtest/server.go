package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	config "praktikum/config"
	controller "praktikum/controllers"
	middleware "praktikum/middlewares"
	model "praktikum/models"
	route "praktikum/routes"
)

var (
    db *gorm.DB = config.InitDB()

    userModel model.IUserModel = model.NewUserModel(db)
    bookModel model.IBookModel = model.NewBookModel(db)

    userController controller.IUserController = controller.NewUserController(userModel)
    bookController controller.IBookController = controller.NewBookController(bookModel)
)

func main() {
    godotenv.Load(".env")

    e := echo.New()

    middleware.Logger(e)

    route.UserRoutes(e, userController)
    route.BookRoutes(e, bookController)


    e.Logger.Fatal(e.Start(":8000"))
}
