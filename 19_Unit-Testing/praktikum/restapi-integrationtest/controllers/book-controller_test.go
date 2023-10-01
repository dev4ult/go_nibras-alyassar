package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	model "praktikum/models"
)

var (
	db = &gorm.DB{}
	MockBookModel = model.NewBookModel(db)
	MockBookController = NewBookController(MockBookModel)
	
	BookExampleData = model.Book{Title: "Book 1", Author: "Nibras", Publisher: "PT. Gremadia"}
)

func TestCreateBook(t *testing.T) {
	var e = echo.New()

	var req = httptest.NewRequest(http.MethodPost, "/", nil)
	var res = httptest.NewRecorder()

	var c = e.NewContext(req, res)
	c.SetPath("/books")

	var tmp = map[string]any{}

	var response = res.Body.Bytes()
	err := json.Unmarshal(response, &tmp)

	assert.Equal(t, 200, res.Code)
	assert.Error(t, err)
}