package routes

import (
	"go-api-project/pkg/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/test/api/users", handlers.HandleGet).Methods("GET")
	router.HandleFunc("/test/api/users", handlers.HandlePost).Methods("POST")

	return router
}
