package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	helper "praktikum/helpers"
	model "praktikum/models"
)

type IBookController interface {
	GetBooks() echo.HandlerFunc
	CreateBook() echo.HandlerFunc
	GetBook() echo.HandlerFunc
	EditBook() echo.HandlerFunc
	RemoveBook() echo.HandlerFunc
}

type BookController struct {
	model model.IBookModel
}

func NewBookController(model model.IBookModel) IBookController {
	return &BookController{
		model: model,
	}
}

func (bc *BookController) GetBooks() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		books := bc.model.SelectAllBook()
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"message": "Books Listed!",
			"books": books,
		})
	}
}

func (bc *BookController) CreateBook() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var book model.Book
	
		ctx.Bind(&book)
	
		result := bc.model.InsertBook(book)
	
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"status": 200,
			"message": "Book Created!",
			"book": book,
		})
	}
}

func (bc *BookController) GetBook() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		result := bc.model.FindBook(ctx.Param("id"))
	
		if result["status"] != http.StatusOK {
			return ctx.JSON(result["status"].(int), result)
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"status": 200,
			"message": "Book Found!",
			"book": result["book"],
		})
	}
}

func (bc *BookController) EditBook() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		book := bc.model.FindBook(ctx.Param("id"))
	
		if book["status"] != http.StatusOK {
			return ctx.JSON(book["status"].(int), book)
		}
	
		var newBook model.Book
	
		ctx.Bind(&newBook)
	
		update := bc.model.UpdateBook(book["id"].(int), newBook)
	
		if !update {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, helper.Response(200, "Book Updated!"))
	}
}

func (bc *BookController) RemoveBook() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		book := bc.model.FindBook(ctx.Param("id"))
	
		if book["status"] != http.StatusOK {
			return ctx.JSON(book["status"].(int), book)
		}
	
		delete := bc.model.DeleteBook(book["id"].(int))
	
		if !delete {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}

		return ctx.JSON(http.StatusOK, helper.Response(200, "book Deleted!"))
	}
}