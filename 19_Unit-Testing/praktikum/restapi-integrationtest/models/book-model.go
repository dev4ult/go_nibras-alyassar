package models

import (
	"log"
	"strconv"

	"gorm.io/gorm"

	helper "praktikum/helpers"
)

type Book struct {
	gorm.Model
	Id        int    `json:"id" form:"id" gorm:"type:int(11)"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(100)"`
	Author    string `json:"author" form:"author" gorm:"type:varchar(255)"`
	Publisher string `json:"publisher" form:"publisher" gorm:"type:varchar(255)"`
}

type IBookModel interface {
	FindBook(paramId string) map[string]interface{}
	InsertBook(newBook Book) *Book
	SelectAllBook() []Book
	UpdateBook(bookId int, newBook Book) bool
	DeleteBook(bookId int) bool
}

type BookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) IBookModel {
	return &BookModel{
		db: db,
	}
}

func (bm *BookModel) FindBook(paramId string) map[string]interface{} {
	var book Book

	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		return helper.Response(400, "Bad Request!")
	}

	result := bm.db.First(&book, bookId)

	if result.RowsAffected < 1 {
		return helper.Response(404, "Not Found!")
	}

	return map[string]interface{} {
		"status": 200,
		"book":   book,
		"id":     bookId,
	}
}

func (bm *BookModel) InsertBook(newBook Book) *Book {
	if err := bm.db.Create(&newBook).Error; err != nil {
		log.Fatal("Cannot Create New Book! Err: " + err.Error())
		return nil
	} 

	return &newBook
}

func (bm *BookModel) SelectAllBook() []Book {
	var books []Book
	bm.db.Find(&books)

	if len(books) == 0 {
		log.Fatal("No Book Found!")
		return nil
	}

	return books
}

func (bm *BookModel) UpdateBook(bookId int, newBook Book) bool {
	if result := bm.db.Table("books").Where("id", bookId).Updates(newBook); result.RowsAffected == 0 {
		log.Fatal("No Book Updated!")
		return false
	}

	return true
}

func (bm *BookModel) DeleteBook(bookId int) bool {
	if result := bm.db.Table("books").Where("id", bookId).Delete(&Book{}, bookId); result.RowsAffected == 0 {
		log.Fatal("No Book Updated!")
		return false
	}

	return true
}