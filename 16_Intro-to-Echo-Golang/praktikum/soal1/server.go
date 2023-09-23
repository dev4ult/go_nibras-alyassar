package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id int 			`json:"id" form:"id"`
	Name string 	`json:"name" form:"name"`
	Email string 	`json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Response struct {
	Status 	int
	Message string
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}

func GetUserByID(qParam string) (User, error, int) {
	userId, err := strconv.Atoi(qParam)

	var index = -1

	var user User
	
	if err != nil {
		fmt.Println(err.Error())
		return user, err, index
	}


	for indexRow, userRow := range users {
		fmt.Println(userRow.Id, userId)
		if userRow.Id == userId {
			user = userRow
			index = indexRow
		}
	}

	return user, nil, index
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here

	user, err, index := GetUserByID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 400, Message: "Bad Request!"})
	}

	if index == -1 {
		return c.JSON(http.StatusNotFound, Response{Status: 404, Message: "User Not Found!"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User Found!",
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	_, err, index := GetUserByID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 400, Message: "Bad Request!"})
	}

	if index == -1 {
		return c.JSON(http.StatusNotFound, Response{Status: 404, Message: "User Not Found!"})
	}

	users = append(users[:index], users[index+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User Deleted!",
		"users": users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	_, err, index := GetUserByID(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 400, Message: "Bad Request!"})
	}

	if index == -1 {
		return c.JSON(http.StatusNotFound, Response{Status: 404, Message: "User Not Found!"})
	}

	var user = User{}

	c.Bind(&user)

	users[index] = user

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "User Updated",
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
// binding data
	user := User{}
	c.Bind(&user)
	
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User Created!",
		"user": user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
