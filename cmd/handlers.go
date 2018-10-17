package main

import (
	"encoding/json"
	"net/http"
	"poc-company/company"

	"github.com/go-chi/chi"
)

// Params default from url
const (
	URLParamID = "id"
)

// versionHandler jusnt ping application
func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// getCompanyHandler get one company by id
func getCompanyHandler(svc *company.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company, _ := svc.Get(chi.URLParam(r, URLParamID))
		json.NewEncoder(w).Encode(company)
	}
}
