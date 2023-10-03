package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	model "praktikum/models"
)

type MockBookModel struct {}

func (bm *MockBookModel) FindBook(paramId string) map[string]interface{} {
	if paramId == "1" {
		return map[string]interface{} {
			"status": 200,
			"message": "Bad Request!",
			"id": 1,
		}	
	}

	if paramId == "satu" {
		return map[string]interface{} {
			"status": 400,
			"message": "Bad Request!",
		}	
	}

	return map[string]interface{} {
		"status": 404,
		"message": "Not Found!",
	}
}

func (bm *MockBookModel) InsertBook(newBook model.Book) *model.Book {
	return nil
}

func (bm *MockBookModel) SelectAllBook() []model.Book {
	return nil
}

func (bm *MockBookModel) UpdateBook(bookId int, newBook model.Book) bool {
	return false
}

func (bm *MockBookModel) DeleteBook(bookId int) bool {
	return false
}

type SuccessMockBookModel struct {}

func (bm *SuccessMockBookModel) FindBook(paramId string) map[string]interface{} {
	return map[string]interface{} {
		"status": 200,
		"book":   map[string]interface{}{
			"id": 1,
			"title": "book 1",
			"author": "sarbin",
			"publisher": "PT. Gremadia",
		},
		"id": 1,
	}
}

func (bm *SuccessMockBookModel) InsertBook(newBook model.Book) *model.Book {
	return &newBook
}

func (bm *SuccessMockBookModel) SelectAllBook() []model.Book {
	var books = []model.Book{{
		Id: 1,
		Title: "Book 1",
		Author: "Sarbin",
		Publisher: "PT. Gremadia",
	}}

	return books
}

func (bm *SuccessMockBookModel) UpdateBook(bookId int, newBook model.Book) bool {
	return true
}

func (bm *SuccessMockBookModel) DeleteBook(bookId int) bool {
	return true
}

func TestGetBooks(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Message string         `json:"message"`
		Books   map[string]any `json:"books"`
	}

	var tmp = Response{}


	t.Run("Fetch All Failed", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.GET("/books", controller.GetBooks())

		var req = httptest.NewRequest(http.MethodGet, "/books", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)

		assert.NotNil(t, tmp)
		
		assert.Nil(t, err)
		assert.Nil(t, tmp.Books)
	})

	t.Run("Fetch All Success", func(t *testing.T) {
		var model = SuccessMockBookModel{}
		var controller = NewBookController(&model)

		e.GET("/books", controller.GetBooks())

		var req = httptest.NewRequest(http.MethodGet, "/books", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var resData = json.NewDecoder(res.Result().Body)
		result := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)

		assert.NotNil(t, tmp)
		assert.NotNil(t, result)

		fmt.Println(tmp.Books)
		
	})
}

func TestCreateBook(t *testing.T) {
	
	var e = echo.New()
	
	t.Run("No Book Created", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)
		
		e.POST("/books", controller.CreateBook())

		var req = httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader([]byte(`{"test": "test"}`)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type Response struct {
			Status int `json:"status"`
			Message string `json:"message"`
		}

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Something went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		
		assert.Nil(t, err)
	})

	t.Run("Success Book Created", func(t *testing.T) {
		var reqData string = `{"title": "Book 1", "author": "Sarbin", "publisher": "PT. Gremadia"}`

		var model = SuccessMockBookModel{}
		var controller = NewBookController(&model)

		e.POST("/books", controller.CreateBook())

		var req = httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type Response struct {
			Status int `json:"status"`
			Message string `json:"message"`
			Book map[string]any `json:"book"`
		}

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "Book Created!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.NotNil(t, tmp.Book)

		assert.Nil(t, err)
	})
}

func TestGetBook(t *testing.T) {
	var e = echo.New()

	type FailedResponse struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.GET("/books/:id", controller.GetBook())

		var req = httptest.NewRequest(http.MethodGet, "/books/satu", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = FailedResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Book Not Found", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.GET("/books/:id", controller.GetBook())

		var req = httptest.NewRequest(http.MethodGet, "/books/999", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = FailedResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Success Get Book", func(t *testing.T) {
		var model = SuccessMockBookModel{}
		var controller = NewBookController(&model)

		e.GET("/books/:id", controller.GetBook())

		var req = httptest.NewRequest(http.MethodGet, "/books/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type SuccessResponse struct {
			Status int `json:"status"`
			Message string `json:"message"`
			Book map[string]any `json:"book"`
		}

		var tmp = SuccessResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "Book Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.NotNil(t, tmp.Book)
		assert.Nil(t, err)
	})
}

func TestEditBook(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	var reqData = `{"title": "Book 1", "author": "Sarbin", "publisher": "PT. Gremadia"}`

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.PUT("/books/:id", controller.EditBook())

		var req = httptest.NewRequest(http.MethodPut, "/books/satu", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Book Not Found", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.PUT("/books/:id", controller.EditBook())

		var req = httptest.NewRequest(http.MethodPut, "/books/999", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Update Failed", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.PUT("/books/:id", controller.EditBook())

		var req = httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, http.StatusInternalServerError, tmp.Status)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Update Success", func(t *testing.T) {
		var model = SuccessMockBookModel{}
		var controller = NewBookController(&model)

		e.PUT("/books/:id", controller.EditBook())

		var req = httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "Book Updated!", tmp.Message)

		assert.Nil(t, err)
	})
}

func TestRemoveBook(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.DELETE("/books/:id", controller.RemoveBook())

		var req = httptest.NewRequest(http.MethodDelete, "/books/satu",nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Book Not Found", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.DELETE("/books/:id", controller.RemoveBook())

		var req = httptest.NewRequest(http.MethodDelete, "/books/999",nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Delete Failed", func(t *testing.T) {
		var model = MockBookModel{}
		var controller = NewBookController(&model)

		e.DELETE("/books/:id", controller.RemoveBook())

		var req = httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, http.StatusInternalServerError, tmp.Status)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Delete Success", func(t *testing.T) {
		var model = SuccessMockBookModel{}
		var controller = NewBookController(&model)

		e.DELETE("/books/:id", controller.RemoveBook())

		var req = httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "Book Deleted!", tmp.Message)

		assert.Nil(t, err)
	})
}