package main

import (
	"Bookstore-Backend-API/pkj/Routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	Routes.SetRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
