package Helpers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetBookFromRequestBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(body, &x)
	if err != nil {
		panic(err)
		return
	}
}

func GetBookIDFromRequestBody(r *http.Request) int64 {
	vars := mux.Vars(r)
	bookID := vars["ID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		panic(err)
		return 0
	}
	return ID
}
