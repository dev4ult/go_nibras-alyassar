package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
	model "praktikum/models"
	util "praktikum/utils"
)

func FindBook(paramId string) map[string]interface{} {
	var book model.Book

	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		return util.Response(400, "Bad Request!")
	}

	result := config.DB.First(&book, bookId)

	if result.RowsAffected < 1 {
		return util.Response(404, "Not Found!")
	}

	return map[string]interface{} {
		"status": 200,
		"book":   book,
		"id":     bookId,
	}
}

func GetBooks(ctx echo.Context) error {
	var books []model.Book

	err := config.DB.Find(&books).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, err.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"message": "Books Listed!",
		"books": books,
	})
}

func CreateBook(ctx echo.Context) error {
	var book model.Book

	ctx.Bind(&book)

	err := config.DB.Create(&book).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.Response(500, err.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "Book Created!",
		"book": book,
	})
}

func GetBook(ctx echo.Context) error {
	book := FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "Book Found!",
		"book": book["book"],
	})
}

func UpdateBook(ctx echo.Context) error {
	book := FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	var newBookData model.Book

	ctx.Bind(&newBookData)

	err := config.DB.Table("books").Where("id", book["id"]).Updates(newBookData).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "Book Updated!"))
}

func DeleteBook(ctx echo.Context) error {
	book := FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	err := config.DB.Delete(&model.Book{}, book["id"]).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "book Deleted!"))
}