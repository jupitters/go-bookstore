package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jupitters/go-bookstore/pkg/models"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
