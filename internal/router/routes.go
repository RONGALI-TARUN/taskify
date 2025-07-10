package router

import (
	"github.com/go-chi/chi/v5"

)

func loadRoutes() {
	router := chi.NewRouter()
	router.HandleFunc("/", testHandler)
	router.Route("/",all)
}
