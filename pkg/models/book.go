package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jupitters/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	books := []Book{}
	db.Find(&books)
	return books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	getBook := Book{}
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
