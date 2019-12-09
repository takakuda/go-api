package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.handleFunc("/home", home).Methods("GET")
	return r
}
