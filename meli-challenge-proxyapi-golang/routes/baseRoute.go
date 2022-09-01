package routes

import (
	"fmt"
	"meli-challenge-proxyapi-golang/controllers"

	"github.com/gorilla/mux"
)

func RegisterHomeHandler(r *mux.Router) {
	fmt.Println("API -> Register '/' route")

	r.HandleFunc("/", controllers.HomeController)
}

func RegisterApiHandler(r *mux.Router) {
	fmt.Println("API -> Register '/api' route")

	RegisterItemsHandler(r)
}
