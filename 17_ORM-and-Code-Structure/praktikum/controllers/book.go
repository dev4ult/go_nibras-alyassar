package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
	model "praktikum/models"
	helper "praktikum/helpers"
)

type BookController struct {
	model model.Book
}

func FindBook(paramId string) map[string]interface{} {
	var book model.Book

	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := config.DB.First(&book, bookId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{}{
		"status": 200,
		"book":   book,
		"id":     bookId,
	}
}

func (b *BookController) GetBooks() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var books []model.Book
	
		err := config.DB.Find(&books).Error
	
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, err.Error()))
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"status": 200,
			"message": "Books Listed!",
			"books": books,
		})
	}
}

func CreateBook(ctx echo.Context) error {
	var book model.Book

	ctx.Bind(&book)

	result := config.DB.Create(&book)

	if result.Error != nil || result.RowsAffected < 1 {
		ctx.JSON(http.StatusInternalServerError, helper.Response(500, result.Error.Error()))
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

	result := config.DB.Table("books").Where("id", book["id"]).Updates(newBookData)

	if result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "Book Updated!"))
}

func DeleteBook(ctx echo.Context) error {
	book := FindBook(ctx.Param("id"))

	if book["status"] != http.StatusOK {
		return ctx.JSON(book["status"].(int), book)
	}

	result := config.DB.Delete(&model.Book{}, book["id"])

	if result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, helper.Response(200, "book Deleted!"))
}