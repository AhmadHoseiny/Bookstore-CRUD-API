package Routes

import (
	"Bookstore-Backend-API/pkj/Controller"
	"github.com/gorilla/mux"
)

var SetRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", Controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{ID}", Controller.GetBookByID).Methods("GET")
	router.HandleFunc("/book", Controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{ID}", Controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{ID}", Controller.UpdateBook).Methods("PUT")
}
