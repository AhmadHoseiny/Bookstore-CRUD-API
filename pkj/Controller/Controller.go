package Controller

import (
	"Bookstore-Backend-API/pkj/Helpers"
	"Bookstore-Backend-API/pkj/Model"
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := Model.GetAllBooks()
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	ID := Helpers.GetBookIDFromRequestBody(r)
	bookToBeReturned := Model.GetBookByID(ID)
	res, _ := json.Marshal(bookToBeReturned)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookToBeCreated Model.Book
	Helpers.GetBookFromRequestBody(r, &bookToBeCreated)
	createdBook := Model.CreateBook(&bookToBeCreated)
	res, _ := json.Marshal(createdBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	ID := Helpers.GetBookIDFromRequestBody(r)
	Model.DeleteBook(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	ID := Helpers.GetBookIDFromRequestBody(r)
	var newBook Model.Book
	Helpers.GetBookFromRequestBody(r, &newBook)
	updatedBook := Model.UpdateBook(ID, newBook)
	res, _ := json.Marshal(updatedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
