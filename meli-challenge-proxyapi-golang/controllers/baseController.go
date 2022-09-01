package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gophers!")
}
