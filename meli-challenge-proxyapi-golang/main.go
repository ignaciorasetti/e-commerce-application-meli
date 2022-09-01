package main

import (
	"log"
	"meli-challenge-proxyapi-golang/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	home := r.PathPrefix("/").Subrouter()
	routes.RegisterHomeHandler(home)

	api := r.PathPrefix("/api").Subrouter()
	routes.RegisterApiHandler(api)

	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()
	headers := handlers.AllowedHeaders([]string{
		"Origin", "Access-Control, Allow-Origin", "Content-Type",
		"Accept", "Authorization", "Origin, Accept", "X-Requested-With",
		"Access-Control-Request-Method", "Access-Control-Request-Header",
	})
	methodsObj := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	server := handlers.CORS(credentials, origins, methodsObj, headers)(r)

	//http.ListenAndServe(":3001", r)
	if err := http.ListenAndServe("localhost:3001", server); err != nil {
		log.Fatal("API -> Serving error")
	}
}
