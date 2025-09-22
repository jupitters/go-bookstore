package routes

import (
	"github.com/gorilla/mux"
	"github.com/jupitters/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{Bookid}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{Bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{Bookid}", controllers.DeleteBook).Methods("DELETE")
}
