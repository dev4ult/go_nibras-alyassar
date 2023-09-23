package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status	int
	Message string
}

type Product struct {
	Id 		int
	Name 	string
	Stock 	int
}

type User struct {
	Username 	string	`json:"username" form:"username"`
	Email 		string	`json:"email" form:"email"`
	password 	string	`json:"password" form:"password"`
}

func main() {
	e := echo.New()

	e.GET("/", homepage)
	e.GET("/dashboard", dashboard)
	e.POST("/register", register)
	e.GET("/product/:id", product)
	e.Logger.Fatal(e.Start(":8000"))
}

func homepage(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World!")
}

func dashboard(e echo.Context) error {
	response := Response{ Status: 200, Message: "This is dashboard page"}

	return e.JSON(http.StatusOK, response)
}

func register(e echo.Context) error {
	user := User{}

	e.Bind(&user)
	
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func product(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	display := e.QueryParam("display")

	if err != nil {
		response := Response{Status: 400, Message: "Id is not parseable"}
		return e.JSON(http.StatusBadRequest, response)
	}

	product := Product{Id: id, Name: "Botol Kecap", Stock: 10}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
		"display": display,
	})
}