package main

import (
	"log"
	"net/http"

	"github.com/hitheshthummala/rest_crud/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
