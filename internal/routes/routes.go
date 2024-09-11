package routes

import (
	"github.com/gorilla/mux"
	"github.com/hitheshthummala/rest_crud/internal/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/items", handlers.GetItems).Methods("Get")
	return r
}
