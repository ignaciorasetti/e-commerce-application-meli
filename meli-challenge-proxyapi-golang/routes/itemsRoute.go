package routes

import (
	"fmt"
	"meli-challenge-proxyapi-golang/controllers"

	"github.com/gorilla/mux"
)

func RegisterItemsHandler(r *mux.Router) {
	fmt.Println("API -> Register '/api/items' routes")

	r.HandleFunc("/items", controllers.GetItemsByQuery).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/items/{id}", controllers.GetItemById).Methods("GET")
}
