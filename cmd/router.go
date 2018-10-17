package main

import (
	"net/http"
	"poc-company/company"

	"github.com/go-chi/chi"
)

func createHandler(service *company.Service) http.Handler {

	router := chi.NewRouter()

	router.Route("/company", func(router chi.Router) {
		router.Get("/version", versionHandler)
		router.Get("/{id}", getCompanyHandler(service))
	})

	return router
}
